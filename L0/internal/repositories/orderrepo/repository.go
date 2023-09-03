package orderrepo

import (
	"context"
	"fmt"

	"github.com/lvmnpkfvmk/L0/config"
	"github.com/lvmnpkfvmk/L0/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	ErrNotFound       = fmt.Errorf("Order not found")
)
type OrderRepository struct {
	db *gorm.DB
	cache map[string]model.Order
}

func NewOrderRepository(ctx context.Context, cfg *config.Config) (*OrderRepository, error) {
	db, err := gorm.Open(postgres.Open(cfg.PgURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error opening gorm: %v", err)
	}

	err = db.AutoMigrate(&model.Order{})
	if err != nil {
		return nil, fmt.Errorf("Error AutoMigrate: %v", err)
	}

	var orders []model.Order
	result := db.Find(&orders)
	if result.Error != nil {
		return nil, fmt.Errorf("Error retrieving orders: %v", err)
	}
	cache := make(map[string]model.Order, result.RowsAffected)
	for _, order := range orders {
		cache[order.OrderUID] = order
	}

	return &OrderRepository{db, cache}, nil
}

func (or *OrderRepository) GetOrder(orderUID string) (*model.Order, error) {
	order, ok := or.cache[orderUID];
	if !ok {
		return nil, ErrNotFound
	}
	return &order, nil
}

func (or *OrderRepository) CreateOrder(order *model.Order) error {
	result := or.db.Create(*order)
	if result.Error != nil {
		return fmt.Errorf("Error creating order: %v", result.Error)
	}
	or.cache[order.OrderUID] = *order;
	return nil
}