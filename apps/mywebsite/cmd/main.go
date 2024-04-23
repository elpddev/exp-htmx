package main

import (
	"html/template"
	"io"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
  templates  *template.Template
}

func (t * Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
  return &Templates{
    templates: template.Must(template.ParseGlob("views/*.html")),
  }
}

type Contact struct {
  Name string
  Email string
}

func newContact(name string, email string) Contact {
  return Contact{
    Name: name,
    Email: email,
  }
}

type Contacts = []Contact

type Data struct {
  Contacts Contacts
}

func newData() Data {
  return Data{
    Contacts: Contacts{
      newContact("John Doe", "jon@dow") ,
      newContact("Jane Doe", "jane@dow"),
    },
  }
}

func main() {
  e := echo.New()
  e.Use(middleware.Logger())

  e.Renderer = newTemplate()

  data := newData()

  e.GET("/",  func(c echo.Context) error {
    return c.Render(200, "index", data)
  })

  e.POST("/contacts",  func(c echo.Context) error {
    name := c.FormValue("name")
    email := c.FormValue("email")
    contact := newContact(name, email)
    data.Contacts = append(data.Contacts, contact)
    return c.Render(200, "display", data)
  })

  e.Logger.Fatal(e.Start(":42069"))
}
