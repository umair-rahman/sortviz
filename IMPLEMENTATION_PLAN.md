# SortViz — Implementation Plan (Zero → Fully Functional)

> Phase-wise build plan for the SortViz terminal app (Code Olympics 2026).
> Each phase is independently runnable/testable so we never sit in a broken state.
> **Two hard rules enforced in every phase:** exactly **3 states**, total **≤ 650 lines**.

---

## Guiding Principles

1. **Always runnable.** Every phase ends with `go run .` working. We never commit a broken build.
2. **Vertical slices.** Build the skeleton end-to-end first (all 3 modes navigable with fake data),
   then deepen each mode. This guarantees a demoable product even if we run out of time.
3. **Watch the budget.** Run `make count` (line counter) after every phase. Stay ≤ 650.
4. **Reliability first.** Input validation + graceful exit go in early, not bolted on at the end.
5. **Determinism.** Seed-driven data from day one so animations are reproducible.

---

## Line Budget Tracker (target per phase, cumulative)

| Phase | Adds | Cumulative | Cap |
|---|---|---|---|
| 0 Setup | ~15 | ~15 | 650 |
| 1 Skeleton FSM | ~120 | ~135 | 650 |
| 2 Data + render engine | ~150 | ~285 | 650 |
| 3 Algorithms + recorder | ~180 | ~465 | 650 |
| 4 RUN animation | ~70 | ~535 | 650 |
| 5 STATS + racing chart | ~60 | ~595 | 650 |
| 6 Polish + Windows VT + exit | ~20 | ~615 | 650 |
| 7 Hardening/tests | 0 (test files don't count toward app) | ~615 | 650 |

> If any phase pushes us over, we trim before moving on (the cheapest cut is the optional 5th
> algorithm or verbose help text).

---

## Phase 0 — Project Setup (15 min)

**Goal:** empty but runnable Go module.

Tasks:
- `go mod init sortviz` (Go 1.21+).
- Create `main.go` with a `func main()` that prints `SortViz` and exits.
- Add `README.md` with run instructions.
- Add a tiny line-counter script / `Makefile` target so we can track the budget:
  - `make run`   → `go run .`
  - `make build` → `go build -o sortviz .`
  - `make count` → counts non-blank, non-comment `.go` lines (exclude `_test.go`).
- `git init`, first commit.

**Exit check:** `go run .` prints a banner. `make count` returns a number.

---

## Phase 1 — Skeleton FSM (the spine) (1–1.5 hr)

**Goal:** All **3 modes** exist and you can navigate SELECT → RUN → STATS → SELECT with
**placeholder** content. This locks the core constraint in on day one.

Files:
- `model.go`
  - `type Mode int` + `const (ModeSelect, ModeRun, ModeStats)`.
  - `type AppState struct` holding `mode`, config (`algo`, `size`, `speed`, `seed`), and a
    placeholder for steps/stats.
- `main.go`
  - The FSM loop: `for { switch state.mode { ... } }`.
  - Input goroutine: reads lines from stdin → sends on a `chan string`.
  - Mode handlers (stubbed): each prints its name + menu and reads a command to transition.
  - Global `Q` → quit; `Esc`/`b` → back where defined.
- `ui.go` (stub)
  - `drawSelect()`, `drawStats()` print static placeholder boxes for now.

Reliability baked in here:
- Unknown command → re-prompt, no crash.
- Quit path restores terminal (cursor on, colors reset) — even if it's just a `fmt.Print(reset)`.

**Exit check:** Launch → SELECT box → type run → RUN placeholder → it returns to STATS placeholder
→ back to SELECT → `Q` quits cleanly. The 3-state contract is now provable.

---

## Phase 2 — Data Model + ASCII Render Engine (2–2.5 hr)

**Goal:** Draw a real colored ASCII bar chart from an array. This is the Visual Creation core.

Files:
- `model.go` (extend)
  - `type Step struct { snapshot []int; kind StepKind; a, b int }` (kind = compare/swap/lock).
  - `type Stats struct { algo string; comparisons, swaps int; dur time.Duration }`.
- `render.go` (new)
  - `colors` constants (reset/cyan/yellow/red/green).
  - `seededArray(size, seed int) []int` — deterministic shuffle of `1..size`.
  - `drawBars(a []int, highlight map/indices, height int)` — vertical bars, top-down, scaled,
    color per element state.
  - Frame control helpers: `home()` (`\x1b[H`), `clearEnd()` (`\x1b[J`), `hideCursor()/showCursor()`.
  - `drawHBar(label string, value, max, width int)` — one horizontal bar (reused later in STATS).

Wire-in:
- SELECT now shows a static preview of the seeded array as bars (proves the engine works).

**Exit check:** SELECT renders a real, colored, correctly-scaled bar chart of the seeded array.
Resize-safe (fixed width). `make count` still well under cap.

---

## Phase 3 — Algorithms + Recorder (2.5–3 hr)

**Goal:** Five instrumented sorting algorithms that produce a replayable `[]Step` and metrics.

Files:
- `sorts.go` (new)
  - `type recorder func(Step)` plus a small struct/closure that also counts comparisons & swaps.
  - Helper constructors: `compareStep`, `swapStep`, `lockStep` (each deep-copies the current array
    slice into the snapshot so playback is independent of later mutations).
  - Implement: `bubbleSort`, `insertionSort`, `selectionSort`, `quickSort`, `mergeSort` — each
    taking `(a []int, rec recorder)`.
  - `runAlgo(name string, a []int) (steps []Step, stats Stats)` dispatcher.

Correctness checks (informal, via a scratch `_test.go` that won't count toward budget):
- After running, the array is sorted.
- comparisons/swaps counts are > 0 and sane.

**Exit check:** Calling `runAlgo("quick", arr)` returns a non-empty step list and correct stats.
Sort output verified correct for all 5 algos. (Test file excluded from line count.)

---

## Phase 4 — RUN Mode Animation (1.5–2 hr)

**Goal:** The live show. Replay recorded steps on a ticker with color-coded highlights.

Tasks (in `main.go` RUN handler + small render helper):
- On entering RUN: build seeded array, `runAlgo(...)` → `steps`, reset `stepIndex = 0`.
- `time.Ticker` with interval from `speed` (Slow=120ms / Normal=60ms / Fast=20ms).
- Main loop uses `select`:
  - `case <-ticker.C:` → advance `stepIndex`; redraw frame `steps[stepIndex]` with the right
    highlight colors (compare=yellow, swap=red, locked=green, rest=cyan); update live counters.
  - `case cmd := <-inputCh:` → `Enter` skips to last step; `Q` quits.
- When `stepIndex == len(steps)-1`: stop ticker, record stats into the session map, transition to
  STATS automatically.

Reliability:
- Stop the ticker on exit (no leak). Guard empty step list. Final frame shows all-green (sorted).

**Exit check:** Pick Bubble → watch bars animate with correct colors and moving counters → it
finishes and lands on STATS. Try Quick/Merge too. Deterministic across re-runs with same seed.

---

## Phase 5 — STATS Mode + Same-Array Racing Chart (1.5 hr)

**Goal:** The proof screen and the innovation centerpiece.

Tasks (in `ui.go` STATS handler):
- Show last run's metrics: comparisons, swaps, time, stability, Big-O row.
- Maintain `session map[string]Stats` — every algorithm run this session stores its comparison
  count for the current seed/size.
- Render comparative horizontal bar chart using `drawHBar` (reuse from Phase 2): one bar per algo
  that has been run, scaled to the max, labeled with the count.
- Optional nicety: if an algo hasn't been run yet, show it greyed with `--`.
- `Enter` → back to SELECT; `Q` → quit.

**Exit check:** Run 2–3 different algorithms; STATS shows each one's bar growing/shrinking
relative to others on the same dataset. The Big-O difference is now visually obvious.

---

## Phase 6 — Cross-Platform Polish + Graceful Exit (1 hr)

**Goal:** Runs cleanly on Linux/macOS/Windows and never leaves the terminal broken.

Files:
- `term_windows.go` (`//go:build windows`) — `enableVT()` calls `kernel32.SetConsoleMode` via
  stdlib `syscall` to turn on Virtual Terminal Processing.
- `term_other.go` (`//go:build !windows`) — `enableVT()` is a no-op.
- In `main.go`:
  - Call `enableVT()` at startup.
  - `signal.Notify` for `os.Interrupt`; on signal → `showCursor()`, reset colors, clear, `os.Exit(0)`.
  - `defer` a cleanup that restores the terminal on any normal exit too.

**Exit check:** Build and run on Windows (your machine) — colors render, `Ctrl-C` exits cleanly,
cursor visible, no leftover escape junk. Confirm `go vet ./...` is clean.

---

## Phase 7 — Hardening, Verification & Submission Prep (1–1.5 hr)

**Goal:** Lock down reliability (40%) and presentation.

Tasks:
- **Edge-case pass:** size below 10 / above 60, non-numeric size, unknown algo, empty input,
  spamming Enter, immediate quit in each mode. Fix any crash; everything should re-prompt safely.
- **Determinism check:** same seed twice → identical step count and identical final stats.
- **Line budget check:** `make count` ≤ 650. Trim if needed (drop optional 5th algo or compress
  help text). Record the final number in README.
- **`gofmt -l .` and `go vet ./...`** → both clean.
- **Constraint self-audit** (mirror what the automated grader checks):
  - State count == 3 (grep the `Mode` enum; no hidden 4th screen).
  - No external deps: `go.mod` has no `require` block; `go list -m all` shows only the module.
  - Domain == terminal visual: yes.
  - Language == Go: yes.
- **README finalize:** one-command run, screenshots/gif of all 3 modes, line count, "zero deps"
  note, constraint table.
- **Demo script:** a 60-second walkthrough — Select Bubble (slow) → Run → Stats → Select Quick →
  Run → Stats (show the racing chart shrink). Optional asciinema recording for Community Choice.

**Exit check:** Fresh clone on a clean machine → `go run .` works first try. No crashes across the
full edge-case checklist. Budget and constraints verified.

---

## Suggested Schedule (3-Day Window)

| Day | Focus | Phases |
|---|---|---|
| **Day 1** | Skeleton + visual engine working end-to-end | 0, 1, 2 |
| **Day 2** | Real algorithms + the live animation | 3, 4 |
| **Day 3** | Stats/racing chart + cross-platform polish + harden + submit | 5, 6, 7 |

> If time gets tight, the priority order to *cut* is: optional 5th algorithm → racing chart extra
> labels → speed presets. Never cut: 3 working modes, deterministic run, graceful exit.

---

## Definition of Done (Submission Checklist)

- [ ] `go run .` works on a fresh clone with zero setup.
- [ ] Exactly 3 modes (SELECT / RUN / STATS), no more.
- [ ] All 5 (or 4) algorithms animate correctly and finish sorted.
- [ ] Live color highlights: compare=yellow, swap=red, locked=green.
- [ ] STATS shows metrics + same-array comparative bar chart.
- [ ] Same seed → identical reproducible animation.
- [ ] No crash across the full edge-case checklist.
- [ ] `Ctrl-C` / quit restores the terminal cleanly.
- [ ] Zero external dependencies (`go.mod` clean).
- [ ] Total `.go` lines ≤ 650 (record exact count in README).
- [ ] `gofmt` + `go vet` clean.
- [ ] README with run command, constraint table, screenshots.

---

## Risk Register (and mitigations)

| Risk | Impact | Mitigation |
|---|---|---|
| Line budget overflow | Auto-fail on constraint | Track every phase; keep ~35-line buffer; trim list ready |
| Windows ANSI not rendering | Looks broken to judge | `enableVT()` via stdlib syscall; tested on your Windows box |
| Animation flicker | Hurts polish | Cursor-home + clear-to-end redraw, not full clear |
| Idle/boring visual | Weak demo | Seeded shuffled array guarantees a lively, varied animation |
| Goroutine/ticker leak | Reliability ding | Stop ticker + close input path on every mode exit |
| Accidentally adding a 4th "screen" | Constraint break | Help/overlays are drawn *within* a mode, never a new Mode value |
