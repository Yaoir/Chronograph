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
	"fmt"
	"os"
	"time"
	)

var final_time time.Time
var duration time.Duration
var ticker *time.Ticker

func quit(code int) {
	C.tty_reset()
	os.Exit(code)
}

// get and process a keypress

func do_key() {
	var code int = 1
	// wait for any key to be pressed
	c := byte(C.getbyte())

	switch c {
		case 'q', 'Q', 'e', 'E', 0x03, 0x04:	// q, e, Ctrl-C, Ctrl-D
			// stop and exit
			stop()
			fmt.Printf("\n")
			if c == 0x03 || c == 0x04 { code = 2 }
			quit(code)
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
// if hh > 99, use 3 digits; don't reset to 0
	sec -= float64(hh*3600)
	mm = int(sec) / 60
	sec -= float64(mm*60)
	return fmt.Sprintf("%02d:%02d:%04.1f ",hh,mm,sec)
}

func start() {
	// tick every 1/100 second
	ticker = time.NewTicker(10*time.Millisecond)
	go count()
}

var stopped chan bool

func stop() {
	ticker.Stop()			// stop the ticker
	stopped <- true			// notify counter that ticker has stopped
}

// back up over the time string and print a new one over it.

func eraseprint(len int, s string) {
	for i := 0; i < len; i++ { fmt.Printf("\b") }
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
				rem := final_time.Sub(t)
				this = dur2str(rem)
				// Print only if the string has changed since last time
				if this != prev && rem.Seconds() >= 0.0 { eraseprint(len(prev),this) }
				prev = this

				if time.Now().After(final_time) {
					// erase time display
					eraseprint(len(prev),"           \r")
					quit(0)
				}
		}
	}
}

func main() {
	var err error

	if len(os.Args) != 2 {
		fmt.Printf("usage: timer <duration>\n")
		os.Exit(3)
	}

	// parse countdown time

	duration, err = time.ParseDuration(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr,"timer: bad duration %q\n",os.Args[1])
		os.Exit(3)
	}

	final_time = time.Now().Add(duration)

	C.tty_setraw()	// put tty in raw mode (unbuffered)
	stopped = make(chan bool)

	// Run countdown timer

	prev = dur2str(duration)
	fmt.Printf("%s",prev)	// initial display: "hh:mm:ss.d "

	start()			// start measuring/displaying running time
	for { do_key() }	// event loop: handle key presses
}
