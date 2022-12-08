package repositories

import "github.com/jmoiron/sqlx"

func NewPostgresDb(conRow string) (*sqlx.DB, error) {
	//db, err := sqlx.Open("postgres", fmt.Sprintf("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"))
	//cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password))
	db, err := sqlx.Open("postgres", conRow)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
