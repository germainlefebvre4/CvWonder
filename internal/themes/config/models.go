package theme_config

import "net/url"

type GithubRepo struct {
	URL   *url.URL
	Owner string
	Name  string
}
