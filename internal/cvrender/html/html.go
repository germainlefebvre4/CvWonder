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

func RenderFormatHTML(cv model.CV, outputDirectory string, inputFilename string, themeName string) error {
	logrus.Debug("Generating HTML")

	// Theme directory
	currentDirectory, err := os.Getwd()
	utils.CheckError(err)
	themeDirectory := currentDirectory + "/themes"

	// Inject custom functions in template
	funcMap := template.FuncMap{
		"dec":     func(i int) int { return i - 1 },
		"replace": strings.ReplaceAll,
		"join":    strings.Join,
	}

	// Template file
	tmplFile := themeDirectory + "/" + themeName + "/index.html"
	tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles(tmplFile)
	utils.CheckError(err)

	// Output file
	outputDirectory, err = filepath.Abs(outputDirectory)
	utils.CheckError(err)
	outputFilename := filepath.Base(inputFilename) + ".html"
	outputFilePath := outputDirectory + "/" + outputFilename
	outputTmpFilePath := outputFilePath + ".tmp"

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
	} else {
		outputTmpFile, err = os.OpenFile(outputTmpFilePath, os.O_WRONLY, 0644)
		utils.CheckError(err)
		defer outputTmpFile.Close()
	}

	// Generate output
	err = tmpl.ExecuteTemplate(outputTmpFile, "index.html", cv)
	utils.CheckError(err)

	// Copy file content from tmp to final
	// This approach avoid flooding file events in the watcher
	input, err := os.ReadFile(outputTmpFilePath)
	utils.CheckError(err)
	err = os.WriteFile(outputFilePath, input, 0644)
	utils.CheckError(err)

	// Clean the tmp file
	err = os.Remove(outputTmpFilePath)
	utils.CheckError(err)

	logrus.Debug("HTML file generated at:", outputFile)
	logrus.Debug("HTML file generated at:", outputFilePath)

	return err
}
