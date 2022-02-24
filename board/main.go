package main

import (
	"log"
	"context"
	"net"

	"google.golang.org/grpc"
	pb "example.com/othello/board"
)
// create files related grpc
// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative board/board.proto

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

func isOutOfRange(xSearch, ySearch int32) bool {
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

type server struct {
	pb.UnimplementedBoardApiServer
}

func reverse(board *pb.Board) (string) {
	squares := toByteSquare(board.Squares)
	squares = clearP(squares)
	var (
		direction        = []int32{-1, 0, 1}
		xSearch, ySearch int32
		reverseStone     byte
		stone            = []byte(board.Stone)[0]
	)
	if stone == 'b' {
		reverseStone = 'w'
	} else {
		reverseStone = 'b'
	}

	type xy struct {
		x int32
		y int32
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
					xSearch = board.X + xDirection*int32(k)
					ySearch = board.Y + yDirection*int32(k)
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
	return toStringSquare(squares)
}

func isPutable(stone byte, squares [][]byte, i, j int32) bool {
	if s := squares[i][j]; s == 'b' || s == 'w' {
		return false
	}

	var (
		direction        = []int32{-1, 0, 1}
		xSearch, ySearch int32
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
					xSearch = i + xDirection*int32(k)
					ySearch = j + yDirection*int32(k)
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


func putableSearch(board *pb.Board) string {
	squares := toByteSquare(board.Squares)
	for i := 0; i < len(squares); i++ {
		for j := 0; j < len(squares[0]); j++ {
			if isPutable([]byte(board.Stone)[0], squares, int32(i), int32(j)) {
				squares[i][j] = 'p'
			}
		}
	}

	return toStringSquare(squares)
}

func (s *server) Putable(ctx context.Context, in *pb.Board) (*pb.Res, error) {
	squares := putableSearch(in)
	return &pb.Res{Squares: squares}, nil	
}

func (s *server) Reverse(ctx context.Context, in *pb.Board) (*pb.Res, error) {
	squares := reverse(in)
	return &pb.Res{Squares: squares}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBoardApiServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
