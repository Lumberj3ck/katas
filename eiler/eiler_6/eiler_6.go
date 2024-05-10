package eiler_6

import "math"
// import "fmt"

func Solution(n Element) int{
    sum_of_quad := float64(0)
    sum_of_numbers := Element(0)
    for i := Element(1); i <= n; i++{
        sum_of_quad += math.Pow(float64(i), 2)  
        sum_of_numbers += i
    }
    quad_of_sum := math.Pow(float64(sum_of_numbers), 2)
    return int(quad_of_sum - sum_of_quad)
}

// func main(){
//     fmt.Println(solution(100))
// }
