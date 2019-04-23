package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type LandingPage struct {
	vecty.Core
}

func (this *LandingPage) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			elem.Heading5(
				vecty.Text("Hold on!"),
				elem.Break(),
				vecty.Text("Coming soon..."),
			),
		),
	)
}
