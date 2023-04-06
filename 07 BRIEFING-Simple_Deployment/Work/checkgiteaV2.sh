#!/bin/bash

function is_sql_service_ok {
	SQLSTATUS=`systemctl is-active mysql.service 2> /dev/null`
	if [[ "$SQLSTATUS" == "active" ]]; then
		echo "True"
	else
		echo "False"
	fi
}

function is_gitea_service_ok {
	GITEASTATUS=`systemctl is-active gitea.service 2> /dev/null`
	if [[ "$GITEASTATUS" == "active" ]]; then
		echo "True"
	else
		echo "False"
	fi
}

function is_ufw_ok {
	UFWACTIVE=`sudo ufw status`
	if [[ "$UFWACTIVE" == *"active"* ]]; then
		echo "True"
	else
		echo "False"
	fi
}

function is_webinterface_ok {
	WEBPAGE="localhost:3000" # web page to be loaded
	HTTPHEAD=`curl -I --max-time 5 --silent "$WEBPAGE"`
	if [[ "$HTTPHEAD" == "HTTP"*"200"* ]]; then
		echo "True"
	else
		echo "False"
	fi
}

function is_webinterface_ok_v2 {
	WEBPAGE="localhost:3000" # web page to be loaded
	HTTPSTATUS=$(curl -LI $WEBPAGE -o /dev/null -w '%{http_code}\n' -s)
	if [ "$HTTPSTATUS" -eq 200 ]; then
		echo "True"
	else
		echo "False"
	fi
}


function sql_service_status {
	SQLSTATUS=`systemctl is-active mysql.service 2> /dev/null`
	printf "MySQL Service is $SQLSTATUS\n"
}

function gitea_service_status {
	GITEASTATUS=`systemctl is-active gitea.service 2> /dev/null`
	printf "Gitea Service is $GITEASTATUS\n"
}

function webinterface_status {
	if [[ "$(is_webinterface_ok)" == "True" ]]; then
		printf "Webinterface ok\n"
	else
		printf "Webinterface error\n"
	fi
}

function ufw_status {
	UFWSTATUS=`sudo ufw status | head -n1`
	printf "UFW $UFWSTATUS\n"
}

PREVIOUS_STATUS=""
function checkloop {
	while [ true ]; do
	if [[ "$(is_sql_service_ok)" == "False" ]] || [[ "$(is_gitea_service_ok)" == "False" ]] || [[ "$(is_ufw_ok)" == "False" ]] || [[ "$(is_webinterface_ok)" == "False" ]]; then
		#https://mailtrap.io/
		DATE=`date +'%A %m/%d/%Y %H:%M:%S'`
		MESSAGE=`sql_service_status;gitea_service_status;webinterface_status;ufw_status`
		if [[ "$PREVIOUS_STATUS" != "$MESSAGE" ]]; then
			#printf "$MESSAGE / mail -s ERROR on Gitea Server \($DATE\) user@localtest.net\n" 
			printf "$MESSAGE" | mail -s "ERROR on Gitea Server ($DATE)" user@localtest.net
			PREVIOUS_STATUS="$MESSAGE"
		fi
	else
		if [[ "$PREVIOUS_STATUS" != "OK" ]]; then
			DATE=`date +'%A %m/%d/%Y %H:%M:%S'`
			MESSAGE=`sql_service_status;gitea_service_status;webinterface_status;ufw_status`
			#printf "$MESSAGE / mail -s Gitea Server OK \($DATE\) user@localtest.net\n" 
			printf "$MESSAGE" | mail -s "Gitea Server OK ($DATE)" user@localtest.net
			PREVIOUS_STATUS="OK"
		fi
	fi
	sleep 5
	done
}
checkloop

