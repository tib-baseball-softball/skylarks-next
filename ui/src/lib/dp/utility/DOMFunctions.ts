import {micromark} from "micromark";

/**
 * Converts Markdown to HTML using the micromark library.
 *
 * Single utility function to be able to easily switch out the library later.
 *
 * @param markdown the input value, usually output of an RTE component
 */
export function markdownToHTML(markdown: string | Uint8Array<ArrayBufferLike>): string {
  return micromark(markdown);
}

export function stripTags(html: string) {
  let doc = new DOMParser().parseFromString(html, 'text/html');
  return doc.body.textContent || "";
}
