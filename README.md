# Advanced Timer
Advanced Timer is a Go package that provides a multi-thread safe wrapper around the time.Timer struct with a few extra features. It allows you to start, pause, resume, and stop a timer, as well as get the remaining time, maximum duration, and start time.

# Features
- Start() starts the timer.
- Pause() pauses the timer.
- Resume() resumes the timer.
- Stop() stops the timer.
- Finished channel is closed when the timer is stopped.
- Remaining is the remaining time on the timer.
- MaxDuration is the maximum duration of the timer.
- StartedAt is the time the timer was started.
- Paused is a boolean indicating whether the timer is paused or not.

# Limitations
- The remaining time has a margin of error of around 30 milliseconds.
- The finished event is only guaranteed to be fired within 1 nanosecond of the timer being stopped. (lowest allowed sleep time IN GO)

# Usage
## Create a new timer
```go
maxDuration := 10 * time.Second
timer := NewAdvancedTimer(maxDuration)
```
## Start the timer
```go
timer.Start()
```
## Pause the timer
```go
timer.Pause()
```
## Resume the timer
```go
timer.Resume()
```
## Stop the timer
```go
timer.Stop()
```
## Get the remaining time
```go
remaining := timer.Remaining
```
## Get the maximum duration
```go
maxDuration := timer.MaxDuration
```

## Check if the timer is paused
```go
paused := timer.Paused
```
## Wait for the timer to finish
```go
<-timer.Finished
```
## Get the time the timer was started
```go
startedAt := timer.StartedAt
```
## Get a string representation of the timer
```go
str := timer.Stringfy()
```

# Example
```go
package main

import (
	"fmt"
	"time"

	"github.com/user/advanced_timer"
)

func main() {
	maxDuration := 5 * time.Second
	timer := advanced_timer.NewAdvancedTimer(maxDuration)

	timer.Start()

	fmt.Println(timer.Stringfy())

	time.Sleep(2 * time.Second)

	timer.Pause()

	fmt.Println(timer.Stringfy())

	time.Sleep(2 * time.Second)

	timer.Resume()

	fmt.Println(timer.Stringfy())

	<-timer.Finished

	fmt.Println(timer.Stringfy())
}

```

# buy me a coffee
Thank you for using the Advanced Timer Go package! If you found it useful and would like to support further development, you can [buy me a coffee](https://www.buymeacoffee.com/7dwkkxwrgtv) using the link below. Your support will help keep this package up-to-date and ensure that it continues to be a reliable solution for your Go timer needs. Thank you for your support!

[![Buy me a coffee](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/7dwkkxwrgtv)

[My GitHub](https://www.github.com/billnice250)
