package cmd

import (
	"github.com/gladstomych/tokensmith/internal/auth"
	"github.com/spf13/cobra"
)

// devicecodeCmd represents the devicecode command
var devicecodeCmd = &cobra.Command{
	Use:   "devicecode",
	Short: "Request Access & Refresh Tokens using the device code flow.",
	Long: `(Not yet implemented) Using the device code flow to authenticate to Entra ID and get tokens for desired resources and scopes.
Often only used by offensive tooling and hence we consider it opsec unsafe. 
`,
	Run: func(cmd *cobra.Command, args []string) {
         // if len(args) == 0 && cmd.Flags().NFlag() == 0 {
         //    cmd.Help()
         //    return       
        // }
        auth.GetTknFromDevCode()
	},
}

func init() {
	rootCmd.AddCommand(devicecodeCmd)
}
