import { LitElement, css, html } from 'lit';

export class XPost extends LitElement {
    static properties = {
        title: {},
        abstract: {},
    };
    static styles = css`
        :host {
          padding: 1rem;
        }
        p {
            font-size: 1.3rem;
            padding: 2rem;
          }
      `;
    constructor() {
        super();
    }
    render() {
        return html`
        <h2>${this.title}</h2>
        <p>What about this paragraph. Is it like the abstract?</p>
        <slot></slot>
      `
    }
}
customElements.define('x-post', XPost);