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
//			fmt.Printf("\n")
			eraseprint(len(prev),"           \r")
			if c == 0x03 || c == 0x04 { code = 2 }
			quit(code)
		default:	// ignore any other keypress
	}
}

func start() {
	// tick every 1/10 second
	ticker = time.NewTicker(100*time.Millisecond)
	go count()
}

var stopped chan bool

func stop() {
	ticker.Stop()			// stop the ticker
	stopped <- true			// notify counter that ticker has stopped
}

// back up over the time string and print a new one over it.

func eraseprint(len int, s string) {
	for i := 0; i < len; i++ {
		fmt.Printf("\b")
	}
	fmt.Printf("%s",s)
}

func time2str(t time.Time) string {
//
	return fmt.Sprintf("%0d:%02d:%02d ", t.Hour(), t.Minute(), t.Second())
}

var prev string

func count() {
//
	var this string

	for {
		select {
			case <-stopped:
				return	// have this goroutine exit
			case t := <- ticker.C:
				this = time2str(t)
				// Print only if the string has changed since last time
				if this != prev { eraseprint(len(prev),this) }
				prev = this
		}
	}
}

func main() {
//
	if len(os.Args) > 1 {
		fmt.Printf("clock: bad arguments\n")
		os.Exit(3)
	}

	C.tty_setraw()	// put tty in raw mode (unbuffered)
	stopped = make(chan bool)

	// Run clock

	prev = time2str(time.Now())
	fmt.Printf("%s",prev)

	start()			// start displaying running time
	for { do_key() }	// event loop: handle key presses
}
