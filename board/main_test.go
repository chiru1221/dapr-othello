package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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
		input []int
		wantResult bool
	}

	testCases := []testCase {
		{input: []int{0, 7}, wantResult: false},
		{input: []int{0, 8}, wantResult: true},
		{input: []int{8, 8}, wantResult: true},
	}

	for _, tc := range testCases {
		result := isOutOfRange(tc.input[0], tc.input[1])
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

func TestReverseSearch(t *testing.T) {
	type testCase struct {
		input      Board
		wantResult string
	}

	testCases := []testCase{
		{
			input: Board{
				Stone:   "b",
				X: 3,
				Y: 5,
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbbbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
		},
		{
			input: Board{
				Stone:   "b",
				X: 3,
				Y: 5,
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnbnnbwbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "nnnnnnnnnnnnnnnnnnnnnnnnbnnbbbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
		},
	}
	
	for _, tc := range testCases {
		tc.input.reverseSearch()
		if tc.input.Squares != tc.wantResult {
			t.Errorf("Expected: %v, but got: %v", tc.wantResult, tc.input.Squares)
		}
	}
}

func TestPutableSearch(t *testing.T) {
	type testCase struct {
		input      Board
		wantResult string
	}

	testCases := []testCase{
		{
			input: Board{
				Stone:   "b",
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwnnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "nnnnnnnnnnnnnnnnnnnnpnnnnnnbwpnnnnpwbnnnnnnpnnnnnnnnnnnnnnnnnnnn",
		},
		{
			input: Board{
				Stone:   "b",
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn",
		},
		{
			input: Board{
				Stone:   "w",
				Squares: "nnnnnnnnnnnnbnnnnnnnnnnnnnnbwnnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: "nnnnnnnnnnnnbnnnnnnpnnnnnnpbwnnnnnnwbpnnnnnnpnnnnnnnnnnnnnnnnnnn",
		},
	}
	
	for _, tc := range testCases {
		tc.input.putableSearch()
		if tc.input.Squares != tc.wantResult {
			t.Errorf("Expected: %v, but got: %v", tc.wantResult, tc.input.Squares)
		}
	}
}

func TestPutableHandler(t *testing.T) {
	type testCase struct {
		input      Board
		wantResult int
	}

	testCases := []testCase{
		{
			input: Board{
				Stone:   "b",
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwnnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: 200,
		},
	}

	for _, tc := range testCases {
		body, _ := json.Marshal(tc.input)
		r := httptest.NewRequest(http.MethodPost, "https://example.com",
			bytes.NewBuffer(body))
		r.Header.Add("Content-Type", "application/json")
		// server := httptest.NewServer(http.HandlerFunc(putableHandler))
		w := httptest.NewRecorder()
		putableHandler(w, r)
		if w.Code != tc.wantResult {
			t.Errorf("Expected: %v, but got: %v", tc.wantResult, w.Code)
		}
	}
}

func TestReverseHandler(t *testing.T) {
	type testCase struct {
		input      Board
		wantResult int
	}

	testCases := []testCase{
		{
			input: Board{
				Stone:   "b",
				X: 3,
				Y: 5,
				Squares: "nnnnnnnnnnnnnnnnnnnnnnnnnnnbwbnnnnnwbnnnnnnnnnnnnnnnnnnnnnnnnnnn",
			},
			wantResult: 200,
		},

	}

	for _, tc := range testCases {
		body, _ := json.Marshal(tc.input)
		r := httptest.NewRequest(http.MethodPost, "https://example.com",
			bytes.NewBuffer(body))
		r.Header.Add("Content-Type", "application/json")
		// server := httptest.NewServer(http.HandlerFunc(putableHandler))
		w := httptest.NewRecorder()
		reverseHandler(w, r)
		if w.Code != tc.wantResult {
			t.Errorf("Expected: %v, but got: %v", tc.wantResult, w.Code)
		}
	}
}
