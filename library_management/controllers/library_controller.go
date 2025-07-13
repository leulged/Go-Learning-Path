package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library_management/models"
	"library_management/services"
)

func StartLibraryApp() {
	library := services.NewLibrary()

	// Preload a test member with ID 1
	library.Members[1] = models.Member{
		ID:            1,
		Name:          "John Doe",
		BorrowedBooks: []models.Book{},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== Library Management =====")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove an existing book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List all borrowed books by a member")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		input := readLine(reader)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			addBook(reader, library)
		case 2:
			removeBook(reader, library)
		case 3:
			borrowBook(reader, library)
		case 4:
			returnBook(reader, library)
		case 5:
			listAvailableBooks(library)
		case 6:
			listBorrowedBooks(reader, library)
		case 7:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func addBook(reader *bufio.Reader, library *services.Library) {
	fmt.Print("Enter Book ID: ")
	id := readInt(reader)
	fmt.Print("Enter Book Title: ")
	title := readLine(reader)
	fmt.Print("Enter Book Author: ")
	author := readLine(reader)

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}

	library.AddBook(book)
}

func removeBook(reader *bufio.Reader, library *services.Library) {
	fmt.Print("Enter Book ID to remove: ")
	id := readInt(reader)
	library.RemoveBook(id)
}

func borrowBook(reader *bufio.Reader, library *services.Library) {
	fmt.Print("Enter Book ID: ")
	bookID := readInt(reader)
	fmt.Print("Enter Member ID: ")
	memberID := readInt(reader)

	if err := library.BorrowBook(bookID, memberID); err != nil {
		fmt.Println("Error:", err)
	}
}

func returnBook(reader *bufio.Reader, library *services.Library) {
	fmt.Print("Enter Book ID: ")
	bookID := readInt(reader)
	fmt.Print("Enter Member ID: ")
	memberID := readInt(reader)

	if err := library.ReturnBook(bookID, memberID); err != nil {
		fmt.Println("Error:", err)
	}
}

func listAvailableBooks(library *services.Library) {
	books := library.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}

func listBorrowedBooks(reader *bufio.Reader, library *services.Library) {
	fmt.Print("Enter Member ID: ")
	memberID := readInt(reader)
	books := library.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No borrowed books or invalid member ID.")
		return
	}
	for _, b := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
	}
}

func readInt(reader *bufio.Reader) int {
	text := readLine(reader)
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Invalid number. Try again.")
		return readInt(reader)
	}
	return num
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}
