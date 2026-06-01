package main

// Mode is one of the exactly three states SortViz can be in. The Code Olympics
// "Simple-State Creator" constraint allows at most three modes; this enum is the
// single source of truth for that contract — there is intentionally no fourth.
type Mode int

const (
	ModeSelect Mode = iota // 1. configure the run
	ModeRun                // 2. animate the sort
	ModeStats              // 3. show metrics + comparison chart
)

// Configuration bounds for the array size chosen in SELECT mode.
const (
	minSize     = 10
	maxSize     = 60
	defaultSize = 40
)

// validAlgos is the set of algorithms selectable in SELECT mode.
var validAlgos = []string{"bubble", "insertion", "selection", "quick", "merge"}

// validSpeeds is the set of animation speeds (timing wired up in Phase 4).
var validSpeeds = []string{"slow", "normal", "fast"}

// bigO maps each algorithm to its average-case time complexity, shown in STATS
// so the comparison chart connects measured counts to the theory.
var bigO = map[string]string{
	"bubble":    "O(n²)",
	"insertion": "O(n²)",
	"selection": "O(n²)",
	"quick":     "O(n log n)",
	"merge":     "O(n log n)",
}

// StepKind classifies what happened at one recorded animation step.
type StepKind int

const (
	StepCompare StepKind = iota // two elements are being compared
	StepSwap                    // two elements are being swapped
	StepLock                    // one element has reached its final sorted position
	StepPivot                   // element is the current pivot (Quick Sort)
)

// Step is one frame of the animation: a snapshot of the array at a single
// compare/swap/lock event, plus which indices are highlighted and why.
type Step struct {
	snapshot []int    // deep copy of the array at this moment
	kind     StepKind // what kind of operation this step represents
	a, b     int      // primary indices (b unused for StepLock/StepPivot)
}

// RunStats holds the metrics collected during one algorithm run.
// Stored in AppState.session so STATS mode can build the comparison chart.
type RunStats struct {
	algo        string
	comparisons int
	swaps       int
	size        int
	seed        int64
}

// AppState is the entire program state. The FSM loop mutates it and the
// renderers read it. Holding everything in one struct keeps the state machine
// explicit and easy to reason about.
type AppState struct {
	mode    Mode                // current state (one of the three Modes)
	algo    string              // selected algorithm
	size    int                 // array size (minSize..maxSize)
	speed   string              // animation speed name
	seed    int64               // deterministic shuffle seed
	msg     string              // transient feedback shown once on the next SELECT draw
	ran     bool                // whether at least one run has completed
	quit    bool                // set true to exit the FSM loop
	steps   []Step              // recorded steps for the current/last run
	lastRun RunStats            // metrics from the most recent run
	session map[string]RunStats // all runs this session keyed by algo name
}

// contains reports whether v is present in list.
func contains(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}
	return false
}
