package mysql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"libraryService/internal/entity"
	"libraryService/internal/storage"
	"log"
	"os"
	"strconv"
	"time"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() (storage.IStorage, error) {
	db, err := sql.Open("mysql", GetDefaultConfig())
	if err != nil {
		return Storage{}, err
	}
	d, err := strconv.Atoi(os.Getenv("MYSQL_CONN_MAX_LIFETIME"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(time.Duration(d) * time.Millisecond)
	db.SetConnMaxLifetime(time.Duration(d) * time.Millisecond)
	n, err := strconv.Atoi(os.Getenv("MYSQL_MAX_IDLE_CONNS"))
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(n)

	if err = db.Ping(); err != nil {
		return Storage{}, err
	}
	if err != nil {
		log.Fatal(err)
	}
	return Storage{db: db}, nil
}

func (s Storage) GetBooksByAuthorName(ctx context.Context, name string) []entity.Book {
	query := `
		SELECT b.id, b.title, b.page_count, b.publishing_year
			FROM book b
			JOIN author_x_book ba ON b.id = ba.book_id
			JOIN author a ON ba.author_id = a.id
		WHERE a.name = ?;
	`

	rows, err := s.db.QueryContext(ctx, query, name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		book  entity.Book
		books []entity.Book
	)

	for rows.Next() {
		if err := rows.Scan(&book.ID, &book.Title, &book.PageCount, &book.PublishingYear); err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return books
}

func (s Storage) CheckIfBookExists(ctx context.Context, title string) bool {
	query := `SELECT count(*) FROM book WHERE title = ?`
	row := s.db.QueryRowContext(ctx, query, title)

	var bookCount int
	err := row.Scan(&bookCount)
	if err != nil {
		log.Fatal(err)
	}

	if bookCount > 1 {
		log.Println("The database is in an inconsistent state, the result may be incorrect")
	}

	return bookCount == 1
}

func (s Storage) GetAuthorsByBookTitle(ctx context.Context, title string) []entity.Author {
	query := `
		SELECT a.id, a.name, a.book_count, a.birth_date
			FROM book b
			JOIN author_x_book ba ON b.id = ba.book_id
			JOIN author a ON ba.author_id = a.id
		WHERE b.title = ?;
	`

	rows, err := s.db.QueryContext(ctx, query, title)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		author  entity.Author
		authors []entity.Author
	)

	for rows.Next() {
		if err := rows.Scan(&author.ID, &author.Name, &author.BookCount, &author.BirthDate); err != nil {
			log.Fatal(err)
		}
		authors = append(authors, author)
	}
	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return authors
}

func (s Storage) CheckIfAuthorExists(ctx context.Context, name string) bool {
	query := `SELECT count(*) FROM author WHERE name = ?`
	row := s.db.QueryRowContext(ctx, query, name)

	var authorCount int
	err := row.Scan(&authorCount)
	if err != nil {
		log.Fatal(err)
	}

	if authorCount > 1 {
		log.Println("The database is in an inconsistent state, the result may be incorrect")
	}

	return authorCount == 1
}
