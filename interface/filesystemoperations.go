package ui

import (
	"embed"
)

//go:embed "html" "templates"
var Files embed.FS
func main(){
	fileSystem, err := fs.Sub(Files, "html")
	
}
