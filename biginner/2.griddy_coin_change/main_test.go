package main

import "testing"

func countCoins(combo map[int]int) int {
	total := 0
	for _, cnt := range combo {
		total += cnt
	}
	return total
}

func comboAmount(combo map[int]int) int {
	total := 0
	for coin, cnt := range combo {
		total += coin * cnt
	}
	return total
}

func TestMinCoins(t *testing.T) {
	denominations := []int{1, 5, 10, 25, 50}

	tests := []struct {
		name     string
		amount   int
		expected int
	}{
		{name: "zero amount", amount: 0, expected: 0},
		{name: "negative amount", amount: -3, expected: -1},
		{name: "simple amount", amount: 7, expected: 3},  // 5+1+1
		{name: "mixed amount", amount: 33, expected: 5},  // 25+5+1+1+1
		{name: "larger amount", amount: 87, expected: 5}, // 50+25+10+1+1
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MinCoins(tc.amount, denominations)
			if got != tc.expected {
				t.Fatalf("MinCoins(%d) = %d, want %d", tc.amount, got, tc.expected)
			}
		})
	}
}

func TestMinCoinsImpossible(t *testing.T) {
	denominations := []int{5, 10}
	got := MinCoins(3, denominations)
	if got != -1 {
		t.Fatalf("MinCoins(3, [5 10]) = %d, want -1", got)
	}
}

func TestCoinCombination(t *testing.T) {
	denominations := []int{1, 5, 10, 25, 50}

	tests := []struct {
		name          string
		amount        int
		expected      map[int]int
		expectedCoins int
	}{
		{
			name:          "zero amount",
			amount:        0,
			expected:      map[int]int{},
			expectedCoins: 0,
		},
		{
			name:          "amount 42",
			amount:        42,
			expected:      map[int]int{25: 1, 10: 1, 5: 1, 1: 2},
			expectedCoins: 5,
		},
		{
			name:          "amount 99",
			amount:        99,
			expected:      map[int]int{50: 1, 25: 1, 10: 2, 1: 4},
			expectedCoins: 8,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CoinCombination(tc.amount, denominations)

			if len(got) != len(tc.expected) {
				t.Fatalf("CoinCombination(%d) returned %v, want %v", tc.amount, got, tc.expected)
			}

			for coin, wantCount := range tc.expected {
				if got[coin] != wantCount {
					t.Fatalf("CoinCombination(%d)[%d] = %d, want %d", tc.amount, coin, got[coin], wantCount)
				}
			}

			if comboAmount(got) != tc.amount {
				t.Fatalf("coin combination %v does not sum to %d", got, tc.amount)
			}

			if countCoins(got) != tc.expectedCoins {
				t.Fatalf("CoinCombination(%d) uses %d coins, want %d", tc.amount, countCoins(got), tc.expectedCoins)
			}
		})
	}
}

func TestCoinCombinationImpossible(t *testing.T) {
	denominations := []int{4, 6}
	got := CoinCombination(3, denominations)
	if len(got) != 0 {
		t.Fatalf("CoinCombination(3, [4 6]) = %v, want empty map", got)
	}
}
