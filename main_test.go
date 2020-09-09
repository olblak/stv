package stv

import (
	"fmt"
	"reflect"
	"testing"
)

func TestResults(t *testing.T) {

	votes := []map[string]int{
		{
			"Bob":    5,
			"Jeremy": 1,
			"Tom":    4,
			"Alice":  3,
			"John":   2,
		},
		{
			"Alice": 1,
			"Bob":   2,
			"Tom":   3,
		},
		{
			"Jeremy": 4,
			"Alice":  1,
			"Bob":    2,
			"Tom":    5,
			"John":   3,
		},
	}

	type dataset struct {
		choice           int
		expectedWinner   []string
		expectedRejected []string
	}
	data := []dataset{
		{
			choice:           1,
			expectedWinner:   []string{"Alice"},
			expectedRejected: []string{"Jeremy", "John"},
		},
		{
			choice:           2,
			expectedWinner:   []string{"Alice", "Jeremy"},
			expectedRejected: []string{},
		},
		{
			choice:           3,
			expectedWinner:   []string{"Alice", "Jeremy", "Bob"},
			expectedRejected: []string{"John", "Tom"},
		},
		{
			choice:           4,
			expectedWinner:   []string{"Alice", "Jeremy", "Bob", "John"},
			expectedRejected: []string{},
		},
		{
			choice:           5,
			expectedWinner:   []string{"Alice", "Jeremy", "Bob", "John", "Tom"},
			expectedRejected: []string{},
		},
	}

	for _, d := range data {
		result, winner, rejected, err := Process(&votes, d.choice)
		if err != nil {
			fmt.Println(err)
		}
		if !reflect.DeepEqual(winner, d.expectedWinner) {
			t.Errorf("For %d choice, expected winner %v, got %v\n", d.choice, d.expectedWinner, winner)
			t.Errorf("Results: %v\n", result)
		}
		if !reflect.DeepEqual(rejected, d.expectedRejected) {
			t.Errorf("For %v choice, expected rejected %v, got %v\n", d.choice, d.expectedRejected, rejected)
			t.Errorf("Results: %v\n", result)
		}

	}
}
