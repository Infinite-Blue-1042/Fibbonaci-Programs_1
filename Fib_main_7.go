package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func main() {
	userInput := bufio.NewScanner(os.Stdin)
	for userInput.Scan()  {
		listNumber, _ := strconv.Atoi(userInput.Text())
		var fib [] int
		for i := 0; i < listNumber; i++ {
			if i <= 1 {
				fib = append(fib, i)
			} else {
				number := fib[i-1] + fib[i-2]
				fib = append(fib, number)
			}
		}
		fmt.Println(fib)
	}
}
