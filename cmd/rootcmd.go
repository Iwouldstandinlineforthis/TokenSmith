package cmd

import (
	"fmt"
	"os"

	"github.com/gladstomych/tokensmith/internal/classes"
	"github.com/gladstomych/tokensmith/internal/utils"
	"github.com/spf13/cobra"
)

var (
	noBanner    bool
	showVersion bool
	version     = "v0.8"
)

var rootCmd = &cobra.Command{
	Use:   "tokensmith",
	Short: "Tokensmith is an Entra ID token retrieval & redemption tool designed to work with conditional access restrictions.",
	Long: `Tokensmith is an Entra ID token retrieval & redemption tool. Tokens retrieved from Tokensmith are compatible with the majority of popular Entra/Azure offensive tooling out of the box.

Example:
# without intune compliant bypass:
./tokensmith authcode

# with intune compliant device bypass:
./tokensmith authcode -i`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Println(version)
			os.Exit(0) // Exit after printing version
		}

		if !noBanner {
			utils.PrintBanner()
		}

	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 && cmd.Flags().NFlag() == 0 {
			err := cmd.Help()
			if err != nil {
				os.Exit(1)
			}
			return
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
	rootCmd.PersistentFlags().BoolVar(&noBanner, "no-banner", false, "Do not display the banner")
	rootCmd.PersistentFlags().StringVarP(&classes.UserAgent, "useragent", "U", "User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36", "User-Agent string sent in HTTP requests.")
	rootCmd.PersistentFlags().StringVarP(&classes.ClientID, "clientID", "c", "1fec8e78-bce4-4aaf-ab1b-5451cc387264", "client ID to request token for. Default - MS Teams. Reasons for customising this might be to enumerate for an MFA gap, or getting different token scopes. To redeem tokens from a refresh token obtained the intune bypass, use -c 9ba1a5c7-f17a-4de9-a1f1-6178c8d51223.")
	rootCmd.PersistentFlags().StringVarP(&classes.Scope, "scope", "s", "openid offline_access", "scope to request token for.")
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Print version and exit")
}
