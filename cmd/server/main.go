package main

import "fmt"

//App struct which contains pointers to database connections
type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up Our APP")
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
