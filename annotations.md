# 🧠 Annotations

- **Author:** _Daniel Oweibo_ _(@danoweibo)_
- **Date:** _Saturday, 1 Nov 2025_

## Introduction

This **`annotations.md`** file exists as a developer’s personal guide and quick reference.
As someone who writes in multiple programming languages across diverse ecosystems, I use notes like this to **reignite my memory** about the **technical decisions** made while working on a project.

Unlike formal documentation, these annotations focus on the **why** behind certain choices — not just the **how**.

They also serve as a helpful guide for anyone exploring the codebase, whether you’re a **beginner** trying to understand the project’s structure or a **seasoned developer** curious about the reasoning behind specific implementation details.

## ⚙️ Makefile Overview

This Makefile bootstraps common tasks for the project **bankgoals**.  
It serves a similar purpose to a `package.json` file — but for **Golang** or **C++** codebases.

### General Rule

Always prefix each command with an **`@`** symbol.  
**Reason:** It prevents `make` from printing the command before executing it, keeping the terminal output clean and readable.

---

### 🏗️ `build` Command

**Description:**  
Compiles the Go codebase into a system object named **`bankgoals`**, placing the output inside the `bin` folder at the project root.

**Command:**

```makefile
@go build -o bin/bankgoals
```

---

### 🚀 `run` Command

**Description:**  
Runs the `build` command first to ensure the application is rebuilt if there have been changes since the last build.  
After building, it executes the newly built binary.

**Note:**  
You don’t need to prefix `build` with `@` here — it’s already a Makefile target, not a shell command.

**Command:**

```makefile
build
@./bin/bankgoals
```

---

### 🧪 `test` Command

**Description:**  
Executes all tests in the project with verbose output.  
The `-v` flag enables detailed logging for every test across the entire codebase (`./...`).

**Command:**

```makefile
@go test -v ./...
```
