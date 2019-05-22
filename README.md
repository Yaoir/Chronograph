#### News: May 2019 Update

Providing a command to run is now optional for the __timer__ and __alarm__ programs. This was done because I found __timer__ to be a useful replacement for __sleep(1)__. Instead of just pausing, __timer__ displays the time remaining. For example:

```
for num in 1 2 3
do
	echo $num
	timer 5s
done
```

A few late additions to the documentation:

__timer__ exit codes:
```
1 if quit by user
2 if quit by Control-C
3 if other error
```
This is to make __timer__ more useful in scripts. While the timer is running, you can type either a `q` or `Control-C` to quit, resulting in a different exit code. In the script, inspect the error code and handle the condition appropriately.

For example, a `q` might result in the timer ending early while allowing the script to continue running, but a `Control-C` would cause the script to exit with an error condition:

```
timer 30s
# ignore user quit, but exit if interrupted
if [ $? -eq 2 ]
then
	exit 1
fi
```

The distribution includes shell scripts __alarm-gui__, __clock-gui__, __stopwatch-gui__, and __timer-gui__ as examples of how to run the programs within a virtual terminal so they act more like GUI apps. They all depend on __konsole(1)__ for the virtual terminal. You will need to modify them if you want to use another virtual terminal.

The included __Makefile__ contains rules for installing the programs, __*program*-gui__ shell scripts, and manual pages. Set the variables __BINDIR__, __SSDIR__, and __MANDIR__ in __Makefile__ appropriately before running one of

```
make install
```
and/or
```
make install-man
```

to install the programs and/or manual pages, respectively.


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
$ timer 10s echo done
$ alarm <clock_time> echo done
```

### Manual Pages

```

ALARM(1)                         User Commands                        ALARM(1)

NAME
       alarm - alarm clock

SYNOPSIS
       alarm time [ command ]

DESCRIPTION
       alarm(1)  is  a  command  line alarm clock. It displays the time of day
       until reaching the specified time.

       Before exiting, it optionally runs a command.

ARGUMENTS
       The first argument specifies the time the alarm clock is set to go off.

       The rest of the arguments, if present, are treated as a command to  run
       when the alarm triggers.

EXAMPLES
       An  alarm  set to play an mp3 (with the command aplay alarmbell.mp3) at
       2:30 pm (14:30):

       alarm 2:30pm aplay alarmbell.mp3

       An alarm set to print "hello, world" at 9 pm (21:00):

       alarm 21:00 echo "hello, world"

       Reminder to feed the cat at 10am:

       alarm 10am echo feed the cat

BUGS
       The time is always shown with a 24-hour clock, even when the  alarm  is
       set with a 12-hour clock.

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2019 Jay Ts

       Released   under   the   GNU   Public   License,  version  3.0  (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)
```

```

CLOCK(1)                         User Commands                        CLOCK(1)



NAME
       clock - clock

SYNOPSIS
       clock

DESCRIPTION
       clock(1)  is  a  command-line clock that runs in a virtual terminal. It
       displays the running wall clock time and is accurate to about 1/10 sec‐
       ond.

       Type a Control-C to exit.

ARGUMENTS
       None.

BUGS
       The time is always shown with a 24-hour clock.

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2018 Jay Ts

       Released   under   the   GNU   Public   License,  version  3.0  (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)
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

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2018 Jay Ts

       Released  under  the  GNU   Public   License,   version   3.0   (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)
```

```

TIMER(1)                         User Commands                        TIMER(1)

NAME
       timer - countdown timer

SYNOPSIS
       timer duration [ command ]

DESCRIPTION
       timer(1)  is  a  countdown  timer.  It runs an optional command after a
       specified duration of time.

       While the timer is running, the remaining time is displayed.

ARGUMENTS
       The first argument specifies the duration.

       If additional arguments are supplied, they are treated as a command  to
       run when the timer reaches 0.

EXAMPLES
       After  2  minutes  and  30 seconds, play an mp3 (with the command aplay
       alarmbell.mp3):

       timer 2m30s aplay alarmbell.mp3

       An alarm set to print "hello, world" in 21 minutes:

       alarm 21m echo "hello, world"

BUGS
       The time is always shown with a 24-hour clock, even when  the  duration
       is set with a 12-hour clock.

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2019 Jay Ts

       Released   under   the   GNU   Public   License,  version  3.0  (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)
```
