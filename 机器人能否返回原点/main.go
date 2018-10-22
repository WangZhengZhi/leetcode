package main

import "fmt"

/*机器人能否返回原点*/
func main() {
	var moves string
	moves = "URDL"
	fmt.Println(judgeCircle(moves))
}
func judgeCircle(moves string) bool {
	var numrow int = 0 //行
	var numcol int = 0 //列
	for i := 0; i < len(moves); i++ {
		switch string(moves[i]) {
		case "U":
			numrow++
		case "D":
			numrow--
		case "L":
			numcol++
		case "R":
			numcol--
		default:
			return true

		}
	}
	if numcol == 0 && numrow == 0 {
		return true
	} else {
		return false
	}
}
