package cmd

import (
	"fmt"
	"os"

	"github.com/cloudcoreo/cli/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"github.com/cloudcoreo/cli/cmd/content"
	"github.com/cloudcoreo/cli/cmd/util"
)


var meCmd = &cobra.Command{
	Use:   "me",
	Short: "gets me",
	Run: func(cmd *cobra.Command, args []string) {

		//generate config keys based on user profile
		apiKey := util.GetValueFromConfig(fmt.Sprintf("%s.%s", userProfile, content.ACCESS_KEY), false)
		secretKey := util.GetValueFromConfig(fmt.Sprintf("%s.%s", userProfile, content.SECRET_KEY), false)

		if apiKey == content.NONE || secretKey == content.NONE {
			fmt.Println(content.ERROR_MISSING_API_KEY_SECRET_KEY)
			return
		}

		a := client.Auth{APIKey: apiKey, SecretKey: secretKey}
		i := client.Interceptor(a.SignRequest)
		c := client.New(content.ENDPOINT_ADDRESS, client.WithInterceptor(i))
		t, err := c.GetTokens(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		fmt.Printf("%#v", t)
	},
}

func init() {
	RootCmd.AddCommand(meCmd)
}