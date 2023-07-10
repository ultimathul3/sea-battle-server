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

	GetRoleByUuidDTO struct {
		HostNickname string `json:"host_nickname"`
		Uuid         string `json:"uuid"`
	}

	GetFieldDTO struct {
		HostNickname string `json:"host_nickname"`
		Role         Role   `json:"role"`
	}

	ShootDTO struct {
		X uint32 `json:"x"`
		Y uint32 `json:"y"`
	}

	UpdateFieldDTO struct {
		HostNickname string `json:"host_nickname"`
		Role         Role   `json:"role"`
		Field        string `json:"field"`
	}

	GetStatusDTO struct {
		HostNickname string `json:"host_nickname"`
	}

	SetStatusDTO struct {
		HostNickname string `json:"host_nickname"`
		Status       Status `json:"status"`
		KeepTTL      bool   `json:"keep_ttl"`
	}
)

func (c *CreateGameDTO) Validate() error {
	if !IsNicknameValid(c.HostNickname) {
		return ErrInvalidHostNickname
	}
	if !IsFieldValid(c.HostField) {
		return ErrInvalidHostField
	}
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
	if !IsFieldValid(s.OpponentField) {
		return ErrInvalidOpponentField
	}
	return nil
}

func (g *GetRoleByUuidDTO) Validate() error {
	if !IsNicknameValid(g.HostNickname) {
		return ErrInvalidHostNickname
	}
	return nil
}

func (g *GetFieldDTO) Validate() error {
	if !IsNicknameValid(g.HostNickname) {
		return ErrInvalidHostNickname
	}
	return nil
}

func (s *ShootDTO) Validate() error {
	if s.X > 9 || s.Y > 9 {
		return ErrInvalidHitCell
	}
	return nil
}

func (u *UpdateFieldDTO) Validate() error {
	if !IsNicknameValid(u.HostNickname) {
		return ErrInvalidHostNickname
	}
	if !IsFieldValid(u.Field) {
		return ErrInvalidField
	}
	return nil
}
