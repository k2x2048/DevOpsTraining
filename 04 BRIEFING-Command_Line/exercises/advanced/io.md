# Pipes and Redirection

Do you know what [I/O](https://en.wikipedia.org/wiki/Input/output) is? It means **input**/**output** and it's basically what the user sees of the interaction between himself and the computer. The shell as a very straightforward I/O system in which the **input is what you
type** and the **output is the answer you receive**.

Well, the _shell_ is such a *wonderful tool* that it allows you to take an output and send it to a **file**, a **stream of data** or even another **command**.

- research output redirection
- redirect the output of the command `pwd` to a file
- redirect the output of one file to another (add to the file content)
- redirect the output of one file to another (overwrite the file content)
- redirect an output to `stderr`
- research pipes (_man pipe_)
- pipe the output of `ls /bin` into the pager `less`
- try to pipe the output of a command through `grep`

Optionally and if you have time you should:

- research [named pipes](https://hackaday.com/2019/07/12/linux-fu-named-pipe-dreams/) (_FIFO_)
- redirect the output of a file into a command without a _pipe_
- learn about `xargs`
