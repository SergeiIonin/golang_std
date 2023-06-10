package main

import (
	"fmt"
)

func main() {
	in0 := [][]byte{
		{5, 3, '.', '.', 7, '.', '.', '.', '.'}, {6, '.', '.', 1, 9, 5, '.', '.', '.'},
		{'.', 9, 8, '.', '.', '.', '.', 6, '.'}, {8, '.', '.', '.', 6, '.', '.', '.', 3},
		{4, '.', '.', 8, '.', 3, '.', '.', 1}, {7, '.', '.', '.', 2, '.', '.', '.', 6},
		{'.', 6, '.', '.', '.', '.', 2, 8, '.'}, {'.', '.', '.', 4, 1, 9, '.', '.', 5},
		{'.', '.', '.', '.', 8, '.', '.', 7, 9},
	}
	in1 := [][]byte{
		{8, 3, '.', '.', 7, '.', '.', '.', '.'}, {6, '.', '.', 1, 9, 5, '.', '.', '.'},
		{'.', 9, 8, '.', '.', '.', '.', 6, '.'}, {8, '.', '.', '.', 6, '.', '.', '.', 3},
		{4, '.', '.', 8, '.', 3, '.', '.', 1}, {7, '.', '.', '.', 2, '.', '.', '.', 6},
		{'.', 6, '.', '.', '.', '.', 2, 8, '.'}, {'.', '.', '.', 4, 1, 9, '.', '.', 5},
		{'.', '.', '.', '.', 8, '.', '.', 7, 9},
	}
	in2 := [][]byte{
		{'.', 4, 6, '.', '.', '.', 6, '.', '.'},
		{'.', '.', '.', 6, '.', '.', '.', '.', 4},
		{'.', '.', '.', '.', '.', 1, '.', '.', '.'},
		{'.', '.', '.', '.', '.', 7, '.', '.', '.'},
		{5, '.', 7, '.', '.', '.', 4, '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', 3},
		{'.', '.', '.', 7, '.', '.', 1, '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', 1, 2, '.', '.', '.', '.', '.'},
	}
	fmt.Println(isValidSudoku(in0)) // true
	fmt.Println(isValidSudoku(in1)) // false
	fmt.Println(isValidSudoku(in2)) // false

}

func isValidSudoku(board [][]byte) bool {

	mapsCols := make([]map[byte]byte, 9)
	mapsBlocks := make([]map[byte]byte, 9)

	for i := 0; i < 9; i++ {
		mapsBlocks[i] = make(map[byte]byte, 9)
		mapsCols[i] = make(map[byte]byte, 9)
	}

	for i := 0; i < 9; i++ {
		mapRow := make(map[byte]byte, 9)
		for j := 0; j < 9; j++ {
			elem := board[i][j]
			if elem != '.' {
				mapRow[elem]++
				if mapRow[elem] > 1 {
					return false
				}
				mapsCols[j][elem]++
				if mapsCols[j][elem] > 1 {
					return false
				}
				indBlock := 3*(i/3) + j/3
				mapsBlocks[indBlock][elem]++
				if mapsBlocks[indBlock][elem] > 1 {
					return false
				}
			}
		}
	}

	return true
}

// fastest solution on leetcode:
// func checkRow(board []byte, n byte, j int) bool {
// 	for i, v := range board {
// 		if v == n && i != j {
// 			return false
// 		}
// 	}

// 	return true
// }

// func checkColumn(board [][]byte, n byte, index, j int) bool {
// 	for i := 0; i < len(board); i++ {
// 		if board[i][j] == n && index != i {
// 			return false
// 		}
// 	}

// 	return true
// }

// func checkHouse(board [][]byte, n byte, i, j int) bool {
// 	fi := i
// 	fj := j

// 	i /= 3
// 	i *= 3

// 	j /= 3
// 	j *= 3

// 	for ip := 0; ip < 3; ip++ {
// 		for jp := 0; jp < 3; jp++ {
// 			if board[i+ip][j+jp] == n && i+ip != fi && j+jp != fj {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

// func isValidSudoku(board [][]byte) bool {
// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[i]); j++ {
// 			if board[i][j] != '.' && !(checkHouse(board, board[i][j], i, j) && checkRow(board[i], board[i][j], j) && checkColumn(board, board[i][j], i, j)) {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }
