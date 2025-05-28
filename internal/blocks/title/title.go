package title

import (
	"bytes"
	"text/template"
)

type TitleBlock struct {
	template *template.Template
}

type TitleBlockInput struct {
	ProjectTitle string
}

func NewTitleBlock(templatePath string) (*TitleBlock, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}
	newBlock := TitleBlock{template: tmpl}
	return &newBlock, nil
}

func (block *TitleBlock) Render(data TitleBlockInput) (string, error) {
	var buffer bytes.Buffer
	err := block.template.Execute(&buffer, data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
