# GoLaunch CLI

<p align="center">
  <img src="assets/GoLaunchLogo.png" alt="Golaunch Logo">
</p>

**GoLaunch** is a command-line tool that allows you to launch multiple programs with a single command. Whether you're a developer, designer, or power user, GoLaunch simplifies your workflow by letting you define custom commands to open your favorite apps instantly.

---

## Features

- **Custom Commands**: Define your own commands (e.g., `dev`, `design`) to launch multiple programs at once.
- **Command Grouping**: Organizes commands into categories like Setup, Core, State, History, and Miscellaneous.
- **Cross-Platform**: Works on Windows, macOS, and Linux.
- **Easy Setup**: Interactive setup process to add programs and commands.
- **Lightweight**: Built with Go, it’s fast and efficient.
- **Homebrew Support**: Install and update via Homebrew for macOS & Linux.
- **Docker Support**: Run GoLaunch as a containerized app.

---

## Installation

### macOS & Linux (Using Homebrew) – Recommended
```bash
brew tap raufzer/homebrew-golaunch-cli
brew install golaunch-cli
```

### Windows & Other Platforms (Manual Installation)
Download pre-built binaries from the [Releases page](https://github.com/raufzer/golaunch-cli/releases).

### From Source
```bash
git clone https://github.com/raufzer/golaunch-cli.git
cd golaunch-cli
go build -o golaunch ./cmd/golaunch
sudo mv golaunch /usr/local/bin/
```

### Using Docker
```bash
docker pull raufzer/golaunch-cli-docker:latest
docker run --rm raufzer/golaunch-cli-docker golaunch
```

---

## Quick Start

### Initialize the CLI
Run the following command to create the necessary configuration files:
```bash
golaunch start
```

### Set Up Custom Commands
Use the setup command to define custom commands and associate them with programs:
```bash
golaunch setup
```
Follow the prompts to:
- Enter a custom command (e.g., `dev`, `design`).
- Add the paths to the programs you want to launch (e.g., `C:\Program Files\Google\Chrome\Application\chrome.exe`).

### Launch Programs
Use the `open` command to launch programs associated with a custom command:
```bash
golaunch open dev
```

### View Available Commands & Groups
```bash
golaunch help
```
Example output:
```
Usage:
  golaunch [command]

Setup & Initialization:
  setup      Set up custom commands and program paths

Core Functionality:
  open       Launch a predefined command group
  list       Show available custom commands

Program State:
  status     Show the current state of GoLaunch

History & Logging:
  history    Show command execution history

Miscellaneous:
  version    Show the current version of GoLaunch

Use "golaunch [command] --help" for more information about a command.
```

---

## Example Workflow

### Initialize the CLI
```bash
golaunch start
```

### Set Up a Custom Command
```bash
golaunch setup
```
- Enter a custom command (e.g., `dev`, `design`): `dev`
- Enter the path to a program (or press Enter to finish):  
  `C:\Program Files\Google\Chrome\Application\chrome.exe`  
  `C:\Users\<YourUsername>\AppData\Local\Programs\Microsoft VS Code\Code.exe`  
- Setup complete! Use `golaunch open dev` to launch your programs.

### Launch Programs
```bash
golaunch open dev
```
Output:
```
Launched C:\Program Files\Google\Chrome\Application\chrome.exe successfully!
Launched C:\Users\<YourUsername>\AppData\Local\Programs\Microsoft VS Code\Code.exe successfully!
```

---

## Configuration

The configuration is stored in `assets/config.json`. Here’s an example:

```json
{
  "dev": [
    "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
    "C:\\Users\\<YourUsername>\\AppData\\Local\\Programs\\Microsoft VS Code\\Code.exe"
  ],
  "design": [
    "D:\\Adobe\\Adobe Lightroom CC\\lightroom.exe"
  ]
}
```

---

## Uninstall

### Using Homebrew
```bash
brew uninstall golaunch-cli
brew untap raufzer/homebrew-golaunch-cli
```

### Remove Docker Image
```bash
docker rmi raufzer/golaunch-cli-docker
```

---

## Contributing

Contributions are welcome! Here’s how you can help:

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your feature"
   ```
4. Push to the branch:
   ```bash
   git push origin feature/your-feature-name
   ```
5. Open a pull request.

---

## License

Abd Raouf Zerkhef - [zerkhefraouf90@gmail.com](mailto:zerkhefraouf90@gmail.com)

