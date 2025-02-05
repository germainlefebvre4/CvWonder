package themes

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/germainlefebvre4/cvwonder/internal/fixtures"
	theme_config "github.com/germainlefebvre4/cvwonder/internal/themes/config"
	"github.com/stretchr/testify/assert"
)

func NewThemesServicesTest() ThemesService {
	return ThemesService{}
}

func TestListThemes(t *testing.T) {
	testDirectory, _ := os.Getwd()
	baseDirectory, err := filepath.Abs(testDirectory + "/../..")
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		ThemesService ThemesService
	}
	type args struct {
		themeName       string
		ThemeConfigYaml []byte
		CvHtml          []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Should list theme #01",
			fields: fields{NewThemesServicesTest()},
			args: args{
				themeName:       theme_config.ThemeConfigModelGood01.Name,
				ThemeConfigYaml: theme_config.ThemeConfigYamlGood01,
				CvHtml:          fixtures.CvHtmlGood01,
			},
		},
		{
			name:   "Should list theme #02",
			fields: fields{NewThemesServicesTest()},
			args: args{
				themeName:       theme_config.ThemeConfigModelGood02.Name,
				ThemeConfigYaml: theme_config.ThemeConfigYamlGood02,
				CvHtml:          fixtures.CvHtmlGood01,
			},
		},
	}

	// Prepare
	if _, err := os.Stat(baseDirectory + "/themes-test"); os.IsNotExist(err) {
		err := os.Mkdir(baseDirectory+"/themes-test", os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	}

	for _, tt := range tests {
		// Prepare
		if _, err := os.Stat(filepath.Join(baseDirectory, "themes-test", tt.args.themeName)); os.IsNotExist(err) {
			err := os.Mkdir(filepath.Join(baseDirectory, "themes-test", tt.args.themeName), os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}
		}
		err = os.WriteFile(filepath.Join(baseDirectory, "themes-test", tt.args.themeName, "theme.yaml"), tt.args.ThemeConfigYaml, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		err := os.WriteFile(filepath.Join(baseDirectory, "themes-test", tt.args.themeName, "index.html"), tt.args.CvHtml, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		// Run test
		t.Run("Should list themes", func(t *testing.T) {
			output := captureOutput(func() {
				listThemes(baseDirectory, "themes-test")
			})
			// assert.Contains(t, output, "Directory\tName\tDescription\tAuthor\n")
			assert.Contains(t, output, tt.args.themeName)

		})

		// Clean
		err = os.RemoveAll(filepath.Join(baseDirectory, "themes-test", tt.args.themeName))
		if err != nil {
			t.Fatal(err)
		}
	}

	// Clean
	err = os.RemoveAll(filepath.Join(baseDirectory, "themes-test"))
	if err != nil {
		t.Fatal(err)
	}
}

func captureOutput(f func()) string {
	// var buf bytes.Buffer
	// log.SetOutput(&buf)
	// f()
	// log.SetOutput(os.Stderr)
	// return buf.String()

	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out)
}
