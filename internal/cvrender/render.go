package cvrender

import (
	"path"

	render_html "github.com/germainlefebvre4/cvwonder/internal/cvrender/html"
	render_pdf "github.com/germainlefebvre4/cvwonder/internal/cvrender/pdf"
	"github.com/germainlefebvre4/cvwonder/internal/model"
	"github.com/germainlefebvre4/cvwonder/internal/utils"

	"github.com/sirupsen/logrus"
)

type RenderServices struct {
	RenderHTMLService render_html.RenderHTMLInterface
	RenderPDFService  render_pdf.RenderPDFInterface
}

func NewRenderServices(
	renderHTMLInterface render_html.RenderHTMLInterface,
	renderPDFInterface render_pdf.RenderPDFInterface,
) (RenderInterface, error) {
	return &RenderServices{
		RenderHTMLService: renderHTMLInterface,
		RenderPDFService:  renderPDFInterface,
	}, nil
}

// CVRender renders the CV based on html template located at internal/templates/index.html
func (r *RenderServices) Render(cv model.CV, baseDirectory string, outputDirectory string, inputFilePath string, themeName string, exportFormat string) {
	logrus.Debug("Rendering CV")

	inputFilenameExt := path.Base(inputFilePath)
	inputFilename := inputFilenameExt[:len(inputFilenameExt)-len(path.Ext(inputFilenameExt))]

	// Generate HTML
	err := r.RenderHTMLService.RenderFormatHTML(cv, baseDirectory, outputDirectory, inputFilename, themeName)
	utils.CheckError(err)

	if exportFormat == "pdf" {
		// Generate PDF
		r.RenderPDFService.RenderFormatPDF(cv, outputDirectory, inputFilename, themeName)
	}
}
