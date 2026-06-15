package main

import "fmt"

func main() {
	// var p *int32 = new(int32)
	// // var p *int32
	// var i int32
	// fmt.Printf("pポインターのバリューは%v", *p)
	// fmt.Printf("\n iのバリューは：%v", i)
	// p = &i
	// *p = 10
	// fmt.Printf("\np ポインターのバリューは：%v", *p)
	// fmt.Printf("\niのバリューは：%v", i)
	// var k int32 = 2
	// i = k

	// // スライスはポインター入っているため、元のデータ書き換えられる
	// var slice = []int32{1, 2, 3}
	// var sliceCopy = slice
	// sliceCopy[2] = 4
	// fmt.Println(slice)
	// fmt.Println(sliceCopy)

	// 配列をポインタで渡すと、関数内でも元の配列本体を直接書き換えられる
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nthing1配列のメモリ領域は：%p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\n結果は：%v", result)
	fmt.Printf("\nthing1のバリューは：%v", thing1)
	// 実行結果：thing1も[1 4 9 16 25]に変わっている
	// → ポインタ渡しなのでsquare内の書き換えがthing1本体に反映される（値渡しならthing1は変わらない）
}

func square(thing2 *[5]float64) [5]float64 {
	// thing2は*[5]float64（配列へのポインタ）。&thing2はポインタ変数thing2自体のアドレスなので
	// thing1のアドレスとは別の値になる（実行結果でも2つのアドレスが異なる）
	fmt.Printf("\nthing2配列のメモリ領域は：%p", &thing2)
	for i := range thing2 {
		// thing2[i]は(*thing2)[i]と同じ意味。Goが自動でポインタをデリファレンスしてくれる
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2 // *thing2で配列本体の値（コピー）を返す
}
