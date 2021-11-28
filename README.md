# howfast
[![Go Report Card](https://goreportcard.com/badge/github.com/Chris-Behan/howfast)](https://goreportcard.com/report/github.com/Chris-Behan/howfast)


Minimalist Go library for timing your code.

Usage:
```go
timer := howfast.Timer{}

timer.Start()
functionYouWantToTime()
timer.Stop()

seconds := timer.SecondsElapsed()
milliSeconds := tiimer.MilliSecondsElapsed()
microSeconds := timer.MicroSecondsElapsed()
nanoSeconds := timer.NanoSecondsElapsed()
```