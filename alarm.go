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
			fmt.Printf("\n")
			if c == 0x03 || c == 0x04 { code = 2 }
			quit(code)
		default:	// ignore any other keypress
	}
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

func time2str(t time.Time) string {
//
	return fmt.Sprintf("%0d:%02d:%02d ", t.Hour(), t.Minute(), t.Second())
}

var prev string

func count() {
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

				if time.Now().After(alarm_time) {
				//
					// erase time display
					eraseprint(len(prev),"           \r")
					quit(0)
				}
		}
	}
}

var alarm_time time.Time	// full time/date for alarm

func main() {
	var err error
	var now time.Time		// current time, from system clock 
	var input_time time.Time	// alarm setting, from argument
	var timespec string		// alarm setting, possibly with an added "m" or "M"

	var time_formats []string = []string{
		"15",
		"15:04",
		"15:04:05",
		"3pm",
		"3PM",
		"3:04pm",
		"3:04PM",
		"3:04:05pm",
		"3:04:05PM",
		}
		
	if len(os.Args) != 2 {
		fmt.Printf("usage: alarm <time>\n")
		os.Exit(3)
	}

	now = time.Now()

	// take these from the current time:

	year := now.Year()
	month := now.Month()
	day := now.Day()
	nsec := 0
	loc := now.Location()

	// If the alarm setting ends with "p" or "a", add "m" to get "am" or "pm"
	// This allows "5:45p" to be written for "5:45pm"

	timespec = os.Args[1]
	if timespec[len(timespec)-1] == 'a' || timespec[len(timespec)-1] == 'p' { timespec += "m" }
	if timespec[len(timespec)-1] == 'A' || timespec[len(timespec)-1] == 'P' { timespec += "M" }

	// parse alarm setting from the argument into input_time

	for _, format := range time_formats {
		input_time, err = time.Parse(format,timespec)
		if err == nil { break }
	}

	if err != nil {
		fmt.Fprintf(os.Stderr,"alarm: bad alarm setting %q\n",os.Args[1])
		os.Exit(3)
	}

	// and take these from the argument:

	hour := input_time.Hour()
	min := input_time.Minute()
	sec := input_time.Second()

	alarm_time = time.Date(year,month,day,hour,min,sec,nsec,loc)

	if now.After(alarm_time) {
	//
		fmt.Fprintf(os.Stderr,"The specified time has already passed\n")
		os.Exit(3)
	}

	C.tty_setraw()	// put tty in raw mode (unbuffered)
	stopped = make(chan bool)

	// Run alarm clock

	prev = time2str(time.Now())
	fmt.Printf("%s",prev)	// initial display: "hh:mm:ss.d "

	start()			// start measuring/displaying running time
	for { do_key() }	// event loop: handle key presses
}
