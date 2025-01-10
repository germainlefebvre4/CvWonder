package themes

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/sirupsen/logrus"
)

func Install(themeURL string) {
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
	URL2 := strings.ReplaceAll(input, "https://", "")
	URL3 := fmt.Sprintf("%s%s", "https://", URL2)
	u, err := url.Parse(URL3)
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

func parseGitHubURL(themeURL string) GithubRepo {
	logrus.Debug("Parse GitHub URL")
	URL2 := strings.ReplaceAll(themeURL, "https://", "")
	URL3 := fmt.Sprintf("%s%s", "https://", URL2)
	URL, err := url.Parse(URL3)
	if err != nil {
		logrus.Error("Error parsing URL")
	}
	path := strings.Split(URL.Path, "/")
	return GithubRepo{URL: URL, Owner: path[1], Name: path[2]}
}

func downloadTheme(githubRepo GithubRepo) {
	logrus.Debug("Download theme")

	themeConfig := GetThemeConfigFromURL(githubRepo)

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
