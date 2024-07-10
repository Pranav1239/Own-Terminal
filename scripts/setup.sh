#!/bin/bash

# Setup script for Own-terminal

# Exit immediately if a command exits with a non-zero status
set -e

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

echo "Starting setup for Own-terminal..."

# Check if Go is installed
if ! command_exists go; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Display Go version
echo "Go version:"
go version

# Install dependencies
echo "Installing dependencies..."
go mod tidy

# Create necessary directories
echo "Creating necessary directories..."
mkdir -p config

# Create a default config file if it doesn't exist
if [ ! -f config/config.json ]; then
    echo "Creating default config file..."
    cat <<EOL > config/config.json
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
EOL
fi

# Make sure data directory exists
mkdir -p data

echo "Setup completed successfully."

echo "You can now run the application using: ./own-terminal"
