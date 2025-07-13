package services

import (
    "fmt"
    "library_management/models"
)

type Library struct {
    Books   map[int]models.Book
    Members map[int]models.Member
}

type LibraryManager interface {
    AddBook(book models.Book)
    RemoveBook(bookID int)
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberID int) []models.Book
    BorrowBook(bookID, memberID int) error
    ReturnBook(bookID, memberID int) error
}

// Constructor function to initialize Library struct with maps
func NewLibrary() *Library {
    return &Library{
        Books:   make(map[int]models.Book),
        Members: make(map[int]models.Member),
    }
}

// Method to add a book
func (l *Library) AddBook(book models.Book) {
    if _, exist := l.Books[book.ID]; exist {
        fmt.Println("Book with this ID already exists")
        return
    }
    l.Books[book.ID] = book
    fmt.Println("Book added successfully")
}

// Method to remove a book by ID
func (l *Library) RemoveBook(bookID int) {
    if _, exist := l.Books[bookID]; exist {
        delete(l.Books, bookID)
        fmt.Println("Book removed successfully")
    } else {
        fmt.Println("There is no book using this ID")
    }
}

// Method to list all available books
func (l *Library) ListAvailableBooks() []models.Book {
    available := []models.Book{}
    for _, book := range l.Books {
        if book.Status == "Available" {
            available = append(available, book)
        }
    }
    return available
}

// Method to list all borrowed books by a member
func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
    member, exists := l.Members[memberID]
    if !exists {
        fmt.Println("Member not found")
        return nil
    }
    return member.BorrowedBooks
}

// Method to borrow a book
func (l *Library) BorrowBook(bookID, memberID int) error {
    book, bookExists := l.Books[bookID]
    if !bookExists {
        return fmt.Errorf("Book with ID %d does not exist", bookID)
    }

    if book.Status == "Borrowed" {
        return fmt.Errorf("Book with ID %d is already borrowed", bookID)
    }

    member, memberExists := l.Members[memberID]
    if !memberExists {
        return fmt.Errorf("Member with ID %d does not exist", memberID)
    }

    // Update book status to "Borrowed"
    book.Status = "Borrowed"
    l.Books[bookID] = book

    // Add book to member's BorrowedBooks
    member.BorrowedBooks = append(member.BorrowedBooks, book)
    l.Members[memberID] = member

    fmt.Println("Book borrowed successfully")
    return nil
}

// Method to return a book
func (l *Library) ReturnBook(bookID, memberID int) error {
    book, bookExists := l.Books[bookID]
    if !bookExists {
        return fmt.Errorf("Book with ID %d does not exist", bookID)
    }

    member, memberExists := l.Members[memberID]
    if !memberExists {
        return fmt.Errorf("Member with ID %d does not exist", memberID)
    }

    // Remove the book from the member's BorrowedBooks slice
    found := false
    newBorrowed := []models.Book{}
    for _, b := range member.BorrowedBooks {
        if b.ID == bookID {
            found = true
            continue // skip this book (we're removing it)
        }
        newBorrowed = append(newBorrowed, b)
    }

    if !found {
        return fmt.Errorf("Book with ID %d is not borrowed by member %d", bookID, memberID)
    }

    member.BorrowedBooks = newBorrowed
    l.Members[memberID] = member

    // Update the book status back to "Available"
    book.Status = "Available"
    l.Books[bookID] = book

    fmt.Println("Book returned successfully")
    return nil
}
