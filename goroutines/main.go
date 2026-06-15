package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// sync.RWMutex：複数ゴルーチンから共有データ(results)に安全にアクセスするためのロック
// Lock()/Unlock()：書き込み用の排他ロック、RLock()/RUnlock()：読み込み用の共有ロック
var m = sync.RWMutex{}

// sync.WaitGroup：起動したゴルーチンが全部終わるまでメインを待たせる仕組み
// Add(1)で待つ数を+1、Done()で-1、Wait()でカウントが0になるまでブロックする
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		// go dbCall(i)
		go count()
	}
	wg.Wait()
	fmt.Printf("\n合計実行時間：%v", time.Since(t0))
	fmt.Printf("\n結果は：%v", results)
}

// 実行結果メモ：
// 合計実行時間：約93ms
// 結果は：[]
// → count()を5つゴルーチンで並行実行（各1億回ループ）しても、
//   1つずつ順番に実行するより大幅に速い＝並行に動いている証拠
// → count()はresultsに値を追加しないので、結果は空のまま

func dbCall(i int) {
	// データベース呼び出し遅延をシミュレーション
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// m.Lock()
	// fmt.Printf("データベース呼び出しの結果は：%v", dbData[i])
	// results = append(results, dbData[i])
	// m.Unlock()
	wg.Done()
}

// func save(result string) {
// 	m.Lock()
// 	results = append(results, result)
// 	m.Unlock()
// }

// func log() {
// 	m.RLock()
// 	fmt.Printf("\n現在の結果は：%v", results)
// 	m.RUnlock()
// }

// count()：1億回の加算をするだけのゴルーチン
// dbCallの代わりに使い、並行実行による速度差を確認するためのもの
func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res += 1
	}
	wg.Done()
}
