package subscriber

import (
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/config"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/handler"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/hivemq"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/model"
	"github.com/saman2000hoseini/smart-home/internal/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	myDB, err := db.FirstSetup()
	if err != nil {
		logrus.Fatalf("failed to setup db: %s", err.Error())
	}

	userRepo := model.NewSQLUserRepo(myDB)
	publisher := hivemq.NewPublisher(cfg.HiveMQ)

	bathHandler := handler.NewBathHandler(userRepo, publisher)
	hivemq.RunSubscriber(cfg.HiveMQ, hivemq.RFID_TOPIC, bathHandler.HandleBath)
}

// Register registers subscriber command for smart-home binary.
func Register(root *cobra.Command, cfg config.Config) {
	publish := &cobra.Command{
		Use:   "subscribe",
		Short: "subscribes messages from hive",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(publish)
}
