package cvparser

import (
	"os"
	"path/filepath"
	"testing"

	mocks "github.com/germainlefebvre4/cvwonder/internal/cvparser/mocks"
	"github.com/germainlefebvre4/cvwonder/internal/fixtures"
	"github.com/germainlefebvre4/cvwonder/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestConvertFileContentToStruct(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		p       *ParserServices
		args    args
		want    model.CV
		wantErr bool
	}{
		{
			name: "Should return a model.CV simple example",
			p:    &ParserServices{},
			args: args{
				content: fixtures.CvYamlGood01,
			},
			want:    fixtures.CvModelGood01,
			wantErr: false,
		},
		{
			name: "Should return a model.CV full example",
			p:    &ParserServices{},
			args: args{
				content: fixtures.CvYamlGood02,
			},
			want:    fixtures.CvModelGood02,
			wantErr: false,
		},
		{
			name: "Should return an empty model.CV",
			p:    &ParserServices{},
			args: args{
				content: fixtures.CvYamlGood03,
			},
			want:    fixtures.CvModelGood03,
			wantErr: false,
		},
		{
			name: "Should return an error",
			p:    &ParserServices{},
			args: args{
				content: fixtures.CvYamlError01,
			},
			want:    fixtures.CvModelError01,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.convertFileContentToStruct(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertFileContentToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseFile(t *testing.T) {
	testDirectory, _ := os.Getwd()
	baseDirectory, err := filepath.Abs(testDirectory + "/../..")
	if err != nil {
		t.Fatal(err)
	}

	// Mocks
	mock := &mocks.ParserInterfaceMock{}
	mock.On("ParseFile", "test").Return(fixtures.CvModelGood01, nil)

	type args struct {
		content  []byte
		filePath string
	}
	tests := []struct {
		name    string
		p       *ParserServices
		args    args
		want    model.CV
		wantErr bool
	}{
		{
			name: "Should return a model.CV",
			p:    &ParserServices{},
			args: args{
				content:  fixtures.CvYamlGood01,
				filePath: baseDirectory + "/generated-test/TestStartLiveReloader.yaml",
			},
			want:    fixtures.CvModelGood01,
			wantErr: false,
		},
		{
			name: "Should return a model.CV",
			p:    &ParserServices{},
			args: args{
				content:  fixtures.CvYamlGood02,
				filePath: baseDirectory + "/generated-test/TestStartLiveReloader.yaml",
			},
			want:    fixtures.CvModelGood02,
			wantErr: false,
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
		err := os.WriteFile(baseDirectory+"/generated-test/TestStartLiveReloader.yaml", tt.args.content, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.ParseFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Nil(t, err)
		})

		// Clean
		err = os.RemoveAll(baseDirectory + "/generated-test")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestReadFile(t *testing.T) {
	testDirectory, _ := os.Getwd()
	baseDirectory, err := filepath.Abs(testDirectory + "/../..")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		p       *ParserServices
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Should return a file content",
			p:    &ParserServices{},
			args: args{
				filePath: baseDirectory + "/generated-test/TestReadFile.yaml",
			},
			want:    []byte(fixtures.CvYamlGood01),
			wantErr: false,
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
		err := os.WriteFile(baseDirectory+"/generated-test/TestReadFile.yaml", tt.want, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.readFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})

		// Clean
		err = os.RemoveAll(baseDirectory + "/generated-test")
		if err != nil {
			t.Fatal(err)
		}
	}
}
