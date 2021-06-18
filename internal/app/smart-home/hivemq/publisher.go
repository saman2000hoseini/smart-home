package hivemq

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/config"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/model"
	"github.com/sirupsen/logrus"
)

type Publisher struct {
	c mqtt.Client
}

func NewPublisher(cfg config.HiveMQ) *Publisher {
	return &Publisher{c: createHiveMQConnection(cfg)}
}

func (p *Publisher) Publish(user model.User) {
	payload := fmt.Sprintf("%d,%d", user.WaterTemp, user.WaterLevel)

	if token := p.c.Publish("smart-home/bath", 1, false, payload); token.Wait() && token.Error() != nil {
		logrus.Info(token.Error())
	}
}
