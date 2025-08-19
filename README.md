# Auto App Opener

A windows CLI application to manage and open groups of apps automatically. Built with Go.

## Features

-   Discover installed applications
-   Create, edit, and delete groups of apps
-   Open multiple apps or groups with one click
-   Simple terminal UI using Charmbracelet's Huh
-   Persistent storage of app and group data

## Getting Started

### Prerequisites

-   Go 1.20+
-   Windows

### Installation

##### Executable

-   Install the [autoappopener.exe](https://github.com/Rithwik0604/auto-app-opener/blob/845e4e2a46007e29c5d4f2ee99f98ae6479c85e8/build/autoappopener.exe) and run it. To run it from anywhere in a terminal, add it your path.

#### Manual

1. Clone the repository:
    ```bash
    git clone <repo-url>
    cd autoappopener/go
    ```
2. Build and run:
    ```bash
    make run
    ```
    Or directly:
    ```bash
    go run ./cmd/main.go
    ```

## Project Structure

```
go.mod
go.sum
Makefile
cmd/
  main.go
internal/
  data/
    fetchData.go
    openApps.go
  models/
    config.go
    model.go
  storage/
    json_storage.go
  ui/
    createGroupForm.go
    editGroupForm.go
    firstForm.go
    initForm.go
    manageGroupsForm.go
    openAppForm.go
    openGroupForm.go
tmp/
```

## Usage

-   On launch, select to open apps, open a group, manage groups, refetch apps, or quit.
-   Manage groups to create, edit, or delete app groups.
-   Data is stored persistently in JSON format.

## License

MIT
