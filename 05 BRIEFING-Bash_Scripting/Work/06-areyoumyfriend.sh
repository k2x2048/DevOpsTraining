#!/bin/bash

#echo test &>/dev/null

botname="Jarvis"

function randomnumberminmax {
	minimum=$1
	maximum=$2
	#Generate the random number
	randomNumber=$(($minimum + $RANDOM % $maximum))
	#Print the generated number
	echo $randomNumber
}

function calculate {
	operations=$@
	echo `echo $operations | bc`
}

function randomjoke {
	echo `shuf -n 1 ~/script/jokes.txt`
}

 
#Black        0;30     Dark Gray     1;30
#Red          0;31     Light Red     1;31
#Green        0;32     Light Green   1;32
#Brown/Orange 0;33     Yellow        1;33
#Blue         0;34     Light Blue    1;34
#Purple       0;35     Light Purple  1;35
#Cyan         0;36     Light Cyan    1;36
#Light Gray   0;37     White         1;37

function echobot {
	displaystring=$@
	COLOR='\033[0;34m' #Blue
	NC='\033[0m' # No Color
	
	for (( i=0; i<${#displaystring}; i++ )); do
		echo -n -e "${COLOR}${displaystring:$i:1}${NC}"
		sleep 0.00$((0 + $RANDOM % 9))
	done
	echo
}


function echouser {
	displaystring=$@
	COLOR='\033[1;33m' #Yellow
	NC='\033[0m' # No Color
	echo -n -e "${COLOR}$displaystring${NC}"
}

function gettime {
	DATE=`date +'%A %d %B %Y'`
	echo $DATE
}

function getdate {
	TIME=`date +'%H:%M:%S'`
	echo $TIME
}

function weather {
	shopt -s nocasematch
	CITY=`echo $@ | awk -F'weather ' '{print $2}' | cut -d " " -f1`
	if [ ! -z "$CITY" ] ; then
		curl wttr.in/$CITY
	else
		curl wttr.in
	fi
}

function commandeval {
	#set comparison case insensitive
	shopt -s nocasematch
	#set enable sort of regex in case
	#shopt -s extglob
	#set -s extglob
	case $@ in
	hello | "hello"* | hi | "hi "* )
		echobot "Hello sir"
		;;
	bye | seey* | "exit" | quit | end )
		echobot "See you again!"
		exit 0
		;;
	"time" | *date*? )
		echobot `getdate`
		;;
	"date" | *time*? )
		echobot `gettime`
		;;
	joke* | *tel*joke*)
		echobot `randomjoke`
		;;
	*weather* | *meteo*)
		weather $@
		;;
	help* )
		echobot "Commands: help, exit, hello, bye, time, date, help, calc, joke, weather"
	;;
	#"calc "* | [0-9.,]+[[:space:]]*[+-/*^%][[:space:]]*[0-9.,] )
	"calc "* )
		operations=$@
		operations=${operations//"calc "/""}
		echobot `calculate $operations`
	;;
	[0-9]*[!A-Za-z]*[0-9])
	#[\(0-9]*[0-9\)] )
		operations=$@
		echobot `calculate $operations`
	;;
	cmd*)
		CMD=`echo $@ | cut -d " " -f 2-`
		eval "$CMD"
	;;
	*)
		echobot "Sorry, I don't understand"
		;;
  esac
}


name=$@
#if command in parameter
if [ ! -z $1 ] ; then
	commandeval $name
	exit
fi


echobot "Hello, I am $botname"
echobot "How can I help you ?"
USERINPUT=""
# read keyboard input

while :
do
	read -e -p ">" USERCOMMAND
	commandeval $USERCOMMAND
done

PATH="$PATH:~/script"