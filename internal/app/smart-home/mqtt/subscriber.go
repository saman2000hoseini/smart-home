package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

const RFID_TOPIC = "smart-home/id"

type Subscriber struct {
	c mqtt.Client
}

func NewSubscriber(c mqtt.Client) *Subscriber {
	return &Subscriber{
		c: c,
	}
}

func (s *Subscriber) Run(topic string, callback mqtt.MessageHandler) {
	done := make(chan bool)

	if token := s.c.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		logrus.Info(token.Error())
	}
	<-done
}
