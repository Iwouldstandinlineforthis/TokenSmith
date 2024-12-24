package cmd

import (
	"github.com/gladstomych/tokensmith/internal/auth"
	"github.com/gladstomych/tokensmith/internal/classes"
	"github.com/spf13/cobra"
)

var intuneBypass bool

// authCodeCmd represents the authcode command
var authCodeCmd = &cobra.Command{
	Use:   "authcode",
	Short: "(Recommended) Request Access & Refresh tokens from Authorization Code flow. Supports Intune compliant device conditional access bypass.",
	Long: `Requesting Entra ID Access & Refresh tokens using the Authorization Code flow. 
It is the recommended way to get tokens as it blends in with normal OAuth2 SSO logins. 
This option also supports the Intune Compliant Device CAP bypass.
TokenSmith can be run from an any Internet connected device, 
Where the login happens in a browser that can complete an interactive login or has an active M365 session.

Example flow:
[Attacker's Machine] ./tokensmith authcode > generate URL
   |
   V
[Machine with active browser session] use generaetd URL to authenticate > redirects with Authorization Code
   |
   V
[Attacker's Machine] Copy & Paste the redirect URL from browser back to tokensmith redeem Entra tokens
`,
	Run: func(cmd *cobra.Command, args []string) {
		// auth.InvokeAuthTokenFlow(ClientID, ResourceURL, Scope, RedirURI)
		// if len(args) == 0 && cmd.Flags().NFlag() == 0 {
		//    cmd.Help()
		//    return
		// }
		auth.InvokeAuthTokenFlow(intuneBypass)
	},
}

func init() {
	rootCmd.AddCommand(authCodeCmd)
	authCodeCmd.Flags().StringVarP(&classes.RedirURI, "rediruri", "R", "https://login.microsoftonline.com/common/oauth2/nativeclient", "OAuth 2 redirect URI for the client unless the --intune-bypass flag is on. Leave it as default if you are not sure.")
	authCodeCmd.Flags().StringVarP(&classes.ResourceURL, "resource", "r", "https://graph.microsoft.com/.default", "Resource URL to request token for.")
	authCodeCmd.Flags().BoolVarP(&intuneBypass, "intune-bypass", "i", false, "Use a special flow to bypass Intune Complaint Device & Entra Hybrid Joined conditional access.")
}
