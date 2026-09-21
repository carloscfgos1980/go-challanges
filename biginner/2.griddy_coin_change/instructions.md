# Challenge 22: Greedy Coin Change

**Problem Statement**

Implement a coin change algorithm that finds the minimum number of coins needed to make a given amount of change. You'll be using a greedy approach, which works with a specific set of coin denominations.

In this challenge, you'll work with the following coin denominations: [1, 5, 10, 25, 50] (representing penny, nickel, dime, quarter, and half-dollar coins).

You'll implement two functions:

MinCoins - Find the minimum number of coins needed to make change for a given amount.
CoinCombination - Find the specific combination of coins that gives the minimum number.

**Function Signatures**

func MinCoins(amount int, denominations []int) int
func CoinCombination(amount int, denominations []int) map[int]int

Input Format
amount - An integer representing the amount of change needed (in cents).
denominations - A slice of integers representing the available coin denominations, sorted in ascending order.
Output Format
MinCoins - Returns an integer representing the minimum number of coins needed.
CoinCombination - Returns a map where keys are coin denominations and values are the number of coins used for each denomination.

**Requirements**

The MinCoins function should return the minimum number of coins needed to make the given amount.
The CoinCombination function should return a map with the specific combination of coins.
If the amount cannot be made with the given denominations, MinCoins should return -1 and CoinCombination should return an empty map.
Your solution should implement the greedy approach, which always chooses the largest coin possible.