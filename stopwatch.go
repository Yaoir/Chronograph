package main

//#include <stdio.h>
//#include <termios.h>
//#include <unistd.h>
//
//int tcgetattr(int fd, struct termios *termios_p);
//int tcsetattr(int fd, int optional_actions, const struct termios *termios_p);
//ssize_t read(int fd, void *buf, size_t count);
//
//struct termios t1, t2;
//
//void tty_makeraw(struct termios *termios_p)
//{
//	termios_p->c_iflag &= ~(IGNBRK | BRKINT | PARMRK | ISTRIP | INLCR | IGNCR | ICRNL | IXON);
//	termios_p->c_lflag &= ~(ECHO | ECHONL | ICANON | ISIG | IEXTEN);
//	termios_p->c_cflag &= ~(CSIZE | PARENB);
//	termios_p->c_cflag |= CS8;
//}
//
//void tty_setraw(void)
//{
//
//	/* get stdin settings */
//	tcgetattr(0, &t1);
//	/* copy to t2 */
//	t2 = t1;
//	/* make raw settings */
//	tty_makeraw(&t2);
//	/* set stdin settings to raw */
//	tcsetattr(0, 0, &t2);
//}
//
//void tty_reset(void)
//{
//	/* reset stdin */
//	tcsetattr(0, 0, &t1);
//}
//
//char getbyte(void)
//{
//	char c;
//	read(0,&c,(size_t)1);
//	return c;
//}
import "C"

import (
	"flag"
	"fmt"
	"os"
	"time"
	)

var ticker *time.Ticker
// start_time is the system (clock) time that the counter was started
var start_time time.Time
// segment_time is like a trip odometer. It counts the time between starting and stopping
var segment_time time.Duration
// total_time is like an odometer. It stores all of the time measured
var total_time time.Duration

func quit(code int) {
	C.tty_reset()
	os.Exit(code)
}

var paused bool

// get and process a keypress

func do_key() {
	var code int = 1
	var this string
	// wait for any key to be pressed
	c := byte(C.getbyte())

	switch c {
		case ' ', 'p', 'P':		// Space bar or p
			// pause/restart
			if ! paused {
				stop()
				paused = true
			} else {
				start()
				paused = false
			}
		case 'r', 'R':
			// reset
			if ! paused { return }	// enabled only while paused
			total_time = 0
			segment_time = 0
//			fmt.Printf("\r%s",dur2str(total_time))	// zero display: "00:00:00.0 "
			this = dur2str(total_time)
			eraseprint(len(prev),this)
			prev = this
		case '\r', 'q', 'Q', 'e', 'E', 0x03, 0x04:	// Enter/Return, q, e, Ctrl-C, Ctrl-D
			// stop and exit
			if ! paused { stop() }
			fmt.Printf("\n")
			if c == 0x03 || c == 0x04 { code = 2 }
			quit(code)
		case 'l', 'L':
			// lap time
			fmt.Printf("\n")
		default:	// ignore any other keypress
	}
}

func dur2str(d time.Duration) string {
	var mm, hh int
	var sec float64

	// sec = seconds, as float64
	sec = float64(float64(int64(d)) / 1e9)
	hh = int(sec) / 3600
// if hh > 24, we don't count days
// if hh > 99, just use 3 digits; don't reset to 0
	sec -= float64(hh*3600)
	mm = int(sec) / 60
	sec -= float64(mm*60)
	return fmt.Sprintf("%02d:%02d:%04.1f ",hh,mm,sec)
}

func start() {
	// tick every 1/100 second
	ticker = time.NewTicker(10*time.Millisecond)
	start_time = time.Now()
	go count()
}

var stopped chan bool

func stop() {
	if paused { return }		// (protective, should be unnecessary)
	ticker.Stop()			// stop the ticker
	stopped <- true			// notify counter that ticker has stopped
	total_time += segment_time	// add in the time counted
	segment_time = 0		// reset segment counter
}

// back up over the time string and print a new one over it.

func eraseprint(len int, s string) {
	for i := 0; i < len; i++ {
		fmt.Printf("\b")
	}
	fmt.Printf("%s",s)
}

var prev string

func count() {
	var this string
	for {
		select {
			case <-stopped:
				return	// have this goroutine exit
			case t := <- ticker.C:
				segment_time = t.Sub(start_time)
				this = dur2str(total_time+segment_time)
				// Print only if the string has changed since last time
//				if this != prev { fmt.Printf("\r%s",this) }
				if this != prev { eraseprint(len(prev),this) }
				prev = this
		}
	}
}

func main() {
//
	C.tty_setraw()	// put tty in raw mode (unbuffered)
	stopped = make(chan bool)
	start_paused := flag.Bool("p",false,"start paused")
	flag.Parse()
	paused = *start_paused
//	fmt.Printf("%s",dur2str(total_time))	// initial display: "00:00:00.0 "
	prev = dur2str(total_time)
	fmt.Printf("%s",prev)	// initial display: "hh:mm:ss.d "

	if ! paused { start() }	// start measuring/displaying running time (presses stopwatch's start button)
	// event loop: handle key presses
	for { do_key() }
}
