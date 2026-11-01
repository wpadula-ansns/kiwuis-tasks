// Command lab01 serves a page with the student's album number read from NUMER_ALBUMU.
package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

//go:embed templates static
var content embed.FS

var page = template.Must(template.ParseFS(content, "templates/index.html"))

type pageData struct {
	Album    string
	Hostname string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	static, err := fs.Sub(content, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServerFS(static)))
	mux.HandleFunc("GET /{$}", index)
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok\n"))
	})

	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, logRequests(mux)))
}

func index(w http.ResponseWriter, _ *http.Request) {
	hostname, _ := os.Hostname()
	data := pageData{
		Album:    strings.TrimSpace(os.Getenv("NUMER_ALBUMU")),
		Hostname: hostname,
	}
	if err := page.Execute(w, data); err != nil {
		log.Printf("render: %v", err)
	}
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
