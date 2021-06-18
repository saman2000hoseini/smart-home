package hivemq

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/config"
	"github.com/sirupsen/logrus"
)

const RFID_TOPIC = "smart-home/id"

func RunSubscriber(cfg config.HiveMQ, topic string, callback mqtt.MessageHandler) {
	c := createHiveMQConnection(cfg)
	if token := c.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		logrus.Info(token.Error())
	}
}
