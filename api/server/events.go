package server

import (
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/unicod3/horreum/internal/order"
	"github.com/unicod3/horreum/pkg/streamer"
)

// RegisterEventHandlers registers the package's events handlers to streamer package
func (h *Handler) RegisterEventHandlers(s *streamer.Stream) {
	s.RegisterHandler(
		h.OrderService.StreamChannel,
		h.OrderService.StreamTopic,
		h.HandleOrderEvents,
	)
}

func (h *Handler) HandleOrderEvents(msg *message.Message) error {
	fmt.Printf(
		"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	var message streamer.Message
	err := json.Unmarshal(msg.Payload, &message)
	if err != nil {
		return err
	}

	var o order.Order
	byteOrder, _ := json.Marshal(message.Data)
	err = json.Unmarshal(byteOrder, &o)
	if err != nil {
		return err
	}

	fmt.Println("EVENT: ", message.EventName)
	switch message.EventName {
	case order.OrderCreated:
		for _, line := range o.Lines {
			product, _ := h.ProductService.GetById(line.ProductID)
			err := product.DecreaseStockBy(h.ArticleService, int64(line.Quantity))
			if err != nil {
				return err
			}
		}
	case order.OrderDeleted:
		for _, line := range o.Lines {
			product, _ := h.ProductService.GetById(line.ProductID)
			err := product.IncreaseStockBy(h.ArticleService, int64(line.Quantity))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
