CMatrix(1)                  General Commands Manual                 CMatrix(1)

NAME
       CMatrix - simulates the display from "The Matrix"

SYNOPSIS
       cmatrix [-abBflohnsmVx] [-u update] [-C color]

DESCRIPTION
       Shows a scrolling 'Matrix' like screen in Linux

   OPTIONS
       -a     Asynchronous scroll

       -b     Bold characters on

       -B     All bold characters (overrides -b)

       -f     Force the linux $TERM type to be on

       -l     Linux mode (sets "matrix.fnt" font in console)

       -o     Use old-style scrolling

       -h, -? Print usage and exit

       -n     No bold characters (overrides -b and -B)

       -s     "Screensaver" mode, exits on first keystroke

       -m     "Lambda" mode, every character becomes a lambda

       -x     X window mode, use if your xterm is using mtx.pcf

       -V     Print version information and exit

       -u delay
              Screen update delay 0 - 9, default 4

       -C color
              Use  this  color  for  matrix (default green).  Valid colors are
              green, red, blue, white, yellow, cyan, magenta and black.

   KEYSTROKES
       The following keystrokes are available during execution (unavailable in
       -s mode)

       a      Toggle asynchronous scroll

       b      Random bold characters

       B      All bold characters

       n      Turn off bold characters

       0-9    Adjust update speed

       ! @ # $ % ^ & )
              Change  the  color of the matrix to the corresponding color: ! -
              red, @ - green, # - yellow, $ - blue, % - magenta, ^ - cyan, & -
              white, ) - black.

       q      Quit the program

BUGS
       This  program  is very CPU intensive.  Don't be surprised if it eats up
       over 40% of your CPU at times.

HOMEPAGE
       http://www.github.com/abishekvashok/cmatrix

                                Mon May 3 1999                      CMatrix(1)
