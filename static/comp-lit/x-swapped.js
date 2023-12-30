import { LitElement, css, html } from 'lit';

export class XSwapped extends LitElement {
    static styles = css`
        :host {
          display: block;
          border: 3px solid gray;
          border-radius: 1rem;
          padding: 1rem;
          margin: 2rem;
          color: black;
        }
      `;
    constructor() {
        super();
    }
    render() {
        return html`
        <div>hello from outside, inner will be swapped</div>
        <slot></slot>
      `
    }
}
customElements.define('x-swapped', XSwapped);