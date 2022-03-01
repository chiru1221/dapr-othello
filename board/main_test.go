package main

import (
	"context"
	"errors"
	"testing"

	pb "example.com/othello/board"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

/*
	utils methods test
*/

func TestToByteSquare(t *testing.T) {
	type testCase struct {
		input      string
		wantResult [][]byte
	}

	testCases := []testCase{
		{
			input: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwnnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			wantResult: [][]byte{
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 98, 119, 110, 110, 110},
				{110, 110, 110, 119, 98, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
			},
		},
	}

	for _, tc := range testCases {
		result := toByteSquare(tc.input)

		for i := 0; i < len(tc.wantResult); i++ {
			for j := 0; j < len(tc.wantResult[0]); j++ {
				if result[i][j] != tc.wantResult[i][j] {
					t.Errorf("Expected: %v, but got: %v", tc.wantResult[i][j], result[i][j])
				}
			}
		}
	}
}

func TestToStringSquare(t *testing.T) {
	type testCase struct {
		input      [][]byte
		wantResult string
	}

	testCases := []testCase{
		{
			input: [][]byte{
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 98, 119, 110, 110, 110},
				{110, 110, 110, 119, 98, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
			},
			wantResult: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwnnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
		},
		{
			input: [][]byte{
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 110, 112, 110, 110, 110},
				{110, 110, 110, 98, 119, 112, 110, 110},
				{110, 110, 112, 119, 98, 110, 110, 110},
				{110, 110, 110, 112, 110, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
				{110, 110, 110, 110, 110, 110, 110, 110},
			},
			wantResult: "nnnnnnnnnnnnnnnnnnnnpnnnnnnbwpnnnnpwbnnnnnnpnnnnnnnnnnnnnnnnnnnn",
		},
	}

	for _, tc := range testCases {
		result := toStringSquare(tc.input)
		if tc.wantResult != result {
			t.Errorf("Expected: %v, but got: %v", tc.wantResult, result)
		}
	}
}

func TestIsOutOfRange(t *testing.T) {
	type testCase struct {
		input      []int
		wantResult bool
	}

	testCases := []testCase{
		{input: []int{0, 7}, wantResult: false},
		{input: []int{0, 8}, wantResult: true},
		{input: []int{8, 8}, wantResult: true},
	}

	for _, tc := range testCases {
		result := isOutOfRange(int32(tc.input[0]), int32(tc.input[1]))
		if result != tc.wantResult {
			t.Errorf("Expected: %v, but got: %v", tc.wantResult, result)
		}
	}
}

func TestClearP(t *testing.T) {
	type testCase struct {
		input      [][]byte
		wantResult [][]byte
	}

	testCases := []testCase{
		{
			input: [][]byte{
				{110, 112, 112},
				{112, 110, 110},
			},
			wantResult: [][]byte{
				{110, 110, 110},
				{110, 110, 110},
			},
		},
	}

	for _, tc := range testCases {
		result := clearP(tc.input)

		for i := 0; i < len(tc.wantResult); i++ {
			for j := 0; j < len(tc.wantResult[0]); j++ {
				if result[i][j] != tc.wantResult[i][j] {
					t.Errorf("Expected: %v, but got: %v", tc.wantResult[i][j], result[i][j])
				}
			}
		}
	}
}

/*
	main process test
*/

func TestReverse(t *testing.T) {
	type testCase struct {
		input      *pb.Board
		wantResult string
	}

	testCases := []testCase{
		{
			input: &pb.Board{
				Stone:   "b",
				X:       3,
				Y:       5,
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbbbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
		},
		{
			input: &pb.Board{
				Stone:   "b",
				X:       3,
				Y:       5,
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnbnnbwbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "nnnnnnnnnnnnnnnnnnnnnnnnbnnbbbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
		},
		{
			input: &pb.Board{
				Stone:   "b",
				X:       0,
				Y:       0,
				Squares: "bwwnwwwbnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "bwwnwwwbnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn",
		},
		{
			input: &pb.Board{
				Stone:   "b",
				X:       0,
				Y:       0,
				Squares: "bwbnwwwbnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "bbbnwwwbnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn",
		},
	}

	for _, tc := range testCases {
		result := reverse(tc.input)
		if result != tc.wantResult {
			t.Errorf("Expected: %v, but got: %v", tc.wantResult, tc.input.Squares)
		}
	}
}

func TestPutableSearch(t *testing.T) {
	type testCase struct {
		input      *pb.Board
		wantResult string
	}

	testCases := []testCase{
		{
			input: &pb.Board{
				Stone:   "b",
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwnnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "nnnnnnnnnnnnnnnnnnnnpnnnnnnbwpnnnnpwbnnnnnnpnnnnnnnnnnnnnnnnnnnn",
		},
		{
			input: &pb.Board{
				Stone:   "b",
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn",
		},
		{
			input: &pb.Board{
				Stone:   "w",
				Squares: "nnnnnnnnnnnnbnnnnnnnnnnnnnnbwnnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "nnnnnnnnnnnnbnnnnnnpnnnnnnpbwnnnnnnwbpnnnnnnpnnnnnnnnnnnnnnnnnnn",
		},
		{
			input: &pb.Board{
				Stone:   "b",
				Squares: "nnnnnnnnnnbnnnnnnnnbnnnnnnnwbbnnnnnwbwnnnnnnwnnnnnnnnwnnnnnnnnnn",
			},
			wantResult: "nnnnnnnnnnbnnnnnnnpbnnnnnnpwbbnnnnpwbwpnnnppwppnnnnnpwnnnnnnnnnn",
		},
	}

	for _, tc := range testCases {
		result := putableSearch(tc.input)
		if result != tc.wantResult {
			t.Errorf("Expected: %v, but got: %v", tc.wantResult, tc.input.Squares)
		}
	}
}

func TestPutableServer(t *testing.T) {
	type testCase struct {
		input      *pb.Board
		wantResult error
	}

	testCases := []testCase{
		{
			input: &pb.Board{
				Stone:   "b",
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwnnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: nil,
		},
		{
			input: &pb.Board{
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwnnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: status.Error(codes.InvalidArgument, "Stone is empty"),
		},
		{
			input: &pb.Board{
				Stone: "b",
			},
			wantResult: status.Error(codes.InvalidArgument, "Squares is empty"),
		},
	}

	s := server{}
	for _, tc := range testCases {
		_, err := s.Putable(context.Background(), tc.input)
		if !errors.Is(err, tc.wantResult) {
			t.Errorf("Expected: %v, but got: %v", tc.wantResult, err)
		}
	}
}

func TestReverseServer(t *testing.T) {
	type testCase struct {
		input      *pb.Board
		wantResult error
	}

	testCases := []testCase{
		{
			input: &pb.Board{
				Stone:   "b",
				X:       3,
				Y:       5,
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: nil,
		},
		{
			input: &pb.Board{
				X:       3,
				Y:       5,
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: status.Error(codes.InvalidArgument, "Stone is empty"),
		},
		{
			input: &pb.Board{
				Stone: "b",
				X:     3,
				Y:     5,
			},
			wantResult: status.Error(codes.InvalidArgument, "Squares is empty"),
		},
	}

	s := server{}
	for _, tc := range testCases {
		_, err := s.Reverse(context.Background(), tc.input)
		if !errors.Is(err, tc.wantResult) {
			t.Errorf("Expected: %v, but got: %v", tc.wantResult, err)
		}
	}
}
