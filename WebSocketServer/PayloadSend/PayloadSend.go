package PayloadSend

import (
	"fmt"
	"sync"

	"github.com/fatih/color"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"MultiRPStoGolang/Helper"
	WebSocketPacket "MultiRPStoGolang/ProtoBuf/ProtoBufPacket"
	WData "MultiRPStoGolang/WebSocketServer/Data"
)

// connInfos는 공용자원이기 때문에 접근 할 때 오류가 없게 하기 위해 connInfos에 대한 전용 뮤텍스 설정
var connMutex sync.Mutex // 쓰기작업 할 때

// 기본 형태의 Send함수, (payload를 그대로 Send 해주기)
func SendWebSocket(conn *websocket.Conn, payload *WebSocketPacket.PayloadClass) {
	// payload가 없다면 클라이언트로 데이터 못보내니까 체크
	if payload == nil {
		color.Red("Error Func SendWebSocket is payload nil")
		return
	}

	color.White("데이터 Send : %s", payload) // 콘솔에 메시지 출력

	// ProtoBuf로 직렬화
	data, err := proto.Marshal(payload)
	if err != nil {
		color.Red("%s Error: %v", Helper.GetFuncName(), err)
		return
	}

	// TODO : 오류가 나지않게하기위해 이게 최선인지 생각해보기
	connMutex.Lock()         // 수정 전에 락을 건다
	defer connMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	// 클라이언트로 응답 데이터 보내기
	err = conn.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		color.Red("%s Error: %v", Helper.GetFuncName(), err)
		return
	}
}

// payload를 생성해서 클라이언트에 Send하는 형식 (data없을 경우 nil 넣어주기, group없을 경우 -1 넣어주기)
func SendWebSocketPayloadBroadCastGroup(payload *WebSocketPacket.PayloadClass) bool {
	// payload가 없다면 클라이언트로 데이터 못보내니까 체크
	if payload == nil {
		color.Red("Error Func SendWebSocket is payload nil")
		return false
	}

	color.White("데이터 Send : %s", payload) // 콘솔에 메시지 출력

	// ProtoBuf로 직렬화
	data, err := proto.Marshal(payload)
	if err != nil {
		color.Red("%s Error: %v", Helper.GetFuncName(), err)
		return false
	}

	fmt.Println("payload : ", payload)

	// TODO : 오류가 나지않게하기위해 이게 최선인지 생각해보기
	connMutex.Lock()         // 수정 전에 락을 건다
	defer connMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	// 진행된 게임의 결과 방에 있는 모두에게 보내주기
	for _, r_player := range WData.GetPlayersByRoomInfo(payload.BroadCastGroup) {
		// 없는 플레이어는 다시for문 처음으로
		if r_player.PlayerSeq == -1 {
			continue
		}
		conn := WData.GetConnInfo(r_player.PlayerSeq)

		// 클라이언트로 응답 데이터 보내기
		err = conn.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			color.Red("%s Error: %v", Helper.GetFuncName(), err)
		}
	}

	return true
}
