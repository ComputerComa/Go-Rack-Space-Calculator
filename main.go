package main

import (
	"fmt"

	tsize "github.com/kopoli/go-terminal-size"
)

var cornerCharTopLeft = "╔"
var cornerCharTopRight = "╗"
var cornerCharBottomLeft = "╚"
var cornerCharBottomRight = "╝"
var lineCharHorizontal = "═"
var lineCharVertical = "║"
var joinCharLeft = "╠"
var joinCharRight = "╣"
var uCount = 8
func MakeEmptyStr(len int) string {
	out := ""
	for i := 0; i < len; i++ {
		out += " "
	}
	return out
}
func middleLine(len int) string {
	out := ""
	for i := 0; i < len; i++ {
		out += lineCharHorizontal
	}
	return out
}

func DrawTop(len int) {
	out := ""
	mid := middleLine(len - 2)
	out = cornerCharTopLeft + mid + cornerCharTopRight
	println(out)
}
func DrawMiddles(len int) {
	out := ""
	mid := middleLine(len - 2)
	out = joinCharLeft + mid + joinCharRight
	println(out)

}

func padText(str string, len int) {
	out := ""
	out += lineCharVertical
	out += centerString(str,len -2)
	out += lineCharVertical
}

func DrawBottom(len int) {
	out := ""
	mid := middleLine(len - 2)
	out = cornerCharBottomLeft + mid + cornerCharBottomRight
	println(out)
}

func drawTopBox(str string){

}

function DrawBlock(idex int, str string){
	switch idez {
	case 0:
		drawTopBox(str)
	case < uCount -1:
		
		
	}
}

func centerString(str string, Width int) string {
	len := len(str)
	diff := Width - len
	fmt.Println(diff)
	centerSize := ((Width - len) / 2)
	fmt.Println(centerSize)
	padding := MakeEmptyStr(centerSize)
	outStr := padding + str + padding

	return outStr
}

func main() {

	fmt.Println("Rack Generator created by Walker Dick")
	var s tsize.Size
	s, err := tsize.GetSize()
	if err == nil {
		width := s.Width
		if width > 400 {
			width = 400
		}
		fmt.Println(s.Width)
		fmt.Println(width)
		test := centerString("Test", width)
		fmt.Println(test)
	}

}
