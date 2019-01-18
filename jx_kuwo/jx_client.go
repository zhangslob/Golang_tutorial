package main

import (
	"fmt"
	"github.com/henrylee2cn/pholcus/common/simplejson"
	"github.com/zaoshu/jx_kuwo/bot"
	"log"
)

var (
	baseSingerUrl = "http://zhiboserver.kuwo.cn/proxy.p?src=android_jx&cmd=gethallmobile3&type=1&page=%d"
)

func getLiveRoom()  {
	b, err := bot.HTTPGet(fmt.Sprintf(baseSingerUrl, 1), bot.DefaultHTTPHeaders)
	if err != nil {
		err = fmt.Errorf("http get failed, %v", err)
		return
	}
	parseRoom(b)

	js, _ := simplejson.NewJson(b)
	pageCount, _ := js.Get("roomlist").GetIndex(0).Get("singercnt").Int()
	page := pageCount / 30 + 2
	fmt.Println(page)

	for i:= 1; i < page; i++ {
		b, err := bot.HTTPGet(fmt.Sprintf(baseSingerUrl, i), bot.DefaultHTTPHeaders)
		if err != nil {
			err = fmt.Errorf("http get failed, %v", err)
			return
		}
		parseRoom(b)
	}
}

func parseRoom(body []byte)  {
	defer func() {
		if err:=recover();err!=nil{
			log.Println("")
				}
	}()
	js, _ := simplejson.NewJson(body)

	roomList := js.Get("roomlist")

	for i := range roomList.MustArray() {
		index := roomList.GetIndex(i)
		singerList := index.Get("singerlist")
		for j := range singerList.MustArray() {
			fmt.Println(singerList.GetIndex(j).Get("id").MustInt())
		}
	}
}

func main() {
	getLiveRoom()
}