package seed

import (
	"github.com/pfumarola/ld-01/database/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func getAuthorByName(firstName string, lastName string) *models.Author {
	var author models.Author = models.Author{}
	db.Where("first_name = ? AND last_name = ?", firstName, lastName).Find(&author)
	return &author
}

func seedBooks(dbPointer *gorm.DB) {
	db = dbPointer
	db.Create(getBooks())
}

func getBooks() []models.Book {
	var books = []models.Book{
		{
			ISBN:              "978-0451524935",
			Title:             "1984",
			AuthorID:          getAuthorByName("George", "Orwell").AuthorID,
			YearOfPublication: "1949",
			AvailableCopies:   15,
		},
		{
			ISBN:              "978-0061120084",
			Title:             "To Kill a Mockingbird",
			AuthorID:          getAuthorByName("Harper", "Lee").AuthorID,
			YearOfPublication: "1960",
			AvailableCopies:   12,
		},
		{
			ISBN:              "978-0141187761",
			Title:             "The Great Gatsby",
			AuthorID:          getAuthorByName("F. Scott", "Fitzgerald").AuthorID,
			YearOfPublication: "1925",
			AvailableCopies:   20,
		},
		{
			ISBN:              "978-0452284234",
			Title:             "Animal Farm",
			AuthorID:          getAuthorByName("George", "Orwell").AuthorID,
			YearOfPublication: "1945",
			AvailableCopies:   15,
		},
		{
			ISBN:              "978-0060929879",
			Title:             "The Catcher in the Rye",
			AuthorID:          getAuthorByName("J.D.", "Salinger").AuthorID,
			YearOfPublication: "1951",
			AvailableCopies:   12,
		},
		{
			ISBN:              "978-0062315007",
			Title:             "To Kill a Mockingbird",
			AuthorID:          getAuthorByName("Harper", "Lee").AuthorID,
			YearOfPublication: "2015",
			AvailableCopies:   10,
		},
		{
			ISBN:              "978-0679764024",
			Title:             "The Picture of Dorian Gray",
			AuthorID:          getAuthorByName("Oscar", "Wilde").AuthorID,
			YearOfPublication: "1890",
			AvailableCopies:   15,
		},
		{
			ISBN:              "978-0486282084",
			Title:             "Dracula",
			AuthorID:          getAuthorByName("Bram", "Stoker").AuthorID,
			YearOfPublication: "1897",
			AvailableCopies:   10,
		},
		{
			ISBN:              "978-0451526342",
			Title:             "Brave New World",
			AuthorID:          getAuthorByName("Aldous", "Huxley").AuthorID,
			YearOfPublication: "1932",
			AvailableCopies:   20,
		},
		{
			ISBN:              "978-0060850524",
			Title:             "Fahrenheit 451",
			AuthorID:          getAuthorByName("Ray", "Bradbury").AuthorID,
			YearOfPublication: "1953",
			AvailableCopies:   15,
		},
	}
	return books
}
