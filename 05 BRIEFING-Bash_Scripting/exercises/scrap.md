# Exercise Bash Introduction - Give me your data

The web is full of useful data that you might want to process on your system. The most common and reliable way to get it being [APIs](https://www.mulesoft.com/resources/api/what-is-an-api#:~:text=API%20is%20the%20acronym%20for,you're%20using%20an%20API), but they don't always exists. There is an alternative however (_although not always legal_), a practice called [web scraping](https://www.techopedia.com/definition/5212/web-scraping) that can help you in case you are unable to get the data by other means, this exercise will have you explore it.

After learning how to use the [curl](https://curl.haxx.se/) command to **get a page's HTML**. You should write a script to get all the _laptops_ from [this](https://webscraper.io/test-sites/e-commerce/allinone/computers/laptops) webshop page and print each one of them on a line in your terminal. 

The example below shows a possible output format, but the important thing is that for each laptop you have: **the name**, **the price** and **the description**.

``` bash
## Example of output for your script (one line)
Aspire E1-510 | 15.6", Pentium N3520 2.16GHz, 4GB, 500GB, Linux | $306.99
```

Follow the instructions below to guide you through the process.

- Research and learn about the `curl` command.
- Research and learn about **web scraping with the shell**.
- Actually write the script (Tips: you could use `sed`, `cut` or/and `awk`).

> **NOTE**: Web scraping with shell scripting is not the simplest of things and in the future you will most likely use languages like _Python_ or _Nim-Lang_ with already existing libraries to do it, but you should see this challenge as an opportunity to practice the `cut`, `sed`, and/or `awk` commands.
