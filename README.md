# Go Version Manager (gov)

[Go Version Manager (gov)](https://gitlab.com/bjerke-tek/go-version-manager) is a simple command-line tool for managing Go versions on your system.

## Table of contents

-   [Installation](#installation)
-   [Usage](#usage)
    -   [List available supported Go versions](#list-available-supported-go-versions)
    -   [List all available Go versions](#list-all-available-go-versions)
    -   [Install a Go version](#install-a-go-version)
    -   [Reinstall a Go version](#reinstall-a-go-version)
    -   [Use a Go version](#use-a-go-version)
    -   [Display the currently used Go version](#display-the-currently-used-go-version)
    -   [List installed Go versions](#list-installed-go-versions)
    -   [Remove a Go version](#remove-a-go-version)
    -   [Prune](#prune)
    -   [Show version](#show-version)
    -   [Self update](#self-update)
    -   [Uninstall gov](#uninstall-gov)
    -   [Show help](#show-help)
-   [Compiling from Source](#compiling-from-source)
-   [License](#license)

## Installation

### Linux and macOS

Using curl (recommended):

```bash
curl -fsSL https://gov.bjerkepedia.com/gov.sh | bash
```

Using wget:

```bash
wget -O- https://gov.bjerkepedia.com/gov.sh | bash
```

## Usage

### List available supported Go versions

**Long:**

```bash
gov list-supported
```

**Short:**

```bash
gov s
```

### List all available Go versions

**Long:**

```bash
gov list-all
```

**Short:**

```bash
gov a
```

### Install a Go version

Install a Go _version_, or the latest release using the "latest" tag, e.g:

**Long:**

```bash
gov install 1.21.1
```

```bash
gov install latest
```

**Short:**

```bash
gov i 1.21.1
```

```bash
gov i latest
```

### Reinstall a Go version

Remove and then install again a Go _version_, e.g:

**Long:**

```bash
gov reinstall 1.21.1
```

**Short:**

```bash
gov r 1.21.1
```

### Use a Go version

The first time gov will set the env var for you and you might need to restart your shell/terminal for the configuration to take effect. Subsequent changes should work seamlessly.

**Long:**

```bash
gov use 1.21.1
```

**Short:**

```bash
gov u 1.21.1
```

### Display the currently used Go version

**Long:**

```bash
gov current
```

**Short:**

```bash
gov c
```

### List installed Go versions

**Long:**

```bash
gov list
```

**Short:**

```bash
gov l
```

### Remove a Go version

**Long:**

```bash
gov remove 1.21.1
```

**Short:**

```bash
gov x 1.21.1
```

### Prune

Remove all installed Go versions except the currently used one, e.g:

**Long:**

```bash
gov prune
```

**Short:**

```bash
gov p
```

### Show version

Show the installed gov version and extras.

**Long:**

```bash
gov version
```

**Short:**

```bash
gov v
```

### Self update

Update gov to its latest version

**Long:**

```bash
gov self-update
```

**Short:**

```bash
gov e
```

### Uninstall gov

This will only remove gov, not your installed Go versions.

After uninstallation, gov will provide instructions to remove Go and leftovers.

```bash
gov sayonara
```

### Show help

```bash
gov --help
```

```bash
gov -h
```

```bash
gov help
```

```bash
gov h
```

```bash
# or make a mistake ¯\_(ツ)_/¯
```

## Compiling from Source

If you prefer to compile gov from source or want to contribute, you can follow these steps:

1. Clone the gov repository:

    ```bash
    git clone https://gitlab.com/bjerke-tek/go-version-manager
    ```

2. Change to the gov directory:

    ```bash
    cd gov
    ```

3. Build gov using the build scripts:

    1. Make script executable (if needed)

        ```bash
        sudo chmod +x ./build.sh
        ```

    2. Build only the supported platforms (recommended)

        ```bash
        ./build.sh <new_gov_version>
        ```

    3. Build for every platform and architecture. I mean, why not?

        ```bash
        ./build-all.sh <new_gov_version>
        ```

4. Optionally, move the gov binary to a directory included in your system's PATH for easy access.

### Building notes

**If you build for Windows note that env vars and zip extraction won't work.
You're welcome to work on Windows support. I just can't be arsed with
Windows, it puts me in a bad mood every time I use it.**

## License

This project is licensed under the GPLv3 License - see the [LICENSE](https://gitlab.com/bjerke-tek/go-version-manager/-/blob/master/LICENSE) file for details.

## Copyright notice

Copyright (c) 2023 [Erik Bjerke](https://erikbjerke.com). All rights reserved.

This software is not affiliated with or endorsed by the Go Project.
