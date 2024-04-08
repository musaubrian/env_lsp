package analysis

import (
	"log"
	"os"
	"strings"

	"github.com/musaubrian/env_lsp/lsp"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	filetype := strings.Split(uri, ".")[1]
	s.Documents["filetype"] = filetype

	return getDiagnosticsForFile(s, text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text

	return getDiagnosticsForFile(s, text)
}

func (s *State) TextDocumentCompletion(id int, params lsp.CompletionParams, lg *log.Logger) lsp.CompletionResponse {
	uri := params.TextDocument.URI
	line := params.Position.Line
	content := s.Documents[uri]

	lines := strings.Split(content, "\n")
	lineContent := lines[line]
	completions := []string{}
	completionTtems := []lsp.CompletionItem{}

	envLocation, err := checkIfEnvExists(uri)
	if err != nil {
		lg.Println("`.env` not found ignoring...", err)
	}
	contents, err := readContents(envLocation)
	if err != nil {
		lg.Println("Could not read `.env` contents")
	}

	if err := loadEnvs(contents); err != nil {
		lg.Println("Could not load env variables: ", err)
	}

	for _, envs := range os.Environ() {
		completions = append(completions, envs)
	}
	prefixes := []string{"os.Getenv(\""}

	if s.Documents["filetype"] == "py" {
		prefixes = []string{"os.getenv(\"", "os.environ[\"", "os.environ.get(\""}
	}

	for _, completion := range completions {
		val := strings.Split(completion, "=")
		if strings.Contains(val[1], "REXIQI") {
			completionTtems = append(completionTtems, lsp.CompletionItem{
				Label:         val[0],
				Detail:        "VALUE: " + obfuscate(val[1]),
				Documentation: "Note: This is not the full length of the value obvi",
			})
		}
	}

	for _, prefix := range prefixes {
		if strings.Contains(lineContent, prefix) {
			return lsp.CompletionResponse{
				Response: lsp.Response{
					RPC: "2.0",
					ID:  &id,
				},
				Result: completionTtems,
			}
		}
	}

	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: []lsp.CompletionItem{},
	}
	return response
}

func getDiagnosticsForFile(s *State, text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	prefixes := []string{"os.Getenv("}

	if s.Documents["filetype"] == "py" {
		prefixes = []string{"os.getenv(", "os.environ[", "os.environ.get("}
	}

	for row, line := range strings.Split(text, "\n") {
		for _, prefix := range prefixes {
			if strings.Contains(line, prefix) {
				idx := strings.Index(line, prefix)
				diagnostics = append(diagnostics, lsp.Diagnostic{
					Range:    LineRange(row, idx, idx+len(prefix)),
					Severity: 2,
					Source:   "Nice",
					Message:  "This is great",
				})
			}
		}
	}

	return diagnostics
}
func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}
