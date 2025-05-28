package generator

import (
	"fmt"
	"strings"

	"github.com/Robert076/readme-generator/internal/blocks"
	"github.com/Robert076/readme-generator/internal/blocks/bash_command"
	"github.com/Robert076/readme-generator/internal/blocks/description"
	"github.com/Robert076/readme-generator/internal/blocks/title"
)

var (
	titleBlock       *title.TitleBlock
	descriptionBlock *description.DescriptionBlock
	bashCmdBlock     *bash_command.BashCommandBlock
)

func init() {
	var err error
	titleBlock, err = blocks.GetRandomTitleBlock()
	if err != nil {
		panic(err)
	}
	descriptionBlock, err = blocks.GetDescriptionBlock()
	if err != nil {
		panic(err)
	}
	bashCmdBlock, err = blocks.GetBashCommandBlock()
	if err != nil {
		panic(err)
	}
}

func GenerateReadme(data map[string]interface{}) (string, error) {
	var sb strings.Builder

	// Render Title
	if titleVal, ok := data["title"].(string); ok {
		titleRendered, err := titleBlock.Render(title.TitleBlockInput{
			ProjectTitle: titleVal,
		})
		if err != nil {
			return "", fmt.Errorf("render title: %w", err)
		}
		sb.WriteString(titleRendered)
		sb.WriteString("\n\n")
	}

	// Render Description
	if descVal, ok := data["description"].(string); ok {
		descRendered, err := descriptionBlock.Render(description.DescriptionBlockInput{
			ProjectDescription: descVal,
		})
		if err != nil {
			return "", fmt.Errorf("render description: %w", err)
		}
		sb.WriteString(descRendered)
		sb.WriteString("\n\n")
	}

	// Render Commands
	if cmds, ok := data["commands"].([]interface{}); ok {
		for i, c := range cmds {
			if cmdMap, ok := c.(map[string]interface{}); ok {
				titleText, _ := cmdMap["title"].(string)
				cmdText, _ := cmdMap["text"].(string)
				bashRendered, err := bashCmdBlock.Render(bash_command.BashCommandBlockInput{
					CommandNumber: fmt.Sprintf("%d", i+1),
					CommandTitle:  titleText,
					CommandText:   cmdText,
				})
				if err != nil {
					return "", fmt.Errorf("render bash command #%d: %w", i+1, err)
				}
				sb.WriteString(bashRendered)
				sb.WriteString("\n\n")
			}
		}
	}

	return sb.String(), nil
}
