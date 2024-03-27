package analysis

import (
	"strings"

	"github.com/musaubrian/env_lsp/lsp"
)

type State struct {
	// Map of file names to contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	// return getDiagnosticsForFile(text)
	return []lsp.Diagnostic{}
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	// return getDiagnosticsForFile(text)
	return []lsp.Diagnostic{}
}

func (s *State) TextDocumentCompletion(id int, params lsp.CompletionParams, completions []string) lsp.CompletionResponse {
	uri := params.TextDocument.URI
	line := params.Position.Line
	content := s.Documents[uri]

	lines := strings.Split(content, "\n")
	lineContent := lines[line]

	// TODO:: Python broke somehow, will rethink it
	prefix := "os.Getenv(\""
	if strings.HasPrefix(lineContent, prefix) {
		return lsp.CompletionResponse{
			Response: lsp.Response{
				RPC: "2.0",
				ID:  &id,
			},
			Result: []lsp.CompletionItem{},
		}
	}
	items := []lsp.CompletionItem{}

	// Ask your static analysis tools to figure out good completions
	for _, v := range completions {
		val := strings.Split(v, "=")
		items = append(items, lsp.CompletionItem{
			Label:         val[0],
			Detail:        "VALUE: " + Obfuscate(val[1]),
			Documentation: "",
		})

	}

	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}

	return response
}
