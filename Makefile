# Makefile for Chronograph programs
#
# Written for GNU make, but will probably work
# with any Unix/POSIX-compatible make program
#
# Requirements:
#	Compile (build): You must have Go installed.
#	To make and view the manual page: ronn, gzip, man

# Modify these three settings to fit your needs:

# 1. Where to install the binary programs:
BINDIR=/home/jay/.bin/elf
# Or perhaps one of:
#BINDIR=/usr/local/bin
#BINDIR=/usr/bin

# 2. Where to install the shell scripts:
SSDIR=/home/jay/.bin
# Or perhaps one of:
#SSDIR=/usr/local/bin
#SSDIR=/usr/bin

# 3. Where to install the manual page:
MANDIR=/usr/local/man/man1
# or maybe one of:
#MANDIR=/usr/local/share/man/man1
#MANDIR=/usr/share/man/man1

all: alarm clock stopwatch timer

alarm: alarm.go
	@go build alarm.go

# Make the manual page

alarm-man: alarm.1.ronn
	@ronn --roff --manual="User Commands" --organization="Jay Ts" --date="2018-12-25" alarm.1.ronn > man1/alarm.1
	@gzip -f alarm.1
	@mv alarm.1.gz man1
	@man -l man1/alarm.1.gz

# Display the manual page

showman-alarm:
	@man -l man1/alarm.1.gz

clock: clock.go
	@go build clock.go

stopwatch: stopwatch.go
	@go build stopwatch.go

timer: timer.go
	@go build timer.go

test_timer:
	@timer 10s paplay /home/jay/.bin/audio/cup.wav
#	./timer 10s echo done

timer-man: timer.1.ronn
	@ronn --roff --manual="User Commands" --organization="Jay Ts" --date="2018-12-25" timer.1.ronn > man1/timer.1
	@gzip -f timer.1
	@mv timer.1.gz man1
	@man -l man1/timer.1.gz

# Display the manual page

showman-timer:
	@man -l man1/timer.1.gz

timeofday: timeofday.go
	@go build timeofday.go

install:
	@cp alarm clock stopwatch timer $(BINDIR)

install-man:
	cp man1/*.1.gz $(MANDIR)

backup back bak:
	@cp *.go *.ronn *-win README.md Makefile TODO .bak
