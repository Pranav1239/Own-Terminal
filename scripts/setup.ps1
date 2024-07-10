# Setup script for Own-terminal

# Exit immediately if a command exits with a non-zero status
$ErrorActionPreference = "Stop"

# Function to check if a command exists
function Command-Exists {
    param (
        [string]$command
    )
    $commandPath = Get-Command $command -ErrorAction SilentlyContinue
    return $null -ne $commandPath
}

Write-Host "Starting setup for Own-terminal..."

# Check if Go is installed
if (-not (Command-Exists "go")) {
    Write-Host "Go is not installed. Please install Go and try again."
    exit 1
}

# Display Go version
Write-Host "Go version:"
go version

# Install dependencies
Write-Host "Installing dependencies..."
go mod tidy

# Create necessary directories
Write-Host "Creating necessary directories..."
if (-not (Test-Path "config")) {
    New-Item -ItemType Directory -Path "config" | Out-Null
}

# Create a default config file if it doesn't exist
$configPath = "config/config.json"
if (-not (Test-Path $configPath)) {
    Write-Host "Creating default config file..."
    @"
{
    "usernameDisplayOptions": {
        "color": "blue",
        "style": "bold"
    },
    "ssh": {
        "defaultHost": "",
        "defaultPort": 22
    },
    "todo": {
        "storagePath": "data/todos.json"
    }
}
"@ | Set-Content -Path $configPath
}

# Make sure data directory exists
if (-not (Test-Path "data")) {
    New-Item -ItemType Directory -Path "data" | Out-Null
}

Write-Host "Setup completed successfully."

Write-Host "You can now run the application using: ./own-terminal"
