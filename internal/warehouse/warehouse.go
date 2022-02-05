package warehouse

import "strconv"

type Warehouse struct {
	ID   uint64 `json:"id" uri:"id"`
	Name string `json:"name"`
}

func NewWarehouse(id uint64, name string) *Warehouse {
	return &Warehouse{
		ID:   id,
		Name: name,
	}
}

func GetDummyWarehouses() []Warehouse {
	var warehouses []Warehouse
	for i := 0; i < 5; i++ {
		w := NewWarehouse(uint64(i+1), "Dummy Warehouse "+strconv.Itoa(i+1))
		warehouses = append(warehouses, *w)
	}
	return warehouses
}

func GetWarehouse(id uint64) *Warehouse {
	return NewWarehouse(id, "Dummy Warehouse")
}

func UpdateWarehouse(w Warehouse) *Warehouse {
	return NewWarehouse(w.ID, w.Name)
}

func CreateWarehouse(w Warehouse) *Warehouse {
	return NewWarehouse(w.ID, w.Name)
}

func DeleteWarehouse(w Warehouse) *Warehouse {
	return NewWarehouse(w.ID, w.Name)
}
