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
	// Обработка запросов для файлов из встроенной HTML-директории
	http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.FS(fileSystem))))

	// Обработка запросов для файлов из встроенной Static-директории
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(Files))))

	// Ваша дополнительная логика здесь

	// Запуск сервера
	fmt.Println("Server is running on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
