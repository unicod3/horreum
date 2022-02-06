package warehouse

import "github.com/unicod3/horreum/pkg/dbclient"

// Handler holds db struct
type Handler struct {
	WarehouseService *Service
}

// NewHandler returns a new Handler
func NewHandler(client *dbclient.Client) *Handler {
	return &Handler{
		WarehouseService: &Service{
			client.NewDataStore("warehouses"),
		},
	}
}
