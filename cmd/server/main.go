package main

import (
	"fmt"
	"net/http"

	"github.com/ignacio-araujo/go-rest-api-course/internal/comment"
	"github.com/ignacio-araujo/go-rest-api-course/internal/database"
	transportHTTP "github.com/ignacio-araujo/go-rest-api-course/internal/transport/http"
)

// App struct which contains pointers to database connections
type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up our APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
