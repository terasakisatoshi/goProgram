package main

import (
    "fmt"
    "os"
)

func main() {
    //省略変数宣言
    s, sep := "", ""
    //used blank identifier
    //Goでは使われないローカル変数の存在を許さない．
    //この場合ブランク識別子 _ を用いて解決する．
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}

//省略変数宣言の仕方
//下記はどれも同じ．
/*
s:=""
var s string
var s = ""
var s string = ""
*/
