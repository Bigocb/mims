# Technical Design Document: GenAI Command Line Assistant ("Mims")

## 1. Introduction

### Purpose of the Document
This document outlines the design of "Mims," a command-line assistant tailored for developers. The assistant integrates with GenAI models and productivity tools to enhance development workflows. It supports tasks like chat-based assistance, project management, Kubernetes management, and code review automation.

### Overview of the System
Mims is a Go-based command-line tool that operates in multiple modes:
1. **CLI Mode:** A straightforward command-line interface for quick interactions and task execution.
2. **TUI Mode:** An interactive text-based interface powered by BubbleTea, offering a more visual and structured experience for users.

### Goals
- Provide CLI and TUI interfaces for diverse developer needs.
- Act as a project manager by integrating with Jira for story creation and basic task updates.
- Leverage GenAI for summarization, research assistance, and code pairing.
- Enable local storage and history search for contextual enhancements.
- Support complex development tasks like opening PRs and Kubernetes updates.

### Non-Goals
- Replace full-fledged project management tools like Jira.
- Offer a web or desktop application interface.

---

## 2. System Architecture

### High-Level Architecture Diagram
_(Placeholder: A diagram showing CLI/TUI components, integration with Jira, GenAI, and optional Kubernetes services.)_

### Components and Responsibilities
1. **Command-Line Interface (CLI):**
    - User-facing interface built with UrfaveCLI.
    - Handles input parsing, command routing, and output display.

2. **Text User Interface (TUI):**
    - Interactive, structured interface built with BubbleTea.
    - Provides enhanced navigation, visual feedback, and multi-step workflows.

3. **Processing Engine:**
    - Core logic shared by CLI and TUI modes.
    - Handles user input, integrates with GenAI, and executes tasks.

4. **Storage:**
    - **BoltDB/Storm:** Local, lightweight storage for chat history and context caching.
    - **Optional Long-Term Storage:** API-backed database for scalable storage in Kubernetes.

5. **Jira Integration:**
    - API client for interacting with Jira to perform basic project management tasks such as creating and updating stories.
    - Optional Kubernetes service for advanced Jira interactions or connecting to Jira Cloud.

6. **Kubernetes Services (Optional Mode):**
    - API Gateway: Central entry point for storage and search APIs.
    - Search Service: Handles indexed queries for efficient history lookups.
    - Storage Service: Manages persistent storage of conversations and data.

---

## 3. Functional Requirements

### Modes of Operation
- **CLI Mode:** Lightweight and quick interactions via terminal commands.
- **TUI Mode:** Interactive workflows for tasks like project management, PR creation, and Kubernetes updates.

### Core Features
- Process user input and generate responses via GenAI.
- Save responses or full chats locally for future reference.
- Search saved history and summarize results for context.

### Developer-Specific Tasks
- Open PRs with auto-generated summaries based on branch changes.
- Perform Kubernetes updates via declarative commands.
- Assist with research by leveraging past conversations.
- Provide unobtrusive code pairing suggestions.

### Project Management Tasks
- Authenticate and integrate with Jira for basic operations.
- Create Jira stories and update their status.
- Fetch basic project details.
- Summarize project progress using GenAI.

### TUI-Specific Features
- Real-time chat interaction with GenAI.
- Visual history navigation and search.
- Multi-step task wizards for complex operations like story creation and Kubernetes updates.

---

## 4. Non-Functional Requirements

- **Performance:** Fast response times for CLI, TUI, and Jira interactions (< 1s for most operations).
- **Scalability:** Optional Kubernetes mode for large-scale storage and search.
- **Usability:** Intuitive CLI commands, TUI navigation, and Jira integration.
- **Security:** Securely handle Jira API keys and other sensitive data.

---

## 5. Detailed Design

### Key Algorithms and Workflows

#### 1. **Mode Switching**
- Detect mode (`mims cli` or `mims tui`) based on user input or configuration.
- Route interactions to the appropriate interface layer (CLI or TUI).

#### 2. **Jira Integration Workflow**
- **Basic Cloud Connection:**
    - Mims CLI/TUI directly connects to Jira Cloud via REST or GraphQL APIs.
    - Securely store API keys or OAuth tokens in a local encrypted file or environment variables.
- **Optional Kubernetes Service:**
    - Deploy a microservice within the Kubernetes cluster for advanced Jira integration, such as caching and search indexing.

#### 3. **CLI Jira Features**
- `mims jira create story --project PROJECT_KEY --summary "Add TUI support"`
    - Creates a new story in the specified project.
- `mims jira list tasks --project PROJECT_KEY`
    - Lists tasks filtered by project.
- `mims jira update TASK_KEY --status DONE`
    - Updates the status of a specific Jira task.

#### 4. **TUI Jira Features**
- Menu-driven interface for basic task management.
- Use GenAI to summarize Jira data for progress updates.

#### 5. **Optional Kubernetes Mode**
- When deployed in Kubernetes, Mims offloads advanced Jira tasks to a service:
    - **Search Service:** Index Jira data for faster lookups.
    - **Storage Service:** Maintain a cache of Jira project data to minimize API calls.
- Provide a Helm chart or Kubernetes manifest to deploy these services.

### Architecture

- Main logic to handle "business" logic.
  - Process user input
  - Determine workflow for request
  - Execute workflow
  - return result to CLI/TUI
- Interface with CLI
  - CLI and TUI should be responsible for taking in user input, proxying that data to the "brain", and
    displaying the results.
- Interface with TUI

---

## 6. Technology Choices

### Core Technologies
- **Programming Language:** Go
- **CLI Framework:** UrfaveCLI
- **TUI Framework:** BubbleTea
- **Storage:** BoltDB, Storm
- **GenAI Integration:** Ollama, LocalAI

### Additional Tools for Jira Integration
- **Jira API:** REST/GraphQL for accessing Jira Cloud or self-hosted instances.
- **Kubernetes:** Host optional Jira integration service.
- **API Gateway:** Simplifies routing for Jira-related requests in Kubernetes.

---

## 7. Implementation Plan

### Timeline
1. Phase 1: CLI Mode with Basic Features (1-2 weeks).
2. Phase 2: TUI Mode with Interactive Features (2-3 weeks).
3. Phase 3: GenAI Integration and Core Features (2-3 weeks).
4. Phase 4: Optional Kubernetes Mode (3-4 weeks).

### Milestones
1. MVP with CLI and local chat/history functionality.
2. TUI implementation with enhanced workflows and navigation.
3. Feature completion (PR automation, Kubernetes commands, Jira integration).

---

## 8. Testing Strategy

- **Unit Testing:** Core business logic and mode-specific features.
- **Integration Testing:** Interaction between CLI, TUI, and GenAI.
- **End-to-End Testing:** Real-world use cases in both modes.
- **Performance Testing:** Ensure responsiveness of CLI, TUI, and Jira-related operations.

---

## 9. Deployment Plan

- **CLI Mode:** No deployment; users download and run the binary.
- **TUI Mode:** Bundled with CLI in the same binary, accessible via a specific command or flag.
- **Kubernetes Mode:** Deploy services via Helm charts or custom manifests.

---

## 10. Future Enhancements

- Add multi-user support for shared environments.
- Introduce advanced TUI features like visual graphs or dashboards.
- Expand API integrations (e.g., GitHub, Jira, Slack).

---

## 11. Appendices

### Glossary
- **GenAI:** Generative AI.
- **BubbleTea:** TUI framework for Go.
- **BoltDB/Storm:** Lightweight embedded key-value databases.
- **Ollama/LocalAI:** Local GenAI service providers.
