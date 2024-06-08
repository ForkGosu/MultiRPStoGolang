package WebSocketServer

import (
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"

	"MultiRPStoGolang/Helper"
	WebSocketPacket "MultiRPStoGolang/ProtoBuf/ProtoBufPacket"
	WData "MultiRPStoGolang/WebSocketServer/Data"
	WFunc "MultiRPStoGolang/WebSocketServer/Function"
	WSend "MultiRPStoGolang/WebSocketServer/PayloadSend"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 웹소켓 시작하는 부분
func Init() {
	// 클라이언트들 초기화
	WData.SetGame()

	WData.InitClients()
	WData.InitConnInfos()

	// Room 초기화
	WData.InitRoomInfos()
	// BGame.InitGameTables()

	// WFunc.InitRooms(WData.DB)

	// 웹소킷 초기화
	http.HandleFunc("/", HandleConnection)

	serverAddr := ":48055"
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

		//* 해당 위치부터 payload 사용
		// App버전 확인
		if payload.RequestCode == WebSocketPacket.PayloadType_AppVersionCheck {
			if payload.AppVersion != WData.Game_version {
				payload.ResultMessage = "Version_Fail"
				WSend.SendWebSocket(conn, payload)
				continue
			}
			if payload.AppName != WData.Game_name {
				payload.ResultMessage = "Name_Fail"
				WSend.SendWebSocket(conn, payload)
				continue
			}
			payload.ResultMessage = "Complete"
			WSend.SendWebSocket(conn, payload)
			continue
		}

		// 로그인 이후 UUID도 업로드됨
		if payload.RequestCode == WebSocketPacket.PayloadType_UserEnter {
			// * 역직렬화
			reciveData := &WebSocketPacket.PayloadClassUserEnter{}
			err := proto.Unmarshal(payload.RequestData, reciveData)
			if err != nil {
				color.Red("%s 역직렬화 Error: %v", Helper.GetFuncName(), err)
			}

			seq, uuid := WData.PlusUserSeq(), Helper.GetUUID()
			// seq와 uuid가 무조건 재할당되어서 로그인 매번 새롭게됨
			payload.UserUUID = uuid
			payload.ResultMessage = "Complete"
			clientSeq = seq

			// 로그인 되었다면 전역변수에 uuid를 값으로 가진 Client 생성(이제 clientUUID를 통해 마음대로(?) 접근가능)
			WData.AddClient(payload.UserUUID, clientSeq, reciveData.UserName)
			WData.AddConnInfo(clientSeq, conn)

			WSend.SendWebSocket(conn, payload)
			continue
		}

		///
		if payload.RequestCode == WebSocketPacket.PayloadType_RoomCreate {
			seq := WData.PlusRoomSeq()
			// seq가 무조건 재할당되어서 방 만들때 매번 새롭게됨
			payload.ResultMessage = "Complete"

			if payload.ResultMessage == "Complete" {
				// 클라이언트의 로컬 변수에 seq저장
				WFunc.RoomCreate(payload, seq)

				color.Red("%v", WData.GetRoomInfo(seq))
			}
			WSend.SendWebSocket(conn, payload)
			continue
		}

		///
		if payload.RequestCode == WebSocketPacket.PayloadType_RoomsSearch {
			tempInfos := WFunc.RoomsSearch(payload)
			payload.ResultMessage = "Complete"

			color.Red("%v", tempInfos)
			for _, info := range tempInfos {
				payload.RequestData, _ = proto.Marshal(info)
				WSend.SendWebSocket(conn, payload)
			}
			continue
		}

		///
		if payload.RequestCode == WebSocketPacket.PayloadType_RoomJoin {
			WFunc.RoomJoin(payload, clientSeq, WData.GetClient(payload.UserUUID).Nick)
			WSend.SendWebSocket(conn, payload)
			continue
		}

		///
		if payload.RequestCode == WebSocketPacket.PayloadType_RoomQuit {

			reciveData := &WebSocketPacket.PayloadClassRoomQuit{}
			err := proto.Unmarshal(payload.RequestData, reciveData)
			if err != nil {
				color.Red("%s 역직렬화 Error: %v", Helper.GetFuncName(), err)
				payload.ResultMessage = "Error"
			}
			WFunc.RoomQuit(payload, reciveData.RoomSeq, clientSeq)
			WSend.SendWebSocket(conn, payload)
			payload.BroadCastGroup = reciveData.RoomSeq
			payload.RequestCode = WebSocketPacket.PayloadType_RPSGamePlayersInfo
			WFunc.RPSGamePlayerInfo(payload, payload.BroadCastGroup)
			// TODO : 방전체 보내기
			WSend.SendWebSocketPayloadBroadCastGroup(payload)
			continue
		}

		/// RPS 시작
		if payload.RequestCode == WebSocketPacket.PayloadType_RPSGamePlayersInfo {
			WFunc.RPSGamePlayerInfo(payload, payload.BroadCastGroup)
			// TODO : 방전체 보내기
			WSend.SendWebSocketPayloadBroadCastGroup(payload)
			continue
		}

		///
		if payload.RequestCode == WebSocketPacket.PayloadType_RPSGameTableInfo {
			WFunc.RPSGameTableInfo(payload, payload.BroadCastGroup)
			// TODO : 방전체 보내기
			WSend.SendWebSocketPayloadBroadCastGroup(payload)
			continue
		}

		///
		if payload.RequestCode == WebSocketPacket.PayloadType_RPSGamePlaying {
			// 시작 안했다며
			if WData.GetRoomInfo(payload.BroadCastGroup).R_process == 0 {
				payload.ResultMessage = "NoGo"
				WSend.SendWebSocket(conn, payload)
				continue
			}

			// 이미 선택했다면
			if WData.GetPlayerAction(payload.BroadCastGroup, clientSeq) != -1 {
				payload.ResultMessage = "Already"
				WSend.SendWebSocket(conn, payload)
				continue
			}

			// * 역직렬화
			reciveData := &WebSocketPacket.PayloadClassRPSGamePlaying{}
			err := proto.Unmarshal(payload.RequestData, reciveData)
			if err != nil {
				color.Red("%s 역직렬화 Error: %v", Helper.GetFuncName(), err)
			}
			rps := reciveData.RpsGamePlay

			WData.SetPlayerAction(payload.BroadCastGroup, clientSeq, rps)

			isComplete := true

			players := WData.GetRoomInfo(payload.BroadCastGroup).R_players
			for i := int32(0); i < int32(len(players)); i++ {
				if players[i].PlayerAction == -1 {
					isComplete = false
				}
			}

			// 최대 인원수 만큼 방에 앉을 PlayerList를 -1로 초기화
			roomPlayers := make([]string, WData.GetRoomInfo(payload.BroadCastGroup).R_maxPersonnel)
			roomPlayerRPSs := make([]int32, WData.GetRoomInfo(payload.BroadCastGroup).R_maxPersonnel)
			for i := int32(0); i < int32(len(players)); i++ {
				roomPlayers[i] = players[i].PlayerNick
				roomPlayerRPSs[i] = players[i].PlayerAction
			}

			// 패킷 데이터 만들기
			tableInfo := &WebSocketPacket.PayloadClassRPSGameTableInfo{
				RpsTableStatus: WData.GetRoomInfo(payload.BroadCastGroup).R_process,
				RpsGamePlays:   roomPlayerRPSs,
			}

			payload.RequestData, _ = proto.Marshal(tableInfo)
			payload.RequestCode = WebSocketPacket.PayloadType_RPSGameTableInfo
			payload.ResultMessage = "Complete"
			WSend.SendWebSocketPayloadBroadCastGroup(payload)

			// 다 냈다면
			if isComplete {
				WData.SetRoomProcess(payload.BroadCastGroup, 0)

				for i := int32(0); i < int32(len(players)); i++ {
					WData.SetPlayerAction(payload.BroadCastGroup, players[i].PlayerSeq, -1)
				}
			}
			continue
		}

		///
		if payload.RequestCode == WebSocketPacket.PayloadType_RPSGameGo {
			// WFunc.RPSGamePlayerInfo(payload, payload.BroadCastGroup)
			count := WData.CountRoomPlayers(payload.BroadCastGroup)
			if count == 2 {
				// Go 승공
				WData.SetRoomProcess(payload.BroadCastGroup, 1)
				payload.ResultMessage = "Complete"
				WSend.SendWebSocketPayloadBroadCastGroup(payload)
			} else {
				payload.ResultMessage = "Error"
				WSend.SendWebSocketPayloadBroadCastGroup(payload)
			}
			continue
		}

		/// TODO : 이건 플레이할때 보낼꺼니 필요 없을뜻
		// if payload.RequestCode == WebSocketPacket.PayloadType_RPSGameProgress {
		// 	WFunc.RoomQuit(payload, clientSeq)
		// 	WSend.SendWebSocket(conn, payload)
		// 	continue
		// }
	}
}
