package amqp_subscriber

import (
	"fmt"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/config"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/handler"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/model"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/mqtt"
	"github.com/saman2000hoseini/smart-home/internal/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	fmt.Println()
	myDB, err := db.FirstSetup()
	if err != nil {
		logrus.Fatalf("failed to setup db: %s", err.Error())
	}

	client := mqtt.CreateMQTTConnection(cfg.MQTT)
	userRepo := model.NewSQLUserRepo(myDB)

	bathHandler := handler.NewAMQPBathHandler(userRepo, client, mqtt.NewPublisher(client))
	subscriber := mqtt.NewSubscriber(client)

	subscriber.Run(mqtt.RFID_TOPIC, bathHandler.HandleAMQPBath)
}

// Register registers coap-server command for smart-home binary.
func Register(root *cobra.Command, cfg config.Config) {
	publish := &cobra.Command{
		Use:   "amqp-subscribe",
		Short: "run amqp subscriber",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(publish)
}
