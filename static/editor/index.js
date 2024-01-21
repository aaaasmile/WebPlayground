import Yace from "https://unpkg.com/yace?module";

// import { highlight, defaultRenderers } from "mdhl";
// import hljs from "highlight.js";


// import "mdhl/mdhl.css";
// import "highlight.js/styles/github-gist.css";

// const renderer = {
//   ...defaultRenderers,
//   codeInFences: (code, language) => {
//     try {
//       // highlight.js throw error when language is not found
//       return hljs.highlight(language, code).value;
//     } catch (error) {
//       return defaultRenderers.codeInFences(code, language);
//     }
//   }
// };

// function highlighter(value) {
//   return highlight(value, renderer);
// }

const editor = new Yace("#editor", {
  value:
    "# Markdown Editor\n\nHighlighting code in code blocks\n\n```js\nfunction sum(a, b) {\n  return a + b\n}\n```\n",
  styles: {
    fontSize: "18px"
  },
  //highlighter,
  lineNumbers: true
});

editor.textarea.spellcheck = true;
