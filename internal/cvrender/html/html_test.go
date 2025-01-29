package render_html

import (
	"os"
	"testing"
	"text/template"

	"github.com/germainlefebvre4/cvwonder/internal/model"
	"github.com/stretchr/testify/assert"
)

var CvTest = model.CV{
	Company: model.Company{
		Name: "Company",
		Logo: "logo.png",
	},
	Person: model.Person{
		Name:        "John Doe",
		Depiction:   "I am a dummy Software Engineer for test.",
		Profession:  "Software Engineer",
		Location:    "Paris",
		Citizenship: "French",
		Email:       "john@doe.fr",
	},
	SocialNetworks: model.SocialNetworks{
		Github:        "dummy",
		Stackoverflow: "dummy",
	},
	Abstract: []model.Abstract{
		{
			Tr: "I am a dummy Software Engineer for test.",
		},
	},
	Career: model.Career{
		{
			CompanyName: "Company",
			CompanyLogo: "logo.png",
			Duration:    "2019 - 2020",
			Missions: []model.Mission{
				{
					Position:     "Software Engineer",
					Company:      "Company",
					Location:     "Paris",
					Dates:        "2019 - 2020",
					Summary:      "I was a Software Engineer.",
					Technologies: []string{"Go", "Python"},
					Description:  []string{"I was a Software Engineer."},
					Project:      "A project for a dummy company.",
				},
			},
		},
	},
	TechnicalSkills: model.TechnicalSkills{
		Domains: []model.Domain{
			{
				Name: "Development",
				Competencies: []model.Competency{
					{
						Name:  "Go",
						Level: 5,
					},
				},
			},
		},
	},
	SideProjects: []model.SideProject{
		{
			Name:        "Project",
			Description: "A project for a dummy company.",
		},
	},
	Certifications: []model.Certification{
		{
			CompanyName:       "CompanyName",
			CertificationName: "CertificationName",
			Issuer:            "Issuer",
			Date:              "Date",
			Link:              "Link",
			Badge:             "Badge",
		},
	},
	Education: []model.Education{
		{
			SchoolName: "SchoolName",
			SchoolLogo: "SchoolLogo",
			Degree:     "Degree",
			Location:   "Location",
			Dates:      "Dates",
			Link:       "Link",
		},
	},
}

func NewRenderHTMLServicesTest() RenderHTMLServices {
	return RenderHTMLServices{}
}

func TestRenderFormatHTML(t *testing.T) {
	testDirectory, _ := os.Getwd()
	baseDirectory := testDirectory + "/../../.."
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
				cv:              CvTest,
				baseDirectory:   baseDirectory,
				outputDirectory: baseDirectory + "/generated-test",
				inputFilename:   "cv",
				themeName:       "default",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		// Prepare
		// TODO: Add theme files

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
		// err := os.RemoveAll(tt.args.outputDirectory)
		// if err != nil {
		// 	t.Fatal(err)
		// }
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
	baseDirectory := testDirectory + "/../../.."
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
		wantErr error
	}{
		{
			name:   "Should generate template file",
			fields: fields{NewRenderHTMLServicesTest()},
			args: args{
				themeDirectory:    baseDirectory + "/themes/default",
				outputDirectory:   baseDirectory + "/generated-test",
				outputFilePath:    baseDirectory + "/generated-test/cv.html",
				outputTmpFilePath: baseDirectory + "/generated-test/cv.html.tmp",
				cv:                CvTest,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		// Prepare
		// TODO: Add theme files

		// Run test
		t.Run(tt.name, func(t *testing.T) {
			service := NewRenderHTMLServicesTest()
			assert.Equalf(
				t,
				tt.wantErr,
				service.generateTemplateFile(tt.args.themeDirectory, tt.args.outputDirectory, tt.args.outputFilePath, tt.args.outputTmpFilePath, tt.args.cv),
				"generateTemplateFile(%v, %v, %v, %v, %v)",
				tt.args.themeDirectory,
				tt.args.outputDirectory,
				tt.args.outputFilePath,
				tt.args.outputTmpFilePath,
				tt.args.cv,
			)
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
	baseDirectory := testDirectory + "/../../.."
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
