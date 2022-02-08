package streamer

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"time"
)

type Streamer interface {
	NewChannel() *Channel
	RegisterHandler(channel Channel, topicName string, handlerFunc message.NoPublishHandlerFunc)
}

type Stream struct {
	Router *WaterMillRouter
}

type WaterMillRouter struct {
	*message.Router
}

type Channel struct {
	*gochannel.GoChannel
}

func NewStreamer() *Stream {
	return &Stream{
		Router: NewRouter(),
	}
}

var logger = watermill.NewStdLogger(false, false)

func NewRouter() *WaterMillRouter {
	router, err := message.NewRouter(message.RouterConfig{
		CloseTimeout: time.Hour,
	}, logger)
	if err != nil {
		panic(err)
	}

	// SignalsHandler will gracefully shutdown Router when SIGTERM is received.
	// You can also close the router by just calling `r.Close()`.
	router.AddPlugin(plugin.SignalsHandler)

	// Router level middleware are executed for every message sent to the router
	router.AddMiddleware(
		// CorrelationID will copy the correlation id from the incoming message's metadata to the produced messages
		middleware.CorrelationID,

		// The handler function is retried if it returns an error.
		// After MaxRetries, the message is Nacked and it's up to the PubSub to resend it.
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          logger,
		}.Middleware,

		// Recoverer handles panics from handlers.
		// In this case, it passes them as errors to the Retry middleware.
		middleware.Recoverer,
	)

	return &WaterMillRouter{
		router,
	}
}

func NewChannel() Channel {
	return Channel{
		gochannel.NewGoChannel(gochannel.Config{}, logger),
	}
}

func (s *Stream) RegisterHandler(channel Channel, topicName string, handlerFunc message.NoPublishHandlerFunc) {
	fmt.Println(topicName)
	s.Router.AddNoPublisherHandler(
		topicName+"_"+time.Now().UTC().Format(time.RFC3339),
		topicName,
		channel,
		handlerFunc,
	)
}

type Message struct {
	EventName string
	Data      interface{}
}

func NewMessage(m *Message) (*message.Message, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	msg := message.NewMessage(watermill.NewUUID(), data)
	middleware.SetCorrelationID(watermill.NewUUID(), msg)
	return msg, nil
}

func PublishMessage(channel Channel, topic string, msg *message.Message) {
	if err := channel.Publish(topic, msg); err != nil {
		logger.Error("Couldn't publish the message", err, watermill.LogFields{
			"CorrelationID": msg.UUID,
		})
	}
}
