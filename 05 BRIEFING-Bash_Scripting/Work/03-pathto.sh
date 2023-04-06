#!/bin/bash
function DisplayText {
	if [ -d $path ] ; then
		echo Listing text files in \"$path\":
		find . -type f | perl -lne 'print if -T'
	elif [ -f $path ] ; then
		FILETYPE=`file $path`
		if [[ "$FILETYPE" == *"text"* ]] ; then
			cat $path
		else
			echo \'$path\' is not a text file.
			exit 1
		fi
	else
		echo Please specify a valid file or directory.
		exit 1
	fi
}

path=$@
if [ -z $path ] ; then
	echo Usage: $0 \<file\|directory\>.
	exit 1
fi
DisplayText $path
