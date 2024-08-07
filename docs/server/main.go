package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	dirPath := "./docs"

	// Verificar la existencia del directorio y algunos archivos clave
	checkDirectoryAndFiles(dirPath, []string{"index.html", "play.wasm", "404.html"})

	fs := http.FileServer(http.Dir(dirPath))
	http.Handle("/", logRequests(fs, dirPath))

	log.Println("Listening on :8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func checkDirectoryAndFiles(dirPath string, files []string) {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		log.Fatalf("Directory does not exist: %s", dirPath)
	} else if err != nil {
		log.Fatalf("Error checking directory: %v", err)
	} else {
		log.Println("Directory exists:", dirPath)
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Printf("File does not exist: %s\n", filePath)
		} else if err != nil {
			log.Printf("Error checking file: %v\n", err)
		} else {
			log.Println("File exists:", filePath)
		}
	}
}

func logRequests(h http.Handler, dirPath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join(dirPath, r.URL.Path)
		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(dirPath, "index.html"))
			return
		}
		h.ServeHTTP(w, r)
	})
}
