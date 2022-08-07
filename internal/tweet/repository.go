package tweet

import (
	"context"
	"errors"
	"time"

	"github.com/ervera/tdlc-gin/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Save(ctx context.Context, t domain.Tweet) (domain.Tweet, error)
	GetAllByUserId(ctx context.Context, ID string, page int64) ([]domain.Tweet, error)
	DeleteOne(ctx context.Context, ID string, UserID string) error
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

func (r *repository) GetAllByUserId(ctx context.Context, ID string, page int64) ([]domain.Tweet, error) {
	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := r.db.Database("tdlc")
	col := db.Collection("tweet")

	var tweets []domain.Tweet

	condicion := bson.M{
		"userid": ID,
	}
	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((page) * 20)

	cursor, err := col.Find(localCtx, condicion, opciones)

	if err != nil {
		return tweets, err
	}
	for cursor.Next(context.TODO()) {
		var registro domain.Tweet
		err := cursor.Decode(&registro)
		if err != nil {
			return tweets, nil
		}
		tweets = append(tweets, registro)
	}

	return tweets, nil
}

func (r *repository) DeleteOne(ctx context.Context, ID string, UserID string) error {
	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := r.db.Database("tdlc")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    objID,
		"userid": UserID,
	}
	result, err := col.DeleteOne(localCtx, condicion)
	if err != nil {
		return err
	}
	if result.DeletedCount < 1 {
		return errors.New("no se elimino ninguno")
	}
	return err
}
