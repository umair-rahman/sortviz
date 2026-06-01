package main

import (
	"sort"
	"testing"
)

// TestAllAlgosSortCorrectly verifies every algorithm produces a correctly
// sorted final array, and that the recorded steps' last snapshot is sorted.
func TestAllAlgosSortCorrectly(t *testing.T) {
	sizes := []int{1, 2, 10, 40, 60}
	seeds := []int64{1, 42, 999, -5}
	for _, algo := range validAlgos {
		for _, size := range sizes {
			for _, seed := range seeds {
				steps, stats := runAlgo(algo, size, seed)

				// The working array result is the last snapshot.
				if len(steps) == 0 && size > 1 {
					t.Fatalf("%s n=%d seed=%d: no steps recorded", algo, size, seed)
				}
				if len(steps) > 0 {
					last := steps[len(steps)-1].snapshot
					if !sort.IntsAreSorted(last) {
						t.Errorf("%s n=%d seed=%d: final snapshot not sorted: %v", algo, size, seed, last)
					}
					// Verify it's a permutation of 1..size.
					seen := make(map[int]bool)
					for _, v := range last {
						seen[v] = true
					}
					for v := 1; v <= size; v++ {
						if !seen[v] {
							t.Errorf("%s n=%d seed=%d: missing value %d", algo, size, seed, v)
						}
					}
				}
				if stats.algo != algo {
					t.Errorf("stats.algo = %s, want %s", stats.algo, algo)
				}
			}
		}
	}
}

// TestDeterminism verifies the same seed produces identical step counts.
func TestDeterminism(t *testing.T) {
	for _, algo := range validAlgos {
		s1, st1 := runAlgo(algo, 40, 42)
		s2, st2 := runAlgo(algo, 40, 42)
		if len(s1) != len(s2) {
			t.Errorf("%s: step count differs across runs: %d vs %d", algo, len(s1), len(s2))
		}
		if st1.comparisons != st2.comparisons || st1.swaps != st2.swaps {
			t.Errorf("%s: metrics differ across runs", algo)
		}
	}
}

// TestSeededArrayIsPermutation verifies seededArray returns 1..n in some order.
func TestSeededArrayIsPermutation(t *testing.T) {
	a := seededArray(40, 7)
	if len(a) != 40 {
		t.Fatalf("size = %d, want 40", len(a))
	}
	sort.Ints(a)
	for i := 0; i < 40; i++ {
		if a[i] != i+1 {
			t.Errorf("not a permutation: pos %d = %d", i, a[i])
		}
	}
}

// TestComparisonOrdering sanity-checks that on the same array, O(n log n)
// algorithms do fewer comparisons than O(n^2) ones for a reasonable size.
func TestComparisonOrdering(t *testing.T) {
	_, bubble := runAlgo("bubble", 60, 42)
	_, quick := runAlgo("quick", 60, 42)
	_, merge := runAlgo("merge", 60, 42)
	if quick.comparisons >= bubble.comparisons {
		t.Errorf("quick (%d) should beat bubble (%d)", quick.comparisons, bubble.comparisons)
	}
	if merge.comparisons >= bubble.comparisons {
		t.Errorf("merge (%d) should beat bubble (%d)", merge.comparisons, bubble.comparisons)
	}
}
