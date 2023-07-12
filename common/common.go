package common

// GetAllBookIDs 得到所有bookID
func GetAllBookIDs() ([]int, error) {
	books, err := API{}.GetAllBooks()
	if err != nil {
		return nil, err
	}
	var ret []int
	for _, data := range books.Data {
		for _, book := range data.Books {
			ret = append(ret, book.ID)
		}
	}
	return ret, nil
}

// GetAllDocIDs 得到所有文章ID
func GetAllDocIDs(bookID int) ([]int, error) {
	docs, err := API{}.GetAllDocs(bookID)
	if err != nil {
		return nil, err
	}

	var ret []int
	for _, data := range docs.Data {
		ret = append(ret, data.ID)
	}
	return ret, nil
}
