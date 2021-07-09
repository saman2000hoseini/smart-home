package cmd

import (
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/cmd/coap-server"
	http_server "github.com/saman2000hoseini/smart-home/internal/app/smart-home/cmd/http-server"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/cmd/mqtt-subscriber"
	"github.com/spf13/cobra"

	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/config"
)

// NewRootCommand creates a new smart-home root command.
func NewRootCommand() *cobra.Command {
	var root = &cobra.Command{
		Use: "smart-home",
	}

	cfg := config.Init()

	mqtt_subscriber.Register(root, cfg)
	coap_server.Register(root, cfg)
	http_server.Register(root, cfg)

	return root
}
