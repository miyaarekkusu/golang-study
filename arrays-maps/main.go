package main

import "fmt"

func main() {
	// 配列(array): [...]で要素数を自動推論させ、サイズ固定の配列を作成する
	// この場合は要素が3つなので [3]int32 と同じ意味になる
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	// スライス(slice): 可変長で、配列への参照(ポインタ・長さ・容量)を持つデータ構造
	// {4, 5, 6} で初期化されるので、長さ・容量はともに3になる
	var intSlice []int32 = []int32{4, 5, 6}
	// len(): 現在の要素数, cap(): 内部配列が確保している容量を取得
	fmt.Printf("the length is %v with capacity %v", len(intSlice), cap(intSlice))

	// append(): スライスの末尾に要素を追加する
	// 容量が足りない場合は、より大きい内部配列が新たに確保され、要素がコピーされる
	intSlice = append(intSlice, 7)
	// 容量を超えたため、capが2倍(3 -> 6)に増えていることが確認できる
	fmt.Printf("\nthe length is %v with capacity %v", len(intSlice), cap(intSlice))

	// make(): スライスを生成する組み込み関数
	// make([]型, 長さ, 容量) の順で指定する -> 長さ3・容量8のint32スライスを作成
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// mapデータ構造
	// key valueのペア("key": "value") → キーからデータ取得
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Sarah"])
	var age, ok = myMap2["Jason"]

	if ok {
		fmt.Printf("年齢は%vです。", age)
	} else {
		fmt.Println("無効な名前です。")
	}

	// for range (map): key, valueの順で全要素を取得できる
	// マップの反復順序は保証されない(実行ごとに順番が変わる可能性がある)
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	// for range (配列・スライス): index, valueの順で全要素を取得できる
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// 通常のforループ: 初期化; 条件; 後処理 の3つを指定する(C言語と同様の書き方)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while文相当の書き方: 条件部分のみを書き、初期化・後処理は外側/内側で行う
	// breakで条件を満たした時点でループを抜ける
	// for {
	// 	if i >= 10 {
	// 	break
	// 	}
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

}
