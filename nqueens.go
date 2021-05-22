package main

import "strings"

// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens attack each other.
// Given an integer n, return all distinct solutions to the n-queens puzzle.
// Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.
// Input: n = 4
// Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// https://leetcode.com/explore/challenge/card/may-leetcoding-challenge-2021/601/week-4-may-22nd-may-28th/3752/

func nqueen(n int, row int, cols, dig1, dig2 []bool, board []string, ans *[][]string) {
	if row == n {
		*ans = append(*ans, append([]string{}, board...))
		return
	}
	for col := 0; col < n; col++ {

		if cols[col] || dig1[row+col] || dig2[row-col+(n-1)] {

			continue

		}
		cols[col] = true
		dig1[row+col] = true
		dig2[row-col+(n-1)] = true
		board[row] = board[row][:col] + "Q" + board[row][col+1:]
		nqueen(n, row+1, cols, dig1, dig2, board, ans)

		// backtracking

		cols[col] = false
		dig1[row+col] = false
		dig2[row-col+(n-1)] = false
		board[row] = board[row][:col] + "." + board[row][col+1:]
	}
}

func solveNQueens(n int) [][]string {
	ans := [][]string{}

	col := make([]bool, n)
	diag1 := make([]bool, 2*n) // bottom left -> top right : row + col
	diag2 := make([]bool, 2*n) //top left -> bottom right: row - col
	board := make([]string, n) // string rep of board

	// init the board
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	nqueen(n, 0, col, diag1, diag2, board, &ans)

	return ans
}
