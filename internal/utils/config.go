package utils

type Configuration struct {
	InputFile       string `mapstructure:"INPUT_FILE"`
	OutputDirectory string `mapstructure:"OUTPUT_DIRECTORY"`
	ThemeName       string `mapstructure:"THEME_NAME"`
	Format          string `mapstructure:"FORMAT"`
	Watch           bool   `mapstructure:"WATCH"`
	Verbose         bool   `mapstructure:"VERBOSE"`
	Port            int    `mapstructure:"PORT"`
}

var CliArgs Configuration

// func LoadConfig(path string) (config Configuration, err error) {
// 	// viper.AddConfigPath(path)
// 	// viper.SetConfigName("app")
// 	// viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}

// 	err = viper.Unmarshal(&config)
// 	return
// }
