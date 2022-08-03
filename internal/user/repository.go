package user

import (
	"context"
	"time"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/pkg/encrypt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Save(ctx context.Context, user domain.User) (string, bool, error)
	Exists(ctx context.Context, email string) (domain.User, bool, string)
}

type repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, user domain.User) (string, bool, error) {
	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := r.db.Database("tdlc")
	col := db.Collection("users")

	user.Password, _ = encrypt.Password(user.Password)

	result, err := col.InsertOne(localCtx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}

func (r *repository) Exists(ctx context.Context, email string) (domain.User, bool, string) {
	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := r.db.Database("tdlc")
	col := db.Collection("users")

	condicion := bson.M{"email": email}

	var resultado domain.User

	err := col.FindOne(localCtx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
