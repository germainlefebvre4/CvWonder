package themes

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	theme_config "github.com/germainlefebvre4/cvwonder/internal/themes/config"

	"github.com/sirupsen/logrus"
)

func (t *ThemesService) List() {
	logrus.Debug("List themes")
	baseDirectory, _ := os.Getwd()
	themeDir := "themes"

	// List directories in themes directory
	listThemes(baseDirectory, themeDir)
}

func listThemes(baseDirectory string, themeDir string) {
	if themeDir == "" {
		themeDir = "themes"
	}
	dirs, err := os.ReadDir(filepath.Join(baseDirectory, themeDir))
	if err != nil {
		logrus.Fatal("Error reading themes directory: ", err)
	}

	// Print directories in a table
	output := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer output.Flush()
	// Table header
	fmt.Fprintf(output, "%s", strings.ToUpper("Directory\tName\tDescription\tAuthor\n"))
	// Table body
	for _, dir := range dirs {
		if dir.IsDir() {
			printRow(baseDirectory, themeDir, dir, output)
			continue
		} else if dir.Type() == os.ModeSymlink {
			if _, err := os.Stat(filepath.Join(baseDirectory, themeDir, dir.Name())); err == nil {
				printRow(baseDirectory, themeDir, dir, output)
			} else {
				logrus.Warn("Symlink to non-existing directory: ", dir.Name())
			}
		} else {
			logrus.Warn("Non-directory file in themes directory: ", dir.Name())
		}
	}
}

func printRow(baseDirectory string, themeDir string, dir os.DirEntry, output *tabwriter.Writer) {
	themeConfig := theme_config.GetThemeConfigFromDir(filepath.Join(baseDirectory, themeDir, dir.Name()))
	fmt.Fprintf(output, "%s\t%s\t%s\t%s\n", themeConfig.Slug, themeConfig.Name, themeConfig.Description, themeConfig.Author)
}
