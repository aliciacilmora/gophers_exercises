package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	cyoa "github.com/aliciacilmora/choose_your_own_adventure"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gophers.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)

	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(storyTmpl)) // template

	h := cyoa.NewHandler(story, cyoa.WithTempelate(tpl), cyoa.WithPathFunc(pathFn)) // Handler setup
	mux := http.NewServeMux()
	mux.Handle("/story/", h) // Register the handler with mux
	fmt.Printf("Starting the server at port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)) // Pass mux instead of h
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro" // Default chapter is "intro"
	}
	return path[len("/story/"):] // Remove the leading "/"
}

var storyTmpl = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Choose Your Own Adventure</title>
	</head>
	<body>
		<section class="page">
			<h1>{{.Title}}</h1>
			{{range .Paragraphs}}
			<p>{{.}}</p>
			{{end}}
			{{if .Options}}
				<ul>
					{{range .Options}}
					<li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
					{{end}}
				</ul>
			{{else}}
				<h3>...</h3>
			{{end}}
		</section>
		<style>
			body {
				font-family: Helvetica, Arial, sans-serif;
			}
			h1 {
				text-align: center;
				position: relative;
			}
			.page {
				width: 80%;
				max-width: 500px;
				margin: auto;
				margin-top: 40px;
				margin-bottom: 40px;
				padding: 80px;
				background: #FFFCF6;
				border: 1px solid #eee;
				box-shadow: 0 10px 6px -6px #777;
			}
			ul {
				border-top: 1px dotted #ccc;
				padding: 10px 0 0 0;
				-webkit-padding-start: 0;
			}
			li {
				padding-top: 10px;
			}
			a,
			a:visited {
				text-decoration: none;
				color: #6295b5;
			}
			a:active,
			a:hover {
				color: #7792a2;
			}
			p {
				text-indent: 1em;
			}
		</style>
	</body>
	</html>`
