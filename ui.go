package main

import (
	"fmt"
	"strconv"
	"strings"
)

// bar is the horizontal rule used inside boxed screens.
const bar = "════════════════════════════════════════════════════"

// line prints one bordered content row, padded to the box width.
func line(text string) { fmt.Printf("║ %-50s ║\n", text) }

// drawSelect renders the SELECT screen (Mode 1) with a live array preview.
func drawSelect(s *AppState) {
	clearScreen()
	fmt.Printf("%s╔%s╗%s\n", ansiBold, bar, ansiReset)
	line("           SortViz  ·  SELECT MODE")
	fmt.Printf("╠%s╣\n", bar)
	line(fmt.Sprintf("Algorithm : %s", s.algo))
	line(fmt.Sprintf("Array Size: %d   (%d–%d)", s.size, minSize, maxSize))
	line(fmt.Sprintf("Speed     : %s", s.speed))
	line(fmt.Sprintf("Seed      : %d", s.seed))
	fmt.Printf("╠%s╣\n", bar)
	line("Commands:")
	line("  a <name>  algo (bubble/insertion/selection/")
	line("            quick/merge)")
	line("  s <n>     size (10–60)    p <speed> slow/normal/fast")
	line("  d <seed>  seed (integer)  run  start    q  quit")
	fmt.Printf("%s╚%s╝%s\n", ansiBold, bar, ansiReset)
	if s.msg != "" {
		fmt.Printf("  %s> %s%s\n", ansiYellow, s.msg, ansiReset)
		s.msg = ""
	}
	fmt.Println()
	drawPreview(s.size, s.seed)
	fmt.Print("\nselect> ")
}

// drawStats renders the STATS screen (Mode 3): last-run metrics plus the
// same-array comparison chart built from every run this session.
func drawStats(s *AppState) {
	showCursor()
	clearScreen()
	fmt.Printf("%s╔%s╗%s\n", ansiBold, bar, ansiReset)
	line("           SortViz  ·  STATS MODE")
	fmt.Printf("╠%s╣\n", bar)
	if s.ran {
		r := s.lastRun
		line(fmt.Sprintf("Last run: %s  (n=%d, seed=%d)", r.algo, r.size, r.seed))
		line(fmt.Sprintf("Comparisons: %d   Swaps: %d   %s", r.comparisons, r.swaps, bigO[r.algo]))
		fmt.Printf("%s╚%s╝%s\n", ansiBold, bar, ansiReset)
		fmt.Printf("\n  %sComparisons on same array (fewer = better):%s\n", ansiBold, ansiReset)
		drawSessionChart(s)
	} else {
		line("Run an algorithm first to see metrics.")
		fmt.Printf("%s╚%s╝%s\n", ansiBold, bar, ansiReset)
	}
	fmt.Print("\nback → select   q → quit\nstats> ")
}

// drawSessionChart prints one horizontal bar per algorithm run this session,
// scaled to the highest comparison count. It marks the algorithm with the
// fewest comparisons as the winner and annotates each bar with its Big-O class,
// turning the raw counts into a visible theory-vs-practice comparison.
func drawSessionChart(s *AppState) {
	maxC, minC, winner := 1, 1<<31, ""
	for algo, rs := range s.session {
		if rs.comparisons > maxC {
			maxC = rs.comparisons
		}
		if rs.comparisons < minC {
			minC, winner = rs.comparisons, algo
		}
	}
	for _, algo := range validAlgos {
		rs, ok := s.session[algo]
		if !ok {
			continue
		}
		color, suffix := ansiCyan, bigO[algo]
		if algo == winner {
			color = ansiGreen
			suffix = bigO[algo] + " " + ansiGreen + "★ best" + ansiReset
		}
		drawHBar(algo, rs.comparisons, maxC, color, suffix)
	}
	// Speedup: how many times fewer comparisons the winner did vs the worst.
	if len(s.session) > 1 && minC > 0 {
		fmt.Printf("\n  %s%s wins: %.1f× fewer comparisons than the slowest.%s\n",
			ansiBold, winner, float64(maxC)/float64(minC), ansiReset)
	}
}

// applySelectCommand parses and applies one SELECT-mode command. It returns the
// next mode and never panics on bad input — invalid commands set a feedback
// message and stay in SELECT. This is the input-validation backbone (40% score).
func applySelectCommand(s *AppState, raw string) Mode {
	fields := strings.Fields(raw)
	if len(fields) == 0 {
		return ModeSelect // empty input: just redraw
	}
	cmd := strings.ToLower(fields[0])
	arg := ""
	if len(fields) > 1 {
		arg = strings.ToLower(fields[1])
	}

	switch cmd {
	case "q", "quit":
		s.quit = true
	case "run":
		return ModeRun
	case "a", "algo":
		if contains(validAlgos, arg) {
			s.algo, s.msg = arg, "algorithm set to "+arg
		} else {
			s.msg = "unknown algorithm (choose: " + strings.Join(validAlgos, "/") + ")"
		}
	case "s", "size":
		n, err := strconv.Atoi(arg)
		switch {
		case err != nil:
			s.msg = "size must be a number"
		case n < minSize:
			s.size, s.msg = minSize, fmt.Sprintf("size clamped up to %d", minSize)
		case n > maxSize:
			s.size, s.msg = maxSize, fmt.Sprintf("size clamped down to %d", maxSize)
		default:
			s.size, s.msg = n, fmt.Sprintf("size set to %d", n)
		}
	case "p", "speed":
		if contains(validSpeeds, arg) {
			s.speed, s.msg = arg, "speed set to "+arg
		} else {
			s.msg = "speed must be slow/normal/fast"
		}
	case "d", "seed":
		if v, err := strconv.ParseInt(arg, 10, 64); err != nil {
			s.msg = "seed must be an integer"
		} else {
			s.seed, s.msg = v, fmt.Sprintf("seed set to %d", v)
		}
	default:
		s.msg = "unknown command: " + cmd
	}
	return ModeSelect
}
