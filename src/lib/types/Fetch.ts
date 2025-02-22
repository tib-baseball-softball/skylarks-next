/**
 * I am passing a lot of fetch functions around (SvelteKit load functions provide one),
 * and the type definition is pretty long.
 */
export type Fetch = {
  (input: RequestInfo | URL, init?: RequestInit): Promise<Response>;
  (input: string | URL | globalThis.Request, init?: RequestInit): Promise<Response>;
}