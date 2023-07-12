package common

import (
	"fmt"
	"testing"
)

func TestGetAllBookIDs(t *testing.T) {
	ret, err := GetAllBookIDs()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)
}
func TestGetAllDocIDs(t *testing.T) {
	ret, err := GetAllDocIDs(28734442)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)
}
