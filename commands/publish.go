package commands

import (
	"net/url"

	"github.com/spf13/cobra"
	"github.com/verminio/hugo/provider"
)

var publishUrl string

var publish = &cobra.Command{
	Use:   "publish",
	Short: "Publish site to server",
	Long:  "Hugo will publish the site to the address specified",
	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig()
		pub(cmd, args)
	},
}

func init() {
	publish.Flags().StringVar(&publishUrl, "url", "", "URL for the address where site should be published")
}

func pub(cmd *cobra.Command, args []string) {
	target, err := url.Parse(publishUrl)
	if err != nil {
		panic("Unable to parse URL: " + err.Error())
	}
	prov := provider.GetProvider(target.Scheme)
	prov.Publish(target)
}
