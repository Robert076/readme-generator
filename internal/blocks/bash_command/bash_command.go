package bash_command

import (
	"bytes"
	"text/template"
)

type BashCommandBlock struct {
	template *template.Template
}

type BashCommandBlockInput struct {
	CommandTitle  string
	CommandText   string
	CommandNumber string
}

func NewBashCommandBlock(templatePath string) (*BashCommandBlock, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}
	newBlock := BashCommandBlock{template: tmpl}
	return &newBlock, nil
}

func (block *BashCommandBlock) Render(data BashCommandBlockInput) (string, error) {
	var buffer bytes.Buffer
	err := block.template.Execute(&buffer, data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
