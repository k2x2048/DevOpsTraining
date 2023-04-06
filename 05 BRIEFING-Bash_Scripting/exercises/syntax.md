# Bash Introduction - Syntax

It's a good thing to know what a **script** is, it's even better to know how to write one yourself. Hopefully that's exactly what you're going to practice here. Follow this [tutorial](https://www.learnshell.org/) (_or others_) on bash scripting then complete the exercises below to ensure fundamental understanding of the syntax. 

Here are some _tips_ to help you out:

* You can append the extension `.sh` to your script for clarity.
* A [shebang](https://en.wikipedia.org/wiki/Shebang_(Unix)) is needed to make sure the shell knows what interpreter should be used.
* To make a file executable use the [chmod](https://askubuntu.com/questions/229589/how-to-make-a-file-e-g-a-sh-script-executable-so-it-can-be-run-from-a-termi) command.
* One way to execute your script is to type `./<path to file>`.

Each exercise should be solved by a [function](https://linuxize.com/post/bash-functions/) and stored in a single file. If you have time try to make your script use all of these functions in a show case seen when the script is executed.

> **NOTE**: You can also find more exercises on [Codewars](https://www.codewars.com/) by selecting the [shell language](https://www.codewars.com/?language=shell) if you want more (_potentially complex_) practice!

## Exercise I - prompt user

Write a shell script *prompting* the user for his name, then replying `Hello <name>`.

## Exercise II - receive arguments

Write a shell script *receiving* a name as *argument*, then replying `Hello <name>`.

## Exercise III - path to

Write a shell script receiving a *path* as argument, depending on if it's a file or a directory, display or list its content.

You can complexify the script by only displaying file containing text (ex: `.txt`, `.js`) and returning an error for anything else (ex: `.mp3`, `.pdf`).

## Exercise IV - information

Write a shell script to see the current date, time, username, and directory.

## Exercise V - list of groups

Write a shell script checking all the [groups](https://www.cyberciti.biz/faq/linux-show-groups-for-user/) a user is part of and for each one display `<username> is part of the group <group>`.
