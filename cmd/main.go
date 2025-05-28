package main

import (
	"fmt"
	"log"

	"github.com/Robert076/readme-generator/internal/blocks/title"
)

func main() {
	titleBlock, err := title.NewTitleBlock("../internal/blocks/title/title1.tmpl")

	if err != nil {
		log.Fatal("Failed loading template:", err)
	}

	titleString, err := titleBlock.Render(struct{ ProjectTitle string }{ProjectTitle: "Test"})

	if err != nil {
		log.Fatal("Failed loading template:", err)
	}

	fmt.Print(titleString)
}
