package auth_test

// import (
// 	"io"
// 	"log"
// 	"os"
// 	"os/exec"
// 	"strings"
// 	"sync"
// 	"syscall"
// 	"testing"
// )

// const FirestoreEmulatorHost = "FIRESTORE_EMULATOR_HOST"

// func TestMain(m *testing.M) {
// 	// Start the firestore emulator before running the tests
// 	cmd := exec.Command("firebase", "emulators:start", "--only", "firestore")

// 	// Make the emulator task killable
// 	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

// 	// Capture the stderr output to know when emulators start
// 	stdout, err := cmd.StdoutPipe()

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stdout.Close()

// 	// Start the emulator
// 	if err := cmd.Start(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Ensure the process is killed when testing is finished,
// 	// even if an error occurs
// 	var result int
//     defer func() {
//         syscall.Kill(-cmd.Process.Pid, syscall.SIGINT)
//         os.Exit(result)
//     }()

// 	// Wait until it's running to start
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	// Start a separate go routine
// 	go func() {
// 		// reading it's output
// 		buf := make([]byte, 512, 512)
// 		for {
// 			n, err := stdout.Read(buf[:])
// 			if err != nil {
// 				// until it ends
// 				if err == io.EOF {
// 					break
// 				}
// 				log.Fatalf("reading stderr %v", err)
// 			}

// 			if n > 0 {
// 				d := string(buf[:n])

// 				// only required if we want to see the emulator output
// 				log.Printf("%s", d)

// 				// Checking if the app server is started
// 				if strings.Contains(d, "All emulators ready!") {
// 					wg.Done()
// 				}

// 				// and capturing the FIRESTORE_EMULATOR_HOST value to set
// 				pos := strings.Index(d, FirestoreEmulatorHost+"=")
// 				if pos > 0 {
// 					host := d[pos+len(FirestoreEmulatorHost)+1 : len(d)-1]
// 					os.Setenv(FirestoreEmulatorHost, host)
// 				}
// 			}
// 		}
// 	}()

// 	// wait until the running message has been received
// 	wg.Wait()

// 	result = m.Run()
// }
