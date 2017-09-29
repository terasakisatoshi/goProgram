//reference gopl.io ch6
package main

import (
    . "./geometry"
    "fmt"
)

func main() {
    p := Point{0, 0}
    q := Point{2, 1}
    r := Point{2, 0}
    fmt.Println(p, q)
    fmt.Println(p.Distance(q))

    path := Path{p, q, r, p}
    fmt.Println(path.Distance())

    var pt = &Point{0, 0}
    pt.Shift(3)
    fmt.Println(*pt)

    var qt = Point{0, 0}
    (&qt).Shift(3)
    fmt.Println(qt)
}
