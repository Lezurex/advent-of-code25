# Advent of Code 2025 🎄

This repository contains my solutions for the [Advent of Code 2025](https://adventofcode.com/2025) challenges, written
in **Go**. However, they may not be the most concise nor the most optimized.

## 📂 Project Structure

```plaintext
advent-of-code25/
│
├── go.mod
│
├── solutions/             # Daily solutions
│   ├── day01.go           # Solution for Day 1
│   ├── day01_test.go      # Test for Day 1
│   ├── day02.go           # Solution for Day 2
│   ├── day02_test.go      # Test for Day 2
│   ├── ...
│   └── inputs/            # Insert inputs here
│       ├── day01.txt      # Input file for Day 1 
│       ├── day02.txt      # Input file for Day 2
│       └── ...
│
└── utils/                 # Reusable utilities
    └── file.go            # Functions for reading input files
```

## 🚀 Getting Started

1. Clone this repository:
   ```bash
   git clone https://github.com/Lezurex/advent-of-code25.git
   cd advent-of-code24
   ```

2. Prepare your input files:
    - Place your inputs in the `solutions/inputs/` directory.
    - Name them as `dayXX.txt` (e.g., `day01.txt` for Day 1).

3. Run a solution:
    - Run all tests:
      ```bash
      go test ./solutions -v
      ```
    - Run a specific day:
      ```bash
      go test ./solutions -v -run '^TestSolveDayXX$'
      ```

Happy coding! 🎅