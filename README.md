# Leafscript
Leafscript is a lightweight programming language created as a proof of concept by someone with no idea how to write a language.
It's written entirely in Golang, and was inspired by the speed and simplicity of Lua. It can be compiled from source to a binary .exe file using the command:

```go build main.go```

## Usage
Programs can be run from the command line using the Leafscript binary file.

```[NAME OF BINARY FILE] -run [PATH TO .lfs FILE] -debug [SET TO FALSE BY DEFAULT]```

E.g. 

```leafscript -run program.lfs```

Includes a basic debugger that prints a list of all variables every line. E.g. 

```leafscript -run program.lfs -debug true```


## Features
It supports:
- Creating and modifying variables (strings, ints and floats)
- Performing mathematical operations on numeric variables
- String concatenation
- For loops
- If/else statements
- Breaks in for loops
- Nested if/for
- Basic debugging mode that prints all variables every line

Currently working on:
- Error messages

No plans to implement in near future
- Array variables

Language examples are contained within the "Examples" file

***Ensure tabs are used to indent and not spaces. Spaces do not work***
