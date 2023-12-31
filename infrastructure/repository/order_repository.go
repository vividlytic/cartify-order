package repository

import (
	"cartify/order/domain/model"
	"cartify/order/domain/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "orders"
	COLLECTION = "orders"
)

type OrderRepository struct {
	client *mongo.Client
}

func NewOrderRepository(client *mongo.Client) repository.OrderRepository {
	return &OrderRepository{
		client: client,
	}
}

func (o *OrderRepository) ListOrders(ctx context.Context, customerId string) ([]*model.Order, error) {
	coll := o.client.Database(DATABASE).Collection(COLLECTION)
	filter := bson.D{{Key: "customerId", Value: customerId}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	var results []*model.Order
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}

func (o *OrderRepository) GetOrder(ctx context.Context, orderId string) (*model.Order, error) {
	coll := o.client.Database(DATABASE).Collection(COLLECTION)

	objectId, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: objectId}}

	var result *model.Order
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return nil, err
		}
		panic(err)
	}

	return result, nil
}

func (o *OrderRepository) CreateOrder(ctx context.Context, customerId string, customerName string, orderItem []model.OrderItem) (string, error) {
	coll := o.client.Database(DATABASE).Collection(COLLECTION)

	newOrder := &model.Order{
		ID:           primitive.NewObjectID(),
		CustomerId:   customerId,
		CustomerName: customerName,
		OrderItem:    orderItem,
	}

	result, err := coll.InsertOne(context.TODO(), newOrder)
	if err != nil {
		panic(err)
	}

	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		panic(err)
	}

	return objectId.Hex(), nil
}
