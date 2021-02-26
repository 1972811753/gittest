package main
import "fmt"
type person struct {
	name   string // 姓名
	age    int    // 年龄
	gender string // 性别
}
type dove interface {
	gugugu() // 鸽
}
// repeater 复读机
type repeater interface {
	repeat(string) // 复读
}
type lemon interface{
	sour() //柠檬精
}
// sweet_strange 真香怪
type sweet_strange interface{
	fragrant() //真香怪
}
// repeat 复读
func (p *person) repeat(word string) {
	fmt.Println(word)
}
func (p *person) gugugu() {
	fmt.Println(p.name, "又鸽了")
}
func (p *person) sour() {
	fmt.Println(p.name, "酸了")
}
func (p *person) fragrant() {
	fmt.Println(p.name, "真香")
}
func main() {
	p := &person{
		name: "lcm",
		age: 18,
		gender: "male",
	}
	p.gugugu()
	p.sour()
	p.fragrant()
	var r repeater
	r = p
	r.repeat("helloworld")
}