package publisher

import (
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/config"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {

}

// Register registers publisher command for smart-home binary.
func Register(root *cobra.Command, cfg config.Config) {
	publish := &cobra.Command{
		Use:   "publish",
		Short: "publish commands into hive",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(publish)
}
