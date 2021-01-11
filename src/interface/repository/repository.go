package repository

import "github.com/rubrafmex/clean-architecture-go/domain"

type DBHandler interface {
	FindAllBooks() ([]*domain.Book, error)
	SaveBook(book domain.Book) error
	SaveAuthor(author domain.Author) error
}
