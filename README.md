![logo](https://github.com/elpddev/exp-htmx/assets/17731302/32d82564-8bfd-4026-83f5-209abfa10844)

# HTMX - Exploration

Follow along of Frontend master & Primagen course - [FULL Introduction To HTMX Using Golang](https://youtu.be/x7v6SNIgJpE?si=_utehPM_u_phix4M).

* [The Primeagen](https://www.youtube.com/channel/UC8ENHE5xdFSwx71u3fDH5Xw)
* [FrontEnd Masters](https://frontendmasters.com/)

## Follow Along

### Form Replacing Inner - Creating Double 'Display'

Instead of replacing the inner html, the target could be specified to swap the outer html.

[hx-swap](ttps://htmx.org/attributes/hx-swap/)

```html
<form hx-swap="outerHTML">
```

### Adding only one item to the dom - Not replacing all list

[hx-swap-oob](https://htmx.org/attributes/hx-swap-oob/)

The hx-swap-oob attribute allows you to specify that some content in a response should be swapped into the DOM **somewhere other** than the target, that is “Out of Band”. 

This allows you to piggy back updates to other element updates on a response.

For oob swap, we also need to specify what the original target should be replaced with. Otherwise it will be replaced with null.

### Debugging

https://htmx.org/docs/#debugging

On the console:

```
htmx.logAll()
```

### Delete

[hx-delete](https://htmx.org/attributes/hx-delete/)

The hx-delete attribute will cause an element to issue a DELETE to the specified URL and swap the HTML into the DOM using a swap strategy:

```html
  <button hx-delete="/contacts/{{ .Id }}">Delete</button>
```

## Tools & Technologies

### HTMX

htmx gives you access to AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML, using attributes, so you can build modern user interfaces with the simplicity and power of hypertext

https://htmx.org/

### Go

An open-source programming language

https://go.dev/

#### Air

Live reload for Go apps.

https://github.com/cosmtrek/air

#### Echo

High performance, extensible, minimalist Go web framework

https://echo.labstack.com/

#### html/template

Package template (html/template) implements data-driven templates for generating HTML output safe against code injection. It provides the same interface as text/template and should be used instead of text/template whenever the output is HTML.

https://pkg.go.dev/html/template

### Mise

mise is a polyglot tool version manager

https://mise.jdx.dev/
