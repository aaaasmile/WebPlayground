import { LitElement, css, html } from 'lit';

export class XPost extends LitElement {
    static properties = {
        title: {},
        date:{},
        abstract: {},
    };
    static styles = css`
        :host {
          padding: 1rem;
        }
        p {
            font-size: 1.0rem;
            padding: 0.5rem;
        }
        .date {
            color: #b0b0b0;
            font-size: 0.6em;
        }
        ::slotted(p) {
            font-size: 1.2rem;
            padding: 2rem;
        }
      `;
    constructor() {
        super();
    }
    render() {
        return html`
        <article>
        <h2>Title: ${this.title}</h2>
        <p class="date">${this.date}</p>
        <p>What about this paragraph. Is it like the abstract?</p>
        <slot></slot>
        </article>
      `
    }
}
customElements.define('x-post', XPost);