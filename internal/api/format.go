package api

type ReadmeInput struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Commands    []BashCommand `json:"commands"`
}

type BashCommand struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

var FormatPOST = ReadmeInput{
	Title:       "My Cool Project",
	Description: "A short summary of what this project does.",
	Commands: []BashCommand{
		{
			Title: "As many commands",
			Text:  "go mod tidy",
		},
		{
			Title: "as you want",
			Text:  "go run main.go",
		},
	},
}
