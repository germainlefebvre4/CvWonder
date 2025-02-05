package render_pdf

import (
	"github.com/germainlefebvre4/cvwonder/internal/cvserve"
	"github.com/germainlefebvre4/cvwonder/internal/model"
)

type RenderPDFInterface interface {
	RenderFormatPDF(cv model.CV, outputDirectory string, inputFilename string, themeName string)
}

type RenderPDFServices struct {
	ServeService cvserve.ServeInterface
}

func NewRenderPDFServices(
	serveInterface cvserve.ServeInterface,
) (RenderPDFInterface, error) {
	return &RenderPDFServices{
		ServeService: serveInterface,
	}, nil
}
