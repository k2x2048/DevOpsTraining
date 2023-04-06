# Environment variables

In any programming language you can store data within [variables](https://launchschool.com/books/ruby/read/variables). However they only exist as long as the program is executed, but what if you wanted to share information between two applications that are not running at the same time.

Then, you could use **environment variables**, which are _variables_ stored on the [shell](https://bash.cyberciti.biz/guide/What_is_Linux_Shell), allowing you to access them from any place that has access to the it. There are five main commands to deal with environment variables in most shells: `env`, `printenv`, `set`, `unset` and `export`. Open your **terminal** and **read their manuals** to understand them, then follow the instructions below.

> **NOTE**: Some of the given command can achieve the same thing, but they are slightly different in there goals.

- print the value of `$HOME`
- create a variables called _NAME_ containing your name
- seek the differences between _temporary_ and _permanent_ variables
- list all the variables on your system
- create a permanent variables (hint: _bashrc_)
- unset your variable `$NAME`

Optionally and if you have time you should:

- research what's a _.env file_
- write one
- load it
