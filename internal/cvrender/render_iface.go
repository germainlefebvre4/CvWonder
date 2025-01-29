package cvrender

import (
	"github.com/germainlefebvre4/cvwonder/internal/model"
)

type RenderInterface interface {
	Render(cv model.CV, baseDirectory string, outputDirectory string, inputFilePath string, themeName string, exportFormat string)
}
