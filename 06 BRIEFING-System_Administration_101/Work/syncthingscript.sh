#!/bin/bash

function system_infos {
	printf "Linux version:\n\n"
	cat /proc/version #Linux version
	printf "\n\nHW infos:\n\n"
	sudo lshw -short #Some HW infos
	printf "\n\nPartitions, mount points and space:\n\n"
	df -H #partitions, mount points and space.
	printf "\n\nMem infos:\n"
	free #Mem infos
}

function infos_syncthing {
	systemctl status syncthing@$USER.service
}

function is_syncthing_installed {
	INSTALLSTATUS=`apt -qq list syncthing 2> /dev/null`
	if [[ "$INSTALLSTATUS" == *"[installed]" ]]; then
 		echo "True"
 	else
 		echo "False"
	fi
}

function is_syncthing_started {
	ACTIVESTATUS=`systemctl is-active syncthing@$USER.service 2> /dev/null`
	if [[ "$ACTIVESTATUS" == "active" ]]; then
 		echo "True"
 	else
 		echo "False"
	fi
}

function is_ufw_active {
	UFWACTIVE=`sudo ufw status`
	if [[ "$UFWACTIVE" == *"active"* ]]; then
 		echo "True"
 	else
 		echo "False"
	fi
}

function secure_syncthing {
	if [[ "$(is_ufw_active)" == "False" ]]; then
		echo "Configuring UFW."
		sudo ufw default deny incoming
		sudo ufw allow syncthing
		#sudo ufw allow syncthing-gui
		#sudo ufw allow OpenSSH
		#sudo ufw allow from 192.168.1.0/24 to any port 22 proto tcp
		sudo ufw enable
		if [[ "$(is_ufw_active)" == "True" ]]; then
			echo "UFW successfully configured and activated."
		else
			echo "Error while configuring UFW."
			exit 1
		fi
	else
		echo "UFW is already active."		
	fi
	#sudo ufw status
	#sudo ufw status verbose
	echo
	echo "Firewall infos:"
	sudo ufw app info syncthing
	sudo ufw app info syncthing-gui	
}

function install_syncthing {
	if [[ "$(is_syncthing_installed)" == "False" ]]; then
		echo "Installing syncthing ..."
		sudo apt update
		sudo apt install syncthing
		if [[ "$(is_syncthing_installed)" == "True" ]]; then
			echo "Syncthing successfully installed."
		else
			echo "Error while installing syncthing."
			exit 1
		fi
	else
		echo "Syncthing is already installed."
	fi
}


function configure_systemdsyncthing {
	#configure systemd service
	filename="/lib/systemd/system/syncthing@.service"
	if [[ -f "$filename" ]]; then
		echo "Service already configured ($filename)." 
	else
		echo "Downloading service config file ($filename)." 
		sudo wget -P /lib/systemd/system/ https://raw.githubusercontent.com/syncthing/syncthing/main/etc/linux-systemd/system/syncthing%40.service
	fi
}

function start_syncthing {
	if [[ "$(is_syncthing_started)" == "False" ]]; then
		echo "Starting syncthing ..."
		sudo systemctl daemon-reload
		sudo systemctl start syncthing@$USER.service
		sleep 1
		if [[ "$(is_syncthing_started)" == "True" ]]; then
			echo "Syncthing successfully started."
		else
			echo "Error while starting syncthing."
			exit 1
		fi
	else
		echo "Syncthing is already started."
		echo
		infos_syncthing
		echo
		read -p "Press Enter to continue with system infos ..."
		clear
		echo "**************************************************"
		system_infos
	fi
}

USER=`whoami`
install_syncthing
configure_systemdsyncthing
secure_syncthing
start_syncthing
