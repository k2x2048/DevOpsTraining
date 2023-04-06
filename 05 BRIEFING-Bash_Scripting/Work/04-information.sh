#!/bin/bash
function DiplayInfos1 {
	date +'%A %m/%d/%Y'
	date +'%H:%M:%S'
	whoami
	pwd
}

function DiplayInfos2 {
	printf "\t`date +'%A %m/%d/%Y'`\n"
	printf "\t`date +'%H:%M:%S'`\n"
	printf "\t`whoami`\n"
	printf "\t`pwd`\n"
}
function DiplayInfos3 {
	DATE=`date +'%A %m/%d/%Y'`
	TIME=`date +'%H:%M:%S'`
	USERNAME=`whoami`
	DIRECTORY=`pwd`
	printf "\tCurrent date = $DATE\n"
	printf "\tCurrent time = $TIME\n"
	printf "\tCurrent user = $USERNAME\n"
	printf "\tCurrent directory = $DIRECTORY\n"
}

clear
printf "Option1:\n"
DiplayInfos1
printf "\n"
printf "Option2:\n"
DiplayInfos2
printf "\n"
printf "Option3:\n"
DiplayInfos3