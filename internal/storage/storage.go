package storage

import (
	"fmt"
)

const (
	gamesSetKey         = "games:set"
	gameKeyFmt          = "games:data:%s:%s"
	hostNicknameKey     = "host_nickname"
	opponentNicknameKey = "opponent_nickname"
	hostFieldKey        = "host_field"
	opponentFieldKey    = "opponent_field"
	hostUuidKey         = "host_uuid"
	opponentUuidKey     = "opponent_uuid"
	statusKey           = "status"
)

func getGameKey(hostNickname string, key string) string {
	return fmt.Sprintf(gameKeyFmt, hostNickname, key)
}
