import Yace from "https://unpkg.com/yace?module";

import { highlight, defaultRenderers } from "https://unpkg.com/mdhl@0.0.6/dist/mdhl.esm.js?module";
import hljs from 'https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/es/highlight.min.js';
import go from "https://unpkg.com/@highlightjs/cdn-assets@11.9.0/es/languages/go.min.js"
hljs.registerLanguage('go', go);

const renderer = {
    ...defaultRenderers,
    codeInFences: (code, language) => {
        const validLanguage = hljs.getLanguage(language) ? language : "plaintext";
        return hljs.highlight(validLanguage, code).value;
    },
};

function highlighter(value) {
    return highlight(value, renderer);
}

const editor = new Yace("#editor", {
    value:
        "# Markdown Editor\n\nHighlighting code in code blocks\n\n```go\nfunc sum(a, b int) {\n  return a + b\n}\n```\n",
    styles: {
        fontSize: "18px"
    },
    highlighter,
    lineNumbers: true
});

editor.textarea.spellcheck = true;
