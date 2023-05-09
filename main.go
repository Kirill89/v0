package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Query().Get("filename")

		content, err := os.ReadFile("public/" + filename)

		if err != nil {
			http.Error(w, "File not found", 404)
			return
		}

		fmt.Fprintf(w, "%s", content)
	})

	http.ListenAndServe(":3000", nil)
}
