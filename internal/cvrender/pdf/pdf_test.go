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
