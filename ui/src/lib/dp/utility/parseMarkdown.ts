import {micromark} from "micromark";

/**
 * Converts Markdown to HTML using the micromark library.
 *
 * Used on after a form submit action to save HTML in the database
 * so it doesn't need to be parsed on the client on each render.
 *
 * Single utility function to be able to easily switch out the library later.
 *
 * @param markdown the input value, usually output of an RTE component
 */
export function parseMarkdown(markdown: string | Uint8Array<ArrayBufferLike>): string {
  return micromark(markdown);
}
