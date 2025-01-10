## Overview
Mims is a command-line assistant built in Go for developers. It offers two modes of operation: **CLI** for quick tasks and **TUI** for an interactive, text-based interface. It integrates with GenAI and productivity tools like Jira, assisting developers with tasks such as project management, code review, and Kubernetes operations.

## Features
- **CLI Mode:** Fast terminal commands for basic interactions.
- **TUI Mode:** Interactive, structured interface with visual feedback.
- **Jira Integration:** Create and update stories, list tasks, and manage project status.
- **GenAI Assistance:** Chat-based help, project summaries, and code pairing.
- **Local Storage:** Save and search chat history for contextual insights.
- **Kubernetes Management:** Update Kubernetes resources and manage PRs.

## Modes of Operation
- **CLI Mode:** Simple commands for quick tasks.
- **TUI Mode:** Enhanced, multi-step workflows with visual elements for more complex operations.

## Core Features
- Generate responses and summarize content with GenAI.
- Create and manage Jira stories and tasks.
- Store conversations locally for future reference.
- Use GenAI for research assistance and code review automation.

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
3. **Phase 3:** Integrate GenAI and extend functionality (PR automation, Jira).
4. **Phase 4:** Kubernetes mode for advanced integrations.

## Installation
Mims is a Go-based tool, and you can install it by downloading the binary or building it from source.

## Future Enhancements
- Multi-user support for shared environments.
- Advanced TUI features like dashboards.
- Integrate with additional APIs (GitHub, Slack, etc.).
- Integrate with voice assistants for voice mode

## Contributing
Feel free to contribute to the project via pull requests or issues.

## Full Design Document
For the complete design documentation, check out the [design.md](design.md) file.
