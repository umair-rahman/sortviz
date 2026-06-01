# ⚡ SortViz

> **Watch sorting algorithms fight it out — live, in your terminal.**

```
  SortViz · RUN  |  quick   comparisons: 91  swaps: 55   step 142/188

  █                        █
  █        █          █    █
█ █ ██  █ ███   █ █   █  █ █
█ ████  █████ ███ █   █  ███
████████████████████ ████████
─────────────────────────────
  █ unsorted  █ comparing  █ swapping  █ sorted
```

A **zero-dependency** terminal sorting visualizer built in pure Go.
Pick an algorithm, watch it sort in real time as color-coded ASCII bars,
then see the hard numbers that prove *why* one algorithm crushes another.

**No npm. No pip. No Docker. Just Go.**

---

## ✨ Why SortViz?

Most sorting visualizers are heavy web apps that need a browser, an internet
connection, and a dozen dependencies. SortViz needs **nothing** — clone it,
run it, watch it work.

- 🎯 **Real-time animation** — every compare, swap, and lock rendered live
- 📊 **Same-array racing chart** — run 5 algorithms on identical data, see who wins
- 🔢 **Big-O made tangible** — "merge is 3.7× faster than bubble" is more convincing than any textbook
- 🌱 **Zero dependencies** — pure Go standard library, works offline, forever
- 🔁 **Deterministic** — same seed = same animation, every single time
- 💻 **Cross-platform** — Windows, Mac, Linux, one binary

---

## 🚀 Quick Start

**Prerequisites:** [Go 1.21+](https://go.dev/dl/) installed

```bash
git clone https://github.com/umair-rahman/sortviz.git
cd sortviz
go run .
```

That's it. No `go get`. No config. No internet after clone.

### Or build a portable binary

```bash
go build -o sortviz .
./sortviz        # Mac / Linux
sortviz.exe      # Windows
```

Share the binary with anyone — they don't even need Go installed.

---

## 🎮 How to Use — Step by Step

### Step 1 — SELECT screen (configure your run)

When you launch SortViz, you land on the **SELECT** screen.
Type a command and press **Enter**.

```
╔════════════════════════════════════════════════════╗
║            SortViz  ·  SELECT MODE                 ║
╠════════════════════════════════════════════════════╣
║ Algorithm : bubble                                 ║
║ Array Size: 40   (10–60)                           ║
║ Speed     : normal                                 ║
║ Seed      : 42                                     ║
╚════════════════════════════════════════════════════╝
```

| Command | What it does | Example |
|---|---|---|
| `a <name>` | Choose algorithm | `a quick` |
| `s <n>` | Set array size (10–60) | `s 30` |
| `p <speed>` | Set animation speed | `p slow` |
| `d <seed>` | Set seed (any integer) | `d 42` |
| `run` | **Start the visualization** | `run` |
| `q` | Quit | `q` |

**Available algorithms:** `bubble` · `insertion` · `selection` · `quick` · `merge`

**Tip:** A live **array preview** is shown below the menu so you can see
the shuffled data before you start sorting.

---

### Step 2 — RUN screen (watch the magic)

Type `run` and press Enter. The animation starts immediately.

```
  SortViz · RUN  |  bubble   comparisons: 156  swaps: 89   step 245/525

  █                      
  █        █             
█ █ ██    ███            
█ █ ██  █ ███   █ █      
████████████████████ ████
─────────────────────────
  █ unsorted  █ comparing  █ swapping  █ sorted
```

**Color guide:**

| Color | Meaning |
|---|---|
| 🔵 Cyan | Unsorted — not touched yet |
| 🟡 Yellow | Being **compared** right now |
| 🔴 Red | Being **swapped** / current pivot |
| 🟢 Green | **Sorted** — locked in final position |

**Controls during animation:**

| Key | Action |
|---|---|
| `Enter` | Skip to the end instantly |
| `q` | Quit |

> The animation finishes → **STATS screen opens automatically.** No key needed.

---

### Step 3 — STATS screen (the proof)

After every run, SortViz shows you the numbers:

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

Every algorithm you run this session appears on the chart — **all on the same
seeded array** — so the comparison is perfectly fair.

| Command | Action |
|---|---|
| `back` | Return to SELECT (run another algorithm) |
| `q` | Quit |

---

## 🏆 The Full Experience (recommended first run)

Do this to see SortViz at its best:

```
# 1. Launch
go run .

# 2. In SELECT — set up
a bubble        ← Enter
s 25            ← Enter
p slow          ← Enter
run             ← Enter   (watch bubble struggle)

# 3. In STATS — go back
back            ← Enter

# 4. Run quick sort on the SAME array
a quick         ← Enter
run             ← Enter   (watch it fly)

# 5. In STATS — go back again
back            ← Enter

# 6. Run merge sort
a merge         ← Enter
run             ← Enter

# 7. STATS now shows all three side by side
# You'll see something like:
#   bubble  ████████████████████  300  O(n²)
#   quick   ██████                 91  O(n log n)
#   merge   █████                  81  O(n log n) ★ best
#   merge wins: 3.7× fewer comparisons
```

**This is the moment Big-O stops being theory and becomes a number you can see.**

---

## 📐 Algorithms

| Algorithm | Best | Average | Worst | Stable | When to use |
|---|---|---|---|---|---|
| **Bubble** | Ω(n) | Θ(n²) | O(n²) | ✅ | Learning only |
| **Insertion** | Ω(n) | Θ(n²) | O(n²) | ✅ | Small / nearly sorted data |
| **Selection** | Ω(n²) | Θ(n²) | O(n²) | ❌ | Minimizing swaps |
| **Quick** | Ω(n log n) | Θ(n log n) | O(n²) | ❌ | General purpose (fast) |
| **Merge** | Ω(n log n) | Θ(n log n) | O(n log n) | ✅ | Guaranteed performance |

---

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────┐
│                  main.go  (FSM loop)                 │
│   SELECT ──────► RUN ──────────────► STATS          │
│      ▲                                  │            │
│      └──────────── back ────────────────┘            │
└──────────────────────┬──────────────────────────────┘
                       │
        ┌──────────────┼──────────────┐
        ▼              ▼              ▼
   sorts.go       render.go        ui.go
  (5 algos +    (ASCII bars +    (screens +
   recorder)     ANSI colors)     input)
```

**Key design decisions:**

- **Record-then-replay** — algorithms run to completion first, recording every
  compare/swap/lock as a `Step`. RUN mode replays on a `time.Ticker`.
  Result: deterministic, flicker-free, decoupled from rendering.

- **Zero dependencies** — ANSI colors are raw escape codes. Windows VT mode
  is enabled via stdlib `syscall` (no external package needed).

- **Exactly 3 modes** — `ModeSelect`, `ModeRun`, `ModeStats`. The enum in
  `model.go` is the single source of truth. No fourth state exists.

---

## 🎯 Code Olympics 2026 — Constraint Compliance

| Constraint | Requirement | Result |
|---|---|---|
| Simple-State Creator | Exactly 2–3 modes | ✅ **3 modes** (SELECT / RUN / STATS) |
| Enterprise Creator | ≤ 650 lines | ✅ **619 / 650 lines** |
| Visual Creation | ASCII art / terminal UI | ✅ Live bar chart + comparison chart |
| Go | Pure Go | ✅ **Zero external dependencies** |

---

## 📁 Project Structure

```
sortviz/
├── main.go          ← FSM loop, animation engine, Ctrl-C handler
├── model.go         ← Mode enum, Step, RunStats, AppState
├── sorts.go         ← 5 instrumented sorting algorithms + recorder
├── render.go        ← ASCII bar engine, ANSI colors, seeded array
├── ui.go            ← SELECT / STATS screens, input parser
├── term_windows.go  ← Windows VT processing (stdlib syscall)
├── term_other.go    ← No-op for Mac/Linux
├── sorts_test.go    ← Correctness + determinism tests
└── go.mod           ← Module file (no require block)
```

---

## 🧪 Run Tests

```bash
go test ./...
```

Tests verify:
- All 5 algorithms produce correctly sorted output (100 combinations)
- Same seed always produces identical step counts (determinism)
- O(n log n) algorithms do fewer comparisons than O(n²) on the same data

---

*Built with pure Go. No frameworks. No libraries. Just fundamentals.*
