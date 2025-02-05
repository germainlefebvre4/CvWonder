package model

type CV struct {
	// Site struct {
	// 	URL string `yaml:"url"`
	// } `yaml:"site"`
	Company         Company         `yaml:"company"`
	Person          Person          `yaml:"person"`
	SocialNetworks  SocialNetworks  `yaml:"socialNetworks"`
	Abstract        []Abstract      `yaml:"abstract"`
	Career          []Career        `yaml:"career"`
	TechnicalSkills TechnicalSkills `yaml:"technicalSkills"`
	SideProjects    []SideProject   `yaml:"sideProjects"`
	Certifications  []Certification `yaml:"certifications"`
	Education       []Education     `yaml:"education"`
}

type Company struct {
	Name string `yaml:"name"`
	Logo string `yaml:"logo"`
}

type Person struct {
	Name        string `yaml:"name"`
	Depiction   string `yaml:"depiction"`
	Profession  string `yaml:"profession"`
	Location    string `yaml:"location"`
	Citizenship string `yaml:"citizenship"`
	Email       string `yaml:"email"`
	Site        string `yaml:"site"`
	Phone       string `yaml:"phone"`
}

type SocialNetworks struct {
	Github        string `yaml:"github"`
	Stackoverflow string `yaml:"stackoverflow"`
	Linkedin      string `yaml:"linkedin"`
	Twitter       string `yaml:"twitter"`
}

type Abstract struct {
	Tr string `yaml:"tr"`
}

type Career struct {
	CompanyName string    `yaml:"companyName"`
	CompanyLogo string    `yaml:"companyLogo"`
	Duration    string    `yaml:"duration,omitempty"`
	Missions    []Mission `yaml:"missions"`
}

type Mission struct {
	Position     string   `yaml:"position"`
	Company      string   `yaml:"company"`
	Location     string   `yaml:"location"`
	Dates        string   `yaml:"dates"`
	Summary      string   `yaml:"summary"`
	Technologies []string `yaml:"technologies"`
	Description  []string `yaml:"description"`
	Project      string   `yaml:"project,omitempty"`
}

type TechnicalSkills struct {
	Domains []Domain `yaml:"domains"`
}

type Domain struct {
	Name         string       `yaml:"name"`
	Competencies []Competency `yaml:"competencies"`
}

type Competency struct {
	Name  string `yaml:"name"`
	Level int    `yaml:"level"`
}

type SideProject struct {
	Name        string `yaml:"name"`
	Position    string `yaml:"position"`
	Description string `yaml:"description"`
	Link        string `yaml:"link"`
	Type        string `yaml:"type"`
	Langs       string `yaml:"langs"`
	Color       string `yaml:"color"`
}

type Certification struct {
	CompanyName       string `yaml:"companyName"`
	CertificationName string `yaml:"certificationName"`
	Issuer            string `yaml:"issuer"`
	Date              string `yaml:"date"`
	Link              string `yaml:"link"`
	Badge             string `yaml:"badge"`
}

type Education struct {
	SchoolName string `yaml:"schoolName"`
	SchoolLogo string `yaml:"schoolLogo"`
	Degree     string `yaml:"degree"`
	Location   string `yaml:"location"`
	Dates      string `yaml:"dates"`
	Link       string `yaml:"link"`
}
