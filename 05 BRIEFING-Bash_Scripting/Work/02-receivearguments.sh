#!/bin/bash
function HelloArg {
	name=$@
	if [ -z $name ] ; then
		name="anonymous"
	fi
	echo Hello $name.
}

HelloArg $@