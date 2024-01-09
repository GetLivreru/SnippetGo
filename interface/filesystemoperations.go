package ui

import (
	"embed"
)

//go:embed "html" "templates"
var Files embed.FS
func main(){
	fileSystem, err := fs.Sub(Files, "html")
	if err != nil {
		fmt.Println("Failed to create file system:", err)
	return
}
