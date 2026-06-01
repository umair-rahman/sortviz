# SortViz — Live Algorithm Visualizer (Terminal)

> **Code Olympics 2026 submission**
> A zero-dependency, terminal-based sorting algorithm visualizer written in pure Go.
> Watch sorting algorithms come alive as color-coded ASCII bars, then compare how
> they actually perform — all inside your terminal, in exactly 3 modes.

---

## 1. One-Liner

**SortViz** turns abstract sorting algorithms into a live, color-coded ASCII animation —
you pick an algorithm, watch it sort in real time, then see the hard numbers
(comparisons, swaps, time) that prove *why* one algorithm beats another.

---

## 2. The Problem It Solves

Every CS student and self-taught dev "learns" sorting algorithms from static textbook
diagrams or YouTube videos they can't interact with. The core insight — *why* Quick Sort
crushes Bubble Sort — stays abstract because you never **see and measure** it on the same data.

Existing visualizers are almost all heavy web apps (React + animation libraries). There is
no clean, instant, **offline, zero-setup** tool that:

- Animates the algorithm step-by-step in the terminal
- Highlights exactly which elements are being compared vs swapped
- **Quantifies** the difference (comparison/swap counts) on the *same* dataset

SortViz fills that gap: one Go binary, one command, instant visual + measurable proof.
It is a genuine learning/teaching utility, not a toy.

---

## 3. Constraint Compliance (4/4) ✅

| Constraint | Requirement | How SortViz Satisfies It |
|---|---|---|
| **Core: Simple-State Creator** | Exactly 2–3 modes/states | A strict 3-state finite state machine: `SELECT`, `RUN`, `STATS`. No 4th state exists. |
| **Line Budget: Enterprise Creator** | ≤ 650 lines | Estimated ~615 lines across well-separated files (see §10). Pure stdlib keeps it lean. |
| **Domain: Visual Creation** | ASCII art / charts / terminal UIs | The entire app is an ANSI-colored ASCII chart engine — vertical bar animation + comparative bar charts. |
| **Language: Go** | Build in Go | 100% Go, standard library only, idiomatic concurrency (goroutine + channel + ticker + select). |

**Bonus:** True **zero external dependencies** (only Go stdlib). This was specifically praised by
judges in the 2025 edition and strengthens the "back to fundamentals" ethos of the competition.

---

## 4. The 3 Modes (Finite State Machine)

The whole program is a single FSM. There are **exactly three** states and the transitions
between them are the only control flow the user experiences.

```
                 ┌──────────────────────────────────────────────┐
                 │                                                │
        start    ▼                                                │
      ┌──────────────────┐   Enter / run    ┌──────────────────┐  │ Enter
      │   1. SELECT      │ ───────────────► │    2. RUN        │  │ (new run)
      │  (configure)     │                  │  (animate sort)  │  │
      └──────────────────┘ ◄─────────────── └──────────────────┘  │
              ▲   ▲          Esc (back)              │             │
              │   │                                  │ sort done   │
              │   │ Enter (run again)                ▼             │
              │   │                         ┌──────────────────┐  │
              │   └─────────────────────────│   3. STATS       │──┘
              │      Esc (back to select)   │  (metrics+chart) │
              └─────────────────────────────└──────────────────┘
                        ('Q' from any state quits)
```

State enum (the contract):

```go
type Mode int

const (
    ModeSelect Mode = iota // 1. configure algorithm, size, speed, seed
    ModeRun                // 2. live ASCII animation of the sort
    ModeStats              // 3. metrics + comparative bar chart
)
```

### Mode 1 — SELECT (configure the run)

```
╔══════════════════════════════════════════════════════╗
║                 SortViz — SELECT MODE                  ║
╠══════════════════════════════════════════════════════╣
║  Algorithm :  ► Bubble    Insertion   Selection        ║
║                 Quick      Merge                       ║
║                                                        ║
║  Array Size:  [ 40 ]      (10 – 60)                    ║
║  Speed     :  [ Normal ]  (Slow / Normal / Fast)       ║
║  Seed      :  [ 42 ]      (deterministic shuffle)      ║
╠══════════════════════════════════════════════════════╣
║  Type choice + Enter:  a)lgo  s)ize  p)speed  d)seed   ║
║  [Enter] Run    [Q] Quit                               ║
╚══════════════════════════════════════════════════════╝
```

- User configures: **algorithm, array size, speed, random seed**.
- Seed makes every run **deterministic** — judges reproduce the exact animation.
- Input is validated (size clamped to 10–60, unknown algo rejected). No crash on bad input.

### Mode 2 — RUN (the show)

```
  SortViz — RUN  |  Quick Sort  |  comparisons: 138  swaps: 41  step 142/300

 100 ┤                          █
     │              █     █     █        █
     │        █     █  █  █  █  █  █     █
     │     █  █  █  █  █  █  █  █  █  █  █
     │  █  █  █  █  █  █  █  █  █  █  █  █  █
   0 ┴───────────────────────────────────────
        ▲pivot     ▲▲ comparing      ███ sorted

 Legend:  ░ unsorted   █ comparing(yellow)   █ swapping(red)   █ sorted(green)
 [Enter] skip to end
```

- Vertical ASCII bars, each scaled to the chart height.
- **Color-coded per element state every frame**:
  - default/cyan = untouched
  - yellow = currently being compared
  - red = being swapped / active (pivot)
  - green = locked in final sorted position
- A `time.Ticker` advances one recorded step per tick; speed is the tick interval.
- Live counters update (comparisons, swaps, step index).
- When the sort finishes, the FSM auto-transitions to **STATS**.

### Mode 3 — STATS (the proof)

```
╔══════════════════════════════════════════════════════╗
║                 SortViz — STATS MODE                   ║
╠══════════════════════════════════════════════════════╣
║  Last run: Quick Sort  (n=40, seed=42)                 ║
║    comparisons : 138        swaps : 41                 ║
║    time        : 0.61 s     stable: no                 ║
║    complexity  : best Ω(n log n)  avg Θ(n log n)        ║
║                 worst O(n²)                            ║
╠══════════════════════════════════════════════════════╣
║  Comparison on SAME array (seed=42, n=40):             ║
║                                                        ║
║   Bubble    ████████████████████████████  1521        ║
║   Insertion ██████████████                  742        ║
║   Selection ██████████████████████████████ 1560        ║
║   Quick     ███                             138        ║
║   Merge     █████                           215        ║
║            (bars = comparison count, fewer = better)   ║
╠══════════════════════════════════════════════════════╣
║  [Enter] back to Select    [Q] Quit                    ║
╚══════════════════════════════════════════════════════╝
```

- Shows metrics for the run that just finished.
- **The killer feature:** a comparative ASCII bar chart of **comparison counts across every
  algorithm run this session on the same seeded array**. This makes the abstract
  "O(n²) vs O(n log n)" *visible and measurable* — the real educational payoff.
- Reuses the same bar-drawing code as the live chart (DRY, saves lines).

---

## 5. How It Works Internally (Architecture)

The core trick: **separate "computing the sort" from "showing the sort."** Each algorithm is
instrumented to **record** every meaningful operation as a `Step`. RUN mode then simply
**plays back** the recorded steps on a timer. This gives three big wins:

1. **Deterministic** — same seed → identical step list → identical animation every time.
2. **Reliable** — the renderer never touches sorting logic; nothing can deadlock mid-sort.
3. **Lean** — one renderer serves both live animation and final stats.

```
        ┌──────────────────────────────────────────────────────────┐
        │                      main.go (FSM loop)                    │
        │   reads app state, dispatches to the active mode renderer  │
        └───────────────┬───────────────────────┬──────────────────┘
                        │                        │
        ┌───────────────▼──────┐     ┌───────────▼───────────────┐
        │  input goroutine     │     │   time.Ticker (RUN only)   │
        │  reads stdin → chan  │     │   1 tick = advance 1 step  │
        └───────────────┬──────┘     └───────────┬───────────────┘
                        │   select { case cmd / case tick }      │
                        └───────────────┬────────────────────────┘
                                        ▼
        ┌──────────────────────────────────────────────────────────┐
        │   sorts.go : instrumented algorithms → []Step (recorder)  │
        └───────────────────────────┬──────────────────────────────┘
                                     ▼
        ┌──────────────────────────────────────────────────────────┐
        │   render.go : draw []int as colored ASCII bars + frame    │
        └───────────────────────────┬──────────────────────────────┘
                                     ▼
                              Terminal (ANSI output)
```

### Data flow, step by step

1. **SELECT** → user finalizes `{algo, size, speed, seed}`.
2. On "Run", we build the array from the seed, deep-copy it, and call the chosen algorithm
   with a `record(step Step)` callback. The algorithm sorts the copy and emits a `Step`
   snapshot at every compare/swap/lock event. Result: a `[]Step`.
3. **RUN** → a `time.Ticker` fires every *speed* ms. Each tick advances `stepIndex`, and the
   renderer draws `steps[stepIndex]`. An input goroutine listens for "skip"/"quit" on a channel;
   `select` multiplexes ticker vs input.
4. When `stepIndex` hits the end → record the run's metrics into a session map → switch to **STATS**.
5. **STATS** → render last-run metrics + comparative bar chart from the session map.

---

## 6. Algorithms Included

Five classics, each instrumented to emit `Step` snapshots:

| Algorithm | Why included | Best | Avg | Worst | Stable |
|---|---|---|---|---|---|
| **Bubble** | the "obviously bad" baseline | Ω(n) | Θ(n²) | O(n²) | yes |
| **Insertion** | great on near-sorted data | Ω(n) | Θ(n²) | O(n²) | yes |
| **Selection** | constant swaps, many compares | Ω(n²) | Θ(n²) | O(n²) | no |
| **Quick** | the divide-and-conquer star | Ω(n log n) | Θ(n log n) | O(n²) | no |
| **Merge** | stable, predictable n log n | Ω(n log n) | Θ(n log n) | O(n log n) | yes |

Each is a normal Go function plus an `emit` callback, e.g.:

```go
// record is called on every compare / swap so RUN can replay it later.
func bubbleSort(a []int, rec recorder) {
    n := len(a)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-1-i; j++ {
            rec(compareStep(a, j, j+1))        // highlight compared pair
            if a[j] > a[j+1] {
                a[j], a[j+1] = a[j+1], a[j]
                rec(swapStep(a, j, j+1))       // highlight swap
            }
        }
        rec(lockStep(a, n-1-i))                // mark element as sorted
    }
}
```

The `recorder` also tallies `comparisons` and `swaps`, so metrics fall out for free.

---

## 7. ASCII Rendering Engine

The visual heart of the project (Domain = Visual Creation).

- The chart area is a fixed grid: `H` rows tall (e.g. 18) × `n` columns wide (one per element).
- Each value is scaled: `barHeight = value * H / maxValue`.
- We draw **top-down**: for each row `r` from `H` down to `1`, for each bar, print a block char
  `█` if `barHeight >= r`, else a space — wrapped in the right ANSI color for that bar's state.
- **Frame redraw** uses ANSI cursor-home (`\x1b[H`) + clear-to-end (`\x1b[J`) instead of full
  clear, to avoid flicker.
- Colors via ANSI SGR codes:

```go
const (
    reset  = "\x1b[0m"
    cyan   = "\x1b[36m"   // unsorted
    yellow = "\x1b[33m"   // comparing
    red    = "\x1b[31m"   // swapping / pivot
    green  = "\x1b[32m"   // sorted/locked
)
```

The **same** bar primitive renders the horizontal comparative chart in STATS (rotated logic),
so we get two visuals from one piece of code.

---

## 8. Cross-Platform & Zero-Dependency Strategy

Goal: **pure Go standard library, runs on Linux / macOS / Windows.**

- **ANSI on Unix**: works out of the box.
- **ANSI on Windows**: modern terminals need Virtual Terminal Processing enabled. We do this with
  the **stdlib** `syscall` package calling `kernel32.SetConsoleMode` — *no external module*.
  Done in a build-tagged file so it compiles cleanly per OS:

```
term_windows.go   // //go:build windows  → enables VT via syscall+kernel32
term_other.go     // //go:build !windows → no-op
```

- **Input model (deliberate reliability choice):** SortViz uses **line-buffered stdin** (read a
  short command + Enter) rather than raw single-key mode. This avoids per-OS termios/raw-mode
  syscalls entirely → maximum reliability (40% of score) with zero deps. The animation itself
  auto-plays, so interactivity lives in SELECT (before the run) where typing is natural.
- **Graceful exit:** on quit / `Ctrl-C` (caught via `os/signal`), we restore the cursor, reset
  colors, and clear the screen so the judge's terminal is never left in a broken state.

> Optional upgrade (documented, not default): swapping to `golang.org/x/term` enables slick
> single-keypress controls and live pause/step. We keep it **off** by default to preserve the
> zero-dependency claim, which scores better here.

---

## 9. Go Concepts Used (Idiomatic Showcase)

| Concept | Where it's used |
|---|---|
| **Goroutine** | Background stdin reader so the UI never blocks on input |
| **Channel** | Input commands flow from the reader goroutine to the main loop |
| **`select`** | Multiplexes the animation ticker against keyboard commands |
| **`time.Ticker`** | Drives frame timing in RUN; interval = chosen speed |
| **Closures / callbacks** | The `recorder` closure tallies metrics while capturing steps |
| **Build tags** | Per-OS terminal setup with no external dependency |
| **Slices & deep copy** | Step snapshots; same seeded array reused across algorithms |

---

## 10. File Structure & Line Budget

```
sortviz/
├── main.go            ← entry point + FSM loop + mode dispatch     ~90
├── model.go           ← Mode enum, Step, Stats, AppState structs   ~50
├── sorts.go           ← 5 instrumented algorithms + recorder       ~180
├── render.go          ← ASCII bar engine, colors, frame control    ~140
├── ui.go              ← SELECT menu + STATS screen + input parse    ~120
├── term_windows.go    ← enable VT on Windows (syscall, build-tag)   ~25
├── term_other.go      ← no-op for non-Windows (build-tag)           ~10
├── go.mod             ← module file (no require block)              (n/a)
└── README.md          ← how to run (not counted toward budget)     (n/a)
                                                          TOTAL  ≈ 615 / 650 ✅
```

The budget is intentionally kept ~35 lines under the cap as safety margin for edge-case handling.

---

## 11. Innovative Features (Innovation 10% — the tie-breaker)

1. **Same-array algorithm racing in STATS** — the comparative bar chart runs every algorithm on
   the *identical seeded dataset*, turning Big-O theory into a number you can *see*. This is the
   genuinely novel angle most visualizers miss.
2. **Record-then-replay engine** — deterministic, reproducible animations from a seed. Judge runs
   it twice, gets the exact same show. Rare and reliability-friendly.
3. **One renderer, two visuals** — the live vertical animation and the horizontal stats chart share
   the same primitive — elegant constraint-mastery within the line budget.
4. **True zero-dependency cross-platform terminal graphics** — including the Windows VT shim via
   stdlib `syscall`. Most teams reach for a TUI library; we don't need one.

---

## 12. Reliability & Edge Cases (Functionality 40%)

- Array size clamped to `[10, 60]`; invalid input re-prompts, never crashes.
- Unknown algorithm/command → friendly re-prompt.
- Empty / whitespace input ignored safely.
- `Ctrl-C` and normal quit both restore the terminal (cursor, colors, screen).
- Duplicate values handled (stable vs unstable behavior shown honestly).
- Tiny arrays (n=10) and full arrays (n=60) both fit the chart via scaling.
- No goroutine leak: input goroutine and ticker are stopped on mode exit / quit.

---

## 13. How To Run (Judge Experience)

```bash
git clone <repo>
cd sortviz
go run .            # runs immediately, zero setup

# or build a single portable binary
go build -o sortviz .
./sortviz
```

**Zero dependencies. One command. Instant visual.** No `go get`, no config, no network.

---

## 14. Scoring Map — How We Win

| Criteria | Weight | Our Edge |
|---|---|---|
| **Functionality & Reliability** | 40% | 3 fully working modes, deterministic replay, restores terminal, never crashes on bad input |
| **Constraint Mastery** | 30% | Exactly 3-state FSM, true zero-dep, ~615/650 lines, one renderer for two visuals |
| **Code Quality** | 20% | Idiomatic Go, clean package layout, concurrency used purposefully, readable |
| **Innovation** | 10% | Same-array algorithm racing + deterministic record/replay engine |

---

## 15. Optional Stretch (only if line budget & time allow)

- Live pause / single-step (requires `x/term` — would cost the zero-dep claim, so guard it).
- A `--seed`/`--algo` CLI flag for instant headless demo.
- Extra algorithm (Heap) if lines remain — but never add a 4th *mode*.

> Hard rule: **never exceed 3 states and never exceed 650 lines.** Both are graded automatically.
