package cvparser

import (
	"fmt"
	"os"

	"github.com/germainlefebvre4/cvwonder/internal/model"
	"github.com/germainlefebvre4/cvwonder/internal/utils"

	"github.com/goccy/go-yaml"
	"github.com/sirupsen/logrus"
)

func (p *ParserServices) ParseFile(filePath string) (model.CV, error) {
	logrus.Debug("Parsing YAML file")
	fileContent, err := p.readFile(filePath)
	utils.CheckError(err)

	dataContent, err := p.convertFileContentToStruct(fileContent)
	utils.CheckError(err)

	return dataContent, nil
}

func (p *ParserServices) convertFileContentToStruct(content []byte) (model.CV, error) {
	cvOutput := model.CV{}
	err := yaml.Unmarshal([]byte(content), &cvOutput)
	if err != nil {
		logrus.Fatal("Error on reading YAML file. It may not be a valid YAML file.")
	}
	return cvOutput, err
}

func (p *ParserServices) readFile(filePath string) ([]byte, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		logrus.Fatal(fmt.Sprintf("Error while reading file: %s", filePath))
	}
	return content, err
}
