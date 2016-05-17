# Wattpad Coding Challenge

## Problem

Wattpad automatically flags comments and messages that are deemed offensive. This is done by detecting key phrases in the text and assigning it a score. If the score is over a certain threshold, it is flagged as offensive.


Write a program in Go or PHP that reads input files of potentially offensive text and writes to an output file with scores for each of the text files (details below).


You are given two files with lists of offensive phrases. One file contains "low risk" phrases and the other, "high risk" phrases, one phrase per line. You are also given a set of 15 input files, each one containing some possibly offensive text that your program will score. The offensive score is defined as:


(number of low risk phrases) + (number of high risk phrases * 2)


Your program should write out one output file containing the scores of each input file in order, in the format:


<input-filename-1>:<score-1>
<input-filename-2>:<score-2>
...


To not offend you, we've adjusted the actual words we look for with nicer words :)


Here are the files:
https://drive.google.com/drive/folders/0B9xNhpxBeiiOfmxUZzBESTJSZ0hFalYtZ2VHVndMOTEzVnRjQlRPUGVwcXppTDZXSWdad1k

## Project Structure

- /handlers: contains the handler dealing with result model
- /input_files: contains all of 15 input file in txt format
- /models: contains result model
- /output_files: contains a output file after execution of project
- /risk_phrases_files: contains high_risk_phrases.txt & low_risk_phrases.txt
- /tests: contains unit test files
- /utils: contains all the utilities we need in this project
- wattpad.go: contains main function

## Unit Test Tool

We use Goconvey as unit test tool in this solution. You can find details of Goconvey in the following link:

- https://github.com/smartystreets/goconvey.git

### Installation
```
$ go get github.com/smartystreets/goconvey
```

### Testing through CLI
Simply use the command below in your testing directory
```
$ go test
```

### Testing through Goconvey web server
Start up the GoConvey web server at $GOPATH:
```
GOPATH/bin/server$ goconvey

```

Then open your browser to:
```
http://localhost:8080
```

## Run the Project

Go to the project directory and run the command below
```
$ go run wattpad.go
``` 

After execution, the output file will be generated in ~/output_files/output.txt