package data

import (
	"fmt"
	"strings"
)

//Result this is necessary to read data into the database via a join
type Result struct {
	BookId int
	NameOfBook string
	NameOfPublisher string

}

//String output info in console
func (R Result) String() string {
	return fmt.Sprintln(R.BookId, strings.TrimSpace(R.NameOfBook), strings.TrimSpace(R.NameOfPublisher))
}
