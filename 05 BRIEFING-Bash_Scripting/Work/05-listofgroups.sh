#!/bin/bash

function DisplayGroups1 {
	#This dont work for groups with spaces
	USER=`whoami`
	USERGROUPS=`whoami | groups`
	for GROUP in $USERGROUPS
	do
		echo $USER is part of the group $GROUP
	done
}

function DisplayGroups2 {
	#This dont work for groups with spaces
	USER=`whoami`
	USERGROUPS=`id -Gnz $USER | tr "\0" ";"`
	IFS=$';'
	for GROUP in $USERGROUPS
	do
		echo $USER is part of the group $GROUP
	done
}

function DisplayGroups3 {
	USER=`whoami`
	id -Gnz $USER | tr "\0" "\n" | sed "s/^/$USER is part of the group /g"
}

function DisplayGroups4 {
	USER=`whoami`
	id -Gnz $USER | tr "\0" "\n" | awk -v USERNAME=$USER '{print USERNAME" is part of the group " $0}'
}

printf "\nOption 1:\n"
DisplayGroups1
printf "\nOption 2:\n"
DisplayGroups2
printf "\nOption 3:\n"
DisplayGroups3
printf "\nOption 4:\n"
DisplayGroups4