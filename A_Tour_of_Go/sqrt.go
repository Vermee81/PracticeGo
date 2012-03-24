package main

import(
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := float64(0)
    z = 1
    for i := 0; i < 11; i++{
        next_z := float64(0)
        next_z = z - (z * z - x) / (2 * z)
        if next_z == z{
            fmt.Println("break")
            fmt.Println(next_z)
            fmt.Println(z)
            break
        }else{
            z = next_z
        }
    }
    return z
}

func main(){
    fmt.Println("My Sqrt(2) = " , Sqrt(2))
    fmt.Println("math.Sqrt(2) = ", math.Sqrt(2))
}

