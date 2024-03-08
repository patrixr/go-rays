/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path"
	"raytracer/internal/core"
	"raytracer/internal/utils"

	"github.com/spf13/cobra"
)

// renderCmd represents the render command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Generate an image using the raytracer",
	Long: `
Basic way of using the raytracer. Specify a scene file and the raytracer will generate an image.
If no scene file is specified, the raytracer will use the default scene file.
A PNG file will be generated in the root directory of the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.Log.Info("Initiating render")

		cwd, err := os.Getwd()

		utils.Boom(err)

		sceneFile, err := cmd.Flags().GetString("scene")

		utils.Boom(err)

		raytracer := core.InitRaytracerFromScene(path.Join(cwd, sceneFile))

		raytracer.PrintDetails()

		img := raytracer.Render()

		outputFile, err := cmd.Flags().GetString("output")

		utils.Boom(err)

		utils.ExportImage(img, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)
	renderCmd.Flags().StringP("scene", "s", "scenes/demo.json", "Path to the scene file")
	renderCmd.Flags().StringP("output", "o", "output.png", "Path to the output file")
}
