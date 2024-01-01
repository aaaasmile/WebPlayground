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
        p {
          text-transform: uppercase;
          font-size: 3rem;
          font-weight: 500;
          line-height: 1;
          border: 5px solid;
          padding: 2rem;
        }
      `;
    constructor() {
        super();
    }
    render() {
        return html`
        <p> I am a paragraph of text that has a few words in it.</p>
        <div>hello from outside, inner will be swapped</div>
        <slot></slot>
      `
    }
}
customElements.define('x-swapped', XSwapped);