package domain

import (
	"regexp"
	"unicode/utf8"
)

const (
	FieldDimension = 10
)

const (
	NicknameRegexp = `^[0-9a-zA-Zа-яА-Я]+$`
	NicknameMinLen = 3
	NicknameMaxLen = 10
)

const (
	EmptyCell               = rune(' ')
	OccupiedCell            = rune('x')
	SingleDeckShipCell      = rune('1')
	DoubleDeckShipDownCell  = rune('2')
	ThreeDeckShipDownCell   = rune('3')
	FourDeckShipDownCell    = rune('4')
	DoubleDeckShipRightCell = rune('@')
	ThreeDeckShipRightCell  = rune('#')
	FourDeckShipRightCell   = rune('$')
	ShipLeftCell            = rune('←')
	ShipLeftEndCell         = rune('◄')
	ShipUpCell              = rune('↑')
	ShipUpEndCell           = rune('▲')
	HitCell                 = rune('o')
	MissCell                = rune('.')
	FrameCell               = rune('~')
)

const (
	RequiredSingleDeckShips = 4
	RequiredDoubleDeckShips = 3
	RequiredThreeDeckShips  = 2
	RequiredFourDeckShips   = 1

	AllShipsCellsNumber = RequiredSingleDeckShips +
		RequiredDoubleDeckShips*2 +
		RequiredThreeDeckShips*3 +
		RequiredFourDeckShips*4
)

type Status uint8

const (
	UnknownStatus Status = iota
	GameCreatedStatus
	WaitingForOpponentStatus
	GameStartedStatus
	HostHitStatus
	HostMissStatus
	OpponentHitStatus
	OpponentMissStatus
	HostWonStatus
	OpponentWonStatus
)

type Role uint8

const (
	UnknownRole Role = iota
	HostRole
	OpponentRole
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

type Event struct {
	Status  Status `json:"status"`
	X       uint32 `json:"x,omitempty"`
	Y       uint32 `json:"y,omitempty"`
	Message string `json:"message,omitempty"`
}

func IsNicknameValid(nickname string) bool {
	return regexp.MustCompile(NicknameRegexp).MatchString(nickname) &&
		utf8.RuneCountInString(nickname) >= NicknameMinLen &&
		utf8.RuneCountInString(nickname) <= NicknameMaxLen
}

func ConvertFieldToRuneMatrix(field string) [][]rune {
	matrix := make([][]rune, 0, FieldDimension+2)

	frame := make([]rune, 0, FieldDimension+2)
	frame2 := make([]rune, 0, FieldDimension+2)
	for i := 0; i < FieldDimension+2; i++ {
		frame = append(frame, FrameCell)
		frame2 = append(frame2, FrameCell)
	}

	matrix = append(matrix, frame)
	for i := 0; i < FieldDimension; i++ {
		row := append([]rune{FrameCell}, []rune(field)[i*FieldDimension:(i+1)*FieldDimension]...)
		row = append(row, FrameCell)
		matrix = append(matrix, row)
	}
	matrix = append(matrix, frame2)

	return matrix
}

func ConvertFieldRuneMatrixToString(matrix [][]rune) string {
	var field string

	for i := 1; i < FieldDimension+1; i++ {
		field += string(matrix[i][1 : FieldDimension+1])
	}

	return field
}

func IsFieldValid(field string) bool {
	if utf8.RuneCountInString(field) != FieldDimension*FieldDimension {
		return false
	}

	matrix := ConvertFieldToRuneMatrix(field)
	var singleDeckShips, doubleDeckShips, threeDeckShips, fourDeckShips int

	for i := 1; i < FieldDimension+1; i++ {
		for j := 1; j < FieldDimension+1; j++ {
			switch matrix[i][j] {
			case EmptyCell, OccupiedCell, ShipLeftCell, ShipLeftEndCell, ShipUpCell, ShipUpEndCell:
				continue
			case SingleDeckShipCell:
				singleDeckShips++
				if !IsSingleDeckShipValid(matrix, i, j) {
					return false
				}
			case DoubleDeckShipRightCell:
				doubleDeckShips++
				if !IsNDeckShipRightValid(2, matrix, i, j) {
					return false
				}
			case ThreeDeckShipRightCell:
				threeDeckShips++
				if !IsNDeckShipRightValid(3, matrix, i, j) {
					return false
				}
			case FourDeckShipRightCell:
				fourDeckShips++
				if !IsNDeckShipRightValid(4, matrix, i, j) {
					return false
				}
			case DoubleDeckShipDownCell:
				doubleDeckShips++
				if !IsNDeckShipDownValid(2, matrix, i, j) {
					return false
				}
			case ThreeDeckShipDownCell:
				threeDeckShips++
				if !IsNDeckShipDownValid(3, matrix, i, j) {
					return false
				}
			case FourDeckShipDownCell:
				fourDeckShips++
				if !IsNDeckShipDownValid(4, matrix, i, j) {
					return false
				}
			default:
				return false
			}
		}
	}

	if singleDeckShips != RequiredSingleDeckShips ||
		doubleDeckShips != RequiredDoubleDeckShips ||
		threeDeckShips != RequiredThreeDeckShips ||
		fourDeckShips != RequiredFourDeckShips {
		return false
	}

	return true
}

func IsSingleDeckShipValid(matrix [][]rune, i, j int) bool {
	return AreLeftCellsOccupied(matrix, i, j) && AreTopAndBottomCellsOccupied(matrix, i, j) &&
		AreRightCellsOccupied(matrix, i, j)
}

func IsNDeckShipRightValid(shipDeck int, matrix [][]rune, i, j int) bool {
	if !AreLeftCellsOccupied(matrix, i, j) {
		return false
	}

	for k := 1; k < shipDeck; k++ {
		if shipDeck > k+1 && matrix[i][j+k] != ShipLeftCell {
			return false
		}
		if !AreTopAndBottomCellsOccupied(matrix, i, j+k) {
			return false
		}
	}

	if !AreRightCellsOccupied(matrix, i, j+shipDeck-1) || matrix[i][j+shipDeck-1] != ShipLeftEndCell {
		return false
	}

	return true
}

func IsNDeckShipDownValid(shipDeck int, matrix [][]rune, i, j int) bool {
	if !AreTopCellsOccupied(matrix, i, j) {
		return false
	}

	for k := 1; k < shipDeck; k++ {
		if shipDeck > k+1 && matrix[i+k][j] != ShipUpCell {
			return false
		}
		if !AreLeftAndRightCellsOccupied(matrix, i+k, j) {
			return false
		}
	}

	if !AreBottomCellsOccupied(matrix, i+shipDeck-1, j) || matrix[i+shipDeck-1][j] != ShipUpEndCell {
		return false
	}

	return true
}

func AreTopAndBottomCellsOccupied(matrix [][]rune, i, j int) bool {
	return IsCellOccupied(matrix, i-1, j) && IsCellOccupied(matrix, i+1, j)
}

func AreLeftAndRightCellsOccupied(matrix [][]rune, i, j int) bool {
	return IsCellOccupied(matrix, i, j-1) && IsCellOccupied(matrix, i, j+1)
}

func AreLeftCellsOccupied(matrix [][]rune, i, j int) bool {
	return IsCellOccupied(matrix, i-1, j-1) && IsCellOccupied(matrix, i, j-1) &&
		IsCellOccupied(matrix, i+1, j-1)
}

func AreRightCellsOccupied(matrix [][]rune, i, j int) bool {
	return IsCellOccupied(matrix, i-1, j+1) && IsCellOccupied(matrix, i, j+1) &&
		IsCellOccupied(matrix, i+1, j+1)
}

func AreTopCellsOccupied(matrix [][]rune, i, j int) bool {
	return IsCellOccupied(matrix, i-1, j-1) && IsCellOccupied(matrix, i-1, j) &&
		IsCellOccupied(matrix, i-1, j+1)
}

func AreBottomCellsOccupied(matrix [][]rune, i, j int) bool {
	return IsCellOccupied(matrix, i+1, j-1) && IsCellOccupied(matrix, i+1, j) &&
		IsCellOccupied(matrix, i+1, j+1)
}

func IsCellOccupied(matrix [][]rune, i, j int) bool {
	return matrix[i][j] == OccupiedCell || matrix[i][j] == FrameCell
}

func Shoot(matrix [][]rune, i, j int) {
	i++
	j++

	switch matrix[i][j] {
	case SingleDeckShipCell:
		matrix[i][j] = HitCell
		FloodLeftCells(matrix, i, j)
		FloodRightCells(matrix, i, j)
		FloodTopAndBottomCells(matrix, i, j)
	case DoubleDeckShipDownCell, ThreeDeckShipDownCell, FourDeckShipDownCell,
		ShipUpCell, ShipUpEndCell:
		matrix[i][j] = HitCell
		if IsShipDownFlooded(matrix, i, j) {
			FloodShipDown(matrix, i, j)
		}
	case DoubleDeckShipRightCell, ThreeDeckShipRightCell, FourDeckShipRightCell,
		ShipLeftCell, ShipLeftEndCell:
		matrix[i][j] = HitCell
		if IsShipRightFlooded(matrix, i, j) {
			FloodShipRight(matrix, i, j)
		}
	case EmptyCell, OccupiedCell:
		matrix[i][j] = MissCell
	}
}

func IsShipDownFlooded(matrix [][]rune, i, j int) bool {
	upChecked := false
	for k := 1; k < 4 && !upChecked; k++ {
		switch matrix[i-k][j] {
		case ShipUpCell, DoubleDeckShipDownCell, ThreeDeckShipDownCell, FourDeckShipDownCell:
			return false
		case EmptyCell, OccupiedCell, FrameCell, MissCell:
			upChecked = true
		}
	}

	downChecked := false
	for k := 1; k < 4 && !downChecked; k++ {
		switch matrix[i+k][j] {
		case ShipUpCell, ShipUpEndCell:
			return false
		case EmptyCell, OccupiedCell, FrameCell, MissCell:
			downChecked = true
		}
	}

	return true
}

func IsShipRightFlooded(matrix [][]rune, i, j int) bool {
	leftChecked := false
	for k := 1; k < 4 && !leftChecked; k++ {
		switch matrix[i][j-k] {
		case ShipLeftCell, DoubleDeckShipRightCell, ThreeDeckShipRightCell, FourDeckShipRightCell:
			return false
		case EmptyCell, OccupiedCell, FrameCell, MissCell:
			leftChecked = true
		}
	}

	rightChecked := false
	for k := 1; k < 4 && !rightChecked; k++ {
		switch matrix[i][j+k] {
		case ShipLeftCell, ShipLeftEndCell:
			return false
		case EmptyCell, OccupiedCell, FrameCell, MissCell:
			rightChecked = true
		}
	}

	return true
}

func FloodShipDown(matrix [][]rune, i, j int) {
	k := 0
	for IsCellForFlood(matrix, i-k, j) {
		FloodLeftAndRightCells(matrix, i-k, j)
		k++
	}
	FloodTopCells(matrix, i-k+1, j)

	k = 1
	for IsCellForFlood(matrix, i+k, j) {
		FloodLeftAndRightCells(matrix, i+k, j)
		k++
	}
	FloodBottomCells(matrix, i+k-1, j)
}

func FloodShipRight(matrix [][]rune, i, j int) {
	k := 0
	for IsCellForFlood(matrix, i, j-k) {
		FloodTopAndBottomCells(matrix, i, j-k)
		k++
	}
	FloodLeftCells(matrix, i, j-k+1)

	k = 1
	for IsCellForFlood(matrix, i, j+k) {
		FloodTopAndBottomCells(matrix, i, j+k)
		k++
	}
	FloodRightCells(matrix, i, j+k-1)
}

func IsCellForFlood(matrix [][]rune, i, j int) bool {
	return matrix[i][j] != EmptyCell && matrix[i][j] != OccupiedCell &&
		matrix[i][j] != FrameCell && matrix[i][j] != MissCell
}

func FloodTopAndBottomCells(matrix [][]rune, i, j int) {
	matrix[i-1][j] = MissCell
	matrix[i+1][j] = MissCell
}

func FloodLeftAndRightCells(matrix [][]rune, i, j int) {
	matrix[i][j-1] = MissCell
	matrix[i][j+1] = MissCell
}

func FloodLeftCells(matrix [][]rune, i, j int) {
	matrix[i-1][j-1] = MissCell
	matrix[i][j-1] = MissCell
	matrix[i+1][j-1] = MissCell
}

func FloodRightCells(matrix [][]rune, i, j int) {
	matrix[i-1][j+1] = MissCell
	matrix[i][j+1] = MissCell
	matrix[i+1][j+1] = MissCell
}

func FloodTopCells(matrix [][]rune, i, j int) {
	matrix[i-1][j-1] = MissCell
	matrix[i-1][j] = MissCell
	matrix[i-1][j+1] = MissCell
}

func FloodBottomCells(matrix [][]rune, i, j int) {
	matrix[i+1][j-1] = MissCell
	matrix[i+1][j] = MissCell
	matrix[i+1][j+1] = MissCell
}

func IsWin(matrix [][]rune) bool {
	hitCells := 0

	for i := 1; i < FieldDimension+1; i++ {
		for j := 1; j < FieldDimension+1; j++ {
			if matrix[i][j] == HitCell {
				hitCells++
				if hitCells == AllShipsCellsNumber {
					return true
				}
			}
		}
	}

	return false
}
