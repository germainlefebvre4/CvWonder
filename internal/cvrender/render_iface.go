package cvrender

import (
	"github.com/germainlefebvre4/cvwonder/internal/model"
)

type RenderInterface interface {
	Render(cv model.CV, outputDirectory string, inputFilePath string, themeName string, exportFormat string)
}
