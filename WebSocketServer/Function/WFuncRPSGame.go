package WebSocketFunction

import (
	"github.com/fatih/color"
	"google.golang.org/protobuf/proto"

	WebSocketPacket "MultiRPStoGolang/ProtoBuf/ProtoBufPacket"
	WData "MultiRPStoGolang/WebSocketServer/Data"
)

// B게임 방 플레이어 정보 전달
func RPSGamePlayerInfo(payload *WebSocketPacket.PayloadClass, r_seq int32) bool {
	players := WData.GetRoomInfo(r_seq).R_players
	// 자리에 잘 착석 못하면 오류
	// if players == nil {
	// 	payload.ResultMessage = "Error"
	// 	return false
	// }

	// 최대 인원수 만큼 방에 앉을 PlayerList를 -1로 초기화
	roomPlayers := make([]string, WData.GetRoomInfo(r_seq).R_maxPersonnel)
	for i := int32(0); i < int32(len(players)); i++ {
		roomPlayers[i] = players[i].PlayerNick
	}
	color.Red("%v", r_seq)
	color.Red("%v", roomPlayers)
	// 패킷 데이터 만들기
	playerInfo := &WebSocketPacket.PayloadClassRPSGamePlayersInfo{
		RpsGamePlayerNames: roomPlayers,
	}

	payload.RequestData, _ = proto.Marshal(playerInfo)

	payload.ResultMessage = "Complete"

	return true
}

// B게임 방 정보 전달
func RPSGameTableInfo(payload *WebSocketPacket.PayloadClass, r_seq int32) bool {
	players := WData.GetRoomInfo(r_seq).R_players
	// 자리에 잘 착석 못하면 오류
	if players == nil {
		color.Red("착석못했나?")
		payload.ResultMessage = "Error"
		return false
	}

	// 최대 인원수 만큼 방에 앉을 PlayerList를 -1로 초기화
	roomPlayers := make([]string, WData.GetRoomInfo(r_seq).R_maxPersonnel)
	roomPlayerRPSs := make([]int32, WData.GetRoomInfo(r_seq).R_maxPersonnel)
	for i := int32(0); i < int32(len(players)); i++ {
		roomPlayers[i] = players[i].PlayerNick
		roomPlayerRPSs[i] = players[i].PlayerAction
	}

	// 패킷 데이터 만들기
	tableInfo := &WebSocketPacket.PayloadClassRPSGameTableInfo{
		RpsTableStatus: WData.GetRoomInfo(r_seq).R_process,
		RpsGamePlays:   roomPlayerRPSs,
	}

	payload.RequestData, _ = proto.Marshal(tableInfo)

	payload.ResultMessage = "Complete"

	return true
}
