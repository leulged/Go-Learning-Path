package models

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book  // <-- remove 'models.' prefix here
}
