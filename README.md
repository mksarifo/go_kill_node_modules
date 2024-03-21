# GO Kill Node Modules - Node Modules Directory Scanner

This application scans your file system for `node_modules` directories and prints out their locations and sizes. The intended use is for developers who frequently work with node.js and accumulate a large number of `node_modules` directories that take up disk space.

## Features
1. Scans a given directory for `node_modules` folders.
2. Calculates the size of each `node_modules` found.
3. Prints the path and size of each `node_modules` folder.

## Known Issues
(Feel free to contribute)
- [ ] When deleting multiple directories at the same time, it crashes
- [ ] The UI needs some improvements
- [ ] Reclaimed space (after delete) is not being calculated


## Installation
Assuming you have a working Go environment set up:

`git clone https://github.com/mksarifo/go_kill_node_modules.git`

`cd repository`

`go build`

## Usage
After installation, you can run the application with:
`./nmkill /path/to/scan`

It will scan the provided path and show a list of results.

## Warning
This tool only reports on the size of `node_modules` directories. It does NOT automatically delete them. Please be cautious and make sure you no longer need the content before deleting any `node_modules` directory.

## Contributing
If you have suggestions for how to improve the scanner, open an issue in the Issues queue for this project. Contributions are welcome, and they are greatly appreciated!