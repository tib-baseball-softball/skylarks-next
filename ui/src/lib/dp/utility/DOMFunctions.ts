import { parse_markdown, init } from "maerkchen";

/**
 * Converts Markdown to HTML using a library.
 *
 * Single utility function to be able to easily switch out the library later.
 *
 * @param markdown the input value, usually output of an RTE component
 */
export async function markdownToHTML(markdown: string): Promise<string> {
  await init();
  return parse_markdown(markdown);
}

export function stripTags(html: string) {
  let doc = new DOMParser().parseFromString(html, "text/html");
  return doc.body.textContent || "";
}
