package cmd

import (
	"os"
	"text/template"

	"github.com/j-sv/readable-time/time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() {
	rootCmd := &cobra.Command{
		Short: "Print the current time in a human readable format",
		RunE: func(cmd *cobra.Command, args []string) error {
			tmpl := template.Must(template.New("format").Parse(viper.GetString("format")))

			now := time.Now()
			if err := tmpl.Execute(os.Stdout, now); err != nil {
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

	if err := viper.BindPFlag("format", rootCmd.PersistentFlags().Lookup("format")); err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
