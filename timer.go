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

// Execute a command specified by the arguments.
// For example:  exec echo hello world

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
	)

var final_time time.Time
var duration time.Duration
var ticker *time.Ticker

func quit() {
	C.tty_reset()
	os.Exit(0)
}

// get and process a keypress

func do_key() {
	// wait for any key to be pressed
	c := byte(C.getbyte())

	switch c {
		case 'q', 'Q', 'e', 'E', 0x03, 0x04:	// q, e, Ctrl-C, Ctrl-D
			// stop and exit
			stop()
			fmt.Printf("\n")
			quit()
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

var command_args []string
var execfile string

func count() {
//
	var this, prev string
	var out bytes.Buffer

	for {
		select {
			case <-stopped:
				return	// have this goroutine exit
			case t := <- ticker.C:
				rem := final_time.Sub(t)
				this = dur2str(rem)
				// Print only if the string has changed since last time
				if this != prev && rem.Seconds() >= 0.0 { fmt.Printf("\r%s",this) }
				prev = this

				if time.Now().After(final_time) {
				//
					cmd := exec.Command(execfile,command_args...)
					cmd.Stdout = &out
					err := cmd.Run()
					if err != nil { fmt.Printf("err = %v\n",err) }
					fmt.Printf("\r           \r")
					fmt.Printf("%s", out.String())
					quit()
				}
		}
	}
}

func main() {
//
	var err error

	if len(os.Args) < 3 {
	//
		fmt.Printf("timer: need arguments\n")
		os.Exit(2)
	}

	// parse countdown time

	duration, err = time.ParseDuration(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr,"timer: bad duration %q\n",os.Args[1])
		os.Exit(2)
	}

	final_time = time.Now().Add(duration)

	C.tty_setraw()	// put tty in raw mode (unbuffered)
	stopped = make(chan bool)

	execfile = os.Args[2]	// program to execute

	// check that execfile exists in $PATH

	_, err = exec.LookPath(execfile)
	if err != nil {
		fmt.Fprintf(os.Stderr,"Cannot find command %s\n",execfile)
		quit()
	}

	// prepare argument list

	for i := 3; i < len(os.Args); i++ { command_args = append(command_args,os.Args[i]) }

	// Run countdown timer

	fmt.Printf("%s",dur2str(duration))	// initial display: "hh:mm:ss.d "

	start()			// start measuring/displaying running time
	for { do_key() }	// event loop: handle key presses
}
