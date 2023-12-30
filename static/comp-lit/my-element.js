import { LitElement, html } from 'lit';
import './x-swapped.js'

export class PageIndex extends LitElement {
    constructor() {
        super();
    }
    render() {
        return html`
            <button>trigger swap</button>
            <x-swapped hx-post="/main/toswap" hx-trigger="click from:button" hx-swap="innerHTML"></x-swapped>
    `;
    }
}
customElements.define('my-element', PageIndex);