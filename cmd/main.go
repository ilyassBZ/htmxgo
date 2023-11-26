package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
    templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer,name string, data interface{},c echo.Context)error  {
  return t.templates.ExecuteTemplate(w,name,data)
  
}

func main()  {
  tmpls,err:=template.ParseGlob("public/views/*.html")
  if err!=nil{
    log.Fatal(http.StatusBadGateway,err)
  }
  e:=echo.New()
  e.Use(middleware.Logger())
  e.Renderer = &TemplateRenderer{
    templates: tmpls,

  }
  e.Static("/dist","./dist")
  e.GET("/",func(c echo.Context) error {
    return c.Render(200,"index.html",nil)
  })
  
  e.Logger.Fatal(e.Start(":42069"))
}
