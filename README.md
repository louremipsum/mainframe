# Mainframe

## Project Structure

mainframe/
├── cmd/
│   └── mainframe/       # Main application entry point
├── internal/
│   ├── ui/              # UI components (Lipgloss styling)
│   ├── models/          # Business logic and state management
│   └── agents/          # AI agents and protocol implementation
├── assets/              # Static assets (mascot art, config files)
├── pkg/
│   ├── styles/          # Shared styling (Lipgloss)
│   ├── terminal/        # Terminal interaction logic
│   └── utils/           # Utility functions (logging, helper methods)
├── config/              # Configurations for models, API keys, etc.
├── Makefile             # Build and run commands
└── go.mod               # Go module file
