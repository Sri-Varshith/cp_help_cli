package ai

import "fmt"

func BuildAnalyzePrompt(code string) string {
	return fmt.Sprintf(`
Analyze the following code and provide:

1. Time Complexity
2. Space Complexity
3. Bottlenecks
4. Explanation

Code:

%s
`, code)
}

func BuildDebugPrompt(code string) string {
	return fmt.Sprintf(`
Find bugs, runtime errors, edge cases and logic issues.

Code:

%s
`, code)
}

func BuildImprovePrompt(code string) string {
	return fmt.Sprintf(`
Suggest improvements for:

1. Performance
2. Readability
3. Maintainability
4. Best Practices

Code:

%s
`, code)
}