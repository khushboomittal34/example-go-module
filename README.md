# [:] Example Go project using Go modules

An example Go project that uses Go modules to demonstrate Veracode agent-based scans.

## Install and activate SourceClear
Follow the instructions in [Setting Up Veracode Agent-Based Scan](https://help.veracode.com/reader/hHHR3gv0wYc2WbCclECf_A/t9Dav~Bju~GJGoKWKe1WKA) to install and activate the CLI scan agent.

## Scan this project
There are 2 ways to scan this project.

### 1. Using url option
`srcclr scan --url https://github.com/srcclr/example-go-modules`

### 2. On local path
```
1. git clone https://github.com/srcclr/example-go-modules
2. cd example-go-modules
3. srcclr scan .
```
