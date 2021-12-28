package main

import "fmt"

/*
	go语言结构体
	结构体类型变量使用struct关键字定义

*/

type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	var Book1 Book
	var Book2 Book

	Book1.title = "go 教程"
	Book1.author = "www.w3cSchool.cn"
	Book1.subject = "go 语言教程"
	Book1.bookId = 6495407

	Book2.title = "python 教程"
	Book2.author = "www.w3cSchool.cn"
	Book2.subject = "python 语言教程"
	Book2.bookId = 6495700

	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 bookId : %d\n", Book1.bookId)

	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 bookId : %d\n", Book2.bookId)

	printBook(Book1)

	printBook(Book2)

}

/*
	结构体作为函数参数
*/
func printBook(book Book) {
	fmt.Printf("Book  title : %s\n", book.title)
	fmt.Printf("Book  author : %s\n", book.author)
	fmt.Printf("Book  subject : %s\n", book.subject)
	fmt.Printf("Book  bookId : %d\n", book.bookId)
}
