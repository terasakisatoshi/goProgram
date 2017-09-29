/*
solve linear equation with respect to x,y,z,w
9x+2y+z+w=20
2x+8y-2z+w=16
-x-2y+7z-2w=8
x-y-2z+6w=17
*/

package main

import (
    "fmt"
    "math"
)

const N = 4

func jacobi(mat [N][N]float64, vec [N]float64, xs [N]float64) [N]float64 {
    for {
        xs_new := [N]float64{0, 0, 0, 0}
        finish := true
        for i := 0; i < N; i++ {
            xs_new[i] = vec[i]
            for j := 0; j < N; j++ {
                if j != i {
                    xs_new[i] -= mat[i][j] * xs[j]
                }
            }

            xs_new[i] /= mat[i][i]
            if math.Abs(xs_new[i]-xs[i]) > 10e-7 {
                finish = false
            }
        }

        xs = xs_new
        if finish {
            break
        }
    }
    return xs
}

func main() {
    var mat = [N][N]float64{
        {9, 2, 1, 1},
        {2, 8, -2, 1},
        {-1, -2, 7, -2},
        {1, -1, -2, 6}}

    var vec = [N]float64{20, 16, 8, 17}
    var init = [N]float64{0, 0, 0, 0}

    var solution = jacobi(mat, vec, init)
    //print result
    for _, value := range solution {
        fmt.Printf("%f\t", value)
    }
}
