package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RoomInfo struct {

}

func getHeader() http.Header {
	var header = http.Header{}
	header.Set("Pragma", "no-cache")
	header.Set("Origin", "http://jx.kuwo.cn")
	header.Set("Accept-Encoding", "gzip, deflate")
	header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7",)
	header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	header.Set("Upgrade", "websocket")
	header.Set("Sec-WebSocket-Extensions", "permessage-deflate; client_max_window_bits")
	header.Set("Cache-Control", "no-cache")
	header.Set("Connection", "Upgrade")
	header.Set("Sec-WebSocket-Version", "13")
	return header
}

func getRoomInfo(rid string)  {
	// rid: 683593
	uri := fmt.Sprintf("http://zhiboserver.kuwo.cn/proxy.p?src=web&cmd=enterroom&rid=%s&auto=1", rid)
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout:timeout,
	}
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	js, err := json.Marshal(string(body))
	fmt.Println(js)
}

func main() {
	//url  := "ws://60.29.226.18:18133/"
	//dialer := &websocket.Dialer{
	//	NetDial: func(network, addr string) (net.Conn, error) {
	//		return net.DialTimeout(network, addr, 10*time.Second)
	//	},
	//}
	//conn, _, err := dialer.Dial(url, getHeader())
	//if err != nil {
	//	fmt.Printf("dial to server %s failed, %v", url, err)
	//}
	getRoomInfo("683593")
}
