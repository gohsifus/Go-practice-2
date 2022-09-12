package inPostgres

import (
	"database/sql"
	"task11/domain/entity"
	"time"
)

type postgresEventRepo struct{
	store *sql.DB
}

func NewPostgresEventRepo(connectionString string) (*postgresEventRepo, error){
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresEventRepo{store: conn}, nil
}

func (p *postgresEventRepo) Create(event *entity.Event) (*entity.Event, error){return nil, nil}
func (p *postgresEventRepo) Update(id int, event *entity.Event) error{return nil}
func (p *postgresEventRepo) Delete(id int) error{return nil}
func (p *postgresEventRepo) GetEventsByDateInterval(from, to time.Time) ([]entity.Event, error){return nil, nil}


