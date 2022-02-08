package order

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
)

func (service *OrderService) printMessages(msg *message.Message) error {
	fmt.Printf(
		"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)

	return nil
}
