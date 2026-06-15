# ここでは、Golangについて覚えたことをメモします。

## フォルダ構成

- `cmd/tutorial1/` : エラーハンドリングや基本文法を学んだチュートリアル
- `arrays-maps/` : 配列・マップを学ぶためのフォルダ（まだ未実装・空ファイル）
- `goroutines/` : ゴルーチン（並行処理）、`sync.WaitGroup`、`sync.RWMutex`を学ぶためのフォルダ

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
