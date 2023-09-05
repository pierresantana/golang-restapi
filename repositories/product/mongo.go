package product

import (
	"context"

	"github.com/pierresantana/golang-restapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName = "products"
)

type MongoProductRepository struct {
	collection *mongo.Collection
}

func NewMongoProductRepository(db *mongo.Database) *MongoProductRepository {
	return &MongoProductRepository{
		collection: db.Collection(collectionName),
	}
}

func (m *MongoProductRepository) GetAll(ctx context.Context) ([]*models.Product, error) {
	var product []*models.Product
	cur, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &product); err != nil {
		return nil, err
	}
	return product, nil
}

func (m *MongoProductRepository) GetByID(ctx context.Context, id string) (*models.Product, error) {
	var product *models.Product
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return product, err
	}
	err = m.collection.FindOne(ctx, bson.M{"_id": bson.M{"$eq": oid}}).Decode(&product)
	return product, err
}

func (m *MongoProductRepository) Insert(ctx context.Context, product *models.Product) (*models.Product, error) {
	r, err := m.collection.InsertOne(ctx, &product)
	if err != nil {
		return nil, err
	}
	product.ID = r.InsertedID.(primitive.ObjectID)
	return product, err
}

func (m *MongoProductRepository) Update(ctx context.Context, id string, product *models.Product) (*models.Product, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	product.ID = oid
	_, err = m.collection.
		UpdateByID(
			ctx,
			oid,
			bson.M{
				"$set": &product,
			},
		)
	return product, err
}

func (m *MongoProductRepository) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = m.collection.DeleteOne(ctx, bson.M{"_id": bson.M{"$eq": oid}})
	return err
}
