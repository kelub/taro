package util


import (
	"testing"
	"fmt"
)
func TestPinyin(test *testing.T){
	s := Hanzi2pinyin("汉字")

	fmt.Println(s)

}