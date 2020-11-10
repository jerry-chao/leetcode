package leetcode

import (
	"log"
	"strconv"
)

/*
 * @lc app=leetcode id=529 lang=golang
 *
 * [529] Minesweeper
 *
 * https://leetcode.com/problems/minesweeper/description/
 *
 * algorithms
 * Medium (59.94%)
 * Likes:    724
 * Dislikes: 576
 * Total Accepted:    70.8K
 * Total Submissions: 117.8K
 * Testcase Example:  '[["E","E","E","E","E"],["E","E","M","E","E"],["E","E","E","E","E"],["E","E","E","E","E"]]\n' +
  '[3,0]'
 *
 * Let's play the minesweeper game (Wikipedia, online game)!
 *
 * You are given a 2D char matrix representing the game board. 'M' represents
 * an unrevealed mine, 'E' represents an unrevealed empty square, 'B'
 * represents a revealed blank square that has no adjacent (above, below, left,
 * right, and all 4 diagonals) mines, digit ('1' to '8') represents how many
 * mines are adjacent to this revealed square, and finally 'X' represents a
 * revealed mine.
 *
 * Now given the next click position (row and column indices) among all the
 * unrevealed squares ('M' or 'E'), return the board after revealing this
 * position according to the following rules:
 *
 *
 * If a mine ('M') is revealed, then the game is over - change it to 'X'.
 * If an empty square ('E') with no adjacent mines is revealed, then change it
 * to revealed blank ('B') and all of its adjacent unrevealed squares should be
 * revealed recursively.
 * If an empty square ('E') with at least one adjacent mine is revealed, then
 * change it to a digit ('1' to '8') representing the number of adjacent
 * mines.
 * Return the board when no more squares will be revealed.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 *
 * [['E', 'E', 'E', 'E', 'E'],
 * ⁠['E', 'E', 'M', 'E', 'E'],
 * ⁠['E', 'E', 'E', 'E', 'E'],
 * ⁠['E', 'E', 'E', 'E', 'E']]
 *
 * Click : [3,0]
 *
 * Output:
 *
 * [['B', '1', 'E', '1', 'B'],
 * ⁠['B', '1', 'M', '1', 'B'],
 * ⁠['B', '1', '1', '1', 'B'],
 * ⁠['B', 'B', 'B', 'B', 'B']]
 *
 * Explanation:
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 *
 * [['B', '1', 'E', '1', 'B'],
 * ⁠['B', '1', 'M', '1', 'B'],
 * ⁠['B', '1', '1', '1', 'B'],
 * ⁠['B', 'B', 'B', 'B', 'B']]
 *
 * Click : [1,2]
 *
 * Output:
 *
 * [['B', '1', 'E', '1', 'B'],
 * ⁠['B', '1', 'X', '1', 'B'],
 * ⁠['B', '1', '1', '1', 'B'],
 * ⁠['B', 'B', 'B', 'B', 'B']]
 *
 * Explanation:
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * The range of the input matrix's height and width is [1,50].
 * The click position will only be an unrevealed square ('M' or 'E'), which
 * also means the input board contains at least one clickable square.
 * The input board won't be a stage when game is over (some mines have been
 * revealed).
 * For simplicity, not mentioned rules should be ignored in this problem. For
 * example, you don't need to reveal all the unrevealed mines when the game is
 * over, consider any cases that you will win the game or flag any squares.
 *
 *
*/

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
	nr := len(board)
	if nr < 1 {
		return [][]byte{}
	}
	nc := len(board[0])
	directionMines = [][2]int{
		{1, 1},
		{1, 0},
		{1, -1},
		{0, 1},
		{0, -1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	}
	visitedMineSweeper = map[string]bool{}
	dfsUpdateBoard(board, click, nr, nc)
	return board
}

var directionMines [][2]int
var visitedMineSweeper map[string]bool

func dfsUpdateBoard(borad [][]byte, click []int, maxr, maxc int) {
	dr := click[0]
	dc := click[1]
	// terminator
	if _, ok := visitedMineSweeper[strconv.Itoa(dr)+strconv.Itoa(dc)]; ok {
		return
	}
	if dr < 0 || dc < 0 || dr >= maxr || dc >= maxc {
		return
	}
	visitedMineSweeper[strconv.Itoa(dr)+strconv.Itoa(dc)] = true
	if borad[dr][dc] == 'M' {
		// set x and return
		borad[dr][dc] = 'X'
	} else if borad[dr][dc] == 'E' {
		mCount := 0
		for i := 0; i < len(directionMines); i++ {
			dfsDr := dr + directionMines[i][0]
			dfsDc := dc + directionMines[i][1]
			log.Println(dfsDr, dfsDc)
			if dfsDc >= 0 && dfsDc < maxc && dfsDr >= 0 && dfsDr < maxr {
				if borad[dfsDr][dfsDc] == 'M' {
					mCount++
				}
			}
			dfsUpdateBoard(borad, []int{dfsDr, dfsDc}, maxr, maxc)
		}
		if mCount == 0 {
			borad[dr][dc] = 'B'
		} else {
			borad[dr][dc] = byte('0' + mCount)
		}
	}
}

// @lc code=end
