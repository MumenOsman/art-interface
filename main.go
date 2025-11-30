package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	h "interface/helpers" // h is the alias for the package helpers
)

type PageData struct {
	Input       string
	Result      string
	Error       string
	Mode        string
	StatusCode  int
	StatusClass string
}

func main() {
	args := os.Args[1:]

	if len(args) > 0 && args[0] == "--web" {
		startServer()
		return
	}

	if len(args) == 0 {
		fmt.Println("Error")
		return
	}

	isMulti := false
	isEncode := false
	var uInput string

	for _, arg := range args { // the loop here goes through the range of args so with every round it checks the next argument, so in the case of multi encode it will still reads both
		if arg == "--help" || arg == "-h" {
			fmt.Println("Help message")
			return
		} else if arg == "--Encode" {
			isEncode = true
		} else if arg == "--Multi" {
			isMulti = true
		} else {
			uInput = arg
		}
	}

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
	renderTemplate(w, PageData{})
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

	// Basic handling of literal \n for single line inputs if they are typed that way
	// But usually textarea sends actual newlines.
	// The helpers seem to handle "Multi" vs "Single".
	// Let's assume if there's a newline char, it's multi.

	if mode == "encode" {
		if isMulti {
			output, err = h.MultiLineEncode(text)
		} else {
			output, err = h.SingleLineEncode(text)
		}
	} else {
		// Decode
		if isMulti {
			output, err = h.MultiDecode(text)
		} else {
			output, err = h.SingleDecode(text)
		}
	}

	data := PageData{
		Input: text,
		Mode:  mode,
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
