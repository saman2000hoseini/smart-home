package mqtt_subscriber

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
	fmt.Println("                           _        _                           \n  ___ _ __ ___   __ _ _ __| |_     | |__   ___  _ __ ___   ___  \n / __| '_ ` _ \\ / _` | '__| __|____| '_ \\ / _ \\| '_ ` _ \\ / _ \\ \n \\__ \\ | | | | | (_| | |  | ||_____| | | | (_) | | | | | |  __/ \n |___/_|_|_| |_|\\__,_|_|   \\__|    |_| |_|\\___/|_| |_| |_|\\___| \n | |__ (_)_   _____ _ __ ___   __ _                             \n | '_ \\| \\ \\ / / _ \\ '_ ` _ \\ / _` |                            \n | | | | |\\ V /  __/ | | | | | (_| |                            \n |_| |_|_| \\_/ \\___|_| |_| |_|\\__, |                            \n            _                   _|_|                            \n  ___ _   _| |__  ___  ___ _ __(_) |__   ___ _ __               \n / __| | | | '_ \\/ __|/ __| '__| | '_ \\ / _ \\ '__|              \n \\__ \\ |_| | |_) \\__ \\ (__| |  | | |_) |  __/ |                 \n |___/\\__,_|_.__/|___/\\___|_|  |_|_.__/ \\___|_|                 \n                                                                ")
	myDB, err := db.FirstSetup()
	if err != nil {
		logrus.Fatalf("failed to setup db: %s", err.Error())
	}

	client := mqtt.CreateMQTTConnection(cfg.MQTT)
	userRepo := model.NewSQLUserRepo(myDB)

	bathHandler := handler.NewMQTTBathHandler(userRepo, mqtt.NewPublisher(client))
	subscriber := mqtt.NewSubscriber(client)

	subscriber.Run(mqtt.RFID_TOPIC, bathHandler.HandleMQTTBath)
}

// Register registers mqtt-subscriber command for smart-home binary.
func Register(root *cobra.Command, cfg config.Config) {
	publish := &cobra.Command{
		Use:   "mqtt-subscribe",
		Short: "subscribes messages from hive",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(publish)
}
