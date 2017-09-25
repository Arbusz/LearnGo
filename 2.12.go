package main
import (
    "fmt"
   // "math"
    "strconv"
   // "strings"
)

var tenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f"}

func main() {
    fmt.Println(decimalToAny(9999, 9))

}

// 10进制转任意进制
func decimalToAny(num, n int) string {
    new_num_str := ""
    var remainder int
    var remainder_string string
    for num != 0 {
        remainder = num % n
        if 17 > remainder && remainder > 9 {
            remainder_string = tenToAny[remainder]
        } else {
            remainder_string = strconv.Itoa(remainder)
        }
        new_num_str = remainder_string + new_num_str
        num = num / n
    }
    return new_num_str
}

