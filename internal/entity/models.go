package entity

type Author struct {
	ID         int    `json:"author_id,omitempty" db:"author_id"`
	FullName   string `json:"full_name,require" db:"full_name"`
	Nick       string `json:"nick,omitempty" db:"nick"`
	Speciality string `json:"speciality,omitempty" db:"speciality"`
}

type Book struct {
	ID       int    `json:"book_id,omitempty" db:"book_id"`
	Name     string `json:"name,omitempty" db:"name"`
	Genre    string `json:"genre,omitempty" db:"genre"`
	CodeISBN string `json:"code_isbn,omitempty" db:"code_isbn"`
	Author
}

type Reader struct {
	ID       int    `json:"id,omitempty" db:"reader_id"`
	FullName string `json:"full_name" db:"full_name"`
}

type ReaderBookList struct {
	Book   Book   `json:"book"`
	Author Author `json:"author"`
}

type TempReaderBookList struct {
	Book
	Author
}

func (t *TempReaderBookList) ToReaderBookList() ReaderBookList {
	var readerBookList ReaderBookList
	readerBookList.Book = t.Book
	readerBookList.Author = t.Author
	return readerBookList
}
