# XAMPP Manager Runner

This is a simple Go program that runs the XAMPP manager on Linux systems.

## Prerequisites

- Go programming language installed on your system
- XAMPP installed in the `/opt/lampp` directory
- Sudo privileges

## Installation

1. Clone this repository or download the `run_xampp_manager.go` file.

2. Open a terminal and navigate to the directory containing the Go file.

## Usage

1. Compile the program:
   ```
   go build go-lampp.go
   ```

2. Run the compiled program with sudo:
   ```
   sudo ./go-lampp
   ```
   
3. OR Just install program into your system:
  ```
  go install go-lampp
  ```

## What it does

This program:
1. Changes the current working directory to `/opt/lampp`
2. Runs the XAMPP manager using the command `sudo ./manager-linux-x64.run`

## Note

This program requires sudo privileges to run the XAMPP manager. Make sure you have the necessary permissions before running it.

## Troubleshooting

If you encounter any errors, check that:
- XAMPP is installed in the correct directory (`/opt/lampp`)
- You have the necessary permissions to access the XAMPP directory and run the manager
- The XAMPP manager executable (`manager-linux-x64.run`) exists and is executable

For any other issues, please refer to the XAMPP documentation or seek assistance from the XAMPP community forums.