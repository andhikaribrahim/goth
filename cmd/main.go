package main

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"goth/cmd/web"
	"goth/internals/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
    tmpl *template.Template
}

func ParseTemplates() (*template.Template, error) {
    templateBuilder := template.New("")
    if t, _ := templateBuilder.ParseGlob("views/*/*/*/*/*.html"); t != nil {
        templateBuilder = t
    }
    if t, _ := templateBuilder.ParseGlob("views/*/*/*/*.html"); t != nil {
        templateBuilder = t
    }
    if t, _ := templateBuilder.ParseGlob("views/*/*/*.html"); t != nil {
        templateBuilder = t
    }
    if t, _ := templateBuilder.ParseGlob("views/*/*.html"); t != nil {
        templateBuilder = t
    }
    return templateBuilder.ParseGlob("views/*.html")
}

func newTemplate() *Template {
    tmpl, err := ParseTemplates()

    if err != nil {
        log.Fatalf("Error loading templates: %v", err)
    }

    log.Printf("Parsed templates: %+v", tmpl.DefinedTemplates())

    return &Template{
        tmpl: tmpl,
    }
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    log.Printf("Rendering template: %s", name)
    return t.tmpl.ExecuteTemplate(w, name, data)
}

func main() {
    // Check if views directory is accessible
    viewsDir, err := filepath.Abs("views")
    if err != nil {
        log.Fatalf("Error getting absolute path for views directory: %v", err)
    }
    if _, err := os.Stat(viewsDir); os.IsNotExist(err) {
        log.Fatalf("Views directory not found at: %s", viewsDir)
    } else {
        log.Printf("Views directory found at: %s", viewsDir)
    }

    e := echo.New()
    e.Static("/static", "static")
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Debug = true

    e.Renderer = newTemplate()

    // Initialize the controllers
    homeController := controllers.NewHomeController()
    countController := controllers.NewCountController()
    componentController := controllers.NewComponentController()
    todoController := controllers.NewTodoController()


    // Register routes
    web.RegisterRoutes(e, homeController, countController, componentController, todoController)

    log.Println("Starting server on :6969")
    e.Logger.Fatal(e.Start(":6969"))
}
