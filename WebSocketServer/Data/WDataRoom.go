package WebSocketData

import (
	"sort"
	"sync"
)

// 방에 대한 Info를 DB로 가지않고 데이터에 저장해서 볼 수 있게 함
type RoomType int32

const (
	RPSGame RoomType = 0
)

type RoomPlayer struct {
	PlayerSeq          int32
	PlayerNick         string
	PlayerAction       int32
	PlayerDisconnPoint int32
}

// Golang 구조체의 변수는 대문자로 시작해야 public이 됨
type RoomInfo struct {
	R_seq          int32
	R_type         RoomType
	R_players      []*RoomPlayer // 플레이어(관전 가능하게 하기)
	R_maxPersonnel int32
	R_process      int32
	R_lock         bool
}

// 각 방의 R_seq를 기반으로 connInfo를 조회하기 위해 U_seq를 하나하나 집어 넣을 것
var roomInfos map[int32]RoomInfo
var roomInfosMutex map[int32]*sync.Mutex

// 초기화 함수
func InitRoomInfos() {
	roomInfos = make(map[int32]RoomInfo)
	roomInfosMutex = make(map[int32]*sync.Mutex)
}

func GetEmptyRoomPlayer() *RoomPlayer {
	gamePlayer := &RoomPlayer{
		PlayerSeq:          int32(-1),
		PlayerDisconnPoint: int32(0),
	}

	return gamePlayer
}

func SetRoomProcess(roomSeq int32, process int32) {
	roomInfosMutex[roomSeq].Lock()
	defer roomInfosMutex[roomSeq].Unlock()

	temp := roomInfos[roomSeq]
	temp.R_process = process
	roomInfos[roomSeq] = temp

}

// 방 추가 함수
func AddRoomInfo(roomSeq int32, roomInfo RoomInfo) {
	// personnel만큼 List 초기화
	players := make([]*RoomPlayer, roomInfo.R_maxPersonnel)
	for i := int32(0); i < roomInfo.R_maxPersonnel; i++ {
		players[i] = GetEmptyRoomPlayer()
	}
	roomInfo.R_players = players

	// 방 정보 넣기
	roomInfos[roomSeq] = roomInfo
	roomInfosMutex[roomSeq] = &sync.Mutex{}
}

// 방 삭제 함수 (사용되고 있는 게임이나 로직이 있으니 신중해서 삭제 할 것)
func RemoveRoomInfo(roomSeq int32) {
	delete(roomInfos, roomSeq)      // 방 삭제
	delete(roomInfosMutex, roomSeq) // 방 삭제
}

// RoomInPlayers가 존재하는지 검사
func ExistsRoomInfo(roomSeq int32) bool {
	// roomInfosMutex[roomSeq].Lock()
	// defer roomInfosMutex[roomSeq].Unlock()

	_, exists := roomInfos[roomSeq]

	// 존재한다면 true
	return exists
}

// 방의 정보에 대한 구조체를 복사해서 전달해줌
func GetRoomInfo(roomSeq int32) RoomInfo {
	// roomInfosMutex[roomSeq].Lock()
	// defer roomInfosMutex[roomSeq].Unlock()

	return roomInfos[roomSeq]
}

// ? 방의 정보에 대한 구조체를 복사해서 전달해줌 > 쓰기를 2개의 코루틴에서만 안하면 공유자원 오류는 안뜸
func GetRoomInfosByType(roomType RoomType) []RoomInfo {
	var tempInfos []RoomInfo
	for _, info := range roomInfos {
		if info.R_type == roomType {
			tempInfos = append(tempInfos, info)
		}
	}

	// 정렬해서 보내기
	sort.Slice(tempInfos, func(i, j int) bool {
		return tempInfos[i].R_seq < tempInfos[j].R_seq
	})

	return tempInfos
}

// ? 방의 정보에 대한 구조체를 복사해서 전달해줌 > 쓰기를 2개의 코루틴에서만 안하면 공유자원 오류는 안뜸
func GetRoomInfosByJoinPlayer(userSeq int32) []RoomInfo {
	var tempInfos []RoomInfo
	for _, info := range roomInfos {

		// 플레이어 확인
		for _, player := range info.R_players {
			if player.PlayerSeq == userSeq {
				tempInfos = append(tempInfos, info)
			}
		}
	}

	// 정렬해서 보내기
	sort.Slice(tempInfos, func(i, j int) bool {
		return tempInfos[i].R_seq < tempInfos[j].R_seq
	})

	return tempInfos
}

// 방에 있는 플레이어들 보내기
func GetPlayersByRoomInfo(roomSeq int32) []*RoomPlayer {
	return roomInfos[roomSeq].R_players
}

func SetPlayerAction(roomSeq int32, userSeq int32, userAction int32) {
	roomInfosMutex[roomSeq].Lock()
	defer roomInfosMutex[roomSeq].Unlock()

	player := GetPlayer(roomSeq, userSeq)

	player.PlayerAction = userAction
}

func GetPlayerAction(roomSeq int32, userSeq int32) int32 {
	player := GetPlayer(roomSeq, userSeq)

	return player.PlayerAction
}

// 방에 있는 플레이어들 보내기
func GetPlayer(roomSeq int32, userSeq int32) *RoomPlayer {
	// 만약 roomInPlayers[roomSeq] 안에 있는 리스트가 -1이라면 들어갈 수 있음
	for i, player := range roomInfos[roomSeq].R_players {
		if player.PlayerSeq == userSeq {
			return roomInfos[roomSeq].R_players[i] // Player가 들어갔기 때문에 자리 있어서 거기 넣음 true
		}
	}

	return nil
}

// 방에 플레이어 추가 함수
func AddRoomPlayer(roomSeq int32, userSeq int32, userNick string) bool {
	roomInfosMutex[roomSeq].Lock()
	defer roomInfosMutex[roomSeq].Unlock()

	// 만약 roomInPlayers[roomSeq] 안에 있는 리스트가 -1이라면 들어갈 수 있음
	for i, player := range roomInfos[roomSeq].R_players {
		if player.PlayerSeq == -1 {
			roomInfos[roomSeq].R_players[i].PlayerSeq = userSeq
			roomInfos[roomSeq].R_players[i].PlayerNick = userNick
			roomInfos[roomSeq].R_players[i].PlayerAction = -1
			roomInfos[roomSeq].R_players[i].PlayerDisconnPoint = 0
			return true // Player가 들어갔기 때문에 자리 있어서 거기 넣음 true
		}
	}

	return false // Player로 들어갈 공간이 없기 때문에 false
}

// 방에서 플레이어 제거 함수
func RemoveRoomPlayer(roomSeq int32, userSeq int32) bool {
	roomInfosMutex[roomSeq].Lock()
	defer roomInfosMutex[roomSeq].Unlock()

	// 플레이어 제거
	for i, player := range roomInfos[roomSeq].R_players {
		if player.PlayerSeq == userSeq {
			roomInfos[roomSeq].R_players[i].PlayerSeq = -1
			roomInfos[roomSeq].R_players[i].PlayerNick = ""
			roomInfos[roomSeq].R_players[i].PlayerAction = -1
			roomInfos[roomSeq].R_players[i].PlayerDisconnPoint = 0
			return true // Player가 삭제되었기 때문에 true
		}
	}

	return false // Player가 삭제되지 않았기 때문에 false
}

// 방에서 플레이어 제거 함수
func CountRoomPlayers(roomSeq int32) int32 {
	// roomInfosMutex[roomSeq].Lock()
	// defer roomInfosMutex[roomSeq].Unlock()

	count := int32(0)

	// 플레이어 제거
	for _, player := range roomInfos[roomSeq].R_players {
		if player.PlayerSeq != -1 {
			count++
		}
	}

	return count // Player가 삭제되지 않았기 때문에 false
}

// 방 최대 인원
func MaxPersonnel(roomSeq int32) int32 {
	return roomInfos[roomSeq].R_maxPersonnel // Player가 삭제되지 않았기 때문에 false
}

// 방에 플레이어 끊겨있으면 끊겨있었다고 1증가
func SetRoomPlayerDisconnPointUp(roomSeq int32, userSeq int32) bool {
	roomInfosMutex[roomSeq].Lock()
	defer roomInfosMutex[roomSeq].Unlock()

	// 만약 roomInPlayers[roomSeq] 안에 있는 리스트가 -1이라면 들어갈 수 있음
	for i, player := range roomInfos[roomSeq].R_players {
		if player.PlayerSeq == userSeq {
			roomInfos[roomSeq].R_players[i].PlayerDisconnPoint++
			return true // Player가 들어갔기 때문에 자리 있어서 거기 넣음 true
		}
	}

	return false // Player로 들어갈 공간이 없기 때문에 false
}

// 방에 플레이어 끊겨있지 않으니 0으로 초기화
func SetRoomPlayerDisconnPointZero(roomSeq int32, userSeq int32) bool {
	roomInfosMutex[roomSeq].Lock()
	defer roomInfosMutex[roomSeq].Unlock()

	// 만약 roomInPlayers[roomSeq] 안에 있는 리스트가 -1이라면 들어갈 수 있음
	for i, player := range roomInfos[roomSeq].R_players {
		if player.PlayerSeq == userSeq {
			roomInfos[roomSeq].R_players[i].PlayerDisconnPoint = 0
			return true // Player가 들어갔기 때문에 자리 있어서 거기 넣음 true
		}
	}

	return false // Player로 들어갈 공간이 없기 때문에 false
}

// 방에 플레이어가 있는지 확인 함수
func ExistsRoomPlayer(roomSeq int32, userSeq int32) bool {
	// roomInfosMutex[roomSeq].Lock()
	// defer roomInfosMutex[roomSeq].Unlock()

	// 플레이어 확인
	for _, player := range roomInfos[roomSeq].R_players {
		if player.PlayerSeq == userSeq {
			return true // Player있기 때문에 true
		}
	}

	return false // Player 없음
}
