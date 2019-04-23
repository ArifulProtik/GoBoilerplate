package handlers

import (
	"github.com/GoBoilerplate/app/components"
	"github.com/go-humble/router"
	"github.com/gopherjs/vecty"
)

func landingPage(c *router.Context) {
	vecty.SetTitle("Coming soon...")
	vecty.RenderBody(&components.LandingPage{})
}
