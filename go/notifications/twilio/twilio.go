package TwilioNotification

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	twilio "github.com/carlosdp/twiliogo"
	"github.com/pandemicsyn/stalker/go/stalker"
)

type TwilioNotification struct {
	AccountSid string
	AuthToken  string
	FromNumber string
	Recipients []string
	tc         *twilio.TwilioClient
}

func New(sid, token, from string, recipients []string) *TwilioNotification {
	return &TwilioNotification{sid, token, from, recipients, twilio.NewClient(sid, token)}
}

func (tn *TwilioNotification) Fail(check stalker.StalkerCheck) {
	for _, v := range tn.Recipients {
		message, err := twilio.NewMessage(tn.tc, tn.FromNumber, v, twilio.Body(fmt.Sprintf("%s on %s is down", check.Check, check.Hostname)))
		if err != nil {
			log.Printf("Error sending notification to %s via twilio: %s\n", v, err.Error())
		} else {
			log.Printf("Sent twilio notification: %+v\n", message)
		}
	}
}

func (tn *TwilioNotification) Clear(check stalker.StalkerCheck) {
	for _, v := range tn.Recipients {
		message, err := twilio.NewMessage(tn.tc, tn.FromNumber, v, twilio.Body(fmt.Sprintf("%s on %s is down", check.Check, check.Hostname)))
		if err != nil {
			log.Printf("Error sending notification to %s via twilio: %s\n", v, err.Error())
		} else {
			log.Printf("Sent twilio notification: %+v\n", message)
		}
	}
}