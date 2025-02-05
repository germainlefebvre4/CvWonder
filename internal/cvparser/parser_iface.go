package cvparser

import "github.com/germainlefebvre4/cvwonder/internal/model"

type ParserInterface interface {
	ParseFile(filePath string) (model.CV, error)
	convertFileContentToStruct(content []byte) model.CV
	readFile(filePath string) []byte
}

type ParserServices struct{}

func NewParserServices() (ParserInterface, error) {
	return &ParserServices{}, nil
}
