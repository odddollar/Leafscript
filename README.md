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

## Examples

***Ensure tabs are used to indent and not spaces. Spaces do not work***

Additional language examples are contained within the "Examples" file

#### Find all the factors of a number
```
var number = inputint "Enter a number to find factors of: "

for x math number+1
	if math number%x == 0
		print concat math x & " is a factor of " & math number
	endif
endfor
```

#### Print all binary numbers from 0 to 15
```
for x 2
	for y 2
		for z 2
			for i 2
				print concat math x & " " & math y & " " & math z & " " & math i
			endfor
		endfor
	endfor
endfor
```

#### Print all primes up to a given number
```
var number = inputint "Enter a number to find primes up to: "
var total = 0

for x math number
	var y = 0
	for i math x
		if math x%i == 0
			var y = math y+1
		endif
	endfor

	if math y < 2
		print concat math x & " is prime"
		var total = math total+1
	endif
endfor

print concat math total & " primes found"
```
