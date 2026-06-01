# SortViz developer tasks.

# Run the app directly.
run:
	go run .

# Build a single portable binary.
build:
	go build -o sortviz .

# Count non-blank, non-comment Go lines in app files (excludes *_test.go).
# This tracks the 650-line Code Olympics budget.
count:
	@powershell -NoProfile -Command "$$total=0; Get-ChildItem -Filter *.go | Where-Object { $$_.Name -notlike '*_test.go' } | ForEach-Object { $$lines = Get-Content $$_.FullName | Where-Object { $$t = $$_.Trim(); $$t -ne '' -and -not $$t.StartsWith('//') }; $$n = ($$lines | Measure-Object).Count; Write-Host ('{0,5}  {1}' -f $$n, $$_.Name); $$total += $$n }; Write-Host ('{0,5}  TOTAL (budget 650)' -f $$total)"

# Format and vet.
check:
	gofmt -l .
	go vet ./...

.PHONY: run build count check
