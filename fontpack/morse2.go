// Package fontpack Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at 2023-08-16 12:53:52.126745 +1000 AEST m=+1.232865667
// using data from https://github.com/xero/figlet-fonts
package fontpack

const Morse2 = `flf2a$ 1 1 30 -1 40
Morse by Glenn Chappell <ggc@uiuc.edu> 10/95
Based on info from "Morse Code and the Phonetic Alphabets"
URL: http://www.soton.ac.uk/~scp93ch/refer/alphabet.html
Includes some ISO Latin-1 characters
Permission is hereby given to modify this font, as long as the
modifier's name is placed on a comment line.

Use of this font allows figlet to convert ASCII to International Morse
Code.

The following substitutions have been made, i.e., if figlet receives a
character on the left as input, it will print the Morse Code for the
character on the right:
!               -> .
;               -> :
[]{}            -> ()
German s-z      -> ss
Latin-1 "prime" -> '
#               -> <Delete last word>
*               -> Ch

Except as noted above, characters not available in Morse Code will print
as "?".

Inter-word spaces have been enlarged to improve readability. To conform
to the standards for automatic transmission of Morse Code, use smushmode
zero ("figlet -f morse -m0"); wait the duration of a dot for each blank
in the output, as well as between any two output characters (dots,
dashes or blanks).

Explanation of first line:
flf2 - "magic number" for file identification
a    - should always be ` + "`" + `a', for now
$    - the "hardblank" -- prints as a blank, but can't be smushed
1    - height of a character
1    - height of a character, not including descenders
30   - max line length (excluding comment lines) + a fudge factor
-1   - default smushmode for this font
40   - number of comment lines

  $@
.-.-.-$@
.-..-.$@
........$@
?$#
?$#
?$#
.----.$@
-.--.-$@
-.--.-$@
----$@
?$#
--..--$@
-....-$@
.-.-.-$@
-..-.$@
-----$@
.----$@
..---$@
...--$@
....-$@
.....$@
-....$@
--...$@
---..$@
----.$@
---...$@
---...$@
?$#
?$#
?$#
..--..$@
?$#
.-$@
-...$@
-.-.$@
-..$@
.$@
..-.$@
--.$@
....$@
..$@
.---$@
-.-$@
.-..$@
--$@
-.$@
---$@
.--.$@
--.-$@
.-.$@
...$@
-$@
..-$@
...-$@
.--$@
-..-$@
-.--$@
--..$@
-.--.-$@
?$#
-.--.-$@
?$#
?$#
?$#
.-$@
-...$@
-.-.$@
-..$@
.$@
..-.$@
--.$@
....$@
..$@
.---$@
-.-$@
.-..$@
--$@
-.$@
---$@
.--.$@
--.-$@
.-.$@
...$@
-$@
..-$@
...-$@
.--$@
-..-$@
-.--$@
--..$@
-.--.-$@
?$#
-.--.-$@
?$#
.-.-$@
---.$@
..--$@
.-.-$@
---.$@
..--$@
...$...$@
0
?$#
160
  $@
171
.-..-.$@
180
.----.$@
187
.-..-.$@
188
.----$-..-.$....-$@
189
.----$-..-.$..---$@
190
...--$-..-.$....-$@
193
.--.-$@
196
.-.-$@
197
.--.-$@
201
..-..$@
209
--.--$@
214
---.$@
220
..--$@
223
...$...$@
225
.--.-$@
228
.-.-$@
229
.--.-$@
233
..-..$@
241
--.--$@
246
---.$@
252
..--$@
`
