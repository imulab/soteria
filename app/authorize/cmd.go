package authorize

import "github.com/spf13/cobra"

func ApiCommand() *cobra.Command {
	return &cobra.Command{
		Use: "authorize",
		Short: "Serves the OAuth 2.0 / Open ID Connect 1.0 authorize API.",
		RunE: func(cmd *cobra.Command, args []string) error {
			api := &authorizeApi{}
			if err := api.setup(); err != nil {
				return err
			}
			return api.startWebServer()
		},
	}
}