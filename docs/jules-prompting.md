# Jules Prompting Guide for AI

## 1. Overview

This document provides a comprehensive guide for AI models to generate effective prompts for Jules, a sophisticated software engineering AI. The goal is to enable other AIs to understand Jules's capabilities, methodologies, and best practices to produce optimal instructions. For more information, refer to the [official Jules documentation](https://jules.google/docs/).

## 2. Core Principles for Prompting Jules

*   **Clarity and Specificity:** Prompts should be unambiguous and provide concrete details.
*   **Completeness:** Include all necessary context, constraints, and expected outcomes.
*   **Structured Format:** Use markdown and clear formatting to structure complex requests.

## 3. Multi-Role Analysis for Prompt Generation

This section details the different roles to consider when constructing prompts for Jules.

### 3.1. AI Prompt Engineer Viewpoint
-   **Key Insights and Concerns:**
    -   **Prompt Ambiguity:** Vague or incomplete prompts are the primary source of error. Jules will attempt to fill in the gaps, but this can lead to solutions that don't align with the user's intent.
    -   **Token Economy:** While Jules is capable, verbose prompts with excessive conversational filler can obscure the core task. Efficiency is key.
    -   **Implicit vs. Explicit Instructions:** Jules relies on explicit instructions. Assuming Jules "knows" about unstated context or preferences is a common pitfall.
-   **Specific Recommendations:**
    -   **Use Structured Data Formats:** For complex tasks, use markdown, JSON, or YAML to structure the prompt. This provides a clear, machine-readable format for Jules to parse.
    -   **Provide "Negative Constraints":** Clearly state what Jules *should not* do. This is often more effective than only describing the desired outcome.
    -   **Reference Specific Files and Line Numbers:** When modifying existing code, provide exact file paths and line numbers. This reduces ambiguity and search time.
    -   **Chain of Thought Priming:** For complex problem-solving, instruct Jules to "think step-by-step" or "create a plan."

### 3.2. Senior Software Engineer Viewpoint
-   **Key Insights and Concerns:**
    -   **Code Quality and Maintainability:** The generated code must adhere to existing coding standards, be well-documented, and easy for human developers to maintain.
    -   **Architectural Cohesion:** Changes should respect the existing architecture and design patterns of the project. Avoid introducing new patterns without explicit instruction.
    -   **Testing:** All new features or bug fixes must be accompanied by appropriate tests.
-   **Specific Recommendations:**
    -   **Specify Coding Standards:** If the project has a style guide, reference it in the prompt.
    -   **Provide Architectural Context:** When requesting a change, briefly explain its place in the overall architecture.
    -   **Explicitly Mention Testing:** The prompt should always include a requirement for testing.
    -   **Reference `AGENTS.md`:** If the repository contains an `AGENTS.md` file, instruct Jules to consult it for project-specific conventions and instructions.

### 3.3. AI System Architect Viewpoint
-   **Key Insights and Concerns:**
    -   **Tool Integration:** Jules operates through a specific set of tools. Prompts must be framed in a way that is achievable with these tools.
    -   **State Management:** Jules does not have a persistent memory of the codebase between tasks. Each new prompt should be self-contained.
    -   **Sequential Operation:** Jules performs one action at a time. Complex tasks need to be broken down into a logical sequence of steps in a plan.
-   **Specific Recommendations:**
    -   **Task Decomposition:** For large tasks, the prompt should guide Jules to create a detailed, step-by-step plan.
    -   **Explicit Verification Steps:** Instruct Jules to verify the outcome of each step.
    -   **Tool-Oriented Instructions:** Phrase requests in terms of the available tools.
    -   **Memory Directives:** Explicitly mention the available memory features when relevant.

### 3.4. Technical Writer (for AI consumption) Viewpoint
-   **Key Insights and Concerns:**
    -   **Parsing and Tokenization:** The prompt's structure directly impacts how it is parsed. Clear, declarative statements are best.
    -   **Signal-to-Noise Ratio:** Every token should contribute to the task's definition.
    -   **Consistency of Terminology:** Use consistent terms for the same concept.
-   **Specific Recommendations:**
    -   **Use Markdown Extensively:** Use headings, lists, code blocks, and bolding to create a clear structure.
    -   **Adopt a "Given-When-Then" Structure:** For behavior-driven tasks, frame the request in this format.
    -   **Isolate Code and Commands:** Always place file paths, code snippets, and shell commands in code blocks.

## 4. Cross-Role Integration & Synthesis

-   **Areas of Agreement:** All roles converge on the necessity of prompts being clear, specific, complete, and structured. Decomposing complex tasks into smaller, verifiable steps is a universally endorsed best practice.
-   **Conflicts and Tensions:**
    -   **Brevity vs. Completeness:** The AI Prompt Engineer's desire for token economy can conflict with the Senior Software Engineer's need for detailed technical context.
    -   **Mechanical vs. Natural Language:** The AI System Architect's focus on tool-oriented commands can be overly rigid, while the Technical Writer aims for a balance of precision and clarity.
-   **Synthesis and Unified Recommendations:** A successful prompt harmonizes these perspectives. It should be:
    1.  **Goal-Oriented:** Start with a high-level objective.
    2.  **Context-Rich, but Structured:** Provide architectural and technical context using markdown lists, code snippets, and references to existing files.
    3.  **Plan-Driven:** Explicitly instruct Jules to create a step-by-step plan with verification at each stage.
    4.  **Tool-Aware:** Frame instructions in a way that naturally maps to Jules's available tools without being overly robotic.
    5.  **Test-Inclusive:** Testing requirements should be a non-negotiable part of any code-altering prompt.
    6.  **Explicitly Constrained:** Use negative constraints to define boundaries.

## 5. Capabilities and Tooling

### 5.1. Core Competencies
-   Code generation, modification, and debugging.
-   Test creation and execution.
-   Filesystem manipulation.
-   Web research and information synthesis.

### 5.2. Tool Reference with Examples
-   **`list_files`**: Lists files and directories. `{"tool_code": "list_files(path='src/')"}`
-   **`read_file`**: Reads the content of a file. `{"tool_code": "read_file(filepath='src/main.go')"}`
-   **`create_file_with_block`**: Creates a new file with content. `{"tool_code": "create_file_with_block(filepath='new_file.txt', content='Hello, world!')"}`
-   **`replace_with_git_merge_diff`**: Performs a targeted search-and-replace using a git merge diff format.
-   **`run_in_bash_session`**: Executes a shell command. `{"tool_code": "run_in_bash_session(command='go test ./...')"}`
-   **`google_search`**: Performs a web search. `{"tool_code": "google_search(query='golang best practices for error handling')"}`
-   **`set_plan`**: Sets the multi-step plan for the task.
-   **`submit`**: Submits the completed work.

## 6. Development Methodology

### 6.1. Planning Phase
-   Jules begins by analyzing the prompt and exploring the codebase using `list_files` and `read_file` to understand the context.
-   A detailed, multi-step plan is created using `set_plan`.

### 6.2. Execution Phase
-   Jules executes each step of the plan sequentially.
-   After each action, Jules uses a read-only tool to verify the change was successful (e.g., `read_file` after a `replace_with_git_merge_diff`).

### 6.3. Pre-commit and Submission Phase
-   Before submitting, Jules runs pre-commit checks, which typically involve running tests and linters.
-   The change is submitted with a descriptive commit message.

## 7. Best Practices and Examples

### 7.1. Prompting for Feature Implementation

**Objective:** Create a prompt that instructs Jules to add a new REST endpoint.

```markdown
Your task is to implement a new `GET /api/v1/health` endpoint in the Go application.

**1. High-Level Goal:**
Create a health check endpoint that returns a JSON response indicating the application's status.

**2. Technical Specifications:**
- **File:** `internal/handlers/health.go` (create this file)
- **Function Signature:** `func HealthHandler(w http.ResponseWriter, r *http.Request)`
- **Response Body (Success):** `{"status": "ok"}`
- **HTTP Status Code (Success):** `200 OK`
- **Routing:** The new handler must be registered in `cmd/server/main.go` to handle the `/api/v1/health` route.

**3. Testing Requirements:**
- Create a new test file `internal/handlers/health_test.go`.
- Add a unit test that calls the `HealthHandler` and asserts that the status code is `200` and the response body is `{"status": "ok"}`.
- Use the `PEA_GEMINI_MOCK=true` environment variable when running tests.

**4. Plan of Action:**
Please create a plan that includes the following steps:
1. Create the `internal/handlers/health.go` file with the `HealthHandler` function.
2. Create the `internal/handlers/health_test.go` file with the unit test.
3. Modify `cmd/server/main.go` to register the new route.
4. Run the tests to verify the implementation.
5. Complete the pre-commit steps.
6. Submit the changes.
```

### 7.2. Prompting for Bug Fixes

**Objective:** Create a prompt for fixing a bug.

```markdown
There is a bug in the `calculate_total` function where it panics if the input slice is empty.

**1. Bug Description:**
- **File:** `internal/util/calculator.go`
- **Function:** `calculate_total(items []int) int`
- **Problem:** When `items` is empty, the function accesses `items[0]`, causing a panic.

**2. Expected Behavior:**
If the input slice is empty, the function should return `0`.

**3. Task:**
1.  Modify the `calculate_total` function in `internal/util/calculator.go` to handle the empty slice case correctly.
2.  Add a new unit test in `internal/util/calculator_test.go` that specifically tests the empty slice case.
3.  Ensure all existing tests continue to pass.
```
