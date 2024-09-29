package main

import (
	"fmt"
	"net/http"

	urlshort "github.com/aliciacilmora/url_shortener"
)

func main() {
	mux := defaultMux()

	// MapHandler using mux as the fallback
	pathsToUrls := map[string]string{
		"/portfolio":  "https://ashutoshanand.work",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /portfolio
  url: https://ashutoshanand.work
- path: /linkedin
  url: https://linkedin.com/in/ashutoshanand-work/
- path: /github
  url: https://github.com/aliciacilmora/
`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
