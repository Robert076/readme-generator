package blocks

type MarkdownBlock interface {
	Render(data interface{}) (string, error)
}
