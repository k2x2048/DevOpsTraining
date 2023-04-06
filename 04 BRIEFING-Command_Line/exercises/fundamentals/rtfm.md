# RTFM - Read the Frickin' Manual

The terminal can be intimidating with all these commands and text flying around. It's an all new world, one of textual interface (_TUI_) with rules very different from a more familiar graphical interface (_GUI_). Hopefully, to guide you in this new environment there are tools like `man`.

- open your **terminal emulator** (_that's where the CLI lives_)
- type `man man`
- hit the `h` key
- leave the manual

You might have guessed it, `man` opens the manual for the wanted page and although it can be technical it remains one of the best place for information. Still, if you simply want a description let's try to use `whatis`.

- type `whatis man`
- type `whatis whatis`

Another great, but more concise, way to understand how to use a command is the help [flag](http://www.tldp.org/LDP/abs/html/standard-options.html).

* type `whatis --help`

You won't always find a manual page, but you will usually be able to get information via the help flag. Practice by learning about `ls`, `cp` and `pwd`.

There are multiple alternatives to this method, you can optionally look into command like [info]() or better and more comprehensive yet [TLDR](https://tldr.sh/).
