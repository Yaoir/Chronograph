BINDIR=/home/jay/.bin

alarm: alarm.go
	@go build alarm.go

#test_alarm:
#	./alarm 5:45pm aplay $bin/audio/cup.wav
#	./alarm 17:45 aplay $bin/audio/cup.wav

test_timer:
	@timer 10s paplay /home/jay/.bin/audio/cup.wav
#	./timer 10s echo done

clock: clock.go
	@go build clock.go

stopwatch: stopwatch.go
	@go build stopwatch.go

timer: timer.go
	@go build timer.go

timeofday: timeofday.go
	@go build timeofday.go

install:
	@cp alarm clock stopwatch timer timeofday $(BINDIR)

backup back bak:
	@cp *.go Makefile .bak
