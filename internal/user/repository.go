package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/pkg/encrypt"
	"github.com/google/uuid"
)

type Repository interface {
	ExistAndGetByMail(ctx context.Context, email string) (domain.User, error)
	Save(ctx context.Context, user domain.User) (domain.User, error)
	GetById(ctx context.Context, ID uuid.UUID) (domain.User, error)
	Update(ctx context.Context, u domain.User) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetById(ctx context.Context, ID uuid.UUID) (domain.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := r.db.QueryRow(query, ID)
	u := domain.User{}
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.FirstName, &u.LastName, &u.CreatedOn, &u.Avatar, &u.Banner, &u.Biography, &u.Location, &u.Website)
	if err != nil {
		return u, err
	}
	u.Password = ""
	return u, nil
}

func (r *repository) ExistAndGetByMail(ctx context.Context, email string) (domain.User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	row := r.db.QueryRow(query, email)
	u := domain.User{}
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.FirstName, &u.LastName, &u.CreatedOn, &u.Avatar, &u.Banner, &u.Biography, &u.Location, &u.Website)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *repository) Save(ctx context.Context, user domain.User) (domain.User, error) {
	user.Password, _ = encrypt.Password(user.Password)
	currentTime := time.Now()
	query := "INSERT INTO users (first_name, last_name, password, email, created_on, avatar, banner, biography, location, website) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return domain.User{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(&user.FirstName, &user.LastName, &user.Password, &user.Email, &currentTime, &user.Avatar, &user.Banner, &user.Biography, &user.Location, &user.Website)
	if err != nil {
		return domain.User{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.User{}, err
	}
	if rowsAffected == 0 {
		return domain.User{}, errors.New("")
	}
	user.ID = uuid.Nil
	user.Password = ""
	return user, nil
}

func (r *repository) Update(ctx context.Context, u domain.User) error {
	query := "UPDATE users SET email=$1, first_name=$2, last_name=$3, avatar=$4, banner=$5, biography=$6, location=$7, website=$8 WHERE id=$9"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&u.Email, &u.FirstName, &u.LastName, &u.Avatar, &u.Banner, &u.Biography, &u.Location, &u.Website, &u.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// func (r *repository) Save(ctx context.Context, user domain.User) (primitive.ObjectID, bool, error) {
// 	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	db := r.db.Database("tdlc")
// 	col := db.Collection("users")

// 	user.Password, _ = encrypt.Password(user.Password)

// 	result, err := col.InsertOne(localCtx, user)
// 	if err != nil {
// 		return primitive.ObjectID{}, false, err
// 	}

// 	ObjID, _ := result.InsertedID.(primitive.ObjectID)
// 	return ObjID, true, nil
// 	//return ObjID.String(), true, nil
// }

// func (r *repository) Exists(ctx context.Context, email string) (domain.User, bool, string) {
// 	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	db := r.db.Database("tdlc")
// 	col := db.Collection("users")

// 	condicion := bson.M{"email": email}

// 	var resultado domain.User

// 	err := col.FindOne(localCtx, condicion).Decode(&resultado)
// 	ID := resultado.ID.Hex()
// 	if err != nil {
// 		return resultado, false, ID
// 	}
// 	return resultado, true, ID
// }

// func (r *repository) GetOne(ctx context.Context, ID string) (domain.User, error) {
// 	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	db := r.db.Database("tdlc")
// 	col := db.Collection("users")
// 	var user domain.User

// 	objID, _ := primitive.ObjectIDFromHex(ID)

// 	condicion := bson.M{
// 		"_id": objID,
// 	}

// 	err := col.FindOne(localCtx, condicion).Decode(&user)
// 	user.Password = ""

// 	if err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }

// func (r *repository) UpdateSelf(ctx context.Context, u domain.User, ID string) error {
// 	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	db := r.db.Database("tdlc")
// 	col := db.Collection("users")
// 	registro := make(map[string]interface{})
// 	if len(u.Nombre) > 0 {
// 		registro["nombre"] = u.Nombre
// 	}
// 	if len(u.Apellido) > 0 {
// 		registro["apellido"] = u.Apellido
// 	}
// 	registro["fechaNacimiento"] = u.FechaNacimiento
// 	if len(u.Banner.PublicID) > 0 {
// 		registro["banner"] = u.Banner
// 	}
// 	if len(u.Avatar.PublicID) > 0 {
// 		registro["avatar"] = u.Avatar
// 	}
// 	if len(u.Biografia) > 0 {
// 		registro["biografia"] = u.Biografia
// 	}
// 	if len(u.Ubicacion) > 0 {
// 		registro["ubicacion"] = u.Ubicacion
// 	}
// 	if len(u.SitioWeb) > 0 {
// 		registro["sitioWeb"] = u.SitioWeb
// 	}
// 	updString := bson.M{
// 		"$set": registro,
// 	}

// 	objID, _ := primitive.ObjectIDFromHex(ID)
// 	filtro := bson.M{"_id": bson.M{"$eq": objID}}

// 	_, err := col.UpdateOne(localCtx, filtro, updString)
// 	return err
// }

// func (r *repository) SaveUserRelation(ctx context.Context, userRelation domain.UserRelation) error {
// 	localCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	db := r.db.Database("tdlc")
// 	col := db.Collection("relation")

// 	filter := bson.M{"userId": userRelation.UserID, "userRelationId": userRelation.UserRelationId + "evc"}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"userId":         userRelation.UserID,
// 			"userRelationId": userRelation.UserRelationId + "evc",
// 		},
// 	}
// 	opts := options.Update().SetUpsert(true)

// 	//_, err := col.InsertOne(localCtx, userRelation)
// 	_, err := col.UpdateOne(localCtx, filter, update, opts)

// 	return err

// }
