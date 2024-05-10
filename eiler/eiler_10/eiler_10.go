package eiler_10

import "math"

func is_prime(num int) bool{
    for i := 2; i <= int(math.Sqrt(float64(num))); i++{
        if num % i == 0{
            return false
        }
    } 
    return true 
}

// asignment limit 2000_000
func Solution(limit int) int{
    prime_sum := 0
    for num := 2; num < limit; num++{
        if is_prime(num){
                prime_sum += num
        }
    }
    return prime_sum
}

