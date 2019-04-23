package main

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/GoBoilerplate/app/handlers"
)

func main() {
	vecty.AddStylesheet("/assets/app.css")
	fmt.Println("Coming soon!")
	handlers.Start()
}
