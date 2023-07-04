package domain

import (
	"regexp"
	"unicode/utf8"
)

const (
	FieldSize      = 10
	NicknameRegexp = `^[0-9a-zA-Zа-яА-Я]+$`
	NicknameMinLen = 3
	NicknameMaxLen = 10
)

type Status uint8

const (
	Unknown Status = iota
	GameCreated
	WaitingForOpponent
	GameStarted
	HostHit
	HostMiss
	OpponentHit
	OpponentMiss
	HostWon
	OpponentWon
)

type Game struct {
	HostNickname     string `json:"host_nickname"`
	OpponentNickname string `json:"opponent_nickname"`
	HostField        string `json:"host_field"`
	OpponentField    string `json:"opponent_field"`
	HostUuid         string `json:"host_uuid"`
	OpponentUuid     string `json:"opponent_uuid"`
	Status           Status `json:"status"`
}

func IsNicknameValid(nickname string) bool {
	return regexp.MustCompile(NicknameRegexp).MatchString(nickname) &&
		utf8.RuneCountInString(nickname) >= NicknameMinLen &&
		utf8.RuneCountInString(nickname) <= NicknameMaxLen
}
