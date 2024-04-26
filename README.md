![logo](https://github.com/elpddev/exp-htmx/assets/17731302/32d82564-8bfd-4026-83f5-209abfa10844)

# HTMX - Exploration

Follow along of Frontend master & Primagen course - [FULL Introduction To HTMX Using Golang](https://youtu.be/x7v6SNIgJpE?si=_utehPM_u_phix4M).

* [The Primeagen](https://www.youtube.com/channel/UC8ENHE5xdFSwx71u3fDH5Xw)
* [FrontEnd Masters](https://frontendmasters.com/)

## Follow Along

### Go Basic Server

#### main

[Declaring a package](https://go.dev/ref/spec#Package_clause
) [`main`](https://go.dev/ref/spec#Program_initialization)

Program execution begins by initializing the program and then invoking the function main in package main. When that function invocation returns, the program exits. It does not wait for other (non-main) goroutines to complete. 

```
package main
```

#### import

Using [import statement](https://go.dev/ref/spec#Import_declarations) to bring other packages code to use.

An import declaration states that the source file containing the declaration depends on functionality of the imported package (§Program initialization and execution) and enables access to exported identifiers of that package. The import names an identifier (PackageName) to be used for access and an ImportPath that specifies the package to be imported.

```go
import (
	"html/template"
	"io"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
```

#### type / struct

A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique.

https://go.dev/ref/spec#Struct_types

```go
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
```

##### Array

An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the array and is never negative.

https://go.dev/ref/spec#ArrayType

```go
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

```go
type Contacts = []Contact
```

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

```go
e.DELETE("/contacts/:id",  func(c echo.Context) error {
    id := c.Param("id")
    idNum, err := strconv.Atoi(id)

    if (err != nil) { 
      return c.String(400, "invalid id")
    }

    index := page.Data.indexOf(idNum)
    if (index == -1) {
      return c.String(404, "contact not found")
    }

    page.Data.Contacts = append(
      page.Data.Contacts[:index], 
      page.Data.Contacts[index+1:]...
    )

    return c.NoContent(200)
})
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
