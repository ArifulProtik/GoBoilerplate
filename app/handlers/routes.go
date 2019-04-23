package handlers

import "github.com/go-humble/router"

func Start() {
	r := router.New()
	defer r.Start()
	r.HandleFunc("/", landingPage)
}
