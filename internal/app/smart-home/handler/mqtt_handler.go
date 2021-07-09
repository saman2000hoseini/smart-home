package handler

import (
	"encoding/hex"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/hivemq"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/model"
	"github.com/sirupsen/logrus"
)

type MQTTBathHandler struct {
	userRepo  model.UserRepo
	publisher *hivemq.Publisher
}

func NewMQTTBathHandler(repo model.UserRepo, publisher *hivemq.Publisher) *MQTTBathHandler {
	return &MQTTBathHandler{
		userRepo:  repo,
		publisher: publisher,
	}
}

func (b *MQTTBathHandler) HandleMQTTBath(_ mqtt.Client, message mqtt.Message) {
	logrus.Infof("user entered: %s", hex.EncodeToString(message.Payload()))

	user, err := b.userRepo.Find(hex.EncodeToString(message.Payload()))
	if err != nil {
		logrus.Infof("couldnt fetch user(id: %s) info from db: %s", string(message.Payload()), err.Error())
		return
	}

	b.publisher.PublishBathInfo(user)
}
