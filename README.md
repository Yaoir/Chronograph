TODO: Linux/POSIX, due to Cgo

TODO: Using the programs as GUI apps

The distribution includes shell scripts __alarm-gui__, __clock-gui__, __stopwatch-gui__, and __timer-gui__ as examples of how to run the programs within a virtual terminal so they act more like GUI apps. They all depend on __konsole(1)__ for the virtual terminal. You will need to modify them if you want to use another virtual terminal.

TODO: Makefile functions - installation

The included __Makefile__ contains rules for installing the programs, __*program*-gui__ shell scripts, and manual pages. Set the variables __BINDIR__, __SSDIR__, and __MANDIR__ in __Makefile__ appropriately before running one of

```
make install
```
and/or
```
make install-ss
```
and/or
```
make install-man
```

to install the programs, shell scripts, and/or manual pages, respectively.
==============================================================================

### Introduction

Chronograph is a group of command line programs that implement functions found in a chronograph watch.

The programs are:

```
alarm - alarm clock
clock - wall clock
stopwatch - stopwatch
timer - countdown timer
```

### Quick Start

In the example commands, `$` is used to indicate a shell prompt.

Compile:

```
$ go build alarm.go
$ go build clock.go
$ go build stopwatch.go
$ go build timer.go
```
or if you have GNU make installed:
```
$ make
```

Run:

```
$ clock
$ stopwatch
$ timer 10s
$ alarm <clock_time>
```

### Manual Pages

```
ALARM(1)                         User Commands                        ALARM(1)



NAME
       alarm - alarm clock

SYNOPSIS
       alarm time

DESCRIPTION
       alarm(1)  is  a  command  line alarm clock. It displays the time of day
       until reaching the specified time.

ARGUMENTS
       The argument specifies the time that alarm exits. The following formats
       may be used:

           Exactly on the hour:

           3pm
           3PM
           15

           Exactly on the minute:

           3:04pm
           3:04PM
           15:04

           Exactly on the second:

           3:04:05pm
           3:04:05PM
           15:04:05

EXAMPLES
       An  alarm  set to play an mp3 (with the command aplay alarmbell.mp3) at
       2:30 pm (14:30):

       alarm 2:30pm && aplay alarmbell.mp3

       An alarm set to print "hello, world" at 9 pm (21:00):

       alarm 21:00 && echo "hello, world"

       Reminder to feed the cat at 10am:

       alarm 10am && echo feed the cat

EXIT VALUES
       0 Normal command exit (reached alarm time)
       1 Quit by user (with q or Q key)
       2 Quit by Control-C or SIGKILL
       3 Error (bad arguments)

       When using alarm in loop in a shell script, use this code to  have  the
       script  exit  when alarm is interrupted, either with a q or Q keypress,
       or a Control-C or SIGKILL from another source:

           alarm 1pm
           # exit if interrupted
           if [ $? -ne 0 ]
           then
               exit 1
           fi

       The following causes the shell script to exit if it receives an  inter‐
       rupt  signal  (Control-C),  but  if the user types a q or Q, alarm will
       quit early, and the loop inside the shell script will continue to  exe‐
       cute:

           alarm 2pm
           # ignore user quit with q or Q key,
           # but exit if interrupted with SIGKILL (Control-C)
           if [ $? -eq 2 ]
           then
               exit 1
           fi

BUGS
       The  time  is always shown with a 24-hour clock, even when the alarm is
       set with a 12-hour clock.

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2019 Jay Ts

       Released  under  the  GNU   Public   License,   version   3.0   (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)



Jay Ts                             June 2019                          ALARM(1)
```

```
CLOCK(1)                         User Commands                        CLOCK(1)



NAME
       clock - clock

SYNOPSIS
       clock

DESCRIPTION
       clock(1)  is  a  command-line clock. It displays the running wall clock
       time and is accurate to about 1/10 second.

       Type a ´q´, ´Q´, or Control-C to exit.

ARGUMENTS
       None.

EXIT VALUES
       1 Quit by user (with q key)
       2 Quit by Control-C or SIGKILL
       3 Error (bad arguments)

BUGS
       The time is always shown with a 24-hour clock.

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2019 Jay Ts

       Released  under  the  GNU   Public   License,   version   3.0   (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)



Jay Ts                             June 2019                          CLOCK(1)
```

```
STOPWATCH(1)                     User Commands                    STOPWATCH(1)



NAME
       stopwatch - stopwatch

SYNOPSIS
       stopwatch [-p]

DESCRIPTION
       stopwatch(1) is a stopwatch that runs in a virtual terminal.

       It is controlled with keys on the keyboard as follows:

       SPACE, p, P

           Pause/restart the stopwatch.

       l, L

           Lap Timer. The current timing is printed, and counting continues on the following line.

       r, R

           Reset. Works only while paused.

       q, Q, e, E, Ctrl-C, Ctrl-D, Enter/Return

           Stop the stopwatch and exit.

ARGUMENTS
       When  the  -p  option  is specified, the stopwatch starts in the paused
       state.

EXIT VALUES
       1 Quit by user (with q key)
       2 Quit by Control-C or SIGKILL
       3 Error (bad arguments)

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2019 Jay Ts

       Released  under  the  GNU   Public   License,   version   3.0   (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)



Jay Ts                             June 2019                      STOPWATCH(1)
```

```
TIMER(1)                         User Commands                        TIMER(1)



NAME
       timer - countdown timer

SYNOPSIS
       timer duration

DESCRIPTION
       timer(1) is a countdown timer.

       While  the  timer is running, the remaining time is displayed. The com‐
       mand exits when the time remaining reaches 0.

ARGUMENTS
       The argument specifies the duration. Some examples of  valid  arguments
       follows:

           5s        Five seconds
           22.9s     Twenty-two and 9/10 seconds
           179s      179 seconds (or two minutes and 59 seconds)
           7m        Seven minutes
           32m30s    32½ minutes
           1h45m     One hour and 45 minutes
           2h56m4s   Two hours, 56 minutes, and 4 seconds
           5h        Five hours
           100h      One hundred hours

EXAMPLES
       Pause for 5 seconds, showing the countdown:

       timer 5s

       After  2  minutes  and  30 seconds, play an mp3 (with the command aplay
       alarmbell.mp3):

       timer 2m30s && aplay alarmbell.mp3

       Print "hello, world" after waiting 21 minutes:

       timer 21m && echo "hello, world"

EXIT VALUES
       0 Normal command exit (timer reached 0)
       1 Quit by user (with q key)
       2 Quit by Control-C or SIGKILL
       3 Error (bad arguments)

       When using timer in loop in a shell script, use this code to  have  the
       script  exit  when  the timer is interrupted, either with a q or Q key‐
       press, or a Control-C or SIGKILL from another source:

           timer 600s
           # exit if interrupted
           if [ $? -ne 0 ]
           then
               exit 1
           fi

       The following causes the shell script to exit if it receives an  inter‐
       rupt  signal  (Control-C),  but  if the user types a q or Q, timer will
       quit early, and the loop inside the shell script will continue to  exe‐
       cute:

           timer 600s
           # ignore user quit with q or Q key,
           # but exit if interrupted with SIGKILL (Control-C)
           if [ $? -eq 2 ]
           then
               exit 1
           fi

BUGS
       The  time  is always shown with a 24-hour clock, even when the duration
       is set with a 12-hour clock.

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2019 Jay Ts

       Released  under  the  GNU   Public   License,   version   3.0   (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)


Jay Ts                             June 2019                          TIMER(1)
```
