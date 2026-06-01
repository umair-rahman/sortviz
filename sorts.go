package main

// recorder is a callback the sorting algorithms call on every meaningful event.
// It receives a Step and accumulates comparisons/swaps in the counters.
type recorder struct {
	steps       []Step
	comparisons int
	swaps       int
}

// rec records one step, deep-copying the array snapshot so playback is
// independent of later mutations to the working slice.
func (r *recorder) rec(arr []int, kind StepKind, a, b int) {
	r.steps = append(r.steps, Step{
		snapshot: copySlice(arr),
		kind:     kind,
		a:        a,
		b:        b,
	})
	switch kind {
	case StepCompare:
		r.comparisons++
	case StepSwap:
		r.swaps++
	}
}

// runAlgo builds a seeded array, runs the named algorithm with a recorder,
// and returns the recorded steps plus the collected metrics.
// It is the single entry point used by the FSM (Phase 4) and STATS (Phase 5).
func runAlgo(name string, size int, seed int64) ([]Step, RunStats) {
	arr := seededArray(size, seed)
	rec := &recorder{}

	switch name {
	case "bubble":
		bubbleSort(arr, rec)
	case "insertion":
		insertionSort(arr, rec)
	case "selection":
		selectionSort(arr, rec)
	case "quick":
		quickSort(arr, 0, len(arr)-1, rec)
	case "merge":
		mergeSort(arr, 0, len(arr)-1, rec)
	}

	stats := RunStats{
		algo:        name,
		comparisons: rec.comparisons,
		swaps:       rec.swaps,
		size:        size,
		seed:        seed,
	}
	return rec.steps, stats
}

// ── Bubble Sort ───────────────────────────────────────────────────────────────

func bubbleSort(a []int, r *recorder) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			r.rec(a, StepCompare, j, j+1)
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				r.rec(a, StepSwap, j, j+1)
			}
		}
		r.rec(a, StepLock, n-1-i, -1)
	}
	if n > 0 {
		r.rec(a, StepLock, 0, -1) // lock the last remaining element
	}
}

// ── Insertion Sort ────────────────────────────────────────────────────────────

func insertionSort(a []int, r *recorder) {
	n := len(a)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			r.rec(a, StepCompare, j-1, j)
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
				r.rec(a, StepSwap, j-1, j)
				j--
			} else {
				break
			}
		}
		r.rec(a, StepLock, i, -1)
	}
	if n > 0 {
		r.rec(a, StepLock, 0, -1)
	}
}

// ── Selection Sort ────────────────────────────────────────────────────────────

func selectionSort(a []int, r *recorder) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			r.rec(a, StepCompare, minIdx, j)
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			a[i], a[minIdx] = a[minIdx], a[i]
			r.rec(a, StepSwap, i, minIdx)
		}
		r.rec(a, StepLock, i, -1)
	}
	if n > 0 {
		r.rec(a, StepLock, n-1, -1)
	}
}

// ── Quick Sort ────────────────────────────────────────────────────────────────

func quickSort(a []int, lo, hi int, r *recorder) {
	if lo >= hi {
		if lo == hi {
			r.rec(a, StepLock, lo, -1)
		}
		return
	}
	p := partition(a, lo, hi, r)
	quickSort(a, lo, p-1, r)
	quickSort(a, p+1, hi, r)
}

func partition(a []int, lo, hi int, r *recorder) int {
	pivot := hi
	r.rec(a, StepPivot, pivot, -1)
	i := lo
	for j := lo; j < hi; j++ {
		r.rec(a, StepCompare, j, pivot)
		if a[j] <= a[pivot] {
			a[i], a[j] = a[j], a[i]
			r.rec(a, StepSwap, i, j)
			i++
		}
	}
	a[i], a[hi] = a[hi], a[i]
	r.rec(a, StepSwap, i, hi)
	r.rec(a, StepLock, i, -1)
	return i
}

// ── Merge Sort ────────────────────────────────────────────────────────────────

func mergeSort(a []int, lo, hi int, r *recorder) {
	if lo >= hi {
		if lo == hi {
			r.rec(a, StepLock, lo, -1)
		}
		return
	}
	mid := (lo + hi) / 2
	mergeSort(a, lo, mid, r)
	mergeSort(a, mid+1, hi, r)
	merge(a, lo, mid, hi, r)
}

func merge(a []int, lo, mid, hi int, r *recorder) {
	tmp := make([]int, hi-lo+1)
	i, j, k := lo, mid+1, 0
	for i <= mid && j <= hi {
		r.rec(a, StepCompare, i, j)
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}
	for i <= mid {
		tmp[k] = a[i]
		i++
		k++
	}
	for j <= hi {
		tmp[k] = a[j]
		j++
		k++
	}
	for idx, v := range tmp {
		a[lo+idx] = v
		r.rec(a, StepSwap, lo+idx, -1)
	}
	for idx := lo; idx <= hi; idx++ {
		r.rec(a, StepLock, idx, -1)
	}
}
