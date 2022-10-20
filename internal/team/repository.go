package team

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ervera/tdlc-gin/internal/domain"
	"github.com/ervera/tdlc-gin/pkg/jwt"
	myuuid "github.com/ervera/tdlc-gin/pkg/myUUID"
	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Team, error)
	GetAllByUserId(ctx context.Context, userId uuid.UUID) ([]domain.Team, error)
	Save(ctx context.Context, t domain.Team) (domain.Team, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllByUserId(ctx context.Context, userId uuid.UUID) ([]domain.Team, error) {
	//query := "SELECT * FROM teams WHERE uuid in (SELECT team_uuid from teams_x_users WHERE user_uuid LIKE ?) AND enable = true ORDER BY created_on ASC"
	query := "SELECT T.*, TU.role FROM teams T INNER JOIN teams_x_users TU ON T.uuid = TU.team_uuid WHERE TU.user_uuid = ? AND enable = true ORDER BY created_on ASC"
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	var teams []domain.Team

	for rows.Next() {
		t := domain.Team{}
		_ = rows.Scan(&t.ID, &t.UUID, &t.CreatedOn, &t.Name, &t.Enable, &t.Color, &t.Image, &t.Role)
		fmt.Println(t)
		teams = append(teams, t)
	}

	return teams, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Team, error) {
	query := "SELECT * FROM teams;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var teams []domain.Team

	for rows.Next() {
		t := domain.Team{}
		_ = rows.Scan(&t.ID, &t.UUID, &t.CreatedOn, &t.Name, &t.Enable, &t.Color, &t.Image)
		teams = append(teams, t)
	}

	return teams, nil
}

func (r *repository) Save1(ctx context.Context, t domain.Team) (domain.Team, error) {
	query := "INSERT INTO teams(created_on,name,enable,color,user_id) VALUES (?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	//r.db.QueryRow(query).Scan(&id)
	if err != nil {
		return domain.Team{}, err
	}

	_, err = stmt.Exec(time.Time(t.CreatedOn), t.Name, t.Enable, t.Color)
	if err != nil {
		return domain.Team{}, err
	}
	return t, nil
}

func (r *repository) Save(ctx context.Context, t domain.Team) (domain.Team, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return domain.Team{}, err
	}
	defer tx.Rollback()
	query := "INSERT INTO teams(created_on,name,enable,color,image) VALUES (?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	//r.db.QueryRow(query).Scan(&id)
	if err != nil {
		return domain.Team{}, err
	}
	result, err := stmt.Exec(time.Time(t.CreatedOn), t.Name, t.Enable, t.Color, t.Image)
	if err != nil {
		return domain.Team{}, err
	}
	ID, err := result.LastInsertId()
	if err != nil {
		return domain.Team{}, err
	}
	uuid, err := myuuid.LastInsertUUID(r.db, "teams", ID, ctx)
	if err != nil {
		return domain.Team{}, err
	}
	t.UUID = uuid

	query = "INSERT INTO teams_x_users(created_on,role,user_uuid,team_uuid) VALUES (?,?,?,?)"
	stmt, err = r.db.Prepare(query)
	if err != nil {
		return domain.Team{}, err
	}
	_, err = stmt.Exec(time.Time(t.CreatedOn), "owner", jwt.UserID, t.UUID)
	if err != nil {
		return domain.Team{}, err
	}
	err = tx.Commit()
	if err != nil {
		return domain.Team{}, err
	}
	return t, nil
}
