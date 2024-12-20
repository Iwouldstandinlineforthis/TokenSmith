package cmd

import (
	"github.com/gladstomych/tokensmith/internal/auth"
	"github.com/gladstomych/tokensmith/internal/classes"
	"github.com/spf13/cobra"
)

// reftokenCmd represents the reftoken command
var refTokenCmd = &cobra.Command{
	Use:   "reftoken",
	Short: "Request Access & Refresh tokens from Entra ID Refresh token.",
	Long: `Request Access & Refresh tokens from Entra ID Refresh token.
You can specify an identical, or a different resource or client from initial scope.
This would work if the initial token was from a 'Foci' client. E.g -

Initial tokens scoped to MS Graph, and client is Az PowerShell.
    |
    V
It is possible to request new token for AD Graph. 

For more info on Foci clients refer to: https://github.com/secureworks/family-of-client-ids-research
`,
	Run: func(cmd *cobra.Command, args []string) {
         if len(args) == 0 && cmd.Flags().NFlag() == 0 {
            cmd.Help()
            return       
        }

        auth.GetAccTknFromRefTkn()
	},
}

func init() {
	rootCmd.AddCommand(refTokenCmd)
    refTokenCmd.Flags().StringVarP(&classes.RefResourceURL, "resource", "r", "https://graph.microsoft.com/", "Resource URL to request token for. When renewing or exchanging tokens, the '.default' part is usually dropped for the resource URL.")
    refTokenCmd.Flags().StringVarP(&classes.RefreshToken, "reftoken", "R", "", "The refresh token to use (required).")

}
