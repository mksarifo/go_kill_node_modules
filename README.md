# Node Modules Directory Scanner

This application scans your file system for `node_modules` directories and prints out their locations and sizes. The intended use is for developers who frequently work with node.js and accumulate a large number of `node_modules` directories that take up disk space.

## Features
1. Scans a given directory for `node_modules` folders.
2. Calculates the size of each `node_modules` found.
3. Prints the path and size of each `node_modules` folder.

## Installation
Assuming you have a working Go environment set up:

sh git clone https://github.com/username/repository.git cd repository go build -o scanner

## Usage
After installation, you can run the application with:
sh ./scanner

You will be prompted to input a directory path to start the scan.

## Warning
This tool only reports on the size of `node_modules` directories. It does NOT automatically delete them. Please be cautious and make sure you no longer need the content before deleting any `node_modules` directory.

## Contributing
If you have suggestions for how to improve the scanner, open an issue in the Issues queue for this project. Contributions are welcome, and they are greatly appreciated!