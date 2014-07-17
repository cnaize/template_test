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
		return
	}
	template_path := root_path + "/templates/" + "test"
	text_path := template_path + ".txt"

	fmt.Printf("Is file exists: %v\n", exist(text_path))

	textTmpl, err := template.ParseFiles(text_path)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	var text bytes.Buffer
	args := make(map[string]interface{})
	args["message"] = "Hello there!"

	if err := textTmpl.ExecuteTemplate(&text, "test"+".txt", args); err != nil {
		fmt.Printf("Error: %v", err)
		return
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
