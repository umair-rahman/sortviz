<div align="center">

# ⚡ SortViz

### *Watch sorting algorithms come alive — live, in your terminal.*

```
                                           █
                            █             ██
                           ██           ████
                         ████          █████
                        █████        ███████
                      ███████       ████████
                     ████████     ██████████
                    █████████    ███████████
                  ███████████   █████████████
                 ████████████  ██████████████
              ███████████████ ████████████████
             ████████████████ █████████████████
           ██████████████████ ████████████████████
          ███████████████████████████████████████████
        █████████████████████████████████████████████
       ███████████████████████████████████████████████
     █████████████████████████████████████████████████
                              S O R T V I Z
```

**A zero-dependency terminal sorting visualizer in pure Go.**
No npm. No pip. No Docker. No browser. Just `go run .`

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/)
[![Zero Deps](https://img.shields.io/badge/dependencies-0-success?style=flat-square)](#)
[![Lines](https://img.shields.io/badge/lines-619%2F650-blue?style=flat-square)](#)
[![Modes](https://img.shields.io/badge/states-3-orange?style=flat-square)](#)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)](#)
[![Code Olympics](https://img.shields.io/badge/Code%20Olympics-2026-red?style=flat-square)](https://codeolympics.com/)

[**Quick Start**](#-quick-start) · [**How It Works**](#-how-to-use--in-30-seconds) · [**The Magic Moment**](#-the-magic-moment) · [**Architecture**](#%EF%B8%8F-architecture)

</div>

---

## 🎬 The Pitch

You learned sorting algorithms from a **textbook**.
You memorized "Quick is O(n log n), Bubble is O(n²)."
But you never actually *saw* the difference. You never *felt* it.

**SortViz fixes that.**

Pick an algorithm. Watch it sort, bar by bar, in real time. Then run another
on the **same exact data** and watch them race. Big-O stops being a phrase
on a slide — it becomes a number you can point to.

```
  bubble     ████████████████████████████████████   300  O(n²)
  insertion  ██████████████████████████             222  O(n²)
  selection  ████████████████████████████████████   300  O(n²)
  quick      ██████████                              91  O(n log n)
  merge      █████████                               81  O(n log n) ★ best

  merge wins: 3.7× fewer comparisons than the slowest.
```

**That's it. That's the entire pitch.**

---

## ✨ Features

|  |  |
|---|---|
| 🎨 | **Live color-coded animation** — yellow = compare, red = swap, green = sorted |
| 🏁 | **Same-array racing** — run 5 algorithms on identical data and rank them |
| 📐 | **Big-O made visible** — every result tagged with its complexity class |
| 🌱 | **Zero dependencies** — pure Go standard library. Period. |
| 🔁 | **Deterministic** — same seed = identical animation, every single time |
| 🪶 | **Lightweight** — 619 lines of Go, 3 finite states, one binary |
| 🌐 | **Cross-platform** — Windows / Mac / Linux, single command to run |
| 📡 | **Offline-first** — no network calls, no telemetry, no cloud. Ever. |

---

## 🚀 Quick Start

> **Prereq:** [Go 1.21+](https://go.dev/dl/) installed.

```bash
git clone https://github.com/umair-rahman/sortviz.git
cd sortviz
go run .
```

That's the entire setup. No `go get`. No config file. No internet after clone.

### Build a portable binary

```bash
go build -o sortviz .
./sortviz          # Mac / Linux
sortviz.exe        # Windows
```

Share the binary with anyone — they don't even need Go installed. ✨

---

## 🎮 How to Use — in 30 Seconds

SortViz has **exactly 3 screens**. Type a command, hit Enter, watch the magic.

<table>
<tr>
<td width="33%" align="center">

### 1️⃣ SELECT
*Configure your run*

</td>
<td width="33%" align="center">

### 2️⃣ RUN
*Watch it sort live*

</td>
<td width="33%" align="center">

### 3️⃣ STATS
*See who won*

</td>
</tr>
<tr>
<td valign="top">

```
Algorithm : bubble
Array Size: 40
Speed     : normal
Seed      : 42
```

</td>
<td valign="top">

```
█ █ ██  █ ███   █ █
█ ████  █████ ███ █
████████████████████
─────────────────────
🟡 compare 🔴 swap 🟢 done
```

</td>
<td valign="top">

```
bubble  ████████  300
quick   ██        91
merge   █         81 ★
3.7× faster!
```

</td>
</tr>
</table>

### SELECT screen — what to type

| Type this | What it does |
|---|---|
| `a quick` | Pick algorithm: `bubble` / `insertion` / `selection` / `quick` / `merge` |
| `s 30` | Set array size (10–60) |
| `p slow` | Set speed: `slow` / `normal` / `fast` |
| `d 42` | Set seed (any integer — same seed = same animation) |
| `run` | 🎬 **Start the visualization** |
| `q` | Quit |

### RUN screen — what to do

> **Nothing.** Just watch. The animation runs itself.

| Key | Action |
|---|---|
| `Enter` | Skip to the end instantly |
| `q` | Quit |

When the sort finishes → STATS screen opens **automatically**.

### STATS screen — what to type

| Type this | What it does |
|---|---|
| `back` | Return to SELECT (run another algorithm) |
| `q` | Quit |

---

## 🪄 The Magic Moment

Here's the experience that sells SortViz in 60 seconds. Try it:

```bash
go run .
```

Then in the terminal:

```
a bubble       ⏎    # pick the slow one
s 25           ⏎    # decent size
p slow         ⏎    # let it breathe
run            ⏎    # 🎬 watch bubble struggle through 300 comparisons

back           ⏎    # go back

a quick        ⏎    # pick the fast one
run            ⏎    # 🎬 watch it FLY through 91 comparisons

back           ⏎
a merge        ⏎
run            ⏎    # 🎬 even tighter
```

**Now look at the STATS chart.**

```
  bubble     ████████████████████████████████████   300  O(n²)
  quick      ██████████                              91  O(n log n)
  merge      █████████                               81  O(n log n) ★ best

  merge wins: 3.7× fewer comparisons than the slowest.
```

This is the moment Big-O **stops being theory** and becomes a number you can
point at. Every CS class should end with this screen.

---

## 🎨 Color Guide

During the live animation, every bar is colored by what's happening to it:

| Color | Meaning |
|:---:|---|
| 🔵 **Cyan** | Untouched — still in original position |
| 🟡 **Yellow** | Currently being **compared** |
| 🔴 **Red** | Currently being **swapped** (or pivot in Quick Sort) |
| 🟢 **Green** | **Locked** — in its final sorted position |

You can literally watch the green spread across the chart as the sort
converges. It's weirdly satisfying.

---

## 📐 The Algorithms

Five classics, each instrumented to record every operation:

| Algorithm | Best | Average | Worst | Stable? | Vibe |
|---|---|---|---|:---:|---|
| 🫧 **Bubble** | Ω(n) | Θ(n²) | O(n²) | ✅ | The honest workhorse — slow but readable |
| 📥 **Insertion** | Ω(n) | Θ(n²) | O(n²) | ✅ | Surprisingly fast on near-sorted data |
| 🎯 **Selection** | Ω(n²) | Θ(n²) | O(n²) | ❌ | Few swaps, many comparisons |
| ⚡ **Quick** | Ω(n log n) | Θ(n log n) | O(n²) | ❌ | The chaotic genius. Usually wins. |
| 🔀 **Merge** | Ω(n log n) | Θ(n log n) | O(n log n) | ✅ | Predictable, stable, never breaks down |

---

## 🏗️ Architecture

```
┌──────────────────────────────────────────────────────┐
│                   main.go (FSM loop)                  │
│      SELECT ─────► RUN ──────────────────► STATS      │
│         ▲                                     │       │
│         └──────────── back ───────────────────┘       │
└────────────────────────┬─────────────────────────────┘
                         │
       ┌─────────────────┼─────────────────┐
       ▼                 ▼                 ▼
   sorts.go          render.go           ui.go
  (5 algos +        (ASCII bars +       (screens +
   recorder)         ANSI colors)        input parser)
```

**Three core ideas keep SortViz tight, fast, and reproducible:**

🎬 **Record-then-replay** — algorithms run to completion first, recording
every compare/swap/lock as a `Step`. RUN mode replays those steps on a
`time.Ticker`. Result: deterministic animations, decoupled from sorting logic.

🌱 **Pure stdlib** — ANSI escape codes do all the visual work. Windows
Virtual Terminal mode is enabled via stdlib `syscall`. **No external
packages exist in `go.mod`.**

🎯 **Exactly 3 states** — `ModeSelect`, `ModeRun`, `ModeStats`. The enum
in `model.go` is the single source of truth. No fourth state can sneak in.

---

## 🎯 Built for Code Olympics 2026

SortViz is a competition entry built under deliberately tight constraints.
Here's the scorecard:

| Constraint | Requirement | Result |
|---|---|:---:|
| **Simple-State Creator** | Exactly 2–3 modes | ✅ 3 modes |
| **Enterprise Creator** | ≤ 650 lines | ✅ 619 / 650 |
| **Visual Creation** | ASCII art / terminal UI | ✅ Live bar chart + race chart |
| **Language: Go** | Pure Go | ✅ Zero dependencies |

> *Most hackathons let you reach for any framework. Code Olympics strips
> that away. SortViz is what's left when you only have fundamentals —
> and it turns out, that's enough.*

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

## 🧪 Run the Tests

```bash
go test ./...
```

Verifies:
- ✅ All 5 algorithms produce correctly sorted output (100 combinations)
- ✅ Same seed always produces identical step counts (determinism)
- ✅ O(n log n) algorithms outperform O(n²) on the same data

---

## 🤝 Contributing

This is a hackathon submission with a hard 650-line budget — additions
are welcome but must respect the constraint. Open an issue first to
discuss any new feature.

---

## 📜 License

MIT — do whatever you want, just keep the credit.

---

<div align="center">

### Built with pure Go.
### No frameworks. No libraries. No excuses.

**Just fundamentals.**

⭐ *If this made you smile or learn something, drop a star.*

</div>
