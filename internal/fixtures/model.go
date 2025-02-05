package fixtures

import (
	"github.com/germainlefebvre4/cvwonder/internal/model"
)

var CvYamlGood01 = []byte(`
person:
  name: John Doe
`)

var CvYamlGood02 = []byte(`
---
company:
  name: Zatsit

person:
  name: Germain
  depiction: profile.png
  profession: Platform Engineer
  location: Lille
  citizenship: FR
  email: germain.lefebvre@mycompany.fr
  site: http://germainlefebvre.fr
  phone: +33 6 00 00 00 00

socialNetworks:
  github: germainlefebvre4
  stackoverflow: germainlefebvre4
  linkedin: germainlefebvre4
  twitter: germainlefebvr4

abstract:
  - tr: "I am a Platform Engineer looking for people to share knowledge to each other."
  - tr: "This section can be a multiples lines of text."
  - tr: "This section can be a multiples lines of text again."

career:
  - companyName: Zatsit
    companyLogo: images/zatsit-logo.webp
    duration: 10 mois, aujourd'hui
    missions:
      - position: Platform Engineer
        company: Adeo
        location: Ronchin, France
        dates: 2024, mars - 2024, décembre
        summary: Construire une IDP, plateforme interne de développement, totalement managée pour aider les développeurs à se focaliser sur le code. Sur base du code source, la plateforme provisionne l'infrastructure sous-jacente, les base de données, la construction des artefact et publication sur la registry, le déploiement dans Kubernetes, l'intégration du monitoring avec Datadog et construction des Monitors.
        technologies:
          - ArgoCD
          - Kubernetes
          - K8s Operrator
          - Crossplane
          - Vault
          - Github Actions
          - JFrog Artifactory
          - Backstage
          - Python
          - Golang
        description:
          - Développement de l'operator Kubernetes responsable du provisioning des bases de données
          - Développement des Compositions Crossplane pour provisionner les base de données
          - Développement de l'API de l'IDP en Golang

technicalSkills:
  list: []
  domains:
    - name: Cloud
      competencies:
        - name: AWS
          level: 80
        - name: GCP
          level: 70
        - name: Azure
          level: 40

sideProjects:
  - name: cvwonder
    position: maintainer
    description: A CLI to render your CV from a YAML file.
    link: germainlefebvre4/cvwonder
    type: github
    langs: Go
    color: 3572A5

certifications:
  - companyName: AWS
    certificationName: Solutions Architect Associate
    issuer: Coursera
    date: Mars 2018
    link: https://www.credly.com/badges/dd09dc40-9ef8-43a4-addb-d861d4dadf26/public_url
    badge: images/aws-certified-solutions-architect-associate.png

education:
  - schoolName: IG2I - Centrale
    schoolLogo: images/centrale-lille-logo.webp
    degree: Titre d'ingénieur (BAC+5)
    location: Lens, France
    dates: 2019 - 2014
    link: https://ig2i.centralelille.fr
`)

var CvYamlGood03 = []byte(`
another:
  field: value
`)

var CvYamlError01 = []byte(`
person:
  name: John Doe
    depiction: I am a dummy Software Engineer for test.
 wrong: field
`)

var CvModelGood01 = model.CV{
	Person: model.Person{
		Name: "John Doe",
	},
}

var CvModelGood02 = model.CV{
	Company: model.Company{
		Name: "Zatsit",
	},
	Person: model.Person{
		Name:        "Germain",
		Depiction:   "profile.png",
		Profession:  "Platform Engineer",
		Location:    "Lille",
		Citizenship: "FR",
		Email:       "germain.lefebvre@mycompany.fr",
		Site:        "http://germainlefebvre.fr",
		Phone:       "+33 6 00 00 00 00",
	},
	SocialNetworks: model.SocialNetworks{
		Github:        "germainlefebvre4",
		Stackoverflow: "germainlefebvre4",
		Linkedin:      "germainlefebvre4",
		Twitter:       "germainlefebvr4",
	},
	Abstract: []model.Abstract{
		{
			Tr: "I am a Platform Engineer looking for people to share knowledge to each other.",
		},
		{
			Tr: "This section can be a multiples lines of text.",
		},
		{
			Tr: "This section can be a multiples lines of text again.",
		},
	},
	Career: []model.Career{
		{
			CompanyName: "Zatsit",
			CompanyLogo: "images/zatsit-logo.webp",
			Duration:    "10 mois, aujourd'hui",
			Missions: []model.Mission{
				{
					Position: "Platform Engineer",
					Company:  "Adeo",
					Location: "Ronchin, France",
					Dates:    "2024, mars - 2024, décembre",
					Summary:  "Construire une IDP, plateforme interne de développement, totalement managée pour aider les développeurs à se focaliser sur le code. Sur base du code source, la plateforme provisionne l'infrastructure sous-jacente, les base de données, la construction des artefact et publication sur la registry, le déploiement dans Kubernetes, l'intégration du monitoring avec Datadog et construction des Monitors.",
					Technologies: []string{
						"ArgoCD",
						"Kubernetes",
						"K8s Operrator",
						"Crossplane",
						"Vault",
						"Github Actions",
						"JFrog Artifactory",
						"Backstage",
						"Python",
						"Golang",
					},
					Description: []string{
						"Développement de l'operator Kubernetes responsable du provisioning des bases de données",
						"Développement des Compositions Crossplane pour provisionner les base de données",
						"Développement de l'API de l'IDP en Golang",
					},
				},
			},
		},
	},
	TechnicalSkills: model.TechnicalSkills{
		Domains: []model.Domain{
			{
				Name: "Cloud",
				Competencies: []model.Competency{
					{
						Name:  "AWS",
						Level: 80,
					},
					{
						Name:  "GCP",
						Level: 70,
					},
					{
						Name:  "Azure",
						Level: 40,
					},
				},
			},
		},
	},
	SideProjects: []model.SideProject{
		{
			Name:        "cvwonder",
			Position:    "maintainer",
			Description: "A CLI to render your CV from a YAML file.",
			Link:        "germainlefebvre4/cvwonder",
			Type:        "github",
			Langs:       "Go",
			Color:       "3572A5",
		},
	},
	Certifications: []model.Certification{
		{
			CompanyName:       "AWS",
			CertificationName: "Solutions Architect Associate",
			Issuer:            "Coursera",
			Date:              "Mars 2018",
			Link:              "https://www.credly.com/badges/dd09dc40-9ef8-43a4-addb-d861d4dadf26/public_url",
			Badge:             "images/aws-certified-solutions-architect-associate.png",
		},
	},
	Education: []model.Education{
		{
			SchoolName: "IG2I - Centrale",
			SchoolLogo: "images/centrale-lille-logo.webp",
			Degree:     "Titre d'ingénieur (BAC+5)",
			Location:   "Lens, France",
			Dates:      "2019 - 2014",
			Link:       "https://ig2i.centralelille.fr",
		},
	},
}

var CvModelGood03 = model.CV{}

var CvModelError01 = model.CV{}

var CvHtmlTemplate01 = []byte(`{{ .Person.Name }}`)

var CvHtmlGood01 = []byte(`John Doe`)
