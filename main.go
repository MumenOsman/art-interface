package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	h "interface/helpers" // h is the alias for the package helpers
)

type Template struct {
	Name  string
	Value string
}

type PageData struct {
	Input       string
	Result      string
	Error       string
	Mode        string
	StatusCode  int
	StatusClass string
	Templates   []Template
}

var templates = []Template{
	{
		Name: "House",
		Value: `
   _
  | |
 _| |_
|     |
|_____|
`,
	},
	{
		Name: "Hello",
		Value: `
 _   _      _ _       
| | | |    | | |      
| |_| | ___| | | ___  
|  _  |/ _ \ | |/ _ \ 
| | | |  __/ | | (_) |
\_| |_/\___|_|_|\___/ 
`,
	},
	{
		Name: "Box",
		Value: `
+-------+
|       |
|       |
+-------+
`,
	},
	{
		Name: "Cat",
		Value: `
      |\      _,,,---,,_
ZZZzz /,` + "`" + `.-'` + "`" + `'    -.  ;-;;,_
     |,4-  ) )-,_. ,\ (  ` + "`" + `'-'
    '---''(_/--'  ` + "`" + `-'\_)
`,
	},
	{
		Name: "Dog",
		Value: `
  __      _
o'')}____//
 ` + "`" + `_/      )
 (_(_/-(_/
`,
	},
	{
		Name: "Owl",
		Value: `
  ,_,
 (O,O)
 (   )
 -"-"-
`,
	},
}

func main() {
	web := flag.Bool("web", false, "Start web server")
	encode := flag.Bool("Encode", false, "Encode mode")
	multi := flag.Bool("Multi", false, "Multi line mode")
	flag.Parse()

	if *web {
		startServer()
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error: No input provided")
		return
	}

	uInput := strings.Join(args, " ")
	isEncode := *encode
	isMulti := *multi

	if uInput == "" {
		fmt.Println("Error")
		return
	}

	var output string
	var err error

	if isEncode {
		if isMulti {
			output, err = h.MultiLineEncode(uInput)
		} else {
			output, err = h.SingleLineEncode(uInput)
		}
	} else {
		if isMulti {
			output, err = h.MultiDecode(uInput)
		} else {
			output, err = h.SingleDecode(uInput)
		}
	}

	if err != nil {
		fmt.Println("Error message")
	} else {
		fmt.Println(output)
	}
}

func startServer() {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/decoder", handleDecoder)

	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	renderTemplate(w, PageData{
		Templates: templates,
	})
}

func handleDecoder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	mode := r.FormValue("mode")

	if text == "" {
		renderTemplate(w, PageData{
			Error:       "Please enter some text.",
			StatusCode:  http.StatusBadRequest,
			StatusClass: "error",
		})
		return
	}

	var output string
	var err error
	isMulti := strings.Contains(text, "\n") || strings.Contains(text, "\\n")

	if mode == "encode" {
		if isMulti {
			output, err = h.MultiLineEncode(text)
		} else {
			output, err = h.SingleLineEncode(text)
		}
	} else if mode == "art" {
		output, err = h.GenerateASCII(text)
	} else {
		// Decode
		if isMulti {
			output, err = h.MultiDecode(text)
		} else {
			output, err = h.SingleDecode(text)
		}
	}

	data := PageData{
		Input:     text,
		Mode:      mode,
		Templates: templates,
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data.Error = "Error processing request" // Generic error as per requirements? Or specific?
		// Requirement: "The server must return the HTTP response for Bad Request in the event of a malformed query string."
		// If encoding/decoding fails, it might be bad input.
		data.StatusCode = http.StatusBadRequest
		data.StatusClass = "error"
	} else {
		w.WriteHeader(http.StatusAccepted)
		data.Result = output
		data.StatusCode = http.StatusAccepted
		data.StatusClass = "success"
	}

	renderTemplate(w, data)
}

func renderTemplate(w http.ResponseWriter, data PageData) {
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Template error:", err)
		return
	}
	tmpl.Execute(w, data)
}
