package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example usage of the CalculateChange function
	amountDue := 13.37
	amountPaid := 100.00

	// Cashier's available denominations and their counts
	cashier := map[float64]int{
		100.00: 10,  // $100 bills
		50.00:  20,  // $50 bills
		20.00:  30,  // $20 bills
		10.00:  40,  // $10 bills
		5.00:   50,  // $5 bills
		1.00:   100, // $1 bills
		0.25:   200, // Quarters
		0.10:   300, // Dimes
		0.05:   400, // Nickels
		0.01:   500, // Pennies
	}

	// Calculate the change to be given
	change, err := CalculateChange(amountDue, amountPaid, cashier)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		totalChange := amountPaid - amountDue
		fmt.Println("Change calculation summary")
		fmt.Printf("Amount due:   $%.2f\n", amountDue)
		fmt.Printf("Amount paid:  $%.2f\n", amountPaid)
		fmt.Printf("Total change: $%.2f\n\n", totalChange)

		var bills []float64
		var coins []float64
		for denom := range change {
			if denom >= 1.00 {
				bills = append(bills, denom)
			} else {
				coins = append(coins, denom)
			}
		}

		sort.Sort(sort.Reverse(sort.Float64Slice(bills)))
		sort.Sort(sort.Reverse(sort.Float64Slice(coins)))

		billsSubtotal := 0.0
		coinsSubtotal := 0.0

		fmt.Println("Bills:")
		if len(bills) == 0 {
			fmt.Println("  none")
		} else {
			for _, denom := range bills {
				count := change[denom]
				lineTotal := denom * float64(count)
				billsSubtotal += lineTotal
				fmt.Printf("  $%6.2f x %-3d = $%6.2f\n", denom, count, lineTotal)
			}
		}

		fmt.Println("\nCoins:")
		if len(coins) == 0 {
			fmt.Println("  none")
		} else {
			for _, denom := range coins {
				count := change[denom]
				lineTotal := denom * float64(count)
				coinsSubtotal += lineTotal
				fmt.Printf("  $%6.2f x %-3d = $%6.2f\n", denom, count, lineTotal)
			}
		}

		fmt.Printf("\nBills subtotal: $%.2f\n", billsSubtotal)
		fmt.Printf("Coins subtotal: $%.2f\n", coinsSubtotal)
		fmt.Printf("Grand total:    $%.2f\n", billsSubtotal+coinsSubtotal)

		var cashierBills []float64
		var cashierCoins []float64
		for denom := range cashier {
			if denom >= 1.00 {
				cashierBills = append(cashierBills, denom)
			} else {
				cashierCoins = append(cashierCoins, denom)
			}
		}

		sort.Sort(sort.Reverse(sort.Float64Slice(cashierBills)))
		sort.Sort(sort.Reverse(sort.Float64Slice(cashierCoins)))

		fmt.Println("\nCashier inventory after payment")
		fmt.Println("Bills:")
		for _, denom := range cashierBills {
			fmt.Printf("  $%6.2f -> %d\n", denom, cashier[denom])
		}

		fmt.Println("\nCoins:")
		for _, denom := range cashierCoins {
			fmt.Printf("  $%6.2f -> %d\n", denom, cashier[denom])
		}
	}
}

func CalculateChange(amountDue float64, amountPaid float64, cashier map[float64]int) (map[float64]int, error) {
	// Validate input amounts
	if amountDue < 0 || amountPaid < 0 {
		return nil, fmt.Errorf("amounts cannot be negative")
	}
	// Validate that the amount paid is sufficient
	if amountPaid < amountDue {
		return nil, fmt.Errorf("amount paid is less than amount due")
	}
	// Calculate the change in cents
	changeCents := int(math.Round((amountPaid - amountDue) * 100))
	// Create a map to hold the change distribution
	changeDistribution := make(map[float64]int)

	// Convert the cashier map to a slice of denominations in cents for easier processing
	denominations := make([]int, 0, len(cashier))
	// Create a map to convert back from cents to float64 denominations
	denomByCents := make(map[int]float64, len(cashier))
	// Populate the denominations slice and the conversion map
	for denom := range cashier {
		cents := int(math.Round(denom * 100))
		denominations = append(denominations, cents)
		denomByCents[cents] = denom
	}
	// Sort the denominations in descending order to prioritize larger bills and coins
	sort.Sort(sort.Reverse(sort.IntSlice(denominations)))
	// Iterate through the sorted denominations to calculate the change distribution
	for _, denom := range denominations {
		if changeCents <= 0 {
			break
		}
		// Check if the cashier has any of this denomination and if it can be used for the change
		denomValue := denomByCents[denom]
		// Calculate how many of this denomination can be used
		if cashier[denomValue] > 0 && changeCents >= denom {
			numBillsCoins := changeCents / denom // Calculate the maximum number of this denomination that can be used
			// Ensure we don't use more than what the cashier has
			if numBillsCoins > cashier[denomValue] {
				numBillsCoins = cashier[denomValue]
			}
			// Update the change distribution and the remaining change
			if numBillsCoins > 0 {
				changeDistribution[denomValue] = numBillsCoins
				changeCents -= numBillsCoins * denom
				cashier[denomValue] -= numBillsCoins
			}
		}
	}
	// If there is still change left to give, it means the cashier doesn't have enough denominations to provide the exact change
	if changeCents > 0 {
		return nil, fmt.Errorf("insufficient change available")
	}
	// Return the calculated change distribution
	return changeDistribution, nil
}
