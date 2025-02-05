package main

import (
	"fmt"
	"os"

	cmdThemes "github.com/germainlefebvre4/cvwonder/cmd/cvwonder/themes"
	"github.com/germainlefebvre4/cvwonder/internal/cvparser"
	"github.com/germainlefebvre4/cvwonder/internal/cvrender"
	render_html "github.com/germainlefebvre4/cvwonder/internal/cvrender/html"
	render_pdf "github.com/germainlefebvre4/cvwonder/internal/cvrender/pdf"
	"github.com/germainlefebvre4/cvwonder/internal/cvserve"
	"github.com/germainlefebvre4/cvwonder/internal/model"
	"github.com/germainlefebvre4/cvwonder/internal/utils"
	"github.com/germainlefebvre4/cvwonder/internal/watcher"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Commit  = ""
	Version = "x.x.x"
	Date    = ""
)

func main() {
	var rootCmd = &cobra.Command{
		PreRun: utils.ToggleDebug,
		Use:    "cvwonder [COMMAND] [OPTIONS]",
		Short:  "CV Wonder",
		Long:   `CV Wonder - Generate your CV with Wonder!`,
	}

	var generateCmd = &cobra.Command{
		PreRun:  utils.ToggleDebug,
		Use:     "generate",
		Aliases: []string{"g", "gen"},
		Short:   "Generate the CV",
		Long:    `Generate the CV`,
		Run: func(cmd *cobra.Command, args []string) {
			if utils.CliArgs.ThemeName == "" {
				utils.CliArgs.ThemeName = "default"
			}

			// Build InputFile object
			inputFile := model.BuildInputFile(utils.CliArgs.InputFile)

			// Build OutputDirectory object
			outputDir := model.BuildOutputDirectory(utils.CliArgs.OutputDirectory)

			logrus.Info("CVRender")
			logrus.Info("  Input file: ", inputFile.RelativePath)
			logrus.Info("  Output directory: ", outputDir.RelativePath)
			logrus.Info("  Theme: ", utils.CliArgs.ThemeName)
			logrus.Info("  Format: ", utils.CliArgs.Format)
			logrus.Info("")

			// Parse the CV
			perserService, err := cvparser.NewParserServices()
			utils.CheckError(err)
			content, err := perserService.ParseFile(inputFile.FullPath)
			utils.CheckError(err)

			// Create render services
			serveService, err := cvserve.NewServeServices()
			utils.CheckError(err)
			renderHTMLService, err := render_html.NewRenderHTMLServices()
			utils.CheckError(err)
			renderPDFService, err := render_pdf.NewRenderPDFServices(serveService)
			utils.CheckError(err)
			renderService, err := cvrender.NewRenderServices(renderHTMLService, renderPDFService)
			utils.CheckError(err)

			// Render the CV
			baseDirectory, err := os.Getwd()
			utils.CheckError(err)
			renderService.Render(content, baseDirectory, outputDir.FullPath, inputFile.FullPath, utils.CliArgs.ThemeName, utils.CliArgs.Format)
			utils.CheckError(err)

			logrus.Info("CV generated successfully")
		},
	}

	var serveCmd = &cobra.Command{
		PreRun:  utils.ToggleDebug,
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Generate and serve the CV",
		Long:    `Generate and serve the CV`,
		Run: func(cmd *cobra.Command, args []string) {
			if utils.CliArgs.ThemeName == "" {
				utils.CliArgs.ThemeName = "default"
			}

			// Build InputFile object
			inputFile := model.BuildInputFile(utils.CliArgs.InputFile)

			// Build OutputDirectory object
			outputDir := model.BuildOutputDirectory(utils.CliArgs.OutputDirectory)

			logrus.Info("CVRender")
			logrus.Info("  Input file: ", inputFile.RelativePath)
			logrus.Info("  Output directory: ", outputDir.RelativePath)
			logrus.Info("  Theme: ", utils.CliArgs.ThemeName)
			logrus.Info("  Format: ", utils.CliArgs.Format)
			logrus.Info("  Watch: ", utils.CliArgs.Watch)
			logrus.Info()

			// Parse the CV
			perserService, err := cvparser.NewParserServices()
			utils.CheckError(err)
			content, err := perserService.ParseFile(inputFile.FullPath)
			utils.CheckError(err)

			// Create render services
			serveService, err := cvserve.NewServeServices()
			utils.CheckError(err)
			renderHTMLService, err := render_html.NewRenderHTMLServices()
			utils.CheckError(err)
			renderPDFService, err := render_pdf.NewRenderPDFServices(serveService)
			utils.CheckError(err)
			renderService, err := cvrender.NewRenderServices(renderHTMLService, renderPDFService)
			utils.CheckError(err)

			// Render the CV
			baseDirectory, err := os.Getwd()
			utils.CheckError(err)
			renderService.Render(content, baseDirectory, outputDir.FullPath, inputFile.FullPath, utils.CliArgs.ThemeName, utils.CliArgs.Format)
			utils.CheckError(err)

			if utils.CliArgs.Watch {
				watcherService, err := watcher.NewWatcherServices()
				utils.CheckError(err)
				go watcherService.ObserveFileEvents(renderService, baseDirectory, outputDir.FullPath, inputFile.FullPath, utils.CliArgs.ThemeName, utils.CliArgs.Format)
			}
			// Serve the CV
			serveService.OpenBrowser(outputDir.FullPath, inputFile.FullPath)
			serveService.StartLiveReloader(utils.CliArgs.Port, outputDir.FullPath, inputFile.FullPath)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&utils.CliArgs.InputFile, "input", "i", "cv.yml", "Input file in YAML format (required). Default is 'cv.yml'")
	rootCmd.PersistentFlags().StringVarP(&utils.CliArgs.OutputDirectory, "output", "o", "generated/", "Output directory (optional). Default is 'generated/'")
	rootCmd.PersistentFlags().StringVarP(&utils.CliArgs.ThemeName, "theme", "t", "default", "Name of the theme (optional). Default is 'default'.")
	rootCmd.PersistentFlags().StringVarP(&utils.CliArgs.Format, "format", "f", "html", "Format for the export (optional). Default is 'html'.")
	rootCmd.PersistentFlags().BoolVarP(&utils.CliArgs.Verbose, "verbose", "v", false, "Verbose mode.")
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().BoolVarP(&utils.CliArgs.Watch, "watch", "w", false, "Watch for file changes")
	serveCmd.PersistentFlags().IntVarP(&utils.CliArgs.Port, "port", "p", 3000, "Listening port")

	rootCmd.AddCommand(cmdThemes.ThemesCmd())

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}

}
