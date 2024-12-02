# Advent of Code Template

This repository serves as a template for solving [Advent of Code](https://adventofcode.com) puzzles using Go. It is designed to streamline the process of setting up daily puzzle solutions, utilizing the Cobra CLI to manage different years and days efficiently.  

## Purpose

The goal of this repository is to provide an organized structure for tackling Advent of Code challenges across multiple years and days. It allows for:  

- **Year-based organization**: Each year (e.g., 2023) has its own directory and commands.  
- **Day-based puzzle execution**: You can run solutions for specific days by using command-line arguments.  
- **Easy Running**: It leverages [Cobra](https://github.com/spf13/cobra) to enable CLI commands and provide flexibility for running different puzzles.  

## Installation

Clone the repository and navigate to the root directory:

```bash
git clone https://github.com/iframeskills/aoc-template.git
cd aoc-template
```

You need to have Go installed on your machine. You can download it from the official [Go website](https://golang.org/dl/).  

## Running a given day:
You can run solutions for specific days using the CLI. The structure follows the format:  
```go run main.go  <year> <day>```  
For example, to run the solution for Day 1 of Advent of Code 2023:  
```go run main.go 2023 day1```  

## Credits
All credits go to [Grep Haxs](https://www.youtube.com/watch?v=Od4Zi1HGB7s&t=34s&ab_channel=GrepHaxs)  
(I just followed his setup because I like it, this is the result)
