package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"net/http"
	)
	
type Book struct {
	Title 		string
	Author		string
	Description string
	}

func main() {
	m := martini.Classic()
	
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	
	m.Get("/", func(ren render.Render, r *http.Request) {
	
	ren.HTML(200, "foob", nil)
})

	m.Run()
	
}
	
