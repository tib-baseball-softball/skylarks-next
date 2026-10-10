import { self } from "$app/service-worker";
import { version } from "$app/env";
import { immutable, assets } from "$app/manifest";
import { resolve } from "$app/paths";

import type { Extension } from "../lib/dp/types/ExpandedResponse";
import { PUBLIC_APPLICATION_CONTEXT } from "$app/env/public";
import { ApplicationContext } from "../lib/dp/types/ApplicationContext";

// Create a unique cache name for this deployment
const CACHE = `cache-${version}`;

const ASSETS = [
  ...immutable.map((asset) => resolve(asset.path)), // the Vite output
  ...assets.map((asset) => resolve(asset.path)), // everything in `static`
];

/**
 * Caching mechanism copied verbatim from SvelteKit documentation.
 */
self.addEventListener("install", (event) => {
  // Create a new cache and add all files to it

  async function addFilesToCache() {
    const cache = await caches.open(CACHE);
    await cache.addAll(ASSETS);
  }

  event.waitUntil(addFilesToCache());
});

self.addEventListener("activate", (event) => {
  // Remove previous cached data from disk
  async function deleteOldCaches() {
    for (const key of await caches.keys()) {
      if (key !== CACHE) await caches.delete(key);
    }
  }

  event.waitUntil(deleteOldCaches());
});

type NotificationAction = { action: string; title: string; navigate?: string };
type PushOptions = Extension<
  NotificationOptions,
  { title: string; actions?: NotificationAction[] }
>;

self.addEventListener("push", function (event) {
  const payload = event.data?.json() as PushOptions;
  if (PUBLIC_APPLICATION_CONTEXT !== ApplicationContext.Production) {
    console.log("Received a push message", payload);
  }

  const options: NotificationOptions = {
    body: payload.body,
    tag: payload.tag,
    data: payload.data,
    // @ts-expect-error - the property is not present on NotificationOptions but is definitely used in Chromium
    actions: payload.actions,
  };

  event.waitUntil(self.registration.showNotification(payload.title, options));
});

self.addEventListener("notificationclick", async (event) => {
  if (PUBLIC_APPLICATION_CONTEXT !== ApplicationContext.Production) {
    console.log("Notification clicked, action:", event.action);
    console.log("Notification clicked, notification:", event.notification);
  }

  if (event.action === "link") {
    event.notification.close();
    await self.clients.openWindow(event.notification.data.navigate);
  }
});

self.addEventListener("fetch", (event) => {
  if (event.request.method !== "GET") {
    return;
  }

  if (event.request.url.includes("/api/")) {
    return;
  }

  async function respond() {
    const url = new URL(event.request.url);
    const cache = await caches.open(CACHE);

    // `build`/`files` can always be served from the cache
    // @ts-expect-error - ASSETS is very strictly typed, pathname is just a string
    if (ASSETS.includes(url.pathname)) {
      const response = await cache.match(url.pathname);

      if (response) {
        return response;
      }
    }

    // for everything else, try the network first, but
    // fall back to the cache if we're offline
    try {
      const response = await fetch(event.request);

      // if we're offline, fetch can return a value that is not a Response
      // instead of throwing - and we can't pass this non-Response to respondWith
      if (!(response instanceof Response)) {
        throw new Error("invalid response from fetch");
      }

      if (response.status === 200) {
        cache.put(event.request, response.clone());
      }

      return response;
    } catch (err) {
      const response = await cache.match(event.request);

      if (response) {
        return response;
      }

      // if there's no cache, then just error out
      // as there is nothing we can do to respond to this request
      throw err;
    }
  }

  event.respondWith(respond());
});
