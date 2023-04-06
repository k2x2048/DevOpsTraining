#!/bin/bash

function getwebpage() { 
	echo Downloading \'$1 to \'$2
	curl -o webpage.html $1
}

#must be escaped: !"$&'()*,:;<=>?@[\]^`{|}

function getLTlines() {
	echo Extracting LT lines from \'$1 to \'$2
	EXPL1="<h4 class=\"pull-right price\">\$[0-9.]*</h4>"
	EXPL2="<a href=\"/test-sites/e-commerce/allinone/product/[0-9.]*\" class=\"title\" title=\".*\">.*</a>"
	EXPL3="<p class=\"description\">.*</p>"
	#echo $EXPL1
	#echo $EXPL2
	#echo $EXPL3
	grep -io -e "${EXPL1}" -io -e "${EXPL2}" -io -e "${EXPL3}" $1 > $2
}

function getLTtexts() {
	echo Extracting LT texts from \'$1 to \'$2
	
	#<h4 class="pull-right price">$1399.00</h4>
	#<a href="/test-sites/e-commerce/allinone/product/632" class="title" title="Asus ROG Strix GL702VM-GC146T">Asus ROG Strix G...</a>
	#<p class="description">Asus ROG Strix GL702VM-GC146T, 17.3&quot; FHD, Core i7-7700HQ, 8GB, 1TB + 128GB SSD, GeForce GTX 1060 3GB, Windows 10 Home, Eng kbd</p>
	
	IFS=$'\n' read -d '' -r -a MYARRAY < $1
	echo Laptop list > $2
	length=${#MYARRAY[@]}
	for (( i=0; i<${length}; i+=3 )); do
		PRICE=${MYARRAY[i]}
		PRICE=${PRICE//"<h4 class=\"pull-right price\">"/""}
		PRICE=${PRICE//"</h4>"/""}
		NAME=${MYARRAY[i+1]}
		NAME=${NAME//"<a href=\"/test-sites/e-commerce/allinone/product/"/""}
		NAME=${NAME//"\" class=\"title\" title=\""/": "}
		NAME=${NAME//"</a>"/""}
		NAME=${NAME//"&quot;"/"\""}
		NAME=$( cut -d'"' -f1  <<<"$NAME" )
		DESCRIPTION=${MYARRAY[i+2]}
		DESCRIPTION=${DESCRIPTION//"<p class=\"description\">"/""}
		DESCRIPTION=${DESCRIPTION//"</p>"/""}
		DESCRIPTION=${DESCRIPTION//"&quot;"/"\""}
		echo $NAME \| $DESCRIPTION \| $PRICE >> $2
	done
}

getwebpage https://webscraper.io/test-sites/e-commerce/allinone/computers/laptops webpage.html
getLTlines webpage.html webpageLT.html
getLTtexts webpageLT.html 07-webscraping.txt
cat 07-webscraping.txt