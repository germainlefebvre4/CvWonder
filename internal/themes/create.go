package themes

import (
	"os"

	"github.com/goccy/go-yaml"
	"github.com/mozillazg/go-slugify"
	"github.com/sirupsen/logrus"

	theme_config "github.com/germainlefebvre4/cvwonder/internal/themes/config"
)

func (t *ThemesService) Create(themeName string) {
	logrus.Debug("Create")
	themeSlugName := slugify.Slugify(themeName)

	createThemesDir()
	if _, err := os.Stat("themes/" + themeSlugName); os.IsNotExist(err) {
		// Create theme directory
		createNewThemeDir(themeSlugName)

		// Create theme.yaml
		createThemeConfig(themeName, themeSlugName)

		logrus.Info("Your theme '" + themeName + "' has been created in the directory themes/" + themeSlugName + "/.")
	} else {
		logrus.Error("Theme '" + themeSlugName + "' already exists.")
	}
}

func createThemeConfig(themeName string, themeSlugName string) {
	themeConfig := theme_config.ThemeConfig{
		Name:        themeName,
		Slug:        themeSlugName,
		Description: "Description of the new theme.",
		Author:      "Anonymous",
	}
	err := createThemeConfigFile("themes/"+themeSlugName+"/theme.yaml", themeConfig)
	if err != nil {
		logrus.Fatal("Error creating theme.yaml: ", err)
	}
}

func createThemeConfigFile(filePath string, themeConfig theme_config.ThemeConfig) error {
	// Create theme.yaml file
	file, err := os.Create(filePath)
	if err != nil {
		logrus.Error("Error creating theme.yaml: ", err)
	}
	defer file.Close()

	// Write theme.yaml
	configYaml, err := yaml.Marshal(&themeConfig)
	if err != nil {
		logrus.Error("fail to marshal credentials")
	}
	err = os.WriteFile(filePath, configYaml, 0600)

	return nil
}
