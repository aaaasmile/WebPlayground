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
dei componenti base che abbiano una semantica comprensibile. Quello che non voglio vedere è

    <button class="button button-primary">
ma bensì:

    <Button>
Mentre Button è un componente statico col suo stile, la parte in html che lo conterrà è generata dinamicamente
sul server usando go html/template.

## Stile
Prova a vedere il posts.htm. Non ha nessuno stile, ma bensì contiene i pezzi in html,
che andranno generati sul server, dove questo contenuto html verrà mostrato dal webcomponent
x-post. x-post ha la sezione css che definisce lo stile dei vari elementi. Il contenuto html
di Abstract, va a finire in slot. Questo perchè non è possibile creare html da stringhe.
Lo stile del contenuto, che sarà un hypertesto in paragrafi, avrà uno stile definito
anch'esso nel componente x-post. Il collegamento è ::slotted(p).  

## Preact, Lit, htmx
Ho provato ad usare le coppie Lit/htmx e Preact/htmx senza essere soddisfatto.
Mi va bene il rendering (css e render) di Preact e il web component di Lit (extends LitElement). 
Quindi, per ora, uso tutti e tre.


## htmx
È il cuore che consente un update parziale del dom usando il codice html generato sul server.
Lo stato dell'applicazione è memorizzato sul server, mentre il client, attraverso il browser,
può andare avanti e indietro nella storia. L'unico svantaggio è che, al momento, non si può usare
un shadow DOM, che è quello che usa React per creare la sua responsive app. 

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
- htmx e webcomponent: https://binaryigor.com/htmx-and-web-components-a-perfect-match.html

## Md editor
Per l'editor di un post in formato md, uso l'editor yace (https://petersolopov.github.io/yace/examples/)
Nell'esempio di Yace ci sono delle dipendenze che ho caricato usando un altro header: header_editor.htm.
Esse sono:
 - highlight su https://highlightjs.org/
 - mdhl che dovrebbe essere su: https://github.com/petersolopov/mdhl
Nell'header ho caricato solo i css, mentre i moduli sono caricati nel file index.js.
Ho cambiato codeInFences per non avere il warning che l'API è deprecated.

## Icone
Icone gratuite in formato svg si possono trovare su https://icons.getbootstrap.com/
Basta metterle nel file icons.js e usarke nel componente.