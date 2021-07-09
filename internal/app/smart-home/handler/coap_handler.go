package handler

import (
	"encoding/json"
	"github.com/dustin/go-coap"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/model"
	"github.com/sirupsen/logrus"
	"net"
	"strings"
)

type CoAPBathHandler struct {
	userRepo model.UserRepo
}

func NewCoAPBathHandler(repo model.UserRepo) *CoAPBathHandler {
	return &CoAPBathHandler{
		userRepo: repo,
	}
}

func (c *CoAPBathHandler) HandleCoAPBath(l *net.UDPConn, a *net.UDPAddr, m *coap.Message) *coap.Message {
	logrus.Infof("CoAP message ==> path=%q: %#v from %v", m.Path(), m, a)

	payload := string(m.Payload)
	payload = strings.ReplaceAll(payload, " ", "")
	payload = strings.TrimSpace(payload)

	user, err := c.userRepo.Find(payload)
	if err != nil {
		logrus.Infof("couldnt fetch user(id: %s) info from db: %s", payload, err.Error())

		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Content,
			MessageID: m.MessageID,
			Token:     m.Token,
			Payload:   []byte("couldnt fetch user"),
		}
		res.SetOption(coap.ContentFormat, coap.TextPlain)

		return res
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		logrus.Infof("couldnt marshal user(id: %v): %s", user, err.Error())

		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Content,
			MessageID: m.MessageID,
			Token:     m.Token,
			Payload:   []byte("couldnt marshal json"),
		}
		res.SetOption(coap.ContentFormat, coap.TextPlain)

		return res
	}

	if m.IsConfirmable() {
		res := &coap.Message{
			Type:      coap.Acknowledgement,
			Code:      coap.Content,
			MessageID: m.MessageID,
			Token:     m.Token,
			Payload:   jsonUser,
		}
		res.SetOption(coap.ContentFormat, coap.TextPlain)

		return res
	}

	res := &coap.Message{
		Type:      coap.NonConfirmable,
		Code:      coap.Content,
		MessageID: m.MessageID,
		Token:     m.Token,
		Payload:   jsonUser,
	}
	res.SetOption(coap.ContentFormat, coap.TextPlain)

	return res
}
