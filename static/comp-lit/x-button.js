import { LitElement, css, html } from 'lit';

export class XButton extends LitElement {
    static properties = {
        caption: {},
        icon: {},
    };
    static styles = css`
        :host {
          padding: 1rem;
        }
      `;
    constructor(caption) {
        super();
        this.caption = caption;
    }
    render() {
        return html`
        <button>${this.caption}
        <span>${this.icon}</span></button>
      `
    }
}
customElements.define('x-button', XButton);