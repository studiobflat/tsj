package redis

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/studiobflat/tsj/event"
	"github.com/studiobflat/tsj/logger"
	"github.com/studiobflat/tsj/subscriber"
)

type WatermillSubscriber struct {
	done        chan bool // to end infinite loop
	sub         Subscriber
	consumeFunc subscriber.ConsumeFunc
}

func NewWatermillSubscriber(sub Subscriber, consumeFunc subscriber.ConsumeFunc) subscriber.Subscriber {
	return &WatermillSubscriber{
		consumeFunc: consumeFunc,
		sub:         sub,
		done:        make(chan bool, 1),
	}
}

func (r *WatermillSubscriber) Close() error {
	r.done <- true
	defer close(r.done)

	return r.sub.Close()
}

func (r *WatermillSubscriber) GroupID() string {
	return r.sub.GroupID()
}

func (r *WatermillSubscriber) Start() {
	log := logger.GetLogger("Redis Subcriber")
	defer log.Sync()

	log.With("topic", r.sub.Topic())

	log.Infow("subscription start...")
	messages, err := r.sub.Subscribe(context.Background(), r.sub.Topic())
	if err != nil {
		log.Panicw("subscribe error, panic now", "error", err)
		r.Close()
	}

	for {
		select {
		case <-r.done:
			log.Info("subscription ended")
			return
		default:
			// continue below to fetch message, etc...
		}

		message := <-messages

		if message == nil || message.UUID == "" {
			log.Debugw("empty message id", "topic", r.sub.Topic())
			continue
		}

		if err := r.consumeMessage(context.Background(), message); err != nil {
			log.Errorw("subscription error, fail to consume the message", "error", err, "topic", r.sub.Topic())
			continue
		}
	}
}

func (r *WatermillSubscriber) Topic() string {
	return r.sub.Topic()
}

func (r *WatermillSubscriber) consumeMessage(ctx context.Context, msg *message.Message) error {
	log := logger.GetLogger("consumeMessage")
	defer log.Sync()

	err := r.consumeFunc(ctx, event.EventString(string(msg.Payload)))
	if !msg.Ack() {
		log.Infow("message is already ack", "message", msg.UUID)
	}
	return err
}
