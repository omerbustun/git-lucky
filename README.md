git-lucky 🍀
============

Discover a random GitHub repository from popular programming languages with just one command. Dive into diverse projects, explore new territories, and get inspired.

Features
--------

*   Randomly selects from popular programming languages.
*   Fetches repositories sorted by the latest commit.
*   Provides a simple command-line interface.
*   Configurable via a JSON configuration file.

Prerequisites
-------------

*   [Go](https://golang.org/dl/) (to build the project from source)

Installation
------------

1.  Clone this repository:
    

    
    ```bash
    git clone https://github.com/omerbustun/git-lucky.git
    ```
    
2.  Navigate to the cloned directory and build the project:
    

    
    ```bash
    go build
    ```
    
3.  Run `git-lucky`:
    

    ```bash
    ./git-lucky
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

Contributing
------------

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

License
-------

This project is licensed under the GNU General Public License, version 3 (GPLv3). See the [LICENSE](LICENSE) file for the full license text.
