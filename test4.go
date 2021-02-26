package main
import "fmt"
type person struct {
	name   string // 姓名
	age    int    // 年龄
	gender string // 性别
}
// dove 鸽子
type dove interface {
	gugugu() // 鸽
}
// repeater 复读机
type repeater interface {
	repeat(string) // 复读
}
// repeat 复读
func (p *person) repeat(word string) {
	fmt.Println(word)
}
func (p *person) gugugu() {
	fmt.Println(p.name, "又鸽了")
}
func main() {
	p := &person{
		name: "lcm",
		age: 18,
		gender: "male",
	}
	p.gugugu()
	var r repeater
	r = p
	r.repeat("helloworld")
}