/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"raytracer/internal/core"
	"raytracer/internal/utils"

	"github.com/spf13/cobra"
)

// demoCmd represents the demo command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Generate a demo image using the raytracer",
	Long: `A quick way to generate a demo image using the raytracer.
	A PNG file will be generated in the root directory of the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Log.Info("Initiating render")

		cwd, err := os.Getwd()

		utils.PanicOnError(err)

		raytracer := core.InitRaytracerFromScene(cwd + "/scenes/demo.json")

		raytracer.PrintDetails()

		img := raytracer.Render()

		utils.ExportImage(img, "output.png")
	},
}

func init() {
	rootCmd.AddCommand(demoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// demoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// demoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
