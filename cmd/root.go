package cmd

import (
	"os"
	"text/template"

	"github.com/j-sv/readable-time/when"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() {
	rootCmd := &cobra.Command{
		Short: "Print the current time in a human readable format",
		RunE: func(cmd *cobra.Command, args []string) error {
			tmpl := template.Must(template.New("format").Parse(viper.GetString("format")))

			t, err := when.Parse(viper.GetString("when"))
			if err != nil {
				return err
			}

			if err := tmpl.Execute(os.Stdout, t); err != nil {
				return err
			}
			return nil
		},
	}

	rootCmd.PersistentFlags().StringP(
		"format",
		"f",
		`{{ .Clock }}, {{ .Weekday }}, {{ .Month }} {{ .Day }}`,
		"format to use",
	)

	rootCmd.PersistentFlags().StringP(
		"when",
		"w",
		"now",
		"what timestamp to use",
	)

	if err := viper.BindPFlag("format", rootCmd.PersistentFlags().Lookup("format")); err != nil {
		panic(err)
	} else if err := viper.BindPFlag("when", rootCmd.PersistentFlags().Lookup("when")); err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
