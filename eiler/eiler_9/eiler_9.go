package eiler_9

import "math"

// asignment limit: 1000
func Solution(limit int) (int, int, int, int){
    for a := 1; a < limit; a++{
        for b := a; b < limit - 2*a; b++{
            c := limit - a - b 
            sum_of_quad := math.Pow(float64(a), 2) + math.Pow(float64(b), 2)
            delta := sum_of_quad - math.Pow(float64(c), 2)
            if delta == 0{
                return a, b, c, a * b * c
            }
        }
    }
    return 0, 0, 0, 0
}

