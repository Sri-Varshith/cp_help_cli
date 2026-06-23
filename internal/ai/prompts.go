package ai

import "fmt"

func BuildAnalyzePrompt(code string) string {
	return fmt.Sprintf(`
You are an expert competitive programming and software engineering reviewer.

Analyze the following code and provide:

1. Time Complexity
2. Space Complexity
3. Detailed Explanation
4. Performance Bottlenecks
5. Worst Case Analysis

Format your response as:

TIME COMPLEXITY:
...

SPACE COMPLEXITY:
...

EXPLANATION:
...

BOTTLENECKS:
...

WORST CASE:
...

CODE:

%s
`, code)
}

func BuildDebugPrompt(code string) string {
	return fmt.Sprintf(`
You are an expert debugging assistant.

Analyze the following code and identify:

1. Bugs
2. Runtime Errors
3. Logic Errors
4. Edge Cases
5. Undefined Behavior
6. Potential Crashes

Format your response as:

BUGS:
...

RUNTIME ISSUES:
...

LOGIC ISSUES:
...

EDGE CASES:
...

RECOMMENDED FIXES:
...

CODE:

%s
`, code)
}

func BuildImprovePrompt(code string) string {
	return fmt.Sprintf(`
You are a senior software engineer.

Review the following code and suggest improvements for:

1. Performance
2. Readability
3. Maintainability
4. Best Practices
5. Memory Usage
6. Algorithmic Efficiency

Format your response as:

PERFORMANCE:
...

READABILITY:
...

MAINTAINABILITY:
...

BEST PRACTICES:
...

ALGORITHM IMPROVEMENTS:
...

CODE:

%s
`, code)
}