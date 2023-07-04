package domain

type (
	CreateGameDTO struct {
		HostNickname string `json:"host_nickname"`
		HostField    string `json:"host_field"`
		HostUuid     string `json:"host_uuid"`
	}

	JoinGameDTO struct {
		HostNickname     string `json:"host_nickname"`
		OpponentNickname string `json:"opponent_nickname"`
		OpponentUuid     string `json:"opponent_uuid"`
	}

	StartGameDTO struct {
		HostNickname  string `json:"host_nickname"`
		OpponentField string `json:"opponent_field"`
		OpponentUuid  string `json:"opponent_uuid"`
	}

	GetFieldForShootDTO struct {
		HostNickname string `json:"host_nickname"`
		Uuid         string `json:"uuid"`
	}

	ShootDTO struct {
		HostNickname string `json:"host_nickname"`
		FieldIndex   uint32 `json:"field_index"`
	}
)

func (c *CreateGameDTO) Validate() error {
	if !IsNicknameValid(c.HostNickname) {
		return ErrInvalidHostNickname
	}
	// TODO: validate field
	return nil
}

func (j *JoinGameDTO) Validate() error {
	if !IsNicknameValid(j.HostNickname) {
		return ErrInvalidHostNickname
	}
	if !IsNicknameValid(j.OpponentNickname) {
		return ErrInvalidOpponentNickname
	}
	return nil
}

func (s *StartGameDTO) Validate() error {
	if !IsNicknameValid(s.HostNickname) {
		return ErrInvalidHostNickname
	}
	// TODO: validate field
	return nil
}

func (g *GetFieldForShootDTO) Validate() error {
	if !IsNicknameValid(g.HostNickname) {
		return ErrInvalidHostNickname
	}
	return nil
}

func (s *ShootDTO) Validate() error {
	if !IsNicknameValid(s.HostNickname) {
		return ErrInvalidHostNickname
	}
	if s.FieldIndex > FieldSize*FieldSize-1 {
		return ErrInvalidFieldIndex
	}
	return nil
}
