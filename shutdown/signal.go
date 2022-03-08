package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

// ListenInterrupt listens for OS interrupt signals and calls f on receive.
// Calling ListenInterrupt blocks until a signal is received and thus must
// called inside a goroutine separated from the main program.
func ListenInterrupt(f func()) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-sig
	f()
}
