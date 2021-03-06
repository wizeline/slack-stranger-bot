// Copyright (c) 2018 Alex Pliutau

package bot

import (
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

// APISlack struct
type APISlack struct {
	api *slack.Client
}

// IAPI interface: slack or mock
type IAPI interface {
	newRTM() *slack.RTM
	getUsers() ([]slack.User, error)
	postMsg(string, string) error
}

// NewAPISlack contructor
func NewAPISlack(token string) *APISlack {
	a := new(APISlack)
	a.api = slack.New(token)
	return a
}

func (a *APISlack) newRTM() *slack.RTM {
	return a.api.NewRTM()
}

func (a *APISlack) getUsers() ([]slack.User, error) {
	return a.api.GetUsers()
}

func (a *APISlack) postMsg(channel, text string) error {
	_, _, msgErr := a.api.PostMessage(channel, text, slack.PostMessageParameters{
		AsUser: true,
	})
	if msgErr != nil {
		log.Info("[postMessage] " + msgErr.Error())
	}

	return msgErr
}
