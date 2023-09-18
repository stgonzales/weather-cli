/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/stgonzales/weather-cli/helpers"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "A simple CLI for getting the weather",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		f, e := cmd.Flags().GetString("location")

		if e != nil {
			fmt.Println(e)
			os.Exit(1)
		}

		d := helpers.GetGeoLocation(f)

		fmt.Printf("Weather for %s, %s\n", d.Location.Name, d.Location.Country)
		fmt.Printf("Temperature now is %.fC feels like %.fC\n\n", d.Current.TempC, d.Current.FeelslikeC)

		for _, v := range d.Forecast.Forecastday[0].Hour {
			r, err := time.Parse("2006-01-02 15:04", v.Time)

			if err != nil {
				continue
			}

			if r.Hour() <= time.Now().Hour() {
				continue
			}

			h := fmt.Sprintf("%02d:%02d", r.Hour(), r.Minute())
			fmt.Printf("%s - %.fC - Feels like %.fC - Chances Rain %d%% - Chances snow %d%% \n", h, v.TempC, v.FeelslikeC, v.ChanceOfRain, v.ChanceOfSnow)

		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Flags().StringP("location", "l", "derby", "Location to get the weather for")
	rootCmd.MarkFlagRequired("location")
}
