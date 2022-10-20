package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/pkg/encrypt"
	myuuid "github.com/ervera/tdlc-gin/pkg/myUUID"
	"github.com/google/uuid"
)

type Repository interface {
	ExistAndGetByMail(ctx context.Context, email string) (domain.User, error)
	Save(ctx context.Context, user domain.User) (domain.User, error)
	GetById(ctx context.Context, ID uuid.UUID) (domain.User, error)
	Update(ctx context.Context, u domain.User) error
	UpdatePasswordById(ctx context.Context, u domain.User, newPassword string) error
	UpdateMedia(ctx context.Context, u domain.User) error
	DeleteUserMedia(ctx context.Context, u domain.User, mediaType string) error
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
	query := "SELECT * FROM users WHERE uuid = ?"
	row := r.db.QueryRow(query, ID)
	u := domain.User{}
	err := row.Scan(&u.ID, &u.UUID, &u.Email, &u.Password, &u.FirstName, &u.LastName, &u.CreatedOn, &u.Avatar, &u.Banner, &u.Biography, &u.Location, &u.Occupation, &u.Website)
	if err != nil {
		return u, err
	}
	u.Password = ""
	return u, nil
}

func (r *repository) ExistAndGetByMail(ctx context.Context, email string) (domain.User, error) {
	query := "SELECT * FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)
	u := domain.User{}
	err := row.Scan(&u.ID, &u.UUID, &u.Email, &u.Password, &u.FirstName, &u.LastName, &u.CreatedOn, &u.Avatar, &u.Banner, &u.Biography, &u.Location, &u.Occupation, &u.Website)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *repository) Save(ctx context.Context, user domain.User) (domain.User, error) {
	user.Password, _ = encrypt.Password(user.Password)
	currentTime := time.Now()
	query := "INSERT INTO users (first_name, last_name, password, email, created_on, avatar, banner, biography, location, occupation, website) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return domain.User{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(&user.FirstName, &user.LastName, &user.Password, &user.Email, &currentTime, &user.Avatar, &user.Banner, &user.Biography, &user.Location, &user.Occupation, &user.Website)
	if err != nil {
		return domain.User{}, err
	}
	ID, _ := result.LastInsertId()
	UUID, _ := myuuid.LastInsertUUID(r.db, "users", ID, ctx)
	fmt.Println(UUID)
	fmt.Println(UUID)
	fmt.Println(UUID)
	fmt.Println(UUID)
	fmt.Println(UUID)
	fmt.Println(UUID)
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return domain.User{}, err
	}
	if rowsAffected == 0 {
		return domain.User{}, errors.New("")
	}
	user.Password = ""
	return user, nil
}

func (r *repository) Update(ctx context.Context, u domain.User) error {
	query := "UPDATE users SET first_name=?, last_name=?, biography=?, location=?, occupation=?, website=? WHERE uuid=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	fmt.Println(u)
	res, err := stmt.Exec(&u.FirstName, &u.LastName, &u.Biography, &u.Location, &u.Occupation, &u.Website, &u.UUID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateMedia(ctx context.Context, u domain.User) error {
	query := "UPDATE users SET avatar=?, banner=? WHERE uuid=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&u.Avatar, &u.Banner, &u.UUID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdatePasswordById(ctx context.Context, u domain.User, newPassword string) error {
	newPassword, _ = encrypt.Password(newPassword)

	query := "UPDATE users SET password=? WHERE uuid=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&newPassword, &u.UUID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteUserMedia(ctx context.Context, u domain.User, mediaType string) error {
	emptyString := ""
	query := fmt.Sprintf("UPDATE users SET %s=? WHERE uuid=?", mediaType)
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&emptyString, &u.UUID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
