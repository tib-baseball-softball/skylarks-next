package dp

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/SherClockHolmes/webpush-go"
)

const (
	VAPIDSubscriberEnvName = "VAPID_SUBSCRIBER"
	VAPIDPublicKeyEnvName  = "VAPID_PUBLIC_KEY"
	VAPIDPrivateKeyEnvName = "VAPID_PRIVATE_KEY"
)

// PushMessage represents the raw data for a push notification message that is encoded into the JSON payload.
type PushMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Tag   string `json:"tag"`
}

// PushService defines the app-wide interface for sending push notifications.
type PushService interface {
	handleTestPush(sub *webpush.Subscription) error
	SendPushMessage(msg *PushMessage, sub *webpush.Subscription) error
}

// PushServiceImpl is the default implementation of PushService. It is meant as a singleton, to be instantiated only once.
// The main advantage is that it encapsulates the VAPID credentials and handles the actual sending of push notifications.
type PushServiceImpl struct {
	Subscriber      string
	VAPIDPublicKey  string
	VAPIDPrivateKey string
}

// NewPushService creates a new PushServiceImpl instance, making sure the VAPID credentials are present in the environment.
func NewPushService() PushService {
	s := PushServiceImpl{
		Subscriber:      os.Getenv(VAPIDSubscriberEnvName),
		VAPIDPublicKey:  os.Getenv(VAPIDPublicKeyEnvName),
		VAPIDPrivateKey: os.Getenv(VAPIDPrivateKeyEnvName),
	}

	if s.Subscriber == "" || s.VAPIDPublicKey == "" || s.VAPIDPrivateKey == "" {
		log.Fatal("Missing VAPID environment variables, exiting")
	}
	return &s
}

// SendPushMessage sends a push notification to the given subscription.
func (p PushServiceImpl) SendPushMessage(msg *PushMessage, sub *webpush.Subscription) error {
	blob, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		resp, err := webpush.SendNotificationWithContext(ctx, blob, sub, &webpush.Options{
			Subscriber:      p.Subscriber,
			VAPIDPublicKey:  p.VAPIDPublicKey,
			VAPIDPrivateKey: p.VAPIDPrivateKey,
			TTL:             30,
		})
		if err != nil {
			// TODO: handle unsubscribe for outdated subs here
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
	}()

	return nil
}

// handleTestPush sends a test push notification to the given subscription.
func (p PushServiceImpl) handleTestPush(sub *webpush.Subscription) error {
	msg := PushMessage{
		Title: "Hello there",
		Body:  "Real men test in production",
		Tag:   "Test Tag",
	}

	return p.SendPushMessage(&msg, sub)
}
