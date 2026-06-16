# ここでは、Golangについて覚えたことをメモします。

## フォルダ構成

- `cmd/tutorial1/` : エラーハンドリングや基本文法を学んだチュートリアル
- `arrays-maps/` : 配列・スライス・マップ・forループを学ぶためのフォルダ
- `structs/` : struct・メソッド・interfaceを学ぶためのフォルダ
- `goroutines/` : ゴルーチン（並行処理）、`sync.WaitGroup`、`sync.RWMutex`を学ぶためのフォルダ
- `channels/` : チャンネルと`select`を学ぶためのフォルダ

## 学習内容

### 1. パッケージのインポート
- `errors` パッケージをインポートし、エラーオブジェクトを生成する
- `fmt` パッケージで標準出力（`Println`, `Printf`）を行う

### 2. 変数宣言
- `var 変数名 型 = 値` の形式で明示的に型を指定して宣言できる
  ```go
  var printValue string = "Hello World"
  var numerator int = 11
  ```

### 3. 関数と複数戻り値
- Goの関数は複数の値を返せる（例: `(int, int, error)`）
  ```go
  func intDivision(numerator int, denominator int) (int, int, error) {
      ...
      return result, remainder, err
  }
  ```
- 呼び出し側は `var result, remainder, err = intDivision(...)` のように複数変数で受け取る

### 4. エラーハンドリング
- `errors.New("メッセージ")` でエラーを作成する
- エラーは戻り値の最後に `error` 型で返すのがGoの慣習
- **早期リターン（early return）パターン**
  ```go
  if err != nil {
      fmt.Println(err.Error())
      return // エラー時はここで処理を終了し、下の処理に進ませない
  }
  ```
  → 先にエラーチェックを行い、エラーがあればすぐ関数を抜けることでネストを浅く保てる

### 5. 条件分岐 (if)
- 0除算をチェックし、`denominator == 0` の場合はエラーを返す
- 割り算の余り（remainder）が0かどうかで出力メッセージを分岐する

### 6. switch文
- `switch remainder { case 0: ... case 1, 2: ... default: ... }`
- `case` に複数の値（`1, 2`）をまとめて指定できる
- 各 `case` は自動的にbreakされる（Cのようなフォールスルーはしない）

### 7. 文字列フォーマット
- `fmt.Println(value)` : 改行付きで値を出力する
- `fmt.Printf("...%v...\n", value)` : `%v` を使って値を埋め込んでフォーマット出力する

### 8. ゴルーチンと並行処理 (goroutines / concurrency)

#### goroutineの役割と仕組み
- goroutineはGoランタイムが管理する軽量な実行単位（OSスレッドより軽く、小さいスタックから必要に応じて伸長する）
- `go 関数名()` と書くだけでその関数を新しいgoroutineとして非同期に起動できる（呼び出し元は終了を待たずに次へ進む）
- Goランタイム内のスケジューラが、多数のgoroutineを少数のOSスレッドに割り振って実行する（M個のスレッドにN個のgoroutineを割り当てるM:N方式）
- concurrency（並行）とparallelism（並列）は別物
  - 並行：複数のタスクを切り替えながら同時に「進行」させること（CPU1コアでも可能）
  - 並列：実際に複数CPUコアで同時に「実行」すること
  - GOMAXPROCSの数までは並列実行も行われ、それを超える分は並行に切り替えながら処理される

#### Mutexの使い方
- 複数のgoroutineが同じ変数（共有データ）に同時に読み書きすると、データ競合（race condition）が発生する
- `sync.Mutex` : `Lock()`で他のgoroutineをブロックし、`Unlock()`で解放する。ロック中はその区間に1つのgoroutineしか入れない
- `sync.RWMutex` : 読み書き用のロック
  - 書き込み時：`Lock()` / `Unlock()`（排他、1つだけ）
  - 読み込み時：`RLock()` / `RUnlock()`（複数goroutineが同時に読める。書き込み中は待たされる）
- 例（goroutines/main.go の `save()` / `log()`、現在はコメントアウト）
  - `save()`：`m.Lock()` → `results = append(...)` → `m.Unlock()` で書き込みを保護
  - `log()`：`m.RLock()` → 読み取り → `m.RUnlock()` で読み取り中の安全性を確保

#### sync.WaitGroup
- `wg.Add(1)` で待つ数を+1、ゴルーチン内で `wg.Done()` を呼ぶと-1
- `wg.Wait()` でカウントが0になるまでブロックする

#### 実行結果（count()を5つゴルーチンで並行実行、各回1億回ループ）
- 合計実行時間：約93ms
- → 1つずつ順番に実行するより大幅に速い＝並行に動いている証拠

### 9. 配列・スライス・マップ (arrays-maps/main.go)

#### 配列 (array)
- `[...]型{値, ...}` で要素数を自動推論した固定長配列を作成する
- サイズは型の一部なので `[3]int32` と `[4]int32` は別の型

#### スライス (slice)
- `[]型{値, ...}` で作る可変長のシーケンス。内部では配列への（ポインタ・長さ・容量）を持つ
- `len(s)` : 現在の要素数、`cap(s)` : 内部配列の容量
- `append(s, 値)` で末尾に追加。容量不足時は内部配列を新たに確保してコピー（cap が約2倍になる）
- `make([]型, 長さ, 容量)` で長さと容量を指定して生成できる

#### マップ (map)
- `map[キー型]値型` で定義し、`make(map[キー型]値型)` または `map[キー型]値型{...}` で初期化する
- `v, ok := m[key]` で値取得。`ok` が `false` ならキーが存在しない（ゼロ値を返す）
- `for k, v := range m {}` で全要素を反復できる（順序は保証されない）

#### forループ
- `for i, v := range コレクション {}` : 配列・スライスはインデックスと値、マップはキーと値を取得
- `for i := 0; i < n; i++ {}` : C言語と同じ3式のforループ
- `for { if 条件 { break } }` : 条件のみのwhile相当の書き方

### 10. struct・メソッド・interface (structs/main.go)

#### struct
- 複数フィールドをまとめた新しい型を定義する（他言語のクラスのデータ部分に近い）
  ```go
  type gasEngine struct {
      mpg     uint8
      gallons uint8
  }
  var e gasEngine = gasEngine{25, 15}
  ```
- フィールドの型に別のstructを指定することもできる

#### メソッド（値レシーバ）
- `func (レシーバ 型) メソッド名() 戻り値型` の形で、struct にメソッドを紐づける
- 値レシーバはstructのコピーが渡されるので、メソッド内でフィールドを変更しても元に影響しない
  ```go
  func (e gasEngine) milesLeft() uint8 {
      return e.gallons * e.mpg
  }
  ```

#### interface
- 「どんなメソッドを持つべきか」というシグネチャだけを定義した型
- メソッドのシグネチャを満たしていれば、どんなstructでも自動的にinterfaceを実装したとみなされる（明示的な `implements` 不要）
  ```go
  type engine interface {
      milesLeft() uint8
  }
  func canMakeit(e engine, miles uint8) { ... }
  // gasEngine も eletricEngine も engine を満たすのでそのまま渡せる
  ```

### 11. チャンネル (channels/main.go)

#### チャンネルとは
- goroutine同士が安全にデータをやり取りするためのパイプ
- 送る側・受け取る側が同期するのでmutex不要でデータ競合を防げる

#### 基本的な使い方
```go
c := make(chan int)       // バッファなし（同期）
c := make(chan int, N)    // バッファN個（N個まで貯められる）
c <- 値                   // 送信（受け取り側が準備できるまでブロック）
v := <-c                  // 受信（送り側が送るまでブロック）
close(c)                  // チャンネルを閉じる（rangeループを終わらせるために必要）
for v := range c {}       // close()されるまで受信し続ける
```

#### select
- 複数チャンネルを同時に待ち、最初に届いた一つだけ処理する
  ```go
  select {
  case v := <-ch1:
      // ch1 から受け取った
  case v := <-ch2:
      // ch2 から受け取った（どちらか早い方だけ実行される）
  }
  ```
- `close()` を忘れると `range` ループがブロックし続けてデッドロックになる

### 12. ジェネリクス (Generics)

#### ジェネリクスとは
- 型パラメータを使って「型を問わず使える汎用的なコード」を書く仕組み（Go 1.18〜）
- 同じロジックを `int` 版・`float64` 版・`string` 版…と何度も書かなくて済む

#### 型パラメータの構文
```go
func 関数名[T 制約](引数 T) T { ... }
```

#### 型制約 (constraint)
- 型パラメータに「どんな型が使えるか」を制約で指定する
- `any` : どんな型でも良い（`interface{}` の別名）
- `comparable` : `==` / `!=` で比較できる型（マップのキーに使える）
- 型セットで独自制約を定義できる
  ```go
  type Number interface {
      int | int32 | int64 | float32 | float64
  }
  ```

#### 汎用関数の例
```go
// どんな数値型でも合計を求める
func sumSlice[T Number](slice []T) T {
    var sum T
    for _, v := range slice {
        sum += v
    }
    return sum
}
sumSlice([]int{1, 2, 3})         // 6
sumSlice([]float64{1.1, 2.2})   // 3.3
```

#### 型を返す汎用関数（空かどうかチェック）
```go
func isEmpty[T any](slice []T) bool {
    return len(slice) == 0
}
```

#### ジェネリックな struct
- struct にも型パラメータを付けられる
  ```go
  type Stack[T any] struct {
      items []T
  }
  func (s *Stack[T]) Push(item T) { s.items = append(s.items, item) }
  func (s *Stack[T]) Pop() T {
      item := s.items[len(s.items)-1]
      s.items = s.items[:len(s.items)-1]
      return item
  }
  // 使い方
  var intStack Stack[int]
  intStack.Push(1)
  intStack.Push(2)
  fmt.Println(intStack.Pop()) // 2
  ```

#### 使いどころ
- 同じアルゴリズムを複数の型で使いたいとき（コレクション操作・数値計算など）
- interfaceでは型情報が消えてしまう場合に型安全を保ちたいとき
- 単純なケースは `any + 型アサーション` でも書けるが、ジェネリクスの方が型安全でコンパイル時にエラーを検出できる
