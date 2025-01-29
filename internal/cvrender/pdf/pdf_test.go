package render_pdf

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewRenderPDFServicesTest() RenderPDFServices {
	return RenderPDFServices{}
}

func TestGenerateOutputFile(t *testing.T) {
	testDirectory, _ := os.Getwd()
	baseDirectory, err := filepath.Abs(testDirectory + "/../../..")
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		RenderPDFService RenderPDFServices
	}
	type args struct {
		outputDirectory string
		inputFilename   string
	}
	test := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "Should create and return output directory and file",
			fields: fields{NewRenderPDFServicesTest()},
			args: args{
				outputDirectory: baseDirectory + "/generated-test",
				inputFilename:   "TestGenerateOutputFile",
			},
			want: baseDirectory + "/generated-test/TestGenerateOutputFile.pdf",
		},
	}
	for _, tt := range test {
		// Prepare
		if _, err := os.Stat(baseDirectory + "/generated-test"); os.IsNotExist(err) {
			err := os.Mkdir(baseDirectory+"/generated-test", os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}
		}

		// Run test
		t.Run(tt.name, func(t *testing.T) {
			service := NewRenderPDFServicesTest()
			assert.Equalf(
				t,
				tt.want,
				service.generateOutputFile(tt.args.outputDirectory, tt.args.inputFilename),
				"generateOutputFile(%v, %v)",
				tt.args.outputDirectory,
				tt.args.inputFilename,
			)
		})

		// Clean
		err := os.RemoveAll(baseDirectory + "/generated-test")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestRunWebServer(t *testing.T) {
	testDirectory, _ := os.Getwd()
	baseDirectory, err := filepath.Abs(testDirectory + "/../../..")
	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		RenderPDFService RenderPDFServices
	}
	type args struct {
		port            int
		inputFilename   string
		outputDirectory string
	}
	test := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "Should run web server and return local server URL",
			fields: fields{NewRenderPDFServicesTest()},
			args: args{
				port:            18080,
				inputFilename:   "TestRunWebServer",
				outputDirectory: baseDirectory + "/generated-test",
			},
			want: "http://localhost:18080/TestRunWebServer.html",
		},
	}
	for _, tt := range test {
		// Prepare
		if _, err := os.Stat(baseDirectory + "/generated-test"); os.IsNotExist(err) {
			err := os.Mkdir(baseDirectory+"/generated-test", os.ModePerm)
			if err != nil {
				t.Fatal(err)
			}
		}
		err := os.WriteFile(baseDirectory+"/generated-test/TestRunWebServer.html", []byte("TestRunWebServer"), os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		// Run test
		t.Run(tt.name, func(t *testing.T) {
			service := NewRenderPDFServicesTest()
			assert.Equalf(
				t,
				tt.want,
				service.runWebServer(tt.args.port, tt.args.inputFilename, tt.args.outputDirectory),
				"runWebServer(%v, %v)",
				tt.args.port,
				tt.args.inputFilename,
				tt.args.outputDirectory,
			)
		})

		// Clean
		err = os.RemoveAll(baseDirectory + "/generated-test")
		if err != nil {
			t.Fatal(err)
		}
	}
}
