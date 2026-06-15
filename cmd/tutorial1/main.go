package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)
}

func printMe(printValue string) {
	fmt.Println(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)

	// 1. まずエラーチェックを行い、エラーならここで関数を終わらせる（早期リターン）
	if err != nil {
		fmt.Println(err.Error())
		return // ← これが重要！エラー時はここで終了し、下の処理に進ませない
	}

	// 2. 正常系：あまりがあるかないかで分岐
	if remainder == 0 {
		fmt.Printf("割り算の結果は%vです\n", result)
	} else {
		fmt.Printf("割り算の結果は %v とあまりは %vです。\n", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("割り算は正解です。")
	case 1, 2:
		fmt.Printf("割り算は結果と近いです。")
	default:
		fmt.Printf("ワイ残の結果は近くないです。")
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {

	// エラーメソッドでエラーハンドリング
	var err error
	if denominator == 0 {
		err = errors.New("0と割り算できないです")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
