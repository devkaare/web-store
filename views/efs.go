package views

import (
	"embed"

	_ "github.com/a-h/templ"
)

//go:embed "assets"
var Files embed.FS
