package book

import (
	"log"

	db "github.com/KazuwoKiwame12/book_store_app_api/DB"
)

//Book ...型の構造体
//あとでjson形式にするので、jsonのタグをあらかじめつけておきます
type Book struct {
	ID          int    `json:"id"`
	TITLE       string `json:"title"`
	DESCRIPTION string `json:"description"`
}

//Get ...Book型のデータを渡す
func Get() []Book {
	db := db.Connect()
	defer db.Close()

	//rowを取得
	rows, err := db.Query("SELECT * FROM book")
	if err != nil {
		log.Fatal(err.Error())
	}
	//Book型のスライスに格納します
	bookArgs := make([]Book, 0)
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.TITLE, &book.DESCRIPTION)
		if err != nil {
			log.Fatal(err.Error())
		}
		bookArgs = append(bookArgs, book)
	}
	return bookArgs
}

//Delete ...DBから引数のidを持つbookを削除
func Delete(id int) bool {
	db := db.Connect()
	defer db.Close()

	//idを持つデータの削除
	stmtDelete, err1 := db.Prepare("DELETE FROM book WHERE id=?")
	if err1 != nil {
		panic(err1.Error())
	}
	defer stmtDelete.Close()

	_, err2 := stmtDelete.Exec(id)
	if err2 != nil {
		panic(err2.Error())
	}

	return isnotError(err2)
}

// Add ...DBへとデータを追加する
func Add(title string, description string) bool {
	db := db.Connect()
	defer db.Close()
	_, err := db.Exec("INSERT INTO book(title, description) VALUES (?, ?)", title, description)
	return isnotError(err)
}

func isnotError(err error) bool {
	if err != nil {
		return false
	}
	return true
}
