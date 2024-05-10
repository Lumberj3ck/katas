package eiler_8


import "strconv"
import "fmt"

var numbers string = "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"

func get_digit(str string) int{
    digit, _ := strconv.Atoi(str) 
    return digit 
}

func wrong_solution(length int) (max_start_ind, max_last_ind int){
    curr_index := 0
    last_index := curr_index + length - 1
    product := 1
    max_product := 1
    for last_index < len(numbers){
        if curr_index == 0 {
            for _, num := range numbers[curr_index: last_index]{
                product *= get_digit(string(num))
            }
            fmt.Println(product)
        } else{
            digit := get_digit(string(numbers[curr_index - 1])) 
            last_digit := get_digit(string(numbers[last_index])) 
            if digit > 0 {
                product /= digit
            }
            product *= last_digit 
        }
        // fmt.Println(product, curr_index, last_index)
        if product > max_product{
            max_product = product
            max_start_ind = curr_index
            max_last_ind = last_index 
        } 
        curr_index += 1
        last_index = curr_index + length - 1
    }
    return max_start_ind, max_last_ind
}

func Solution(length int) (max_product int){
    for curr_index := 0; curr_index + length < len(numbers); curr_index++{
        last_index := curr_index + length 
        var product int = 1 
        for _, num := range numbers[curr_index: last_index]{
            product *= get_digit(string(num))
        }
        max_product = max(max_product, product)
    }
    return 
}


// func main(){
//     fmt.Println(Solution(13))
// }
