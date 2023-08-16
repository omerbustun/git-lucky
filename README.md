git-lucky 🍀
============

Discover a random GitHub repository from popular programming languages with just one command. Dive into diverse projects, explore new territories, and get inspired.

Features
--------

*   Randomly selects from popular programming languages or allows you to specify a language of your choice.
*   Provides a simple command-line interface.
*   Option to select a specific programming language using the `-lang` flag.
*   View the help text and options using the `-h` flag.
*   Fetches repositories sorted by the latest commit.
*   Configurable via a JSON configuration file.

Getting Started
---------------

### Download Precompiled Binaries

For those who prefer not to compile the software, precompiled binaries are available:

- [Windows](https://github.com/omerbustun/git-lucky/releases/download/v1.0.0/git-lucky-windows-amd64.exe)
- [macOS](https://github.com/omerbustun/git-lucky/releases/download/v1.0.0/git-lucky-darwin-amd64)
- [Linux](https://github.com/omerbustun/git-lucky/releases/download/v1.0.0/git-lucky-linux-amd64)

### Build from Source

#### Prerequisites

*   [Go](https://golang.org/dl/)

#### Steps

1.  Clone this repository:
    
    ```bash
    git clone https://github.com/omerbustun/git-lucky.git
    ```
    
2.  Navigate to the cloned directory:
    
    ```bash
    cd git-lucky
    ```
    
3.  Build the project:
    
    ```bash
    go build
    ```
Configuration
-------------

1.  Copy the sample configuration file:
    
    ```bash
    cp config.sample.json config.json
    ```
    
2.  Add your GitHub API token (optional but recommended to increase rate limits):
    
    ```json
    {
      "api_token": "YOUR_GITHUB_API_TOKEN"
    }
    ```
Usage
-----

Simply run:

```bash
./git-lucky
```

To specify a programming language:


```bash
./git-lucky -lang=Python
```

For help:


```bash
./git-lucky -h
```

Contributing
------------

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

License
-------

This project is licensed under the GNU General Public License, version 3 (GPLv3). See the [LICENSE](LICENSE) file for the full license text.