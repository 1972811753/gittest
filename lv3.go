package main
import (
	"fmt"
)
type Author struct {
	Name string //名字
	VIP bool //是否是高贵的带会员
	Icon string //头像
	Signature string //签名
	Focus int //关注人数
}
type video struct{
	Author //作者
	url string //地址
	video_name string //视频名字
	cover_picture string //封面图片
	likes int //点赞数
	transponds int //转发数
    coins int //投币数
	collects int //收藏数
}
func (v *video) like(){
	v.likes++
	fmt.Println("已点赞")
}
func (v *video) coin(){
	v.coins++
	fmt.Println("已投币")
}
func (v *video) collect(){
	v.collects++
	fmt.Println("已收藏")
}
func (v *video) Three_even_akey(){
	v.likes++
	v.coins++
	v.collects++
	fmt.Println("已一键三连")
}
func publish(name string, vname string) *video {
	v:=&video{
		Author: Author{
			Name:      name,
			VIP:        true,
			Icon:      "",
			Signature: "",
			Focus:     0,
		},
		video_name:    vname,
		cover_picture: "",
		likes:         0,
		transponds:    0,
		coins:         0,
		collects:      0,
	}
	fmt.Println("发布视频成功")
	return v

}
func main(){
	v := &video{
		Author: Author{
			Name:      "lcm",
			VIP:       true,
			Icon:      "",
			Signature: "",
			Focus:     2000,
		},
		video_name:    "",
		cover_picture: "",
		likes:         999,
		transponds:    999,
		coins:         99,
		collects:      888,
	}
	v.like()
	v.coin()
	v.collect()
	v.Three_even_akey()
	fmt.Println(publish("卢靖","论如何学习go语言"))
}