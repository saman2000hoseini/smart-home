package coap_server

import (
	"fmt"
	"github.com/dustin/go-coap"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/config"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/handler"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/model"
	"github.com/saman2000hoseini/smart-home/internal/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func main(_ config.Config) {
	fmt.Println()
	myDB, err := db.FirstSetup()
	if err != nil {
		logrus.Fatalf("failed to setup db: %s", err.Error())
	}

	userRepo := model.NewSQLUserRepo(myDB)

	bathHandler := handler.NewCoAPBathHandler(userRepo)

	api := coap.NewServeMux()
	api.Handle("/bath", coap.FuncHandler(bathHandler.HandleCoAPBath))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	logrus.Info("coap server started!")
	go coap.ListenAndServe("udp", ":5683", api)

	s := <-sig

	logrus.Infof("signal %s received", s)
}

// Register registers coap-server command for smart-home binary.
func Register(root *cobra.Command, cfg config.Config) {
	publish := &cobra.Command{
		Use:   "coap-server",
		Short: "server for coap",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(publish)
}
