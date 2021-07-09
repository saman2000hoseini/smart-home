package http_server

import (
	"fmt"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/config"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/handler"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/model"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/router"
	"github.com/saman2000hoseini/smart-home/internal/pkg/db"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func main(cfg config.Config) {
	fmt.Println()
	myDB, err := db.FirstSetup()
	if err != nil {
		logrus.Fatalf("failed to setup db: %s", err.Error())
	}

	userRepo := model.NewSQLUserRepo(myDB)

	bathHandler := handler.NewHTTPBathHandler(userRepo)

	e := router.New(cfg)

	e.POST("/bath", bathHandler.HandleHTTPBath)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := e.Start(cfg.HTTPServer.Address); err != nil {
			logrus.Fatalf("failed to start virtual-box management server: %s", err.Error())
		}
	}()
	logrus.Info("http server started!")

	s := <-sig

	logrus.Infof("signal %s received", s)
}

// Register registers http-server command for smart-home binary.
func Register(root *cobra.Command, cfg config.Config) {
	publish := &cobra.Command{
		Use:   "http-server",
		Short: "server for http",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg)
		},
	}

	root.AddCommand(publish)
}
