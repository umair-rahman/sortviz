package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// ── ANSI color codes ──────────────────────────────────────────────────────────
// These are the only "graphics library" we need — pure escape sequences.
const (
	ansiReset  = "\x1b[0m"
	ansiCyan   = "\x1b[36m" // unsorted / default bar
	ansiYellow = "\x1b[33m" // comparing
	ansiRed    = "\x1b[31m" // swapping / pivot
	ansiGreen  = "\x1b[32m" // locked in final sorted position
	ansiBold   = "\x1b[1m"
	ansiDim    = "\x1b[2m"
)

// chartHeight is the number of rows in the vertical bar chart.
// Kept fixed so the layout is stable across all array sizes.
const chartHeight = 16

// ── Terminal helpers ──────────────────────────────────────────────────────────

// hideCursor / showCursor suppress the blinking cursor during animation.
func hideCursor() { fmt.Print("\x1b[?25l") }
func showCursor() { fmt.Print("\x1b[?25h") }

// clearScreen moves the cursor home and clears the screen for a full redraw.
func clearScreen() { fmt.Print("\x1b[H\x1b[J") }

// ── Seeded array ──────────────────────────────────────────────────────────────

// seededArray returns a slice [1..size] shuffled deterministically by seed.
// Same seed always produces the same array — judges get reproducible animations.
func seededArray(size int, seed int64) []int {
	a := make([]int, size)
	for i := range a {
		a[i] = i + 1
	}
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

// copySlice returns a deep copy of src so the recorder can snapshot safely.
func copySlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

// ── Vertical bar chart (RUN mode + SELECT preview) ───────────────────────────

// barColor returns the ANSI color for one bar given the current step context.
// a, b are the two active indices; kind tells us what is happening.
func barColor(idx int, kind StepKind, a, b int, locked map[int]bool) string {
	switch {
	case locked[idx]:
		return ansiGreen
	case kind == StepLock && idx == a:
		return ansiGreen
	case kind == StepPivot && idx == a:
		return ansiRed
	case kind == StepSwap && (idx == a || idx == b):
		return ansiRed
	case kind == StepCompare && (idx == a || idx == b):
		return ansiYellow
	}
	return ansiCyan
}

// drawBars renders the array as a vertical ASCII bar chart, coloring each bar
// by its state (unsorted/comparing/swapping/locked). a,b are the highlighted
// indices for this step; locked holds indices already in final position.
func drawBars(arr []int, kind StepKind, a, b int, locked map[int]bool, showLegend bool) {
	if len(arr) == 0 {
		return
	}

	// Find max value for scaling.
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}

	// Build the chart row by row, top to bottom.
	// Each column is one element; bar height = value * chartHeight / maxVal.
	var sb strings.Builder
	for row := chartHeight; row >= 1; row-- {
		for i, v := range arr {
			barH := v * chartHeight / maxVal
			color := barColor(i, kind, a, b, locked)
			if barH >= row {
				sb.WriteString(color)
				sb.WriteString("█")
				sb.WriteString(ansiReset)
			} else {
				sb.WriteString(" ")
			}
		}
		sb.WriteString("\n")
	}
	// Bottom axis line.
	sb.WriteString(ansiDim)
	sb.WriteString(strings.Repeat("─", len(arr)))
	sb.WriteString(ansiReset)
	sb.WriteString("\n")

	fmt.Print(sb.String())

	if showLegend {
		fmt.Printf("  %s█%s unsorted  %s█%s comparing  %s█%s swapping  %s█%s sorted\n",
			ansiCyan, ansiReset,
			ansiYellow, ansiReset,
			ansiRed, ansiReset,
			ansiGreen, ansiReset,
		)
	}
}

// drawPreview draws a compact bar chart for the SELECT screen so the judge
// can see the seeded array before starting the run.
func drawPreview(size int, seed int64) {
	arr := seededArray(size, seed)
	// No highlights, no locked elements for the static preview.
	locked := map[int]bool{}
	fmt.Printf("  %sArray preview%s (seed=%d, n=%d):\n", ansiBold, ansiReset, seed, size)
	drawBars(arr, StepCompare, -1, -1, locked, false)
}

// ── Horizontal bar (STATS mode comparison chart) ─────────────────────────────

// hBarWidth is the maximum character width of a horizontal bar in STATS.
const hBarWidth = 36

// drawHBar prints one labeled horizontal bar scaled to maxVal, followed by the
// value and an optional suffix (e.g. Big-O label or a winner marker).
func drawHBar(label string, value, maxVal int, color, suffix string) {
	if maxVal == 0 {
		maxVal = 1 // guard against division by zero
	}
	filled := value * hBarWidth / maxVal
	if filled < 1 && value > 0 {
		filled = 1 // always show at least one block for non-zero values
	}
	bar := strings.Repeat("█", filled)
	fmt.Printf("  %-10s %s%-*s%s %5d  %s\n",
		label,
		color, hBarWidth, bar, ansiReset,
		value, suffix,
	)
}
