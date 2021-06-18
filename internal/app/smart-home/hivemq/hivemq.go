package hivemq

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/config"
)

func createHiveMQConnection(cfg config.HiveMQ) mqtt.Client {
	server := fmt.Sprintf("%s://%s:%d", cfg.Connection, cfg.Address, cfg.Port)
	opts := mqtt.NewClientOptions().AddBroker(server).SetClientID(cfg.Client)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return c
}
