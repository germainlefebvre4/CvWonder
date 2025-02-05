package render_html

import "github.com/germainlefebvre4/cvwonder/internal/model"

type RenderHTMLInterface interface {
	RenderFormatHTML(cv model.CV, baseDirectory string, outputDirectory string, inputFilename string, themeName string) error
	generateTemplateFile(themeDirectory string, outputDirectory string, outputFilePath string, outputTmpFilePath string, cv model.CV)
}

type RenderHTMLServices struct{}

func NewRenderHTMLServices() (RenderHTMLInterface, error) {
	return &RenderHTMLServices{}, nil
}
