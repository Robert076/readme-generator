package blocks

import (
	"fmt"
	"math/rand"

	"github.com/Robert076/readme-generator/internal/blocks/bash_command"
	"github.com/Robert076/readme-generator/internal/blocks/description"
	"github.com/Robert076/readme-generator/internal/blocks/title"
)

type MarkdownBlock interface {
	Render(data interface{}) (string, error)
}

func GetRandomTitleBlock() (*title.TitleBlock, error) {
	n := rand.Intn(3) + 1

	path := fmt.Sprintf("internal/blocks/title/templates/title%d.tmpl", n)

	return title.NewTitleBlock(path)
}

func GetDescriptionBlock() (*description.DescriptionBlock, error) {
	return description.NewDescriptionBlock("internal/blocks/description/templates/description.tmpl")
}

func GetBashCommandBlock() (*bash_command.BashCommandBlock, error) {
	return bash_command.NewBashCommandBlock("internal/blocks/bash_command/templates/bash_command.tmpl")
}
