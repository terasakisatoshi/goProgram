package main

/*
main package はライブラリではなく単独で動作する実行可能hなプログラムを定義します．
*/

//fmtは基本的な出力関数の一つである．
import "fmt"

/*
main関数も特別であり，この関数からプログラムの実行が開始される．
何であれmain関数が行うことがそのプログラムが行うことである．
*/
func main() {
	fmt.Println("Hello, World 世界")
}
