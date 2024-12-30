package cvrender

import (
	"path"

	render_html "github.com/germainlefebvre4/cvwonder/internal/cvrender/html"
	"github.com/germainlefebvre4/cvwonder/internal/model"
	"github.com/germainlefebvre4/cvwonder/internal/utils"

	"github.com/sirupsen/logrus"
)

// CVRender renders the CV based on html template located at internal/templates/index.html
func Render(cv model.CV, outputDirectory string, inputFilePath string, themeName string, exportFormat string) {
	logrus.Debug("Rendering CV")

	inputFilenameExt := path.Base(inputFilePath)
	inputFilename := inputFilenameExt[:len(inputFilenameExt)-len(path.Ext(inputFilenameExt))]

	// Generate HTML
	if exportFormat == "html" {
		err := render_html.RenderFormatHTML(cv, outputDirectory, inputFilename, themeName)
		utils.CheckError(err)
	} else if exportFormat == "pdf" {
		// Generate PDF
		err := render_pdf.RenderFormatPDF(cv, outputDirectory, inputFilename, themeName)
		utils.CheckError(err)
	} else {
		logrus.Fatal("Unsupported format")
	}
}
