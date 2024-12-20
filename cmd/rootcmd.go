package cmd

import (
	"os"
    "fmt"

	"github.com/spf13/cobra"
	"github.com/gladstomych/tokensmith/internal/classes"
	"github.com/gladstomych/tokensmith/internal/utils"

)

var (
    noBanner bool
    showVersion bool
    version = "v0.8"
)

var rootCmd = &cobra.Command{
	Use:   "tokensmith",
	Short: "Tokensmith is an Entra ID token retrieval & redemption tool designed to work with conditional access restrictions.",
	Long: `Tokensmith is an Entra ID token retrieval & redemption tool. Current popular tooling often defaults to deviceCode auth which is arguably opsec unsafe, and also require running the authentication tool on a beachhead endpoint. Tokensmith is designed around authorization code flow and decouples browser authentication from the tool so that the consultant can redeem tokens in the comfort of their testing devices. The refresh and access tokens retrieved from Tokensmith are compatible with the majority of popular Entra/Azure offensive tooling, such as Roadrecon, GraphRunner, out of the box.

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

        if !noBanner{
            utils.PrintBanner()
        }

    },
	Run: func(cmd *cobra.Command, args []string) {
         if len(args) == 0 && cmd.Flags().NFlag() == 0 {
            cmd.Help()
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
	rootCmd.PersistentFlags().StringVarP(&classes.UserAgent, "useragent",  "U",   "User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:131.0) Gecko/20100101 Firefox/131.0", "User-Agent string sent in HTTP requests.")
	rootCmd.PersistentFlags().StringVarP(&classes.ClientID, "clientID",  "c", "1fec8e78-bce4-4aaf-ab1b-5451cc387264", "client ID to request token for. Default - MS Teams. Reasons for customising this might be to enumerate for an MFA gap, or getting different token scopes. To redeem tokens from a refresh token obtained the intune bypass, use -c 9ba1a5c7-f17a-4de9-a1f1-6178c8d51223.")
    rootCmd.PersistentFlags().StringVarP(&classes.Scope, "scope", "s", "openid offline_access", "scope to request token for.")
    rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Print version and exit")
}
