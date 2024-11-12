package redis

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/studiobflat/tsj/publisher"
)

type WatermillPublisher struct {
	pub Publisher
}

func NewWatermillPublisher(pub Publisher) (publisher.Publisher, error) {
	return &WatermillPublisher{
		pub: pub,
	}, nil
}

func (w *WatermillPublisher) Close() error {
	return w.pub.Close()
}

func (w *WatermillPublisher) PublishMessage(messages ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()

	msgs := make([]*message.Message, 0)

	for _, v := range messages {
		msgs = append(msgs, message.NewMessage(watermill.NewUUID(), []byte(v)))
	}

	if err := w.pub.Client().XTrimMaxLen(ctx, w.pub.Topic(), w.pub.MaxStreamEntries()).Err(); err != nil {
		return err
	}

	if err := w.pub.Publish(w.pub.Topic(), msgs...); err != nil {
		return err
	}

	return nil
}

func (w *WatermillPublisher) Topic() string {
	return w.pub.Topic()
}
