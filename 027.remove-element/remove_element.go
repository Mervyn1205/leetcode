package main
import "fmt"
func removeElement(nums []int, val int) int {
    i := 0
    for _, v := range nums {
        if val != v {
            nums[i] = v
            i++
        }
    }

    return i
}

func main() {
   arr1 := []int{0,1,2,2,3,0,4,2}
    _ =removeElement(arr1, 2)
    fmt.Println(arr1)
}