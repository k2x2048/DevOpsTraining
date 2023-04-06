# Practice Exercises

## 1. Hello World

Write a program that greets the user by name, or by saying `Hello, World!` if no name is given.

## 2. Pangram 

Determine if a sentence is a pangram. A pangram (Greek: παν γράμμα, pan gramma, "every letter") is a sentence using every letter of the alphabet at least once. The best known English pangram is `The quick brown fox jumps over the lazy dog`.

## 3. Rna Transcription 

Write a program that, given a DNA strand, returns its RNA complement (_per RNA transcription_).

## 4. Word Count 

Write a program that given a phrase can count the occurrences of each word in that phrase.

For example for the input `olly olly in come free`:

```plain
olly: 2
in: 1
come: 1
free: 1
```

## 5. Gigasecond

Write a program that will calculate the date that someone will celebrate their one gigasecond anniversary (_one billion seconds_).

## 6. Run Length Encoding 

Implement run-length encoding and decoding. Run-length encoding (RLE) is a simple form of data compression, where runs (_consecutive data elements_) are replaced by just one data value and count.

For example we can represent the original 53 characters with only 13.

```plain
"WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWB"  ->  "12WB12W3B24WB"
```

RLE allows the original data to be perfectly reconstructed from the compressed data, which makes it a lossless data compression.

```plain
"AABCCCDEEEE"  ->  "2AB3CD4E"  ->  "AABCCCDEEEE"
```

## 7. Difference Of Squares 

Find the difference between the sum of the squares and the square of the sums of the first N natural numbers.

The square of the sum of the first ten natural numbers is,

`(1 + 2 + ... + 10)**2 = 55**2 = 3025`

The sum of the squares of the first ten natural numbers is,

`1**2 + 2**2 + ... + 10**2 = 385`

Hence the difference between the square of the sum of the first ten natural numbers and the sum of the squares is 2640:

`3025 - 385 = 2640`

## 8. Anagram 

Write a program that, given a word and a list of possible anagrams, selects the correct sublist.

Given `listen` and a list of candidates like `enlists` `google` `inlets` `banana` the program should return a list containing `inlets`.

## 9. Allergies 

Write a program that, given a person's allergy score, can tell them whether or not they're allergic to a given item, and their full list of allergies.

An allergy test produces a single numeric score which contains the information about all the allergies the person has (that they were tested for). The list of items (and their value) that were tested are:

* eggs (1)
* peanuts (2)
* shellfish (4)
* strawberries (8)
* tomatoes (16)
* chocolate (32)
* pollen (64)
* cats (128)

So if Tom is allergic to peanuts and chocolate, he gets a score of 34. Now, given just that score of 34, your program should be able to say:

* Whether Tom is allergic to any one of those allergens listed above.
* All the allergens Tom is allergic to.

## 10. Series 

Write a program that will take a string of digits and give you all the contiguous substrings of length `n` in that string.

For example, the string "49142" has the following 3-digit series:

* 491
* 914
* 142

And the following 4-digit series:

* 4914
* 9142

And if you ask for a 6-digit series from a 5-digit string, you deserve whatever you get.

NOTE: that these series are only required to occupy adjacent positions in the input; the digits need not be _numerically consecutive_.

## 11. Robot Simulator 

Write a robot simulator. A robot factory's test facility needs a program to verify robot movements. The robots have three possible movements:

* turn right
* turn left
* advance

Robots are placed on a hypothetical infinite grid, facing a particular direction `[north, east, south, west]` at a set of `{x,y}` coordinates, e.g., `{3,8}`, with coordinates increasing to the north and east.

The robot then receives a number of instructions, at which point the testing facility verifies the robot's new position, and in which direction it is pointing.

The letter-string "RAALAL" means:

* turn right
* advance twice
* turn left
* advance once
* turn left yet again

Say a robot starts at {7, 3} facing north. Then running this stream of instructions should leave it at {9, 4} facing west.

## 12. Atbash Cipher

Create an implementation of the atbash cipher, an ancient encryption system created in the Middle East.

The Atbash cipher is a simple substitution cipher that relies on transposing all the letters in the alphabet such that the resulting alphabet is backwards. The first letter is replaced with the last letter, the second with the second-last, and so on. An Atbash cipher for the Latin alphabet would be as follows:

```plain
Plain:  abcdefghijklmnopqrstuvwxyz
Cipher: zyxwvutsrqponmlkjihgfedcba
```

It is a very weak cipher because it only has one possible key, and it is a simple monoalphabetic substitution cipher. However, this may not have been an issue in the cipher's time.

Ciphertext is written out in groups of fixed length, the traditional group size being 5 letters, and punctuation is excluded. This is to make it harder to guess things based on word boundaries.

For example:

- Encoding `test` gives `gvhg`.
- Decoding `gvhg` gives `test`.
- Decoding `gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt` gives `the quick brown fox jumps over the lazy dog`.

## 13. Sum Of Multiples

Write a program that, given a number, can find the sum of all the multiples of particular numbers up to but not including that number.

## 14. Say

Write a program that will take a number from 0 to 999,999,999,999 and spell out that number in English.

## 15. Kindergarten Garden 

Write a program that, given a diagram, can tell you which plants each child in the kindergarten class is responsible for.

The kindergarten class is learning about growing plants. The teachers thought it would be a good idea to give them actual seeds, plant them in actual dirt, and grow actual plants. They've chosen to grow grass, clover, radishes, and violets.

To this end, they've put little styrofoam cups along the window sills, and planted one type of plant in each cup, choosing randomly from the available types of seeds.

There are 12 children in the class:

* Alice, Bob, Charlie, David,
* Eve, Fred, Ginny, Harriet,
* Ileana, Joseph, Kincaid, and Larry.

Each child gets 4 cups, two on each row. The children are assigned to cups in alphabetical order. The following diagram represents Alice's plants:

```plain
[window][window][window]
VR......................
RG......................
```

So in the row nearest the window, she has a violet and a radish; in the row behind that, she has a radish and some grass. Your program will be given the plants from left-to-right starting with the row nearest the windows. From this, it should be able to determine which plants belong to which students. For example, if it's told that the garden looks like so:

```plain
[window][window][window]
VRCGVVRVCGGCCGVRGCVCGCGV
VRCCCGCRRGVCGCRVVCVGCGCV
```

Then if asked for Alice's plants, it should provide: `Violets, radishes, violets, radishes`.

While asking for Bob's plants would yield: `Clover, grass, clover, clover`.

## 16. Grade School 

Write a small archiving program that stores students' names along with the grade that they are in.

In the end, you should be able to:

- Add a student's name to the roster for a grade.
- Get a list of all students enrolled in a grade.
- Get a sorted list of all students in all grades. Grades should sort as 1, 2, 3, etc., and students within a grade should be sorted alphabetically by name.

> **NOTE**: All the students only have one name. (_It's a small town, what do you want?_).

## 17.Roman Numerals 

Write a function to convert from normal numbers to Roman Numerals.

## 18. Luhn 

Write a program that can take a number and determine whether or not it is valid per the Luhn formula.

The Luhn formula is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers and Canadian Social Insurance Numbers.

The formula verifies a number against its included check digit, which is usually appended to a partial number to generate the full number. This number must pass the following test:

- Counting from rightmost digit (which is the check digit) and moving left, double the value of every second digit.
- For any digits that thus become 10 or more, subtract 9 from the result. `1111 becomes 2121` and `8763 becomes 7733 (from 2×6=12 → 12-9=3 and 2×8=16 → 16-9=7)`.
- Add all these digits together. `1111 becomes 2121 sums as 2+1+2+1 to give a check digit of 6` and `8763 becomes 7733, and 7+7+3+3 is 20`.

If the total ends in 0 (put another way, if the total modulus 10 is congruent to 0), then the number is valid according to the Luhn formula; else it is not valid. So, 1111 is not valid (as shown above, it comes out to 6), while 8763 is valid (as shown above, it comes out to 20).

Write a program that, given a number:

- Can check if it is valid per the Luhn formula. This should treat, for example,`2323 2005 7766 3554` as valid.
- Can return the checksum, or the remainder from using the Luhn method.
- Can add a check digit to make the number valid per the Luhn formula and return the original number plus that digit. This should give `2323 2005 7766 3554` in response to `2323 2005 7766 355`.

### About Checksums

A checksum has to do with error-detection. There are a number of different ways in which a checksum could be calculated.

When transmitting data, you might also send along a checksum that says how many bytes of data are being sent. That means that when the data arrives on the other side, you can count the bytes and compare it to the checksum. If these are different, then the data has been garbled in transmission.

In the Luhn problem the final digit acts as a sanity check for all the prior digits. Running those prior digits through a particular algorithm should give you that final digit.

It doesn't actually tell you if it's a real credit card number, only that it's a plausible one. It's the same thing with the bytes that get transmitted -- you could have the right number of bytes and still have a garbled message. So checksums are a simple sanity-check, not a real in-depth verification of the authenticity of some data. It's often a cheap first pass, and can be used to quickly discard obviously invalid things.

## 19. Etl 

We are going to do the `Transform` step of an Extract-Transform-Load.

Extract-Transform-Load (ETL) is a fancy way of saying, "We have some crufty, legacy data over in this system, and now we need it in this shiny new system over here, so we're going to migrate this." (Typically, this is followed by, "We're only going to need to run this once." That's then typically followed by much forehead slapping and moaning about how stupid we could possibly be.)

We're going to extract some scrabble scores from a legacy system. The old system stored a list of letters per score:

* 1 point: `"A", "E", "I", "O", "U", "L", "N", "R", "S", "T",`
* 2 points: `"D", "G",`
* 3 points: `"B", "C", "M", "P",`
* 4 points: `"F", "H", "V", "W", "Y",`
* 5 points: `"K",`
* 8 points: `"J", "X",`
* 10 points: `"Q", "Z",`

The shiny new scrabble system instead stores the score per letter, which makes it much faster and easier to calculate the score for a word. It also stores the letters in lower-case regardless of the case of the input letters:

* `"a"` is worth 1 point.
* `"b"` is worth 3 points.
* `"c"` is worth 3 points.
* `"d"` is worth 2 points.
* Etc.

Your mission, should you choose to accept it, is to write a program that transforms the legacy data format to the shiny new format.

> **NOTE**: A final note about scoring, Scrabble is played around the world in a variety of languages, each with its own unique scoring table. For example, an "A" is scored at 14 in the Basque-language version of the game while being scored at 9 in the Latin-language version.

## 20. Pig Latin 

Implement a program that translates from English to Pig Latin. Pig Latin is a made-up children's language that's intended to be confusing. It obeys a few simple rules (below), but when it's spoken quickly it's really difficult for non-children (and non-native speakers) to understand.

- Rule 1: If a word begins with a vowel sound, add an "ay" sound to the end of the word.
- Rule 2: If a word begins with a consonant sound, move it to the end of the word, and then add an "ay" sound to the end of the word.

There are a few more rules for edge cases, and there are regional variants too. [See](https://en.wikipedia.org/wiki/Pig_latin) for more details.
