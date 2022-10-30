package types

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestPhaseInflation(t *testing.T) {
}

func TestNextPhase(t *testing.T) {
	minter := DefaultInitialMinter()
	params := DefaultParams()

	blocksPerYear := uint64(100)
	tests := []struct {
		currentBlock, currentPhase, startPhaseBlock, blocksYear, expPhase uint64
	}{
		{1, 0, 0, blocksPerYear, 1},
		{50, 1, 1, blocksPerYear, 1},
		{99, 1, 1, blocksPerYear, 1},
		{100, 1, 1, blocksPerYear, 1},
		{101, 1, 1, blocksPerYear, 2},
		// increase blocks_per_year
		{102, 2, 101, blocksPerYear * 2, 2},
		{201, 2, 101, blocksPerYear * 2, 2},
		{299, 2, 101, blocksPerYear * 2, 2},
		{300, 2, 101, blocksPerYear * 2, 2},
		{301, 2, 101, blocksPerYear * 2, 3},
	}
	for i, tc := range tests {
		minter.Phase = tc.currentPhase
		minter.StartPhaseBlock = tc.startPhaseBlock
		params.BlocksPerYear = tc.blocksYear

		phase := minter.NextPhase(params, tc.currentBlock)

		require.True(t, phase == tc.expPhase,
			"Test Index: %v\nPhase:  %v\nExpected: %v\n", i, phase, tc.expPhase)
	}
}

func TestBlockProvision(t *testing.T) {
	minter := InitialMinter(sdk.NewInt(1))
	params := DefaultParams()

	secondsPerYear := int64(60 * 60 * 8766)

	tests := []struct {
		annualProvisions int64
		expProvisions    int64
	}{
		{secondsPerYear / 5, 1},
		{secondsPerYear/5 + 1, 1},
		{(secondsPerYear / 5) * 2, 2},
		{(secondsPerYear / 5) / 2, 0},
	}
	for i, tc := range tests {
		minter.AnnualProvisions = sdk.NewInt(tc.annualProvisions)
		provisions := minter.BlockProvision(params)

		expProvisions := sdk.NewCoin(params.MintDenom,
			sdk.NewInt(tc.expProvisions))

		require.True(t, expProvisions.IsEqual(provisions),
			"test: %v\n\tExp: %v\n\tGot: %v\n",
			i, tc.expProvisions, provisions)
	}
}

// Benchmarking :)
// previously using sdk.Int operations:
// BenchmarkBlockProvision-4 5000000 220 ns/op
//
// using sdk.Dec operations: (current implementation)
// BenchmarkBlockProvision-4 3000000 429 ns/op
func BenchmarkBlockProvision(b *testing.B) {
	b.ReportAllocs()
	minter := InitialMinter(sdk.NewInt(1))
	params := DefaultParams()

	s1 := rand.NewSource(100)
	r1 := rand.New(s1)
	minter.AnnualProvisions = sdk.NewInt(r1.Int63n(1000000))

	// run the BlockProvision function b.N times
	for n := 0; n < b.N; n++ {
		minter.BlockProvision(params)
	}
}

// Next inflation benchmarking
// BenchmarkPhaseInflation-4 1000000 1828 ns/op
func BenchmarkPhaseInflation(b *testing.B) {
	b.ReportAllocs()
	minter := InitialMinter(sdk.NewInt(1))
	phase := uint64(4)

	// run the PhaseInflationRate function b.N times
	for n := 0; n < b.N; n++ {
		minter.InflationcalculationFn(phase)
	}
}

// Next annual provisions benchmarking
// BenchmarkNextAnnualProvisions-4 5000000 251 ns/op
func BenchmarkNextAnnualProvisions(b *testing.B) {
	b.ReportAllocs()
	// minter := InitialMinter(sdk.NewInt(1))
	// params := DefaultParams()
	// totalSupply := sdk.NewInt(100000000000000)

	// run the NextAnnualProvisions function b.N times
	// for n := 0; n < b.N; n++ {
	// 	minter.params, totalSupply)
	// }
}
