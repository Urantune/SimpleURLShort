package linkStorage

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"

	"SimpleURLShortener/models"
)

var DB *sql.DB

var ErrNotFound = errors.New("not found")

func InitPostgres() {

	dsn := "postgres://urantune:Seigakartisde9@localhost:5432/URLShort?sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("open db error: ", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("ping db error: ", err)
	}
	log.Println("Connected to PostgreSQL")
}

func CodeExists(code string) bool {
	var count int

	_ = DB.QueryRow(`SELECT COUNT(1) FROM links WHERE code = $1`, code).Scan(&count)
	return count > 0
}

func SaveLink(code, originalURL string) error {
	_, err := DB.Exec(`
		INSERT INTO links (code, original_url, visits, created_at, updated_at)
		VALUES ($1, $2, 0, NOW(), NOW())
	`, code, originalURL)
	return err
}

func GetCode(code string) (*models.Link, error) {
	row := DB.QueryRow(`
		SELECT id, code, original_url, visits, created_at, updated_at
		FROM links WHERE code = $1
	`, code)

	var l models.Link
	if err := row.Scan(&l.ID, &l.Code, &l.OriginalURL, &l.Visits, &l.CreatedAt, &l.UpdatedAt); err != nil {
		return nil, ErrNotFound
	}
	return &l, nil
}

func IncreaseVisit(code string) error {
	res, err := DB.Exec(`
		UPDATE links
		SET visits = visits + 1, updated_at = NOW()
		WHERE code = $1
	`, code)
	if err != nil {
		return err
	}
	aff, _ := res.RowsAffected()
	if aff == 0 {
		return ErrNotFound
	}
	return nil
}

func GetAll() ([]models.Link, error) {
	rows, err := DB.Query(`
		SELECT id, code, original_url, visits, created_at, updated_at
		FROM links
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]models.Link, 0)
	for rows.Next() {
		var l models.Link
		_ = rows.Scan(&l.ID, &l.Code, &l.OriginalURL, &l.Visits, &l.CreatedAt, &l.UpdatedAt)
		list = append(list, l)
	}
	return list, nil
}
