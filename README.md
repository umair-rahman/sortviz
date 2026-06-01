# SortViz

A zero-dependency, terminal-based sorting algorithm visualizer written in pure Go.
Watch sorting algorithms come alive as color-coded ASCII bars, then compare how
they actually perform on the same data — all inside your terminal, in exactly 3 modes.

> **Code Olympics 2026** submission by umair-rahman
> Constraints: **3 states** · **≤650 lines** · **terminal Visual Creation** · **Go** · **zero external dependencies**

---

## Run

```bash
git clone https://github.com/umair-rahman/sortviz.git
cd sortviz
go run .
```

Or build a single portable binary (no Go install needed on target machine):

```bash
go build -o sortviz .
./sortviz          # Linux / Mac
sortviz.exe        # Windows
```

**Zero dependencies. One command. Instant visual.**

---

## How It Works

SortViz has exactly **3 modes**:

### Mode 1 — SELECT
Configure your run, then type `run` to start.

| Command | What it does |
|---|---|
| `a <name>` | set algorithm: `bubble` / `insertion` / `selection` / `quick` / `merge` |
| `s <n>` | set array size (10–60) |
| `p <speed>` | set speed: `slow` / `normal` / `fast` |
| `d <seed>` | set seed (any integer — same seed = same animation every time) |
| `run` | start the visualization |
| `q` | quit |

A live **array preview** is shown so you can see the shuffled data before sorting.

### Mode 2 — RUN
Watch the algorithm sort in real time as color-coded ASCII bars:

```
  SortViz · RUN  |  quick   comparisons: 91  swaps: 55   step 142/188

  █                      
  █        █             
█ █ ██  █ ███   █ █      
█ ████  █████ ███ █   █  
████████████████████ ████
─────────────────────────
  █ unsorted  █ comparing  █ swapping  █ sorted
```

| Color | Meaning |
|---|---|
| 🔵 Cyan | unsorted |
| 🟡 Yellow | being compared |
| 🔴 Red | being swapped / pivot |
| 🟢 Green | locked in final position |

Press **Enter** to skip to the end. Press **q** to quit.
Animation finishes → **STATS screen opens automatically**.

### Mode 3 — STATS
See the metrics and a same-array comparison chart:

```
╔════════════════════════════════════════════════════╗
║            SortViz  ·  STATS MODE                  ║
╠════════════════════════════════════════════════════╣
║ Last run: merge  (n=25, seed=42)                   ║
║ Comparisons: 81   Swaps: 118   O(n log n)          ║
╚════════════════════════════════════════════════════╝

  Comparisons on same array (fewer = better):
  bubble     ████████████████████████████████████   300  O(n²)
  insertion  ██████████████████████████             222  O(n²)
  selection  ████████████████████████████████████   300  O(n²)
  quick      ██████████                              91  O(n log n)
  merge      █████████                               81  O(n log n) ★ best

  merge wins: 3.7× fewer comparisons than the slowest.
```

Run multiple algorithms on the same seed to see the **O(n²) vs O(n log n)** gap as a real measured number.

| Command | What it does |
|---|---|
| `back` | return to SELECT |
| `q` | quit |

---

## Algorithms

| Algorithm | Avg Complexity | Stable |
|---|---|---|
| Bubble Sort | O(n²) | Yes |
| Insertion Sort | O(n²) | Yes |
| Selection Sort | O(n²) | No |
| Quick Sort | O(n log n) | No |
| Merge Sort | O(n log n) | Yes |

---

## Design Highlights

- **Record-then-replay engine** — algorithms run to completion first, recording every compare/swap/lock as a `Step`. RUN mode replays those steps on a `time.Ticker`. Same seed = identical animation every time (deterministic).
- **Zero external dependencies** — all visuals are raw ANSI escape codes. Windows VT mode enabled via stdlib `syscall` (see `term_windows.go`).
- **Exactly 3 modes** — the `Mode` enum in `model.go` is the single source of truth. No fourth state exists.
- **619 / 650 lines** — well within the Enterprise Creator budget.

---

## Constraint Compliance

| Constraint | Requirement | Status |
|---|---|---|
| Simple-State Creator | exactly 2–3 modes | ✅ 3 modes (SELECT / RUN / STATS) |
| Enterprise Creator | ≤ 650 lines | ✅ 619 / 650 |
| Visual Creation | ASCII art / terminal UI | ✅ live bar chart + comparison chart |
| Go | pure Go | ✅ zero external dependencies |
