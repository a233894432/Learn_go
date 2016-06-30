//package main

//import (
//	"fmt"
//	"reflect" //用于查看数据类型

//	//	"mymath"
//)

//func main() {
//	//	fmt.Printf("Hello, wssorld.  Sqrt(2) = %v\n", mymath.Sqrt(2))
//	//	test1()
//	//	test2()
//	//	test3()
//	//	test4()
//	//	test5()
//	//	test6()
//	test7()
//}

//func test1() {
//	sum := 1
//	for sum < 10 {
//		sum += sum
//	}
//	fmt.Printf("%v\n", sum)
//}

//func test2() {
//	for index := 10; index > 0; index-- {
//		if index == 2 {
//			//			break // 或者continue
//			continue
//		}
//		fmt.Println(index)
//	}

//}

//func test3() {
//	i := 10
//	switch i {
//	case 1:
//		fmt.Println("i is equal to 1")
//	case 2, 3, 4:
//		fmt.Println("i is equal to 2, 3 or 4")
//	case 10:
//		fmt.Println("i is equal to 10")
//	default:
//		fmt.Println("All I know is that i is an integer")
//	}
//}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b

//}

//func test4() {

//	x := 3
//	y := 4
//	z := 5

//	max_xy := max(x, y) //调用函数max(x, y)
//	max_xz := max(x, z) //调用函数max(x, z)

//	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
//	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
//	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // 也可在这直接调用它

//}

////返回 A+B 和A*B
//func SumAndProduct(a, b int) (int, int) {
//	return a + b, a * b
//}

//func test5() {
//	x := 3
//	y := 4
//	xPlus, xT := SumAndProduct(x, y)
//	fmt.Printf("%d + %d = %d\n", x, y, xPlus)
//	fmt.Printf("%d * %d = %d\n", x, y, xT)
//	fmt.Println("type:", reflect.TypeOf(x)) //查看数据类型

//}

//type person struct {
//	name string
//	age  int
//}

//func test6() {
//	var P person

//	P.name = "diogo"
//	P.age = 25

//	fmt.Printf("The person's name is %s  %v \n", P.name, P.age)
//}

//type Skills []string

//type Human struct {
//	name   string
//	age    int
//	weight int
//}

//type Student struct {
//	Human          // 匿名字段，struct
//	Skills         // 匿名字段，自定义的类型string slice
//	num        int // 内置类型作为匿名字段
//	speciality string
//}

//func test7() {
//	// 初始化学生Jane
//	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
//	// 现在我们来访问相应的字段
//	fmt.Println("Her name is ", jane.name)
//	fmt.Println("Her age is ", jane.age)
//	fmt.Println("Her weight is ", jane.weight)
//	fmt.Println("Her speciality is ", jane.speciality)
//	// 我们来修改他的skill技能字段
//	jane.Skills = []string{"anatomy"}
//	fmt.Println("Her skills are ", jane.Skills)
//	fmt.Println("She acquired two new ones ")
//	jane.Skills = append(jane.Skills, "physics", "golang")
//	fmt.Println("Her skills now are ", jane.Skills)
//	// 修改匿名内置类型字段
//	jane.num = 3
//	fmt.Println("Her preferred number is", jane.num)
//}

//----
//package main

//import "fmt"

//const (
//	WHITE = iota
//	BLACK
//	BLUE
//	RED
//	YELLOW
//)

//type Color byte

//type Box struct {
//	width, height, depth float64
//	color                Color
//}

//type BoxList []Box //a slice of boxes

//func (b Box) Volume() float64 {
//	return b.width * b.height * b.depth
//}

//func (b *Box) SetColor(c Color) {
//	b.color = c
//}

//func (bl BoxList) BiggestColor() Color {
//	v := 0.00
//	k := Color(WHITE)
//	for _, b := range bl {
//		if bv := b.Volume(); bv > v {
//			v = bv
//			k = b.color
//		}
//	}
//	return k
//}

//func (bl BoxList) PaintItBlack() {
//	for i, _ := range bl {
//		bl[i].SetColor(BLACK)
//	}
//}

//func (c Color) String() string {
//	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
//	return strings[c]
//}

//func main() {

//	boxes := BoxList{
//		Box{4, 4, 4, RED},
//		Box{10, 10, 1, YELLOW},
//		Box{1, 1, 20, BLACK},
//		Box{10, 10, 1, BLUE},
//		Box{10, 30, 1, WHITE},
//		Box{20, 20, 20, YELLOW},
//	}

//	fmt.Printf("We have %d boxes in our set\n", len(boxes))
//	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
//	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
//	fmt.Println("The biggest one is", boxes.BiggestColor().String())

//	fmt.Println("Let's paint them all black")
//	boxes.PaintItBlack()
//	fmt.Println("The color of the second one is", boxes[1].color.String())

//	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
//}

package main

import "fmt"

type Rect struct {
}
