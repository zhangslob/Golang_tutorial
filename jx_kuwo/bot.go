package main

import (
	"github.com/zaoshu/huya/bot"
	"github.com/zaoshu/huya/bot/proto"
	"golang.org/x/net/websocket"
	"sync"
	"time"
)

type Bot struct {
	LogData			bool
	LogPrefix		string
	Event			bot.Event
	PingInterval	time.Duration
	LiveInfo		bot.LiveInfo
	UserID			proto.HUYAUserId
	WSCmdHandler	bot.WSCmdHandler
	AutoClose		bool

	coon	*websocket.Conn

	sendCh	chan []byte
	event	chan func()
	stopCh	chan bool

	stop				bool
	living				bool
	lastActivityTime	time.Time
	lock 				sync.Mutex

	wg sync.WaitGroup
}

func NewBot(info bot.LiveInfo) *Bot {
	return &Bot{
		LiveInfo: info,
		UserID: proto.HUYAUserId{
			HuYaUA: proto.HuyaUA,
		},
		sendCh:				make(chan []byte, 64),
		event:				make(chan func(), 512),
		stopCh:				make(chan bool),
		living:				info.TopSid > 0,
		lastActivityTime:	time.Now(),
		PingInterval: 		30 * time.Second,
	}
}


func main() {
	
}
