package main

import "fmt"

/*
	结构体指针
*/

type BookDemo struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	var Book1 BookDemo
	var Book2 BookDemo

	Book1.title = "go 教程"
	Book1.author = "www.w3cSchool.cn"
	Book1.subject = "go 语言教程"
	Book1.bookId = 6495407

	Book2.title = "python 教程"
	Book2.author = "www.w3cSchool.cn"
	Book2.subject = "python 语言教程"
	Book2.bookId = 6495700

	printBooks(&Book1)
	printBooks(&Book2)
}

func printBooks(book *BookDemo) {
	fmt.Printf("Book  title : %s\n", book.title)
	fmt.Printf("Book  author : %s\n", book.author)
	fmt.Printf("Book  subject : %s\n", book.subject)
	fmt.Printf("Book  bookId : %d\n", book.bookId)
}
