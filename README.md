# socketio

A copy of github.com/googollee/go-socket.io with defects fixed and updated to 1.4.5 of socket.io.
Also works with 1.7.4 and 2.0.3 of Socket.IO client Javascript.

socketio is an implementation of [socket.io](http://socket.io) in Go (golang).
This provides the ability to perform real time communication from a browser
to a server and back.  Content can be pushed from the server out to
the browser.

It is compatible with the 1.4.5 version of socket.io in Node.js, and supports room and namespace.
This version will be updated on a regular basis with the latest version of socket.io.

## Install

Install the package with:

```bash
	go get github.com/SnowpMakes/socketio
	go get github.com/pschlump/MiscLib
```

Import it with:

```go
	import "github.com/SnowpMakes/socketio"
```

## Example

Please check the ./examples/chat directory for more comprehensive examples.

```go
	package main

	//
	// Command line arguments can be used to set the IP address that is listened to and the port.
	//
	// $ ./chat --port=8080 --host=127.0.0.1
	//
	// Bring up a pair of browsers and chat between them.
	//

	import (
		"flag"
		"fmt"
		"log"
		"net/http"
		"os"

		"github.com/pschlump/MiscLib"
		"github.com/SnowpMakes/socketio"
	)

	var Port = flag.String("port", "9000", "Port to listen to")                           // 0
	var HostIP = flag.String("host", "localhost", "Host name or IP address to listen on") // 1
	var Dir = flag.String("dir", "./asset", "Direcotry where files are served from")      // 1
	func init() {
		flag.StringVar(Port, "P", "9000", "Port to listen to")                           // 0
		flag.StringVar(HostIP, "H", "localhost", "Host name or IP address to listen on") // 1
		flag.StringVar(Dir, "d", "./asset", "Direcotry where files are served from")     // 1
	}

	func main() {

		flag.Parse()
		fns := flag.Args()

		if len(fns) != 0 {
			fmt.Printf("Usage: Invalid arguments supplied, %s\n", fns)
			os.Exit(1)
		}

		var host_ip string = ""
		if *HostIP != "localhost" {
			host_ip = *HostIP
		}

		// Make certain that the command line parameters are handled correctly
		// fmt.Printf("host_ip >%s< HostIP >%s< Port >%s<\n", host_ip, *HostIP, *Port)

		server, err := socketio.NewServer(nil)
		if err != nil {
			log.Fatal(err)
		}

		server.On("connection", func(so socketio.Socket) {
			fmt.Printf("%sa user connected%s, %s\n", MiscLib.ColorGreen, MiscLib.ColorReset, godebug.LF())
			so.Join("chat")
			so.On("chat message", func(msg string) {
				fmt.Printf("%schat message, %s%s, %s\n", MiscLib.ColorGreen, msg, MiscLib.ColorReset, godebug.LF())
				so.BroadcastTo("chat", "chat message", msg)
			})
			so.On("disconnect", func() {
				fmt.Printf("%suser disconnect%s, %s\n", MiscLib.ColorYellow, MiscLib.ColorReset, godebug.LF())
			})
		})

		server.On("error", func(so socketio.Socket, err error) {
			fmt.Printf("Error: %s, %s\n", err, godebug.LF())
		})

		http.Handle("/socket.io/", server)
		http.Handle("/", http.FileServer(http.Dir(*Dir)))
		fmt.Printf("Serving on port %s, brows to http://localhost:%s/\n", *Port, *Port)
		listen := fmt.Sprintf("%s:%s", host_ip, *Port)
		log.Fatal(http.ListenAndServe(listen, nil))
	}
```

## License

The 3-clause BSD License  - see LICENSE for more details

## History

This code is from an original https://github.com/googollee/go-socket.io .  The following 
have been made:

1. Renamed the package so that the directory structure matches with the package name.
Some outside tools depend on this.
2. Included go-engine.io as a subdirectory.
3. Fixed defect #68 - "Not Thread Safe".  All accesses to maps are synchronized.
4. Documentation improvements.
5. Updated to use the current version of socket.io (1.3.6)
6. Provided a packed(uglified) version of the JavaScript scoket.io library. A non-uglified version is in the 
same directory also.
7. Original defect #95 - Crash occurs when too many arguments are passed - suggested fix used and tested.
8. Fixed a set of continuous connect/disconnect problems
9. #45 - incorrect usage - see correct usage in test/o45 - fixed.
10. #47 - crashing on Windows - unable to reproduce with go 1.3.1 on windows 8.  Appears to be fixed by changes for #68.
11. #83 - see example in test/o83 - fixed.
12. #82 - see example in test/o82 - fixed.
13. #52 - see example in test/o52 - fixed.
14. #67 - see example in test/o67 - fixed.
15. Identified the problem where a emit is sent from client to server and server seems to discard/ignore the emit.  This is caused by an invalid paramter and an ignored error message.  Code review for all discarded/ignored error messages in progress.

This code is from an original https://github.com/mlsquires/socket.io. The following changes have been made:
1. Changed a timeout value to account for slow connections.
2. Changed/removed debug messages/formatting to just use log.

## FAQ

1. Why is this not a fork of the original?  A: I can't figure out how to make a fork and change the
name of the package on github.com.   Since a variety of outside tools hurl over the "-" and ".io" in
the directory name I just made a copy and started at the beginning.   My apologies to anybody
that feels offended by this approach.  


