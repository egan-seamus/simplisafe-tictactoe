package ticTacToe

import (
	"testing"
)

type TestingSet struct {
	Board          []string
	ExpectedWinner string
	IsGameOver     bool
	AnyMovesLeft   bool
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
		IsGameOver:     true,
		AnyMovesLeft:   false,
	},
	{
		Board:          xSquareWin,
		ExpectedWinner: string(X),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          oSquareWin,
		ExpectedWinner: string(O),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          bothHaveSquares,
		ExpectedWinner: string(X),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          xColWin,
		ExpectedWinner: string(X),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          oColWin,
		ExpectedWinner: string(O),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          bothColWin,
		ExpectedWinner: string(X),
		IsGameOver:     true,
		AnyMovesLeft:   false,
	},
	{
		Board:          xRowWin,
		ExpectedWinner: string(X),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          oRowWin,
		ExpectedWinner: string(O),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          bothRowWin,
		ExpectedWinner: string(X),
		IsGameOver:     true,
		AnyMovesLeft:   false,
	},
	{
		Board:          xDiagonalWin,
		ExpectedWinner: string(X),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          oDiagonalWin,
		ExpectedWinner: string(O),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          bothDiagonalWin,
		ExpectedWinner: string(X),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          xCornersWin,
		ExpectedWinner: string(X),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          oCornersWin,
		ExpectedWinner: string(O),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          xCornersOBox,
		ExpectedWinner: string(X),
		IsGameOver:     true,
		AnyMovesLeft:   true,
	},
	{
		Board:          emptyBoard,
		ExpectedWinner: "",
		IsGameOver:     false,
		AnyMovesLeft:   true,
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
	anyMovesLeft := AnyMovesLeft(params.Board)
	if anyMovesLeft != params.AnyMovesLeft {
		t.Fatalf("Expected %t but saw %t", params.AnyMovesLeft, anyMovesLeft)
	}
	isGameOver := IsGameOver(params.Board)
	if isGameOver != params.IsGameOver {
		t.Fatalf("Expected %t but saw %t", params.IsGameOver, isGameOver)
	}

}
