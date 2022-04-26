import "fmt"

func main() {
	var arr = [5]int{3, 4, 5, 1, 6}
	Largest(arr)
}

func Largest(arr [5]int) {
	length := len(arr)
	max := arr[0]
	for i := 0; i < length; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Printf("Largest Element:%d\n", max)

}
