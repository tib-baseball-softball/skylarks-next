import {Carta} from "carta-md";
import DOMPurify from "dompurify";

/**
 * Converts Markdown to HTML using a library.
 *
 * Single utility function to be able to easily switch out the library later.
 *
 * @param markdown the input value, usually output of an RTE component
 */
export function markdownToHTML(markdown: string): string {
  const carta = new Carta({
    sanitizer: DOMPurify.sanitize,
  });
  return carta.renderSSR(markdown);
}

export function stripTags(html: string) {
  let doc = new DOMParser().parseFromString(html, 'text/html');
  return doc.body.textContent || "";
}
