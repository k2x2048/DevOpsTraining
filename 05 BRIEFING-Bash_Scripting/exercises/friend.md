# Are you my friend

Have you ever thought that you'd rather talk to your computer than to the people around you? Then you're in luck as today we will code ourselves a friend.

The idea is to create a sort of _bot_ which has the ability to **interpret some user's input** and **reply** to predefined sentences like: _what's the time_, _tell me a joke_ or _what's 4 + 5_. Your script should also be both **interactive** and **non-interactive**, but what does it mean? Follow, this [link](https://www.tldp.org/LDP/abs/html/intandnonint.html) to find out.

Your script must respect the rules below:

- The script must be able to **tell jokes randomly selected** from a file using the ``shuf` command or the `$RANDOM` environment variable.
- The script must be able to **tell the time** (`date`).
- The script must also be able to **calculate** simple equations.
- Implement both an **interactive** and **non-interactive** interface.

Optionally it could also:

- **Prettify** the output of your command.
- Implement audio output with [espeak](https://itsfoss.com/espeak-text-speech-linux/).
- Implement a **weather** question (Tips: _[wttr.in](https://github.com/chubin/wttr.in) API_).
- Add the script to your `$PATH` to make it executable from everywhere on your system.

It won't be _JARVIS_, but it'll be a friend you can _count on_... Jokes aside, shell scripting allows you to harness the power of the commands installed on your system in a recallable manner. Now that's power!

> **NOTE**: The goal is to try to make your friend as useful as possible, therefore don't hesitate to add new features that are not listed above.
