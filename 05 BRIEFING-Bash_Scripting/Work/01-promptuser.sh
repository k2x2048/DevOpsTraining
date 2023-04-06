#!/bin/bash
function Hello {
	echo "What is your name ?"
	read -e name
	echo Hello $name.
}
Hello