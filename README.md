# Project Title

## Website Status Checker

## Description

The Website Status Checker is a simple Go program designed to check the status of websites by sending HTTP requests and verifying their accessibility. It operates concurrently, ensuring quick checks even for a large number of websites.

## Features

1. Concurrent Website Checking: Utilizes Go's concurrency model (goroutines) for efficient status checking.
2. Custom URL Input: Users have the flexibility to specify their own set of URLs via command-line arguments.
3. Default URL Set: If no URLs are provided, the program will check a predefined list of popular websites

## Usage

First, ensure you have Go installed. If not, you can download and install it from the official Go website.

Clone the repository:

```bash
git clone https://github.com/GiorgiMakharadze/website-status-checker-golang.git
```

```bash
cd path-to-status-checker-directory
```

### Running the Program

```bash
go run main.go
```

To check custom URLs, provide them as command-line arguments:

```bash
go run main.go http://example1.com http://example2.com
```

## Testing:

To run the unit tests:

```bash
go test
```
