#!/bin/sh

# Define the URL of the gov binary
GOV_URL="https://gov.bjerkepedia.com/bin"

# Define the installation directory (where gov will be installed)
INSTALL_DIR="/usr/local/bin"

# Determine the platform (Linux or macOS)
PLATFORM=""
if [ "$(uname)" = "Linux" ]; then
  PLATFORM="linux"
elif [ "$(uname)" = "Darwin" ]; then
  PLATFORM="darwin"
else
  echo "Unsupported operating system. Please use Linux or macOS."
  exit 1
fi

# Determine the architecture for Linux
ARCH=""
if [ "$PLATFORM" = "linux" ]; then
  if [ "$(uname -m)" = "x86_64" ]; then
    ARCH="amd64"
  elif [ "$(uname -m)" = "i386" ]; then
    ARCH="386"
  elif [ "$(uname -m)" = "aarch64" ]; then
    ARCH="arm64"
  elif [ "$(uname -m)" = "armv7l" ]; then
    ARCH="arm"
  else
    echo "Unsupported architecture. Please use x86_64, i386, armv7l, or aarch64."
    exit 1
  fi
fi

# Determine the architecture for macOS
if [ "$PLATFORM" = "darwin" ]; then
  ARCH="$(uname -m)"
fi

# Download the appropriate gov binary
if command -v curl >/dev/null 2>&1; then
  sudo curl -sSL "$GOV_URL/$PLATFORM/$ARCH/gov" -o "$INSTALL_DIR/gov"
elif command -v wget >/dev/null 2>&1; then
  sudo wget -q "$GOV_URL/$PLATFORM/$ARCH/gov" -O "$INSTALL_DIR/gov"
else
  echo "Neither curl nor wget found. Please install one of them and try again."
  exit 1
fi

# Make the gov binary executable
sudo chmod +x "$INSTALL_DIR/gov"

echo "gov has been installed to $INSTALL_DIR/gov."
