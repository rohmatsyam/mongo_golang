package mongo

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rohmatsyam/mongo_golang/domain"
	"github.com/rohmatsyam/mongo_golang/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type cakeRepository struct {
	CakeCollection *mongo.Collection
}

var appJSON = "application/json"

func NewCakeRepository(cakeCollection *mongo.Collection) domain.CakeRepository {
	return &cakeRepository{
		CakeCollection: cakeCollection,
	}
}

func (c cakeRepository) GetCakesRepository() (cakes []*domain.Cake, err error) {
	contex, _ := context.WithTimeout(context.Background(), time.Second*20)
	cursor, err := c.CakeCollection.Find(contex, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(contex)
	for cursor.Next(contex) {
		cake := &domain.Cake{}
		err := cursor.Decode(cake)
		if err != nil {
			return nil, err
		}
		cakes = append(cakes, cake)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return cakes, nil
}

func (c cakeRepository) GetCakeIdRepository(ctx *gin.Context) (cake *domain.Cake, err error) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		log.Println("Invalid id")
	}
	contex, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = c.CakeCollection.FindOne(contex, bson.M{"_id": id}).Decode(&cake)
	if err != nil {
		return nil, err
	}
	return cake, nil
}

func (c cakeRepository) CreateCakeRepository(ctx *gin.Context) (cake *domain.Cake, err error) {
	contentType := helpers.GetContentType(ctx)
	if contentType == appJSON {
		ctx.ShouldBindJSON(&cake)
	} else {
		ctx.ShouldBind(&cake)
	}
	contex, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := c.CakeCollection.InsertOne(contex, bson.M{
		"title":       cake.Title,
		"description": cake.Description,
		"rating":      cake.Rating,
		"image":       cake.Image,
		"created_at":  primitive.NewDateTimeFromTime(time.Now()),
	})
	if err != nil {
		return nil, errors.New("can't insert cake to the database.")
	}
	cake.Id = result.InsertedID.(primitive.ObjectID)
	return cake, nil
}

func (c cakeRepository) UpdateCakeRepository(ctx *gin.Context) (cake *domain.Cake, err error) {
	oldcake, err := c.GetCakeIdRepository(ctx)
	if err != nil {
		return nil, err
	}

	contentType := helpers.GetContentType(ctx)
	if contentType == appJSON {
		ctx.ShouldBindJSON(&cake)
	} else {
		ctx.ShouldBind(&cake)
	}
	contex, _ := context.WithTimeout(context.Background(), time.Second*20)
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	_, err = c.CakeCollection.UpdateOne(contex, bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"title":       cake.Title,
			"description": cake.Description,
			"rating":      cake.Rating,
			"image":       cake.Image,
			"updated_at":  primitive.NewDateTimeFromTime(time.Now()),
		}})
	if err != nil {
		return nil, errors.New("can't update cake to the database.")
	}
	cake.Id = oldcake.Id
	return cake, nil
}

func (c cakeRepository) DeleteCakeRepository(ctx *gin.Context) (err error) {
	_, err = c.GetCakeIdRepository(ctx)
	if err != nil {
		return err
	}
	contex, _ := context.WithTimeout(context.Background(), time.Second*20)
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	_, err = c.CakeCollection.DeleteOne(contex, bson.M{"_id": id})
	if err != nil {
		return errors.New("can't delete cake to the database.")
	}
	return nil
}
