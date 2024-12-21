package render_pdf

import (
	"cvwonder/internal/model"

	"github.com/sirupsen/logrus"
)

func RenderFormatPDF(cv model.CV, outputDirectory string, inputFilename string, themeName string) error {
	logrus.Debug("Generating PDF")

	return nil
}
