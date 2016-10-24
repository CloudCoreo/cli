package cmd

import (
	"fmt"
	"os"

	"github.com/cloudcoreo/cli/client"
	"github.com/cloudcoreo/cli/cmd/util"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var LSCmd = &cobra.Command{
	Use:   "ls",
	Short: "lists files",
	Long:  `reallly long lists files`,
	Run: func(cmd *cobra.Command, args []string) {
		long, err := cmd.LocalFlags().GetBool("long")
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		flags := make([]string, 0, len(args)+1)
		if long {
			flags = append(flags, "-l")
		}

		for _, v := range args {
			flags = append(flags, v)
		}

		out, err := util.Exec(context.Background(), "ls", flags...)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		fmt.Println(out)
	},
}

var TC = &cobra.Command{
	Use:   "tc",
	Short: "gets test",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New("https://demo2715234.mockable.io/")
		t, err := c.GetTest(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}

		fmt.Printf("%#v", t)
	},
}

func init() {
	RootCmd.AddCommand(LSCmd)
	RootCmd.AddCommand(TC)
	LSCmd.Flags().BoolP("long", "l", false, "list files in long format")
}
