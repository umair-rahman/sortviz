# SortViz

A zero-dependency, terminal-based sorting algorithm visualizer written in pure Go.
Pick an algorithm, watch it sort live as color-coded ASCII bars, then compare how
algorithms actually perform on the same data — all in exactly 3 modes.

> Code Olympics 2026 entry.
> Constraints: **3 states** (Simple-State Creator) · **≤650 lines** (Enterprise Creator)
> · **terminal Visual Creation** domain · **Go** · **zero external dependencies**.

## Run

```bash
go run .
```

Or build a single portable binary:

```bash
go build -o sortviz .
./sortviz
```

No `go get`, no config, no network. Just the Go standard library.

## Developer tasks

```bash
make run      # go run .
make build    # build ./sortviz binary
make count    # count Go lines vs the 650 budget
make check    # gofmt + go vet
```

## Status

- [x] Phase 0 — project setup (module, skeleton, tooling)
- [x] Phase 1 — 3-state FSM skeleton (SELECT / RUN / STATS, input validation, graceful exit)
- [x] Phase 2 — data model + ASCII render engine (Step/RunStats structs, seededArray, drawBars, drawHBar, live preview in SELECT)
- [x] Phase 3 — instrumented algorithms + recorder (bubble/insertion/selection/quick/merge, record-then-replay)
- [x] Phase 4 — RUN animation (time.Ticker + goroutine + select, color-coded live frames, auto-advance to stats)
- [x] Phase 5 — STATS + same-array comparison chart (session metrics, racing bar chart)
- [x] Phase 6 — cross-platform polish (Windows VT via stdlib syscall, Ctrl-C terminal restore)
- [x] Phase 7 — hardening + verification (unit tests, E2E tests, gofmt/vet clean, 598/650 lines)

All phases complete. 598 / 650 lines, zero dependencies, exactly 3 modes.

## How to use

1. `go run .` opens the **SELECT** screen with a live preview of the seeded array.
2. Configure the run:
   - `a quick` — set algorithm (bubble / insertion / selection / quick / merge)
   - `s 30` — set array size (10–60)
   - `p fast` — set speed (slow / normal / fast)
   - `d 42` — set seed (any integer; same seed = same animation)
   - `run` — start the visualization
3. **RUN** mode animates the sort as color-coded bars (yellow = comparing,
   red = swapping/pivot, green = locked, cyan = unsorted). Press Enter to skip
   to the end, or `q` to quit. When the sort finishes it auto-advances to STATS.
4. **STATS** mode shows the run's comparisons/swaps plus a comparison chart of
   every algorithm you have run this session **on the same array** — so the
   O(n²) vs O(n log n) gap is visible at a glance. `back` returns to SELECT.

## Design notes

- **Record-then-replay**: algorithms run to completion first, recording each
  compare/swap/lock as a `Step`; RUN mode replays those steps on a ticker. This
  makes animations deterministic and decouples sorting logic from rendering.
- **Zero dependencies**: all visuals are raw ANSI escape codes; Windows VT mode
  is enabled through the stdlib `syscall` package (see `term_windows.go`).
- **Exactly 3 modes**: the `Mode` enum in `model.go` is the single source of
  truth — SELECT, RUN, STATS. No fourth state exists.
