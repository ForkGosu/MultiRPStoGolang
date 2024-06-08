package WebSocketFunction

import (
	"MultiRPStoGolang/Helper"
	WebSocketPacket "MultiRPStoGolang/ProtoBuf/ProtoBufPacket"
	WData "MultiRPStoGolang/WebSocketServer/Data"

	"github.com/fatih/color"
	"google.golang.org/protobuf/proto"
)

func RoomCreate(payload *WebSocketPacket.PayloadClass, r_seq int32) {
	// 패킷 데이터 만들기
	roomInfo := WData.RoomInfo{
		R_seq:          r_seq,
		R_type:         0,
		R_maxPersonnel: 2,
		R_process:      0, // 시작 전
		R_lock:         false,
	}

	WData.AddRoomInfo(roomInfo.R_seq, roomInfo)
}

func RoomsSearch(payload *WebSocketPacket.PayloadClass) []*WebSocketPacket.PayloadClassRoomsSearch {
	// * 역직렬화
	reciveData := &WebSocketPacket.PayloadClassRoomsSearch{}
	err := proto.Unmarshal(payload.RequestData, reciveData)
	if err != nil {
		color.Red("%s 역직렬화 Error: %v", Helper.GetFuncName(), err)
		payload.ResultMessage = "Error"
		return nil
	}

	getRoomInfos := WData.GetRoomInfosByType(WData.RoomType(reciveData.RoomType))

	var tempInfos []*WebSocketPacket.PayloadClassRoomsSearch

	for _, info := range getRoomInfos {
		// 나중에 관전 넣고 싶으면 방의 NowPersonnel을 특정 게임의 정보의 SitPlayer로 해주거나
		// TODO : 또 다른 방법으로 클라이언트에서 특정 게임에 풀방이여도 참여가능하게하고 게임에 Sit만 안하면 됨
		tempInfo := &WebSocketPacket.PayloadClassRoomsSearch{
			RoomSeq:          info.R_seq,
			RoomType:         int32(info.R_type),
			RoomNowPersonnel: int32(WData.CountRoomPlayers(info.R_seq)),
			RoomMaxPersonnel: info.R_maxPersonnel,
		}
		tempInfos = append(tempInfos, tempInfo)
	}

	return tempInfos
}

func RoomJoin(payload *WebSocketPacket.PayloadClass, u_seq int32, u_nick string) bool {
	// * 역직렬화
	reciveData := &WebSocketPacket.PayloadClassRoomJoin{}
	err := proto.Unmarshal(payload.RequestData, reciveData)
	if err != nil {
		color.Red("%s 역직렬화 Error: %v", Helper.GetFuncName(), err)
		payload.ResultMessage = "Error"
		return false
	}

	existsPlayer := WData.ExistsRoomPlayer(reciveData.RoomSeq, u_seq)

	// 이미 방에 있다는 것
	if existsPlayer {
		payload.ResultMessage = "Already"
		return false
	}

	// 방에 플레이어 넣기 (넣을 수 없다면 풀방 이라는 것)
	isJoin := WData.AddRoomPlayer(reciveData.RoomSeq, u_seq, u_nick)

	// 입장할 수 없었다면 방에 사람이 가득 찼다는 것
	if !isJoin {
		payload.ResultMessage = "Full"
		return false
	} else {
		payload.ResultMessage = "Complete"
		return true
	}
}

// 방나가기
func RoomQuit(payload *WebSocketPacket.PayloadClass, r_seq int32, u_seq int32) bool {
	// * 역직렬화

	// 일단 게임에 앉아 있다면 게임에서 일어나야지(돈 들고)
	var isQuit bool = false

	isQuit = WData.RemoveRoomPlayer(r_seq, u_seq)

	if isQuit {
		// MysqlDB.PlayerQuit(WData.DB, reciveData.RoomSeq, u_seq)
		payload.ResultMessage = "Complete"
		return true
	} else {
		// 방에 없다면 오류
		payload.ResultMessage = "Error"
		return false
	}
}

// // TODO : 방 재입장
// func RoomReJoin(payload *WebSocketPacket.PayloadClass, u_seq int32) []*WebSocketPacket.PayloadClassRoomReJoin {
// 	getRoomInfos := WData.GetRoomInfosByJoinPlayer(u_seq)

// 	var tempInfos []*WebSocketPacket.PayloadClassRoomReJoin

// 	for _, info := range getRoomInfos {
// 		tempInfo := &WebSocketPacket.PayloadClassRoomReJoin{
// 			RoomSeq:  info.R_seq,
// 			RoomType: int32(info.R_type),
// 		}
// 		tempInfos = append(tempInfos, tempInfo)
// 	}

// 	return tempInfos
// }

// // TODO : 방 플레이어 인원
// func RoomPlayersSearch(payload *WebSocketPacket.PayloadClass) bool {
// 	reciveData := &WebSocketPacket.PayloadClassBGamePlayersSearch{}
// 	err := proto.Unmarshal(payload.RequestData, reciveData)
// 	if err != nil {
// 		color.Red("%s 역직렬화 Error: %v", Helper.GetFuncName(), err)
// 		payload.ResultMessage = "Error"
// 		return false
// 	}

// 	reciveData.BGamePlayerCount = WData.CountRoomPlayers(payload.BroadCastGroup)
// 	reciveData.BGameRoomPersonnel = WData.MaxPersonnel(payload.BroadCastGroup)

// 	payload.ResultMessage = "Complete"

// 	payload.RequestData, _ = proto.Marshal(reciveData)

// 	return true
// }
