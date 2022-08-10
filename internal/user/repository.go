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
	GetOne(ctx context.Context, ID string) (domain.User, error)
	UpdateSelf(ctx context.Context, u domain.User, ID string) error
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

func (r *repository) GetOne(ctx context.Context, ID string) (domain.User, error) {
	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := r.db.Database("tdlc")
	col := db.Collection("users")
	var user domain.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(localCtx, condicion).Decode(&user)
	user.Password = ""

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) UpdateSelf(ctx context.Context, u domain.User, ID string) error {
	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := r.db.Database("tdlc")
	col := db.Collection("users")
	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellido) > 0 {
		registro["apellido"] = u.Apellido
	}
	registro["fechaNacimiento"] = u.FechaNacimiento
	if len(u.Banner.PublicID) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Avatar.PublicID) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}
	updString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(localCtx, filtro, updString)
	return err
}
