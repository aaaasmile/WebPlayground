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
      `;
    constructor(caption) {
        super();
        this.caption = caption;
    }
    render() {
        return html`
        <div>${this.title}</div>
        <div>${this.abstract}</div>
      `
    }
}
customElements.define('x-post', XPost);