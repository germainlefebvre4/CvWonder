package cvparser

import (
	"os"
	"path/filepath"
	"testing"

	mocks "github.com/germainlefebvre4/cvwonder/internal/cvparser/mocks"
	"github.com/germainlefebvre4/cvwonder/internal/model"
	"github.com/stretchr/testify/assert"
)

var cvByteGiven = []byte(`
person:
  name: Germain Lefebvre
`)

var cvYamlWanted = model.CV{
	Person: model.Person{
		Name: "Germain Lefebvre",
	},
}

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
			name: "Should return a model.CV",
			p:    &ParserServices{},
			args: args{
				content: cvByteGiven,
			},
			want:    cvYamlWanted,
			wantErr: false,
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
	mock.On("ParseFile", "test").Return(cvYamlWanted, nil)

	type args struct {
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
				filePath: baseDirectory + "/generated-test/TestStartLiveReloader.yaml",
			},
			want:    cvYamlWanted,
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
		err := os.WriteFile(baseDirectory+"/generated-test/TestStartLiveReloader.yaml", []byte(cvByteGiven), os.ModePerm)
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
