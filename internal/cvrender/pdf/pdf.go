package render_pdf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/germainlefebvre4/cvwonder/internal/cvserve"
	"github.com/germainlefebvre4/cvwonder/internal/model"
	"github.com/germainlefebvre4/cvwonder/internal/utils"

	"github.com/go-rod/rod"
	"github.com/sirupsen/logrus"
)

func (r *RenderPDFServices) RenderFormatPDF(cv model.CV, outputDirectory string, inputFilename string, themeName string) error {
	logrus.Debug("Generating PDF")

	// Output file
	outputFilePath := r.generateOutputFile(outputDirectory, inputFilename)

	// Run the server to output the HTML
	localServerUrl := r.runWebServer(utils.CliArgs.Port, inputFilename, outputDirectory)

	// Open the browser and convert the page to PDF
	r.convertPageToPDF(localServerUrl, outputFilePath)

	return nil
}

func (*RenderPDFServices) convertPageToPDF(localServerUrl string, outputFilePath string) {
	err := rod.Try(func() {
		rod.New().MustConnect().MustPage(localServerUrl).MustWaitLoad().MustPDF(outputFilePath)
	})
	if err != nil {
		message := fmt.Sprintf("ERROR: Failed to connect to the server %s", localServerUrl)
		logrus.Fatal(message)
	}
}

func (*RenderPDFServices) runWebServer(port int, inputFilename string, outputDirectory string) string {
	if port == 0 {
		port = 8080
	}

	localServerUrl := fmt.Sprintf("http://localhost:%d/%s.html", port, inputFilename)
	logrus.Info("Serve temporary the CV on server at address ", localServerUrl)
	go func() {
		cvserve.StartServer(outputDirectory)

	}()
	return localServerUrl
}

func (*RenderPDFServices) generateOutputFile(outputDirectory string, inputFilename string) string {
	outputDirectory, err := filepath.Abs(outputDirectory)
	utils.CheckError(err)
	outputFilename := filepath.Base(inputFilename) + ".pdf"
	outputFilePath, err := filepath.Abs(outputDirectory + "/" + outputFilename)
	utils.CheckError(err)
	w, err := os.Create(outputFilePath)
	utils.CheckError(err)
	defer w.Close()
	return outputFilePath
}
