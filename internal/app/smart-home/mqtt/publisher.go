package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/model"
	"github.com/sirupsen/logrus"
)

type Publisher struct {
	c mqtt.Client
}

func NewPublisher(c mqtt.Client) *Publisher {
	return &Publisher{c: c}
}

func (p *Publisher) PublishBathInfo(user model.User) {
	if token := p.c.Publish("smart-home/bath/temp", 1, false, string(user.WaterTemp)); token.Wait() && token.Error() != nil {
		logrus.Info(token.Error())
	}

	if token := p.c.Publish("smart-home/bath/level", 1, false, string(user.WaterLevel)); token.Wait() && token.Error() != nil {
		logrus.Info(token.Error())
	}
}

func (p *Publisher) Publish(topic, msg string) {
	if token := p.c.Publish(topic, 1, false, msg); token.Wait() && token.Error() != nil {
		logrus.Info(token.Error())
	}
}
