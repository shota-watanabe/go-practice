package main

import (
	"flag"
	"fmt" // Goの標準パッケージ(文字の出力や生成に使われる頻出パッケージ)
)

func main() {
	flag.Parse()
	arg := flag.Arg(0) // コマンドライン引数を受け取り変数に入れる
	fmt.Printf("Hello %s\n", arg)
}
