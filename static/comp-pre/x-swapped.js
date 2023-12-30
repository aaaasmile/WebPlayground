// this is not working
export class XSwapped extends HTMLElement {
    constructor() {
        super();
    }

    render() {
        return html`
        <style>
        :host {
            display: block;
            border: 3px solid gray;
            border-radius: 1rem;
            padding: 1rem;
            margin: 2rem;
            color: black;
        }
        </style>
        <div>hello from outside, inner will be swapped</div>
        <slot></slot>
      `
    }
}
customElements.define('x-swapped', XSwapped);