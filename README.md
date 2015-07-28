# polar2cartesian
Example program using two communication channels

This example program uses two communication channels and does its processing in a separate Go routine.

The program is an interactive console program that prompts the user to enter two whitespace-separated numbers - a radius and an angle - which the program then uses to compute the equivalent cartesian coordinates. In addition to illustrating one particular approach to concurrency, it also shows some simple structs and how to determine if the program is running on a Unix-like system or on Windows for when the difference matters.

--Programming in Go, Mark Summerfield
