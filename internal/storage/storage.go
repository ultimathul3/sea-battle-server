package storage

import "fmt"

const (
	gamesSetKey     = "games:set"
	gameKeyFmt      = "games:data:%s:%s"
	hostNicknameKey = "host_nickname"
	fieldKey        = "field"
	uuidKey         = "uuid"
)

func getGameKey(hostNickname string, key string) string {
	return fmt.Sprintf(gameKeyFmt, hostNickname, key)
}
