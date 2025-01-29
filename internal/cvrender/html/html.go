package render_html

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/germainlefebvre4/cvwonder/internal/model"
	utils "github.com/germainlefebvre4/cvwonder/internal/utils"

	"github.com/sirupsen/logrus"
)

func (r *RenderHTMLServices) RenderFormatHTML(cv model.CV, outputDirectory string, inputFilename string, themeName string) error {
	logrus.Debug("Generating HTML")

	// Theme directory
	currentDirectory, err := os.Getwd()
	utils.CheckError(err)
	themeDirectory := filepath.Join(currentDirectory, "themes", themeName)

	// Output file
	outputDirectory, err = filepath.Abs(outputDirectory)
	utils.CheckError(err)
	outputFilename := filepath.Base(inputFilename) + ".html"
	outputFilePath := filepath.Join(outputDirectory, outputFilename)
	outputTmpFilePath := outputFilePath + ".tmp"

	// Generate template file
	err = r.generateTemplateFile(themeDirectory, outputDirectory, outputFilePath, outputTmpFilePath, cv)
	utils.CheckError(err)

	// Copy template file to output directory
	err = copyTemplateFileContent(outputTmpFilePath, outputFilePath)
	utils.CheckError(err)

	// Copy theme assets to output directory
	err = utils.CopyDirectory(themeDirectory, outputDirectory)
	utils.CheckError(err)

	return err
}

func getTemplateFunctions() template.FuncMap {
	funcMap := template.FuncMap{
		"dec":     func(i int) int { return i - 1 },
		"replace": strings.ReplaceAll,
		"join":    strings.Join,
	}
	return funcMap
}

func (r *RenderHTMLServices) generateTemplateFile(themeDirectory string, outputDirectory string, outputFilePath string, outputTmpFilePath string, cv model.CV) error {
	// Inject custom functions in template
	funcMap := getTemplateFunctions()

	// Template file
	tmplFile := themeDirectory + "/index.html"
	tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles(tmplFile)
	utils.CheckError(err)

	// Create output file and directory
	if _, err := os.Stat(outputDirectory); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(outputDirectory, os.ModePerm)
		utils.CheckError(err)
	}
	outputFile, err := os.Create(outputFilePath)
	utils.CheckError(err)
	defer outputFile.Close()
	var outputTmpFile *os.File
	if _, err := os.Stat(outputTmpFilePath); errors.Is(err, os.ErrNotExist) {
		outputTmpFile, err = os.Create(outputTmpFilePath)
		utils.CheckError(err)
		defer outputTmpFile.Close()
	}

	// Generate output
	err = tmpl.ExecuteTemplate(outputTmpFile, "index.html", cv)
	if err != nil {
		errFile := os.Remove(outputTmpFilePath)
		utils.CheckError(errFile)
		logrus.Fatal(err)
	}

	logrus.Debug("HTML file generated at:", outputFilePath)

	return nil
}

func copyTemplateFileContent(outputTmpFilePath string, outputFilePath string) error {
	// Note: Copy file content from tmp to final to avoid flooding file events in the watcher
	input, err := os.ReadFile(outputTmpFilePath)
	utils.CheckError(err)
	err = os.WriteFile(outputFilePath, input, 0644)
	utils.CheckError(err)

	// Clean the tmp file
	err = os.Remove(outputTmpFilePath)
	utils.CheckError(err)

	return nil
}
