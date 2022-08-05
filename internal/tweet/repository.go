package tweet

import (
	"context"
	"time"

	"github.com/ervera/tdlc-gin/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Save(ctx context.Context, t domain.Tweet) (domain.Tweet, error)
}

type repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, t domain.Tweet) (domain.Tweet, error) {
	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := r.db.Database("tdlc")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	//_, err := col.InsertOne(localCtx, registro)
	result, err := col.InsertOne(localCtx, registro)
	t.ID = result.InsertedID.(primitive.ObjectID)
	//result, err := col.InsertOne(localCtx, registro)
	// if err != nil {
	// 	return err
	// }

	//ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return t, err
}
