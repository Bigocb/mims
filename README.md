
## Overview
Mims will be a command-line assistant built in Go for developers. 

## Planned Features
- **CLI Mode:** Fast terminal commands for basic interactions.
- **TUI Mode:** Interactive, structured interface with visual feedback.
- **GenAI Assistance:** Chat-based help, project summaries, and code pairing. Choice of OpenAI or a Local Ollama server
- **Local Storage:** Save and search chat history for contextual insights.
- **Workflow:** Update Kubernetes cluster resources, manage PRs, other operations tasks.

## Planned Modes of Operation
- **CLI Mode:** Simple commands for quick tasks.
- **TUI Mode:** Enhanced, multi-step workflows with visual elements for more complex operations.

## Commands

- **research**: Use this command to ask a research question. It takes a `-topic` flag and proxies the question and response to OpenAI.

  Example:
  `mims research -topic "who were the Mayans"`

  Response:
  `The Maya civilization was a Mesoamerican civilization that flourished in present-day Mexico, Guatemala, Belize, Honduras, and El Salvador. The Maya civilization is known for its advancements in art, architecture, mathematics, astronomy, and writing systems. They developed a complex calendar system, built impressive cities with pyramids and temples, and created intricate artwork and sculptures. The Maya were also known for their achievements in agriculture, cultivating crops such as maize, beans, and squash.`

- **chat**: Opens an interactive terminal session with Mims, where you can engage in real-time conversations and get assistance.

  Example:
  `mims chat`

   This command will start an interactive session with Mims, where you can ask questions, get assistance, and explore project management features.

## Technical Details
- **Programming Language:** Go
- **CLI Framework:** UrfaveCLI
- **TUI Framework:** BubbleTea
- **Storage:** BoltDB/Storm (local), Kubernetes for optional scalability
- **GenAI Integration:** Ollama, LocalAI

## Development Plan
1. **Phase 1:** Implement CLI mode with basic functionality.
2. **Phase 2:** Add TUI mode with interactive features.

## Full Design Document
For the complete design documentation, check out the [design.md](design.md) file.
