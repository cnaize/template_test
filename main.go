package main

import (
	"bitbucket.org/kardianos/osext"
	"bytes"
	"fmt"
	"os"
	"text/template"
)

//==============================================================================
func main() {
	fmt.Println("Begin")

	root_path, err := osext.Executable()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	template_path := root_path + "/templates/" + "test"
	text_path := template_path + ".txt"

	textTmpl, err := template.ParseFiles(text_path)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	var text bytes.Buffer
	args := make(map[string]interface{})
	args["message"] = "Hello there!"

	if err := textTmpl.ExecuteTemplate(&text, "test"+".txt", args); err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Println("End")
	fmt.Println("Works!")
}

//==============================================================================
func exist(file_path string) bool {
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		return false
	}

	return true
}

//==============================================================================
