// Command app serves a page with the student's album number read from NUMER_ALBUMU
// and the base system of the container image read from /etc/os-release.
//
// Usage: app [serve | info | healthcheck]
//
//	serve        start the web server on PORT (default 8080), the default without arguments
//	info         print the album number, base system and port, then exit
//	healthcheck  exit 0 when the server on PORT answers /healthz, 1 otherwise
package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//go:embed templates static
var content embed.FS

var page = template.Must(template.ParseFS(content, "templates/index.html"))

type pageData struct {
	Album    string
	Hostname string
	System   string
	Port     string
}

func main() {
	command := "serve"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}
	switch command {
	case "serve":
		serve()
	case "info":
		fmt.Printf("Numer albumu: %s\nSystem bazowy: %s\nPort: %s\n", album(), osRelease(), port())
	case "healthcheck":
		os.Exit(healthcheck())
	default:
		fmt.Fprintf(os.Stderr, "nieznane polecenie %q, dostępne: serve, info, healthcheck\n", command)
		os.Exit(2)
	}
}

func port() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
}

func album() string {
	return strings.TrimSpace(os.Getenv("NUMER_ALBUMU"))
}

func serve() {
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

	log.Printf("listening on :%s", port())
	log.Fatal(http.ListenAndServe(":"+port(), logRequests(mux)))
}

// healthcheck asks the running server for /healthz, so HEALTHCHECK works on any base image without curl.
func healthcheck() int {
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("http://localhost:" + port() + "/healthz")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, resp.Status)
		return 1
	}
	fmt.Println("ok")
	return 0
}

func index(w http.ResponseWriter, _ *http.Request) {
	hostname, _ := os.Hostname()
	data := pageData{
		Album:    album(),
		Hostname: hostname,
		System:   osRelease(),
		Port:     port(),
	}
	if err := page.Execute(w, data); err != nil {
		log.Printf("render: %v", err)
	}
}

// osRelease returns PRETTY_NAME from /etc/os-release, the file every current distribution ships.
func osRelease() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "nieznany"
	}
	for _, line := range strings.Split(string(data), "\n") {
		if value, ok := strings.CutPrefix(line, "PRETTY_NAME="); ok {
			return strings.Trim(value, `"`)
		}
	}
	return "nieznany"
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
