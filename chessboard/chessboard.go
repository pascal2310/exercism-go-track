package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	currentFile := cb[file]
	var count = 0
	for i := range currentFile {
		if currentFile[i] {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var count = 0
	for key := range cb {
		file := cb[key]
		index := rank - 1
		if index < 0 || index > len(file) {
			return 0
		}
		if file[index] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count = 0
	for i := range cb {
		for range cb[i] {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count = 0
	for i := range cb {
		for j := range cb[i] {
			if cb[i][j] {
				count++
			}
		}
	}
	return count
}
