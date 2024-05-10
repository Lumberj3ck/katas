package eiler_7

import "math"
// import "fmt"

func is_prime(num int) bool{
    for i := 2; i <= int(math.Sqrt(float64(num))); i++{
        if num % i == 0{
            return false
        }
    } 
    // fmt.Println(num)
    return true 
}


func Solution(n int) int{
    prime_counter := 1
    last_prime := 2
    for num := 3; prime_counter <= n; num++{
       if is_prime(num){
           prime_counter++
           last_prime = num
       } 
    }
    return last_prime
}

// func main(){
//     res := solution(10001)
//     fmt.Println(res)
// }
