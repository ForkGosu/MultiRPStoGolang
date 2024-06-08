package Helper

import (
	"fmt"
	"runtime"
	"time"

	"github.com/google/uuid"
)

// 현재 함수의 이름을 반환합니다. 호출한 함수의 이름을 얻기 위해 skip=1을 사용합니다.
func GetFuncName() string {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Sprintf("? %v", 0)
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return fmt.Sprintf("%v %v", file, line)
	}

	// return fmt.Sprintf("%v %v %v", file, line, fn.Name()) // 이거하면 함수 이름도 알 수 있지만 디버그가 너무 길어짐
	return fmt.Sprintf("%v %v", file, line)
}

// GetCurrentUTC은 현재 시간을 문자열로 반환합니다.
func GetCurrentUTC() int64 {
	// 현재 UTC 시간을 나노초 단위로 가져옵니다.
	now := time.Now().UTC().UnixMilli() //time.Now().UTC().UnixNano()

	// return fmt.Sprintf("%d", now)
	return now
}

// GetCurrentUnixTime은 현재 Unix 시간을 문자열로 반환합니다.
func GetCurrentUnixTime() string {
	// 현재 UTC 시간을 초 단위 Unix 타임스탬프로 가져옵니다.
	now := time.Now().UTC().Unix()

	return fmt.Sprintf("%d", now)
}

// GetCurrentUnixTimestamp 함수는 현재의 Unix 시간을 int64 형식으로 반환합니다.
func GetCurrentUnixTimestamp() int64 {
	// 현재 시간을 UTC로 변환하고, Unix 타임스탬프(초 단위)를 int64로 반환합니다.
	return time.Now().UTC().Unix()
}

// GetUUID 함수는 uuid
func GetUUID() string {
	randomUUID, _ := uuid.NewRandom()
	return randomUUID.String()
}
