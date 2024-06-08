package WebSocketData

import "sync"

var Game_name string
var Game_version string
var Game_userSeq int32
var Game_roomSeq int32

// gameMutex는 공용자원이기 때문에 접근 할 때 오류가 없게 하기 위해 gameMutex에 대한 전용 뮤텍스 설정
var gameMutex sync.Mutex // 쓰기작업 할 때

func SetGame() {
	Game_name = "MultiRPS"
	Game_version = "0.1.0"
	Game_userSeq = 0
	Game_roomSeq = 0
}

func PlusUserSeq() int32 {
	// TODO : 오류가 나지않게하기위해 이게 최선인지 생각해보기
	gameMutex.Lock()         // 수정 전에 락을 건다
	defer gameMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	Game_userSeq++

	return Game_userSeq
}

func PlusRoomSeq() int32 {
	// TODO : 오류가 나지않게하기위해 이게 최선인지 생각해보기
	gameMutex.Lock()         // 수정 전에 락을 건다
	defer gameMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	Game_roomSeq++

	return Game_roomSeq
}
