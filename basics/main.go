package main

func main() {
	fmt.Println(multiarr(3, 3)
}


func multiarr(xy, arr2 int) ([3][3]int) {
	xy := make([][]int, 3)
		for i := 0; i < 3; i++ {
			innerLen := i + 1
			xy[i] = make([]int, innerLen)
			for j := 0; j < innerLen; j++ {
				xy[i][j] = i + j
			}
		}
