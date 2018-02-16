package business

import (
	"reflect"
	"testing"

	"github.com/mauleyzaola/challenge/domain"
)

func TestWhenEachParser(t *testing.T) {
	type tcase struct {
		input    string
		sample   []domain.Product
		error    bool
		expected map[int]bool
	}
	testCases := []tcase{
		{
			input: "",
			error: true,
		},
		{
			input: "each:3",
			sample: []domain.Product{
				{Code: "1", Name: "One", Price: 10},
				{Code: "2", Name: "Two", Price: 20},
				{Code: "3", Name: "Three", Price: 30},
				{Code: "4", Name: "Four", Price: 40},
				{Code: "5", Name: "Five", Price: 50},
				{Code: "6", Name: "Six", Price: 60},
				{Code: "7", Name: "Seven", Price: 70},
			},
			expected: map[int]bool{
				3: true,
				6: true,
			},
		},
	}

	for _, tc := range testCases {
		callback, err := WhenParser(tc.input)
		if tc.error {
			if err == nil {
				t.Errorf("expected error but got nil with input:%s", tc.input)
			}
		} else {
			if err != nil {
				t.Error(err)
				continue
			}
			result := callback(tc.sample)
			if !reflect.DeepEqual(tc.expected, result) {
				t.Errorf("expected:\n%#v\nbut got instead:\n%#v\n", tc.expected, result)
			}
		}
	}
}

func TestWhenTotalCounter(t *testing.T) {
	type tcase struct {
		input    string
		sample   []domain.Product
		error    bool
		expected map[int]bool
	}
	testCases := []tcase{
		{
			input: "",
			error: true,
		},
		{
			input: "gte:3",
			sample: []domain.Product{
				{Code: "1", Name: "One", Price: 10},
				{Code: "2", Name: "Two", Price: 20},
				{Code: "3", Name: "Three", Price: 30},
				{Code: "4", Name: "Four", Price: 40},
			},
			expected: map[int]bool{
				0: true,
				1: true,
				2: true,
				3: true,
			},
		},
		{
			input: "gte:5",
			sample: []domain.Product{
				{Code: "1", Name: "One", Price: 10},
				{Code: "2", Name: "Two", Price: 20},
				{Code: "3", Name: "Three", Price: 30},
				{Code: "4", Name: "Four", Price: 40},
			},
			expected: map[int]bool{},
		},
	}

	for _, tc := range testCases {
		callback, err := WhenParser(tc.input)
		if tc.error {
			if err == nil {
				t.Errorf("expected error but got nil with input:%s", tc.input)
			}
		} else {
			if err != nil {
				t.Error(err)
				continue
			}
			result := callback(tc.sample)
			if !reflect.DeepEqual(tc.expected, result) {
				t.Errorf("expected:\n%#v\nbut got instead:\n%#v\n", tc.expected, result)
			}
		}
	}
}
