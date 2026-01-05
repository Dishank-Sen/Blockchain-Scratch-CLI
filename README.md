# bloc — Command Line Interface

`bloc` is the command-line interface for interacting with the **Bloc daemon (`blocd`)** and the Bloc network.  
It is used to initialize repositories, manage peer connections, and communicate with the bootstrap server.

---

## Installation

### Build from source

```bash
make build
```

Usage
```bash
bloc [command]
```

init

Initialize a new Bloc repository in the current directory.
```bash
bloc init
```
This creates:

Repository metadata

Local identity information (not committed to Git)

Run this once per project.

connect

Connects to the bootstrap server and starts the Bloc daemon if it is not already running.
```bash
bloc connect
```

What this does:

Ensures blocd is running

Registers the local peer with the bootstrap server

Establishes required IPC and network connections

disconnect

Disconnects from the bootstrap server and gracefully shuts down the daemon.
```bash
bloc disconnect
```

This:

Requests a clean daemon shutdown

Closes active connections

Leaves the repository intact
```bash
peer
```
Perform peer-related operations.
```bash
bloc peer [subcommand]
```

Example:
```bash
bloc peer list
```

Displays the list of known peers connected to the network.

completion

Generate shell autocompletion scripts.

bloc completion bash
bloc completion zsh
bloc completion fish


Follow the printed instructions to enable autocompletion for your shell.

help

Get help for any command.
```bash
bloc help
bloc help connect
```
Global Flags
-h, --help   Show help for bloc or a command

Typical Workflow
bloc init          # initialize repository
bloc connect       # start daemon and connect to network
bloc peer list     # inspect connected peers
bloc disconnect    # stop daemon

Notes

Runtime files (identity, keys, state) are stored locally and must not be committed.

The CLI communicates with blocd via IPC.

Network connections may close automatically when idle; this is expected behavior.

Troubleshooting
Daemon not running
daemon is not running
Run `bloc connect` to start the daemon.


Fix:
```bash
bloc connect
```