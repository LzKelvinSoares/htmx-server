package books

import "embed"

//go:embed edit.html item.html list.html
var FS embed.FS
