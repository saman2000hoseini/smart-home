package cmd

import (
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/cmd/publisher"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/cmd/subscriber"
	"github.com/spf13/cobra"

	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/config"
)

// NewRootCommand creates a new smart-home root command.
func NewRootCommand() *cobra.Command {
	var root = &cobra.Command{
		Use: "smart-home",
	}

	cfg := config.Init()

	publisher.Register(root, cfg)
	subscriber.Register(root, cfg)

	return root
}
