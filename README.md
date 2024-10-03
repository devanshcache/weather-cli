# weather-cli

A simple command-line application that fetches the current weather information for a specified location. This tool retrieves weather data from [wttr.in](https://wttr.in) and displays it in a user-friendly format using colored output.

## Features

- Displays current temperature, humidity, "feels like" temperature, UV index, and weather description for any location.
- Fetches real-time data from `wttr.in`.
- Simple, lightweight, and easy to use.

![image](https://github.com/user-attachments/assets/fa83a535-1616-422f-a16a-00455651f67f)


## Installation

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/yourusername/weather-cli.git
    cd weather-cli
    ```

2. **Build the Application**:
    Make sure you have [Go](https://golang.org/dl/) installed.

    ```bash
    go build -o weather-cli
    ```

3. **Install the Binary**:
    Move the binary to a directory in your `PATH`, for example `$HOME/.local/bin`:

    ```bash
    mv weather-cli ~/.local/bin/
    ```

    Ensure that `~/.local/bin` is in your `PATH`:
    ```bash
    export PATH="$HOME/.local/bin:$PATH"
    ```
## Usage

After installing, simply run the `weather-cli` followed by a location (city name) to get the current weather.

### Example:

```bash
weather-cli Budapest

