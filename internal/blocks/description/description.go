package description

import (
	"bytes"
	"text/template"
)

type DescriptionBlock struct {
	template *template.Template
}

type DescriptionBlockInput struct {
	ProjectDescription string
}

func NewDescriptionBlock(templatePath string) (*DescriptionBlock, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}
	newBlock := DescriptionBlock{template: tmpl}
	return &newBlock, nil
}

func (block *DescriptionBlock) Render(data interface{}) (string, error) {
	var buffer bytes.Buffer
	err := block.template.Execute(&buffer, data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
