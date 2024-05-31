package WebSocketServer

import (
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"

	WebSocketPacket "MultiRPStoGolang/ProtoBuf/ProtoBufPacket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 웹소켓 시작하는 부분
func Init() {

	http.HandleFunc("/", HandleConnection)

	serverAddr := ":48059"
	log.Printf("WebSocket server listening on %s", serverAddr)

	server := &http.Server{
		Addr:    serverAddr,
		Handler: nil,
	}

	log.Fatal(server.ListenAndServe())
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // HTTP 연결을 웹소켓 연결로 업그레이드
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	// 클라이언트의 IP 주소와 포트 번호를 로그로 출력
	clientAddr := conn.RemoteAddr().String()
	color.Cyan("Connection established with %s", clientAddr)
	// Seq를 자신의 local로 저장
	var clientSeq int32
	clientSeq = -1
	// UUID는 payload.UserUUID 에 존재

	defer func() {
		conn.Close() // 웹소켓을 닫아버림 (return)

		color.Red("하나의 Connection 끊김 : %v", clientAddr)
	}()

	for { // 무한 루프를 통해 클라이언트로부터 메시지를 지속적으로 받음
		_, message, err := conn.ReadMessage() // 클라이언트로부터 메시지 읽기

		//* 끊어질 때도 패킷이 보내져서 err가 됨 -> close 1000 (normal)
		if err != nil {
			color.Red("Read error: %v", err)
			break
		}
		color.Cyan("%s의 seq : %v", clientAddr, clientSeq)

		// //* 들어온 메세지 역직렬화
		payload := &WebSocketPacket.PayloadClass{}

		err = proto.Unmarshal(message, payload)
		if err != nil {
			color.Red("Proto Unmarshaling error: %v", err)
			break
		}
	}
}
