package render_html

import (
	"os"
	"path/filepath"
	"testing"
	"text/template"

	"github.com/germainlefebvre4/cvwonder/internal/fixtures"
	"github.com/germainlefebvre4/cvwonder/internal/model"
	"github.com/stretchr/testify/assert"
)

func NewRenderHTMLServicesTest() RenderHTMLServices {
	return RenderHTMLServices{}
}

func TestRenderFormatHTML(t *testing.T) {
	testDirectory, _ := os.Getwd()
	baseDirectory, err := filepath.Abs(testDirectory + "/../../..")
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		RenderHTMLService RenderHTMLServices
	}
	type args struct {
		cv              model.CV
		baseDirectory   string
		outputDirectory string
		inputFilename   string
		themeName       string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "Should render HTML",
			fields: fields{NewRenderHTMLServicesTest()},
			args: args{
				cv:              fixtures.CvModelGood01,
				baseDirectory:   baseDirectory,
				outputDirectory: baseDirectory + "/generated-test",
				inputFilename:   "cv",
				themeName:       "test",
			},
			wantErr: nil,
		},
	}
	// Create the test theme
	if _, err := os.Stat(baseDirectory + "/themes/test"); os.IsNotExist(err) {
		err := os.Mkdir(baseDirectory+"/themes/test", os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, tt := range tests {
		// Prepare
		// Create the theme template file
		if _, err := os.Stat(baseDirectory + "/themes/test"); os.IsNotExist(err) {
			err := os.Mkdir(baseDirectory+"/themes/test", os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}
		}
		err = os.WriteFile(baseDirectory+"/themes/test/index.html", fixtures.CvHtmlTemplate01, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		// Run test
		t.Run(tt.name, func(t *testing.T) {
			service := NewRenderHTMLServicesTest()
			assert.Equalf(
				t,
				tt.wantErr,
				service.RenderFormatHTML(tt.args.cv, tt.args.baseDirectory, tt.args.outputDirectory, tt.args.inputFilename, tt.args.themeName),
				"RenderFormatHTML(%v, %v, %v, %v)",
				tt.args.cv,
				tt.args.outputDirectory,
				tt.args.inputFilename,
				tt.args.themeName,
			)
		})

		// Clean
		err = os.RemoveAll(tt.args.outputDirectory)
		if err != nil {
			t.Fatal(err)
		}
	}
	err = os.RemoveAll(baseDirectory + "/themes/test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTemplateFunctions(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		args args
		want template.FuncMap
	}{
		{
			name: "Should return template functions",
			args: args{},
			want: template.FuncMap{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.IsType(
				t,
				tt.want,
				getTemplateFunctions(),
				"getTemplateFunctions()",
			)
		})
	}
}

func TestGenerateTemplateFile(t *testing.T) {
	testDirectory, _ := os.Getwd()
	baseDirectory, err := filepath.Abs(testDirectory + "/../../..")
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		RenderHTMLService RenderHTMLServices
	}
	type args struct {
		themeDirectory    string
		outputDirectory   string
		outputFilePath    string
		outputTmpFilePath string
		cv                model.CV
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Should generate template file",
			fields: fields{NewRenderHTMLServicesTest()},
			args: args{
				themeDirectory:    baseDirectory + "/themes/default",
				outputDirectory:   baseDirectory + "/generated-test",
				outputFilePath:    baseDirectory + "/generated-test/cv.html",
				outputTmpFilePath: baseDirectory + "/generated-test/cv.html.tmp",
				cv:                fixtures.CvModelGood01,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		// Prepare
		// TODO: Add theme files

		// Run test
		t.Run(tt.name, func(t *testing.T) {
			service := NewRenderHTMLServicesTest()
			service.generateTemplateFile(tt.args.themeDirectory, tt.args.outputDirectory, tt.args.outputFilePath, tt.args.outputTmpFilePath, tt.args.cv)
			assert.DirExists(t, tt.args.outputDirectory)
			assert.FileExists(t, tt.args.outputFilePath)
			assert.FileExists(t, tt.args.outputTmpFilePath)
		})

		// Clean
		err := os.RemoveAll(tt.args.outputDirectory)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestCopyTemplateFileContent(t *testing.T) {
	testDirectory, _ := os.Getwd()
	baseDirectory, err := filepath.Abs(testDirectory + "/../../..")
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "Should copy template file content",
			args:    args{baseDirectory + "/generated-test/TestCopyTemplateFileContent.test.tmp", baseDirectory + "/generated-test/TestCopyTemplateFileContent.test"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		// Prepare
		if _, err := os.Stat(baseDirectory + "/generated-test"); os.IsNotExist(err) {
			err := os.Mkdir(baseDirectory+"/generated-test", os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}
		}
		err := os.WriteFile(tt.args.src, []byte("test"), os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f1, err1 := os.ReadFile(tt.args.src)
		if err1 != nil {
			t.Fatal(err1)
		}

		// Run test
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.wantErr,
				copyTemplateFileContent(tt.args.src, tt.args.dst),
				"copyTemplateFileContent(%v, %v)",
				tt.args.src,
				tt.args.dst,
			)
		})
		assert.FileExists(t, tt.args.dst)

		// Clean
		f2, err2 := os.ReadFile(tt.args.dst)
		if err2 != nil {
			t.Fatal(err2)
		}
		assert.Equal(t, f1, f2)

		err = os.Remove(tt.args.dst)
		if err != nil {
			t.Fatal(err)
		}
	}
}
