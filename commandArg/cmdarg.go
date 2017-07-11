//コマンドライン引数を表示する．

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	/*明示的に値がassignされていないが，この場合，
	型が持っているzero valueに初期化される．
	string であれば""すなわち，空の文字列．
	*/
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
