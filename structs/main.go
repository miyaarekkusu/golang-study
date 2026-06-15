package main

import "fmt"

// ===== structとは =====
// structは複数のフィールド（データ）を1つにまとめて、新しい型として定義する仕組み。
// 他言語の「クラスのデータ部分のみ」に近いイメージ（メソッドは外側で別途定義する）。
// フィールドにはそれぞれ型を指定する。

// gasEngine: ガソリン車のエンジンを表すstruct
// mpg (miles per gallon) : 1ガロンあたり走れる距離
// gallons                : タンクに入っているガソリンの量
type gasEngine struct {
	mpg     uint8
	gallons uint8
	// owner
	// int
}

// structのフィールドの型には、組み込み型だけでなく
// 自分で定義した別のstruct型を指定することもできる。
// type owner struct {
// 	name string
// }

// eletricEngine: 電気自動車のエンジンを表すstruct
// mpkwh (miles per kilowatt-hour) : 1kWhあたり走れる距離
// kwh                              : バッテリーに残っている電力量
type eletricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// ===== メソッド（メソッドレシーバ）とは =====
// Goにはクラスは無いが、struct型に対して関数（メソッド）を紐づけることができる。
// ここでの (e eletricEngine) や (e gasEngine) は「値レシーバ」。
// 値レシーバではstructのコピーが渡されるため、メソッド内でeのフィールドを
func (e eletricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

// ===== interfaceとは =====
// interfaceは「どんなメソッドを持つべきか」というシグネチャだけを定義した型で、
// フィールド（データ）は持たない。
// gasEngineとeletricEngineは全く別のstructだが、
// どちらも `milesLeft() uint8` というメソッドを実装しているため、
// この engine interface を満たしている（implementしている）とみなされる。

type engine interface {
	milesLeft() uint8
}

// canMakeit: 引数の型をengine interfaceにすることで、
// gasEngineでもeletricEngineでも同じ関数で処理できる（ポリモーフィズム）。
// この関数の中では具体的な型（gasEngineかeletricEngineか）は分からないが、
// milesLeft()が呼べることだけはinterfaceによって保証されている。
func canMakeit(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("可能です")
	} else {
		fmt.Println("先に補充しないといけないです")
	}
}

func main() {
	// structリテラル: フィールド定義順に値を渡して初期化できる（mpg=25, gallons=15）
	var myEngine gasEngine = gasEngine{25, 15}

	// myEngineの型はgasEngineだが、milesLeft()を実装しているため
	// engine interfaceを引数に取るcanMakeitにそのまま渡せる
	canMakeit(myEngine, 50)

	// var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}, 10}
	// myEngine.mpg = 20

	// 「変数.フィールド名」「変数.メソッド名()」のようにドット(.)でアクセスする
	fmt.Printf("タンクに残っている距離は%v", myEngine.milesLeft())
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name, myEngine.int)
}
