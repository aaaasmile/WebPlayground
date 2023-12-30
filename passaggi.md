# WebPlayground
Un progetto per provare diverse tecniche per creare Web-App.
Principalmente lo scopo è quelo di spostare più parti dell'applicazione sul server, per evitare
di dover creare un'App che usi vue.js per mostrare il contenuto di un sito tipo invido.it.

## Templ
Templ è un template che compila in go. Serve per avere un sistema robusto del template
con la ricognizione di errori di sintassi nella fase di compilazione (2 pass: creare il file go e poi compilarlo).
Al momento edito costantemente il template hmtl in visual code, che poi vedo immediatamente nel browser.
Se usassi templ, ogni minimo cambiamento (per esempio i tag class) richiederebbe una compilazione.
Per questo motivo sono al momento scettico nel suo utilizzo.

## Web Components
Sono partito usando preact, che non mi sembra male. Però Lit è più snello e al momento ha la mia 
preferenza. Vorrei usare Web Components per isolare lo style. I Web Components sono statici e vengono
scaricati dal server all'inizio senza go template. Questo ha la praticità di editare il file senza
dovere compilare di nuovo. Però non voglio creare una SPA, ma solo isolare lo stile e creare
dei componenti base che abbiano una semantica coprensibile. Quello che non voglio vedere è

    <button class="button button-primary">
ma bensì:

    <Button>
Mentre Button è un componente statico col suo stile, la parte in html che lo conterrà è generata dinamicamente
sul server usando go html/template.

## htmx
È il cuore che consente un update parziale del dom usando il codice html generato sul server.
Lo stato dell'applicazione è memorizzato sul server, mentre il client, attraverso il browser,
può andare avanti e indietro nella storia. 

## Risorse

- Discussione molto interessante su htmx, templ e go: https://news.ycombinator.com/item?id=38597599
- templ riferimento: https://templ.guide/
- templ git: https://github.com/a-h/templ/tree/main
- Lit Web Component: https://lit.dev/docs/components/overview/
- Lit e htmx assieme: https://github.com/brtheo/ServerWebComponent-htmx-lit-bun-elysia/
- Htmx: https://htmx.org/
- Hypermedia Book: https://hypermedia.systems/book/contents/
- Specifiche di HTML: https://html.spec.whatwg.org/multipage/
- go htmx example: https://github.com/joerdav/go-htmx-examples/tree/main