package handler

import (
	"encoding/hex"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/model"
	amqp "github.com/saman2000hoseini/smart-home/internal/app/smart-home/mqtt"
	"github.com/sirupsen/logrus"
)

type AMQPBathHandler struct {
	userRepo  model.UserRepo
	client    mqtt.Client
	publisher *amqp.Publisher
}

func NewAMQPBathHandler(repo model.UserRepo, client mqtt.Client, publisher *amqp.Publisher) *AMQPBathHandler {
	return &AMQPBathHandler{
		userRepo:  repo,
		client:    client,
		publisher: publisher,
	}
}

func (a *AMQPBathHandler) HandleAMQPBath(_ mqtt.Client, message mqtt.Message) {
	logrus.Infof("user entered: %s", hex.EncodeToString(message.Payload()))

	user, err := a.userRepo.Find(hex.EncodeToString(message.Payload()))
	if err != nil {
		logrus.Infof("couldnt fetch user(id: %s) info from db: %s", string(message.Payload()), err.Error())
		return
	}

	a.publisher.PublishBathInfo(user)
}
