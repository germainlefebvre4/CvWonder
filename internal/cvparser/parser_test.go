package cvparser

import (
	"testing"

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
