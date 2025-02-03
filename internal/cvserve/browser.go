package cvserve

import (
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"

	"github.com/germainlefebvre4/cvwonder/internal/utils"

	"github.com/sirupsen/logrus"
)

func (s *ServeServices) OpenBrowser(outputDirectory string, inputFilePath string) {
	logrus.Debug("Opening browser")

	// Input file
	inputFilenameExt := path.Base(inputFilePath)
	inputFilename := inputFilenameExt[:len(inputFilenameExt)-len(path.Ext(inputFilenameExt))]

	// Output file
	outputFilename := filepath.Base(inputFilename) + ".html"

	url := fmt.Sprintf("http://localhost:%d/%s", utils.CliArgs.Port, outputFilename)
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	utils.CheckError(err)
}
