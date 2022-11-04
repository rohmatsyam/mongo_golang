package domain

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cake struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title" binding:"required"`
	Description string             `json:"description" bson:"description"`
	Rating      float32            `json:"rating" bson:"rating"`
	Image       string             `json:"image" bson:"image"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   primitive.DateTime `json:"updated_at" bson:"updated_at,omitempty"`
}

type CakeUseCase interface {
	GetCakesUC() ([]*Cake, error)
	GetCakeIdUC(ctx *gin.Context) (*Cake, error)
	CreateCakeUC(ctx *gin.Context) (*Cake, error)
	UpdateCakeUC(ctx *gin.Context) (*Cake, error)
	DeleteCakeUC(ctx *gin.Context) error
}

type CakeRepository interface {
	GetCakesRepository() ([]*Cake, error)
	GetCakeIdRepository(ctx *gin.Context) (*Cake, error)
	CreateCakeRepository(ctx *gin.Context) (*Cake, error)
	UpdateCakeRepository(ctx *gin.Context) (*Cake, error)
	DeleteCakeRepository(ctx *gin.Context) error
}
