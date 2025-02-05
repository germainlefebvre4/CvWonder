package themes

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"

	theme_config "github.com/germainlefebvre4/cvwonder/internal/themes/config"

	"github.com/go-git/go-git/v5"
	"github.com/sirupsen/logrus"
)

func (t *ThemesService) Install(themeURL string) {
	logrus.Debug("Install")
	verifyTheme(themeURL)
	githubRepo := parseGitHubURL(themeURL)
	createThemesDir()
	downloadTheme(githubRepo)
}

func verifyTheme(themeURL string) {
	logrus.Debug("Verify theme")

	if !isGitHubURL(themeURL) {
		logrus.Error("Not a GitHub URL: ", themeURL)
	}
}

func isGitHubURL(input string) bool {
	formattedURL := fmt.Sprintf("%s%s", "https://", strings.ReplaceAll(input, "https://", ""))
	u, err := url.Parse(formattedURL)
	if err != nil {
		return false
	}
	host := u.Host
	if strings.Contains(host, ":") {
		host, _, err = net.SplitHostPort(host)
		if err != nil {
			return false
		}
	}
	return host == "github.com"
}

func parseGitHubURL(themeURL string) theme_config.GithubRepo {
	logrus.Debug("Parse GitHub URL")
	formattedURL := fmt.Sprintf("%s%s", "https://", strings.ReplaceAll(themeURL, "https://", ""))
	URL, err := url.Parse(formattedURL)
	if err != nil {
		logrus.Error("Error parsing URL")
	}
	path := strings.Split(URL.Path, "/")
	return theme_config.GithubRepo{URL: URL, Owner: path[1], Name: path[2]}
}

func downloadTheme(githubRepo theme_config.GithubRepo) {
	logrus.Debug("Download theme")

	themeConfig := theme_config.GetThemeConfigFromURL(githubRepo)

	themeDirectory := fmt.Sprintf("themes/%s", themeConfig.Slug)
	if _, err := os.Stat(themeDirectory); !os.IsNotExist(err) {
		logrus.Error("Theme '" + themeConfig.Name + "' already exists in " + themeDirectory + "/")
		return
	}
	_, err := git.PlainClone(themeDirectory, false, &git.CloneOptions{
		URL: githubRepo.URL.String(),
		// Progress: os.Stdout,
		Depth: 1,
	})
	if err != nil {
		logrus.Error("Error cloning theme")
	}

	logrus.Info("Theme '" + themeConfig.Name + "' successfully installed in " + themeDirectory + "/")
}
