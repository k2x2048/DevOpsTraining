#!/bin/bash
function writedatetimeloop {
         while [ true ]; do
               message=`date +'%A %m/%d/%Y %H:%M:%S'`
               #printf "$message\n"
               printf "$message\n">>/home/user/scripts/serviscript.log
               sleep 10
         done
}
writedatetimeloop