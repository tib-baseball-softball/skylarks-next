// convert-tailwind-to-css.js
// https://stackoverflow.com/questions/74979020/convert-tailwindcss-to-native-css/79414154#79414154

import fs from "fs"
import postcss from "postcss"
import tailwindcssPlugin from "@tailwindcss/postcss"

// Minify?
const args = process.argv.slice(2)
let minify = false
if (args.includes("--minify")) {
  minify = true
}
// Generated CSS indent spaces count
const indentSpaces = 2
// Generated CSS output file
const outputCSS = "./output.css"

// CSS Prettier (TailwindCSS with LightningCSS returns incorrectly formatted results with the settings optimize: true and minify: false)
const prettier = (css, indent = 2) => {
  const lines = css.split("\n")
  let indentLevel = 0

  return lines
    .map((line) => {
      const trimmed = line.trim()

      if (trimmed.endsWith("}")) {
        indentLevel = Math.max(indentLevel - 1, 0)
      }

      const formattedLine = " ".repeat(indentLevel * indent) + trimmed

      if (trimmed.endsWith("{")) {
        indentLevel++
      }

      return formattedLine
    })
    .join("\n")
}

// Convert Tailwind CSS to native CSS
postcss([
  tailwindcssPlugin({
    optimize: {
      minify, // minify or not?
    },
  }),
])
  //   .process(`
  //   @layer theme, base, components, utilities;
  //   @import "tailwindcss/theme.css" layer(theme);
  //
  //   /* preflight: Not required, it only creates rules for CSS reset. */
  //   /* @import "tailwindcss/preflight.css" layer(base); */
  //
  //   @import "tailwindcss/utilities.css" layer(utilities);
  // `, {from: './src'})

  .process(
    `
    @import 'tailwindcss';
@import "tw-animate-css";
@plugin '@tailwindcss/forms';
@plugin "@tailwindcss/typography";
  `,
    { from: "./src" }
  )

  .then((result) => {
    let formattedCSS
    if (!minify) {
      // Format CSS (The optimized result returns the output with incorrect indentation.)
      formattedCSS = prettier(result.css, indentSpaces)
    }

    // Write the generated CSS to a file
    fs.writeFileSync(outputCSS, formattedCSS || result.css, "utf8")
    console.log(`Native CSS generated: ${outputCSS}`)
  })
  .catch((err) => console.error("An error occurred:", err))
