// =====================================================
// Goチャンネル (Channel) メモ
// =====================================================
// 【役割】
//   goroutine同士が安全にデータをやり取りするためのパイプ。
//   「送る側」と「受け取る側」が同期するので、mutex不要でデータ競合を防げる。
//
// 【基本的な使い方】
//   作成:  c := make(chan 型)          // 同期チャンネル（バッファなし）
//          c := make(chan 型, N)       // バッファあり（N個まで貯められる）
//   送信:  c <- 値                     // 受け取り側が準備できるまでブロック
//   受信:  v := <-c                    // 送り側が送るまでブロック
//   閉じる: close(c)                   // rangeループを終わらせるために必要
//   range: for v := range c {}        // close()されるまで受信し続ける
//
// 【select】
//   複数チャンネルを同時に待ち、最初に届いたものだけ処理する。
//   case website := <-ch1:  // ch1 か ch2 のどちらか早い方を受け取る
//   case website := <-ch2:
// =====================================================

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	// --- 基本例（コメントアウト済み）---
	// var c = make(chan int)
	// go process(c)
	// for i := range c {     // close(c)されるまでループ
	//     fmt.Println(i)
	//     time.Sleep(time.Second * 1)
	// }

	// func process(c chan int) {
	//     defer close(c)         // 忘れるとrangeがブロックし続けてデッドロック
	//     for i := 0; i < 5; i++ {
	//         c <- i
	//     }
	// }

	// --- 実践例: 複数サイトを並列チェックし、条件を満たした最初の結果を使う ---
	var chickenChannel = make(chan string) // チキンの安値を見つけたサイト名を送る
	var tofuChannel = make(chan string)    // 豆腐の安値を見つけたサイト名を送る
	var websites = []string{"wallmart.com", "costco.com", "wholesfood.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel) // 各サイトをgoroutineで並列監視
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel) // どちらか先に見つかった方を通知
}

// 1秒ごとにランダムな価格を生成し、上限以下になったらサイト名をチャンネルへ送信
func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website // 条件達成 → チャンネルに送ってループ終了
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

// selectで2つのチャンネルを同時に待ち、最初に届いた方だけ処理する
func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nテキストを送信しました。相当値段のチキンは%vにあります", website)
	case website := <-tofuChannel:
		fmt.Printf("\nメールを送信しました。相当値段の豆腐は%vにあります", website)
	}
}
