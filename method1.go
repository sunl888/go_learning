package main

import "fmt"

const (
	WHITE  = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box //a slice of boxes

// 返回 box 的容量
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// 設置顏色
func (b *Box) SetColor(c Color) {
	b.color = c
}

// 顏色
func (bl BoxList) BiggestColor() Color {
	v := bl[0].Volume()
	k := Color(WHITE)
	for _, b := range bl {
		// 查找容量最大的
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PaintItBlack() {
	for i := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{1, 1, 1, RED},    // 1
		Box{2, 2, 2, YELLOW}, // 8
		Box{3, 3, 3, BLACK},  // 27
		Box{4, 4, 4, BLUE},   // 64
		Box{5, 5, 5, WHITE},  // 125
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())
	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}
