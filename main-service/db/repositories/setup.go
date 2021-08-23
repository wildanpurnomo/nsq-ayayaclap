package repositories

import "database/sql"

type Repository struct {
	db *sql.DB
}

var Repo *Repository

func InitRepository(db *sql.DB) {
	Repo = &Repository{db: db}
}

func (r *Repository) CloseDB() {
	r.db.Close()
}
