package main

import "fmt"

type Book struct{
	Title string
	Author string
	IsBorrowed bool
}

type Library interface {
	Borrow() string
	Return() string
}

func (b *Book) Borrow() string{
	if b.IsBorrowed {
		return "Book is already borrowed"
	}
	b.IsBorrowed= true
	return fmt.Sprintf("You borrowed '%s' by %s", b.Title, b.Author)
}

func (b*Book) Return() string {
	if !b.IsBorrowed {
		return "Book was not borrowed"
	}
	b.IsBorrowed= false
	return fmt.Sprintf("You returned '%s'", b.Title)
}

func main() {
	book:= &Book{Title: "To kill a mocking bird", Author: "Harper Lee"}
	fmt.Println(book.Borrow())
	fmt.Println(book.Borrow())
	fmt.Println(book.Return())
	fmt.Println(book.Return())
}