package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    //create map its key is string and value is int
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    //usage to break for you should input Ctrl+D
    for input.Scan() {
        line := input.Text()
        counts[line] += 1
    }

    for key, value := range counts {
        if value > 1 {
            fmt.Printf("%d\t%s\n", value, key)
        }
    }
}
