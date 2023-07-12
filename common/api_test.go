package common

import (
	"fmt"
	"testing"
)

func TestAPI_GetAllDocs(t *testing.T) {
	ret, err := API{}.GetAllDocs(28734442)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(ret)
}
func TestAPI_GetAllBooks(t *testing.T) {
	ret, err := API{}.GetAllBooks()
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(ret)
}

func TestAPI_DownloadDoc(t *testing.T) {
	err := API{}.DownloadDoc("https://www.yuque.com/guo-fei/gbwowk/awp9nbmalsr6211y/markdown?attachment=true&latexcode=false&anchor=true&linebreak=true")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAPI_GetDownloadLink1(t *testing.T) {
	got, err := API{}.GetDownloadLink(128385030)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(got)
}
