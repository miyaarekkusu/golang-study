# ここでは、Golangについて覚えたことをメモします。

## フォルダ構成

- `cmd/tutorial1/` : エラーハンドリングや基本文法を学んだチュートリアル
- `arrays-maps/` : 配列・マップを学ぶためのフォルダ（まだ未実装・空ファイル）

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
