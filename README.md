# cp_help_cli

A simple command-line tool built in Go to help review source code using AI.

The tool reads a code file, generates a prompt based on the selected command, and sends it to an AI model for analysis.

## Features

Currently supported commands:

### Analyze

Analyze a source file and get:

* Time complexity
* Space complexity
* General explanation
* Potential bottlenecks

```bash
cphelp analyze solution.cpp
```

### Debug

Review a source file and identify:

* Possible bugs
* Runtime issues
* Logic errors
* Edge cases

```bash
cphelp debug solution.cpp
```

### Improve

Get suggestions for improving code quality, including:

* Readability
* Performance
* Maintainability
* Best practices

```bash
cphelp improve solution.cpp
```

---

## Project Structure

```text
cp_help_cli/
├── cmd/
│   ├── analyze.go
│   ├── debug.go
│   ├── improve.go
│   └── root.go
│
├── internal/
│   ├── ai/
│   │   ├── client.go
│   │   ├── config.go
│   │   └── prompts.go
│   │
│   ├── output/
│   │   └── printer.go
│   │
│   └── parser/
│       └── file.go
│
├── .env
├── .gitignore
├── main.go
├── go.mod
└── go.sum
```


---

## Installation

Clone the repository:

```bash
git clone https://github.com/Sri-Varshith/cp_help_cli.git
cd cp_help_cli
```

Install dependencies:

```bash
go mod tidy
```

Create a `.env` file:

```env
OPENAI_API_KEY=your_api_key
OPENAI_MODEL=gpt-5
```

Run:

```bash
go run . analyze sample.cpp
```

---

## Example Usage

```bash
cphelp analyze solution.cpp

cphelp debug solution.cpp

cphelp improve solution.cpp
```

---

## Technologies Used

* Go
* Cobra CLI
* OpenAI API
* godotenv

---

## Future Improvements

Some ideas planned for future versions:

* Add an `explain` command for code walkthroughs
* JSON output support
* Support for analyzing entire directories
* Better formatting of AI responses
* Generate test cases from source code
* Generate explanations for competitive programming solutions
* Direct Codeforces integration
* Submit solutions directly to Codeforces
* Fetch and display Codeforces problem statements
* Analyze accepted Codeforces submissions
* Store previous analyses locally
* Multi-language support (C++, Python, Java, Go)

---

## Notes

This is a learning project built to explore:

* Go CLI development
* Working with external APIs
* Prompt generation
* Code analysis workflows

The project is still evolving and new features may be added over time.
