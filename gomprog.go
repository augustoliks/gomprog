package gomprog

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	for s := range signals {
		switch s {
		case syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt:
			os.Exit(1)
		}
	}
}

func ConfirmReceivedLogToRsyslog() {
	fmt.Println("OK")
}

// func main() {
// 	// go handleSignals()

// 	in := bufio.NewReader(os.Stdin)

// 	for {
// 		log_line, err := in.ReadString('\n')
// 		log_line = strings.TrimSuffix(log_line, "\n")

// 		confirmReceivedLogToRsyslog()

// 		if log_line == "" || err == io.EOF {
// 			break
// 		}

// 	}
// 	os.Exit(0)
// }
