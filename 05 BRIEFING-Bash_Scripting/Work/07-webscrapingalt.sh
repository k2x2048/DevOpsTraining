#!/bin/bash

url=$@				# variable to take arguments
curl  $url | 			#curl url 
html2text | 			#clean html text 
grep -A4 -F "[item]" |		# matching item and print the 4 line after 
sed  "s/.*\.\.\.//g" |		#replace string finishing by ...
sed ':a;N;$!ba;s/\n/ /g' | 	#replace line break by space 
sed 's/item/\n/g;s/\[//g;s/\]//g' |	#replace item by a break line and rm [] 
sed 's/\*//g' | 		# replace * by nothing
sed 's/\t/ /g' |		# replace tabulation by space 
sed 's/--//g' 