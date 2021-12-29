package product

import (
	"context"
	"go-fiber-rest/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(product *entities.Product) (*entities.Product, error)
	Read() (*[]entities.Product, error)
	Update(product *entities.Product) (*entities.Product, error)
	Delete(ID string) error
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r repository) Create(product *entities.Product) (*entities.Product, error) {
	product.ID = primitive.NewObjectID()
	_, err := r.Collection.InsertOne(context.Background(), product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r repository) Read() (*[]entities.Product, error) {
	var products []entities.Product
	cursor, err := r.Collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var product entities.Product
		_ = cursor.Decode(&product)
		products = append(products, product)
	}

	return &products, nil
}

func (r repository) Update(product *entities.Product) (*entities.Product, error) {
	_, err := r.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": product.ID},
		bson.M{"$set": product})

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r repository) Delete(ID string) error {
	productID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	_, err = r.Collection.DeleteOne(
		context.Background(),
		bson.M{"_id": productID})
	if err != nil {
		return err
	}

	return nil
}
