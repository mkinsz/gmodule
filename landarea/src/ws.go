package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"encoding/json"
)

var addr = flag.String("addr", "localhost:56789", "http service address")

//// 我们需要定义一个 Upgrader
//// 它需要定义 ReadBufferSize 和 WriteBufferSize
//var upgrade = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//
//	// 可以用来检查连接的来源
//	// 这将允许从我们的 React 服务向这里发出请求。
//	// 现在，我们可以不需要检查并运行任何连接
//	CheckOrigin: func(r *http.Request) bool { return true },
//}

var upgrade = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		var msg string
		err = json.Unmarshal([]byte(message), msg)
		if err != nil {
			log.Println("json:", err)
		}

		fmt.Println(mt)
		log.Println(websocket.BinaryMessage, websocket.TextMessage)

		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprintf(w, "parse template error: %s", err.Error())
		return
	}
	t.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}