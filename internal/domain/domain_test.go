package domain

import (
	"testing"
)

func TestFieldValidation(t *testing.T) {
	validTests := []struct {
		field string
	}{
		{
			ConvertFieldRuneMatrixToString([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~ xxxxxxxx1~"),
				[]rune("~ x1x1x1xxx~"),
				[]rune("~ xxxxxxxx ~"),
				[]rune("~  xxxxx2x ~"),
				[]rune("~  x#←◄x▲x ~"),
				[]rune("~xxxxxxxxx ~"),
				[]rune("~x4x3x2xxxx~"),
				[]rune("~x↑x↑x▲xx@◄~"),
				[]rune("~x↑x▲xxxxxx~"),
				[]rune("~x▲xxx     ~"),
				[]rune("~~~~~~~~~~~~"),
			}),
		},
		{
			ConvertFieldRuneMatrixToString([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~1x3x3x  x1~"),
				[]rune("~xx↑x↑x  xx~"),
				[]rune("~4x▲x▲xxxx ~"),
				[]rune("~↑xxxxxx2x ~"),
				[]rune("~↑x2x2xx▲x ~"),
				[]rune("~▲x▲x▲xxxx ~"),
				[]rune("~xxxxxx    ~"),
				[]rune("~          ~"),
				[]rune("~xx      xx~"),
				[]rune("~1x      x1~"),
				[]rune("~~~~~~~~~~~~"),
			}),
		},
		{
			ConvertFieldRuneMatrixToString([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~xx xxxx xx~"),
				[]rune("~1x x@◄x x1~"),
				[]rune("~xxxxxxxxxx~"),
				[]rune("~#←◄x   x@◄~"),
				[]rune("~xxxxx xxxx~"),
				[]rune("~$←←◄x x@◄x~"),
				[]rune("~xxxxx xxxx~"),
				[]rune("~xx      xx~"),
				[]rune("~1xxxxx  x1~"),
				[]rune("~xx#←◄x  xx~"),
				[]rune("~~~~~~~~~~~~"),
			}),
		},
	}

	invalidTests := []struct {
		field string
	}{
		{
			ConvertFieldRuneMatrixToString([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~ xxxxxxx11~"),
				[]rune("~ x1x1x1xxx~"),
				[]rune("~ xxxxxxxx ~"),
				[]rune("~  xxxxx2x ~"),
				[]rune("~  x#←◄x▲x ~"),
				[]rune("~xxxxxxxxx ~"),
				[]rune("~x4x3x2xxxx~"),
				[]rune("~x↑x↑x▲xx@◄~"),
				[]rune("~x↑x▲xxxxxx~"),
				[]rune("~x▲xxx     ~"),
				[]rune("~~~~~~~~~~~~"),
			}),
		},
		{
			ConvertFieldRuneMatrixToString([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~ xxxxxxxx1~"),
				[]rune("~ x1x1x1xxx~"),
				[]rune("~ xxxxxxxx ~"),
				[]rune("~  xxxxx2x ~"),
				[]rune("~  x#←◄x▲x ~"),
				[]rune("~xxxxxxxxx ~"),
				[]rune("~x4x3x2xxxx~"),
				[]rune("~x↑x↑x▲ x@◄~"),
				[]rune("~x↑x▲xxxxxx~"),
				[]rune("~x▲xxx     ~"),
				[]rune("~~~~~~~~~~~~"),
			}),
		},
		{
			ConvertFieldRuneMatrixToString([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~ xxxxxxxx1~"),
				[]rune("~ x1x1x1xxx~"),
				[]rune("~ xxxxxxxx ~"),
				[]rune("~  xxxxx2x ~"),
				[]rune("~  x#←◄x▲x ~"),
				[]rune("~xxxxxxxxx ~"),
				[]rune("~x4x3x2xxxx~"),
				[]rune("~x↑x↑x▲xx@◄~"),
				[]rune("~x↑x▲xxxxxx~"),
				[]rune("~xxxxx     ~"),
				[]rune("~~~~~~~~~~~~"),
			}),
		},
		{
			ConvertFieldRuneMatrixToString([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~ xxxxxxxx1~"),
				[]rune("~ x1x1x1xxx~"),
				[]rune("~ xxxxxxxx ~"),
				[]rune("~  xxxxx2x ~"),
				[]rune("~  x#←◄x▲x ~"),
				[]rune("~xxxxxxxxx ~"),
				[]rune("~x4x3x2xxxx~"),
				[]rune("~x↑x↑x▲xx@◄~"),
				[]rune("~x↑x▲xxxxxx~"),
				[]rune("~x▲xxx   x1~"),
				[]rune("~~~~~~~~~~~~"),
			}),
		},
		{
			ConvertFieldRuneMatrixToString([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~ xxxxxxxx1~"),
				[]rune("~ x1x1x1xxx~"),
				[]rune("~ xxxxxxxx ~"),
				[]rune("~  xxxxx2x ~"),
				[]rune("~  x#←◄x▲x ~"),
				[]rune("~xxxxxxxxx ~"),
				[]rune("~x4x3x2xxxx~"),
				[]rune("~x↑x↑x◄xx@◄~"),
				[]rune("~x↑x▲xxxxxx~"),
				[]rune("~x▲xxx     ~"),
				[]rune("~~~~~~~~~~~~"),
			}),
		},
		{
			ConvertFieldRuneMatrixToString([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~ xxxxxxxx1~"),
				[]rune("~ x1x1x1xxx~"),
				[]rune("~ xxxxxxxx ~"),
				[]rune("~  xxxxx2x ~"),
				[]rune("~  x#←◄x▲x ~"),
				[]rune("~xxxxxxxxx ~"),
				[]rune("~  x3x2xxxx~"),
				[]rune("~  x↑x▲xx@◄~"),
				[]rune("~  x▲xxxxxx~"),
				[]rune("~  xxx     ~"),
				[]rune("~~~~~~~~~~~~"),
			}),
		},
		{
			ConvertFieldRuneMatrixToString([][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~ xxxxxxx  ~"),
				[]rune("~ x1x1x1x  ~"),
				[]rune("~ xxxxxxxx ~"),
				[]rune("~  xxxxx2x ~"),
				[]rune("~  x#←◄x▲x ~"),
				[]rune("~xxxxxxxxx ~"),
				[]rune("~  x3x2xxxx~"),
				[]rune("~  x↑x▲xx@◄~"),
				[]rune("~ xx▲xxxxxx~"),
				[]rune("~ x1xx     ~"),
				[]rune("~~~~~~~~~~~~"),
			}),
		},
	}

	for i, test := range validTests {
		if !IsFieldValid(test.field) {
			t.Errorf("%d field in 'validTests' is not valid", i+1)
		}
	}

	for i, test := range invalidTests {
		if IsFieldValid(test.field) {
			t.Errorf("%d field in 'invalidTests' is valid", i+1)
		}
	}
}

func TestShoot(t *testing.T) {
	type Coords struct {
		x int
		y int
	}
	tests := []struct {
		fieldBeforeShoot [][]rune
		coords           []Coords
		fieldAfterShoot  [][]rune
	}{
		{
			[][]rune{
				[]rune("~~~~~~~~~~~~"),
				[]rune("~ xxxxxxxx1~"),
				[]rune("~ x1x1x1xxx~"),
				[]rune("~ xxxxxxxx ~"),
				[]rune("~  xxxxx2x ~"),
				[]rune("~  x#←◄x▲x ~"),
				[]rune("~xxxxxxxxx ~"),
				[]rune("~x4x3x2xxxx~"),
				[]rune("~x↑x↑x▲xx@◄~"),
				[]rune("~x↑x▲xxxxxx~"),
				[]rune("~x▲xxx     ~"),
				[]rune("~~~~~~~~~~~~"),
			},
			[]Coords{
				{x: 2, y: 1},

				{x: 9, y: 0},

				{x: 4, y: 1},

				{x: 6, y: 1},

				{x: 7, y: 3},
				{x: 7, y: 4},

				{x: 8, y: 7},
				{x: 9, y: 7},

				{x: 5, y: 7},
				{x: 5, y: 6},

				{x: 5, y: 4},
				{x: 3, y: 4},
				{x: 4, y: 4},

				{x: 3, y: 6},
				{x: 3, y: 7},
				{x: 3, y: 8},

				{x: 1, y: 9},
				{x: 1, y: 7},
				{x: 1, y: 8},
				{x: 1, y: 6},
			},
			[][]rune{
				[]rune("~~~~~~~~~..."),
				[]rune("~ ........o."),
				[]rune("~ .o.o.o...."),
				[]rune("~ ........ ~"),
				[]rune("~  .....o. ~"),
				[]rune("~  .ooo.o. ~"),
				[]rune("~......... ~"),
				[]rune("~.o.o.o....."),
				[]rune("~.o.o.o..oo."),
				[]rune("~.o.o......."),
				[]rune("~.o...     ~"),
				[]rune("~...~~~~~~~~"),
			},
		},
	}

	for i, test := range tests {
		for _, coords := range test.coords {
			Shoot(test.fieldBeforeShoot, coords.y, coords.x)
		}

		for j := 0; j < FieldDimension; j++ {
			for k := 0; k < FieldDimension; k++ {
				if test.fieldAfterShoot[j][k] != test.fieldBeforeShoot[j][k] {
					t.Errorf("%d field after shoot is not valid", i+1)
					return
				}
			}
		}
	}
}
