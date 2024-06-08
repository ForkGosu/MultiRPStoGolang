package WebSocketData

import (
	"sync"

	"github.com/gorilla/websocket"
)

// ! clients 영역 (강제 로그아웃 시에 해당 clients만 Remove시킴, 정상적인 로그아웃시에도 삭제, 비정상 종료시에는 삭제 안함-재접속해야되기 때문)
// uuid를 키로 seq를 값으로 설정 할 것
var clients map[string]Client

// 구조체 정의
type Client struct {
	Seq  int32
	Nick string
}

// clients는 공용자원이기 때문에 접근 할 때 오류가 없게 하기 위해 clients에 대한 전용 뮤텍스 설정
var clientsMutex sync.Mutex // 쓰기작업 할 때

// 무조건 초기화 해줘야 함
func InitClients() {
	clients = make(map[string]Client)
}

// SetClient 함수는 새로운 Client 인스턴스를 생성하고 초기화합니다.
func AddClient(uuidKey string, seq int32, nick string) {
	clientsMutex.Lock()         // 수정 전에 락을 건다
	defer clientsMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	// 초기 값 추가
	clients[uuidKey] = Client{Seq: seq, Nick: nick}
	// clients[uuidKey].Seq = seq  // 새 클라이언트를 맵에 추가
	// clients[uuidKey].Nick = seq // 새 클라이언트를 맵에 추가
}

// 로그아웃을 안했다면 remove하지 않아서 클라이언트가 존재하게 됨
func RemoveClient(uuidKey string) {
	clientsMutex.Lock()         // 수정 전에 락을 건다
	defer clientsMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	delete(clients, uuidKey) // 클라이언트 삭제
}

// Client를 이용해 Seq반환
func GetClient(uuidKey string) Client {
	// clientsMutex.Lock()         // 수정 전에 락을 건다
	// defer clientsMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	// seq := clients[uuidKey].Seq

	// 가져오기
	return clients[uuidKey]
}

// Client가 존재하는지 검사
func ExistsClient(uuidKey string) bool {
	// clientsMutex.Lock()         // 수정 전에 락을 건다
	// defer clientsMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	_, exists := clients[uuidKey]

	// 존재한다면 true
	return exists
}

// ! connInfos 영역 (게임에 입장할 때 쓰고 정상적인 로그아웃시에만 삭제, 강제 로그아웃시에는 연결 시켜줘야해서 삭제 안됨)
// seq를 키로 설정 할 것
var connInfos map[int32]*websocket.Conn

// connInfos는 공용자원이기 때문에 접근 할 때 오류가 없게 하기 위해 connInfos에 대한 전용 뮤텍스 설정
var connInfosMutex sync.Mutex // 쓰기작업 할 때

// 무조건 초기화 해줘야 함
func InitConnInfos() {
	connInfos = make(map[int32]*websocket.Conn)
}

// SetConnInfos 함수는 새로운 Client 인스턴스를 생성하고 초기화합니다.
func AddConnInfo(seqKey int32, conn *websocket.Conn) {
	connInfosMutex.Lock()         // 수정 전에 락을 건다
	defer connInfosMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	connInfos[seqKey] = conn // 새 연결정보를 맵에 추가
}

// 로그아웃을 안했다면 remove하지 않아서 클라이언트가 존재하게 됨
func RemoveConnInfo(seqKey int32) {
	connInfosMutex.Lock()         // 수정 전에 락을 건다
	defer connInfosMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	delete(connInfos, seqKey) // 클라이언트 삭제
}

// ConnInfos를 이용해 Conn반환
func GetConnInfo(seqKey int32) *websocket.Conn {
	// connInfosMutex.Lock()         // 수정 전에 락을 건다
	// defer connInfosMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	conn := connInfos[seqKey]

	// 가져오기
	return conn
}

// ConnInfos를 이용해 확인
func ExistsConnInfo(seqKey int32) bool {
	// connInfosMutex.Lock()         // 수정 전에 락을 건다
	// defer connInfosMutex.Unlock() // 함수가 반환될 때 락을 해제한다

	_, exists := connInfos[seqKey]

	// 존재한다면 true
	return exists
}
