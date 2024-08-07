package main

import (
	"testing"
)

func eqAprox(a, b, errRate float64) bool {
	diff := max(a, b) - min(a, b)
	return diff < max(a, b) * errRate
}

func TestSatToGold(t *testing.T) {
	goldOzUsdPrice := 100.00
	btcUsdPrice := 1_000.00
	amountSat := 100_000_000.00
	expOz := 10.0

	r := satToGold(btcUsdPrice, goldOzUsdPrice, amountSat)

	if !eqAprox(expOz, r, 0.0001) {
		t.Fatal("unexpected result:", r)
	}
}

func TestGoldToSat(t *testing.T) {
	amountOz := 10.0
	goldOzUsdPrice := 100.00
	btcUsdPrice := 1_000.00
	expectedSat := 100_000_000.00 

	r := goldToSat(btcUsdPrice, goldOzUsdPrice, amountOz)

	if !eqAprox(expectedSat, r, 0.0001) {
		t.Fatal("unexpected result", r)
	}
}

func TestGoldToSat_1(t *testing.T) {
	amountOz := 23.74
	goldOzUsdPrice := 2403.71
	btcUsdPrice := 57383.00
	expectedSat := 100_000_000.00 

	r := goldToSat(btcUsdPrice, goldOzUsdPrice, amountOz)

	if !eqAprox(expectedSat, r, 0.01) {
		t.Fatal("unexpected result", r)
	}
}

func TestSatToGold_1(t *testing.T) {
	amountSat := 100_000_000.00
	goldOzUsdPrice := 2403.71
	btcUsdPrice := 57383.00
	expectedOz := 23.74

	r := satToGold(btcUsdPrice, goldOzUsdPrice, amountSat)

	if !eqAprox(expectedOz, r, 0.01) {
		t.Fatal("unexpected result", r)
	}
}
