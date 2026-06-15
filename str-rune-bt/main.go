package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "resume"
	// 文字列のインデックスアクセスはruneではなくbyte(uint8)を返す
	// → 出力: 101, uint8 ('e'のASCIIコード)
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)

	// rangeはバイト位置(i)とrune(v, int32)を返す
	// ASCIIなので1文字=1バイト → iは0〜5の連番
	// vは文字ではなく数値(コードポイント)で表示される: 114 101 115 117 109 101
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// len()はバイト数を返す。ASCIIのみなので文字数と一致 → 6
	fmt.Printf("'myString'の長さは %v", len(myString))

	// runeリテラルは数値(コードポイント)として扱われる → 97 ('a')
	var myRune = 'a'
	fmt.Printf("\nmyRune =  %v", myRune)

	// strings.Builderでスライスを効率的に連結 → "abc"
	var strSlice = []string{"a", "b", "c"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)

	// Goの文字列は不変(immutable)なため、インデックスへの代入はコンパイルエラーになる
	// catStr[0] = "a"
	// fmt.Printf("\n%v", catStr)
}
