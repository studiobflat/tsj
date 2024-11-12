package redis

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"
	"github.com/studiobflat/tsj/logger"
)

type Publisher interface {
	Publish(topic string, msgs ...*message.Message) error
	Close() error
	Topic() string
	Client() redis.UniversalClient
	MaxStreamEntries() int64
}

type streamPublisher struct {
	*redisstream.Publisher
	topic            string
	client           redis.UniversalClient
	maxStreamEntries int64
}

func NewPublisher(config *PubConfig, topic string) (Publisher, error) {
	log := logger.GetLogger("NewPublisher")
	defer log.Sync()

	c, err := NewConfig()
	if err != nil {
		return nil, err
	}

	log.Infow("loaded redis publisher config")

	r, err := NewRedis(c)
	if err != nil {
		return nil, err
	}
	log.Infow("redis publisher connected")

	publisher, err := redisstream.NewPublisher(
		redisstream.PublisherConfig{
			Client:     r,
			Marshaller: redisstream.DefaultMarshallerUnmarshaller{},
		},
		watermill.NewStdLogger(config.LoggerDebug, config.LoggerTrace),
	)

	if err != nil {
		return nil, err
	}

	return &streamPublisher{
		Publisher:        publisher,
		topic:            topic,
		client:           r,
		maxStreamEntries: config.MaxStreamEntries,
	}, nil
}

func (r *streamPublisher) Topic() string {
	return r.topic
}

func (r *streamPublisher) Client() redis.UniversalClient {
	return r.client
}

func (r *streamPublisher) MaxStreamEntries() int64 {
	return r.maxStreamEntries
}
