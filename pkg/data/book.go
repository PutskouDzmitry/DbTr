package data

import (
	"database/sql"
	"fmt"
	dbConst "github.com/PutskouDzmitry/DbTr/pkg/const_db"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
)

//Entity in database
type Book struct {
	BookId            int    // primary key
	AuthorId          int    // foreign key
	PublisherId       int    // foreign key
	NameOfBook        string // name of book
	YearOfPublication string // year of publication of the book
	BookVolume        int    // book volume
	Number            int    // number of book
	Price             int
}

//ReadAll output all data with table books
func (B BookData) ReadAll() ([]Book, error) {
	var books []Book
	result := B.db.Find(&books)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read users from database, error: %w", result.Error)
	}
	logrus.Info(books)
	return books, nil
}

//String output data in console
func (B Book) String() string {
	return fmt.Sprintln(B.BookId, B.AuthorId, B.PublisherId, strings.TrimSpace(B.NameOfBook), B.YearOfPublication, B.BookVolume, B.Number)
}

//BookData create a new connection
type BookData struct {
	db *gorm.DB // connection in db
}

//NewBookData it's imitation constructor
func NewBookData(db *gorm.DB) *BookData {
	return &BookData{db: db}
}

//Read read data in db
func (B BookData) Read(id int) ([]Result, error) {
	var results []Result
	result := B.db.Table(dbConst.Publishers).Select(dbConst.SelectBookAndPublisher).
		Joins(dbConst.ReadBookWithJoin).Where("book_id", id).
		Find(&results)
	if result.Error != nil {
		return nil, result.Error
	}
	return results, nil
}

//Add add data in db
func (B BookData) Add(book Book) (int, error) {
	result := B.db.Create(&book)
	if result.Error != nil {
		return -1, fmt.Errorf(dbConst.CantAddDataError, result.Error)
	}
	return book.BookId, nil
}

//Update update number of books by the id
func (B BookData) Update(id int, value int) error {
	result := B.db.Table(dbConst.Books).Where(dbConst.BookId, id).Update("number", value)
	if result.Error != nil {
		return fmt.Errorf(dbConst.CantUpdateDataError, result.Error)
	}
	return nil
}

//Delete delete data in db
func (B BookData) Delete(id int) error {
	result := B.db.Where(dbConst.BookId, id).Delete(&Book{})
	if result.Error != nil {
		return fmt.Errorf(dbConst.CantDeleteDataError, result.Error)
	}
	return nil
}

type sellBook struct {
	Number int
	Price  int
}

type money struct {
	Money int
}

func (B BookData) BuyBook(name string) (int, error) {
	var bookSale sellBook
	var moneyUser money
	err := B.db.Transaction(func(tx *gorm.DB) error {
		tx.Begin(&sql.TxOptions{Isolation: sql.LevelRepeatableRead})
		result := tx.Table("userMoney").Select("*").Find(&moneyUser)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
		result = tx.Table(dbConst.Books).Select("number, price").Where("books.name_of_book = ?", name).Find(&bookSale)
		if bookSale.Price > moneyUser.Money {
			tx.Rollback()
			return fmt.Errorf("not enough mooney")
		}
		if bookSale.Number == 0 {
			tx.Rollback()
			return fmt.Errorf("you cann't buy a book, because we sold all books")
		}
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
		result = tx.Table("userMoney").Where("money = ?", moneyUser.Money).Update("money", moneyUser.Money-bookSale.Price)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
		result = tx.Table(dbConst.Books).Where("books.name_of_book = ?", name).Update("number", bookSale.Number-1)
		if result = tx.Commit(); result.Error != nil {
			tx.Rollback()
			return fmt.Errorf("your transaction has been canceled and your money is saved, because %v", result.Error)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	checkResult := strings.Contains(err.Error(), "transaction has already been committed")
	if !checkResult {
		return -1, err
	}
	logrus.Info("Deal hab been completed successfully!")
	return bookSale.Number, nil
}
