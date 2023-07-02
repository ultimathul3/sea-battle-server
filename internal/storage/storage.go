package storage

import "fmt"

const (
	gamesSetKey      = "games:set"
	gameKeyFmt       = "games:data:%s:%s"
	hostNicknameKey  = "host_nickname"
	hostFieldKey     = "host_field"
	opponentFieldKey = "opponent_field"
	uuidKey          = "uuid"
	statusKey        = "status"
)

const (
	GameCreated int = iota
)

func getGameKey(hostNickname string, key string) string {
	return fmt.Sprintf(gameKeyFmt, hostNickname, key)
}
