package util

import (
	"github.com/mozillazg/go-pinyin"
)

func Hanzi2pinyin(s string)( string){
	//处理非汉子方法
	var Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}
	//newpy := pinyin.NewArgs()
	newpyx := pinyin.Args{0,false,"_",Fallback}
	r := pinyin.Slug(s,newpyx)
	return r
}