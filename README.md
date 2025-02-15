# Golaunch CLI

**Golaunch** is a command-line tool that allows you to launch multiple programs with a single command. Whether you're a developer, designer, or power user, Golaunch simplifies your workflow by letting you define custom commands to open your favorite apps instantly.

---

## Features

- **Custom Commands**: Define your own commands (e.g., `dev`, `design`) to launch multiple programs at once.
- **Cross-Platform**: Works on Windows, macOS, and Linux.
- **Easy Setup**: Interactive setup process to add programs and commands.
- **Lightweight**: Built with Go, it’s fast and efficient.

---

## Installation

### Prerequisites
- [Go](https://golang.org/dl/) (if building from source)
- Git (optional)

### Download the Executable
1. Download the latest release for your operating system from the [Releases page](https://github.com/raufzer/golaunch-cli/releases).
2. Move the executable to a directory in your system’s PATH (e.g., `/usr/local/bin` on macOS/Linux or `C:\Windows\System32` on Windows).

### Build from Source
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/golaunch-cli.git
   cd golaunch-cli
   ```
2. Build the executable:
   ```bash
   go build -o golaunch
   ```
3. Move the executable to a directory in your PATH:

   **Windows:**
   ```bash
   move golaunch.exe C:\Windows\System32\
   ```

   **macOS/Linux:**
   ```bash
   sudo mv golaunch /usr/local/bin/
   ```

---

## Usage

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

Abd Raouf Zerkhef zerkhefraouf90@gmail.com
