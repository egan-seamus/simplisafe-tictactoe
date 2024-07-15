package ticTacToe

import (
	"testing"
)

type TestingSet struct {
	Board          []string
	ExpectedWinner string
}

var xSquareWin = []string{"XX__", "XX__", "____", "____"}
var oSquareWin = []string{"OO__", "OO__", "____", "____"}
var bothHaveSquares = []string{"OO__", "OO__", "_XX_", "_XX_"}

var xRowWin = []string{"XXXX", "____", "____", "____"}
var oRowWin = []string{"OOOO", "____", "____", "____"}
var bothRowWin = []string{"OXOX", "OOOO", "XOXO", "XXXX"}

var xColWin = []string{"XOOO", "X___", "X___", "X__O"}
var oColWin = []string{"OXXX", "O___", "O___", "O__X"}
var bothColWin = []string{"OXOX", "OXOX", "OXOX", "OXOX"}

var xDiagonalWin = []string{"X___", "_X__", "__X_", "___X"}
var oDiagonalWin = []string{"O___", "_O__", "__O_", "___O"}
var bothDiagonalWin = []string{"X__O", "_XO_", "_OX_", "O__X"}

var xCornersWin = []string{"X__X", "____", "____", "X__X"}
var oCornersWin = []string{"O__O", "____", "____", "O__O"}
var xCornersOBox = []string{"X__X", "_OO_", "_OO_", "X__X"}

var drawBoard = []string{"OXOX", "XOXO", "OXOX", "OOXX"}
var emptyBoard = []string{"____", "____", "____", "____"}

var ValuesToTest = []TestingSet{
	{
		Board:          drawBoard,
		ExpectedWinner: "",
	},
	{
		Board:          xSquareWin,
		ExpectedWinner: string(X),
	},
	{
		Board:          oSquareWin,
		ExpectedWinner: string(O),
	},
	{
		Board:          bothHaveSquares,
		ExpectedWinner: string(X),
	},
	{
		Board:          xColWin,
		ExpectedWinner: string(X),
	},
	{
		Board:          oColWin,
		ExpectedWinner: string(O),
	},
	{
		Board:          bothColWin,
		ExpectedWinner: string(X),
	},
	{
		Board:          xRowWin,
		ExpectedWinner: string(X),
	},
	{
		Board:          oRowWin,
		ExpectedWinner: string(O),
	},
	{
		Board:          bothRowWin,
		ExpectedWinner: string(X),
	},
	{
		Board:          xDiagonalWin,
		ExpectedWinner: string(X),
	},
	{
		Board:          oDiagonalWin,
		ExpectedWinner: string(O),
	},
	{
		Board:          bothDiagonalWin,
		ExpectedWinner: string(X),
	},
	{
		Board:          xCornersWin,
		ExpectedWinner: string(X),
	},
	{
		Board:          oCornersWin,
		ExpectedWinner: string(O),
	},
	{
		Board:          xCornersOBox,
		ExpectedWinner: string(X),
	},
	{
		Board:          emptyBoard,
		ExpectedWinner: "",
	},
}

func TestAll(t *testing.T) {
	for _, param := range ValuesToTest {
		ValidateBoard(t, param)
	}
}

func ValidateBoard(t *testing.T, params TestingSet) {
	winnerName := CheckWinner(params.Board)
	if winnerName != params.ExpectedWinner {
		t.Fatalf("Expected %s but saw %s", params.ExpectedWinner, winnerName)
	}
}
