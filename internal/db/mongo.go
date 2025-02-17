package db

import (
	"context"
	"fmt"
	"github.com/Andrew-UA/product-list/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnector struct {
	cfg    *config.Config
	client *mongo.Client
}

// NewMongoConnector створює новий конектор на основі конфігурації
func NewMongoConnector(cfg *config.Config) *MongoConnector {
	return &MongoConnector{cfg: cfg}
}

// Connect створює з'єднання з MongoDB
func (m *MongoConnector) Connect() (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", m.cfg.DbUser, m.cfg.DbPassword, m.cfg.DbHost, m.cfg.DbPort)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("MongoDB is not responding: %w", err)
	}
	return client, nil
}

func (m *MongoConnector) Close(ctx context.Context) error {
	if m.client != nil {
		return m.client.Disconnect(ctx)
	}
	return nil
}
