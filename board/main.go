package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Board struct {
	Stone   string `json:"stone"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Squares string `json:"squares"`
}

type Res struct {
	Squares string `json:"squares"`
}

/*
	utils methods
*/

func toByteSquare(str string) [][]byte {
	var squares [][]byte
	squares = make([][]byte, 8)
	for i := 0; i < 8; i++ {
		squares[i] = []byte(str[8*i : 8*(i+1)])
	}
	return squares
}

func toStringSquare(bytes [][]byte) string {
	var squares string
	for i, _ := range bytes {
		squares += string(bytes[i])
	}
	return squares
}

func isOutOfRange(xSearch, ySearch int) bool {
	if xSearch < 0 || xSearch > 7 {
		return true
	} else if ySearch < 0 || ySearch > 7 {
		return true
	}
	return false
}

func clearP(squares [][]byte) [][]byte {
	for i := range squares {
		for j, sij := range squares[i] {
			if sij == 'p' {
				squares[i][j] = 'n'
			}
		}
	}
	return squares
}

/*
	main processes
*/

func (board *Board) reverseSearch() {
	squares := toByteSquare(board.Squares)
	squares = clearP(squares)
	var (
		direction        = []int{-1, 0, 1}
		xSearch, ySearch int
		reverseStone     byte
		stone            = []byte(board.Stone)[0]
	)
	if stone == 'b' {
		reverseStone = 'w'
	} else {
		reverseStone = 'b'
	}

	type xy struct {
		x int
		y int
	}

	for _, xDirection := range direction {
		for _, yDirection := range direction {
			if xDirection == 0 && yDirection == 0 {
				continue
			}
			xSearch = board.X + xDirection
			ySearch = board.Y + yDirection
			if isOutOfRange(xSearch, ySearch) {
				continue
			}

			history := []xy{}
			if s := squares[xSearch][ySearch]; s == reverseStone {
				for k := 1; k < len(squares); k++ {
					xSearch = board.X + xDirection*k
					ySearch = board.Y + yDirection*k
					if isOutOfRange(xSearch, ySearch) {
						break
					}
					history = append(history, xy{xSearch, ySearch})
					if s := squares[xSearch][ySearch]; s == stone {
						for _, point := range history {
							squares[point.x][point.y] = stone
						}
						break
					} else if s == 'n' {
						break
					}
				}
			}
		}
	}
	board.Squares = toStringSquare(squares)
}

func isPutable(stone byte, squares [][]byte, i, j int) bool {
	if s := squares[i][j]; s == 'b' || s == 'w' {
		return false
	}

	var (
		direction        = []int{-1, 0, 1}
		xSearch, ySearch int
		reverseStone     byte
	)
	if stone == 'b' {
		reverseStone = 'w'
	} else {
		reverseStone = 'b'
	}

	for _, xDirection := range direction {
		for _, yDirection := range direction {
			if xDirection == 0 && yDirection == 0 {
				continue
			}
			xSearch = i + xDirection
			ySearch = j + yDirection
			if isOutOfRange(xSearch, ySearch) {
				continue
			}

			if s := squares[xSearch][ySearch]; s == reverseStone {
				for k := 2; k < len(squares); k++ {
					xSearch = i + xDirection*k
					ySearch = j + yDirection*k
					if isOutOfRange(xSearch, ySearch) {
						break
					}
					// same stone is found -> putable
					// space is found -> no putable
					if s := squares[xSearch][ySearch]; s == stone {
						return true
					} else if s != reverseStone {
						break
					}
				}
			}
		}
	}
	return false
}

func (board *Board) putableSearch() {
	squares := toByteSquare(board.Squares)
	for i := 0; i < len(squares); i++ {
		for j := 0; j < len(squares[0]); j++ {
			if isPutable([]byte(board.Stone)[0], squares, i, j) {
				squares[i][j] = 'p'
			}
		}
	}

	board.Squares = toStringSquare(squares)
}

func putableHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var (
		board        Board
		putableBoard Res
	)

	json.NewDecoder(r.Body).Decode(&board)
	// Validate request
	board.putableSearch()
	putableBoard.Squares = board.Squares
	json.NewEncoder(w).Encode(putableBoard)
}

func reverseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var (
		board        Board
		reverseBoard Res
	)

	json.NewDecoder(r.Body).Decode(&board)
	// Validate request
	board.reverseSearch()
	reverseBoard.Squares = board.Squares
	json.NewEncoder(w).Encode(reverseBoard)
}

func main() {
	router := mux.NewRouter()
	log.Println("start server")
	router.HandleFunc("/putable", putableHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/reverse", reverseHandler).Methods(http.MethodPost, http.MethodOptions)
	log.Fatal(http.ListenAndServe(":8080", router))
}
