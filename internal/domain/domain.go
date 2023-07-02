package domain

type Game struct {
	HostNickname     string
	OpponentNickname string
	HostField        string
	OpponentField    string
	UUID             string
	Status           uint16
}

type CreateGameDTO struct {
	HostNickname string
	HostField    string
	UUID         string
}
