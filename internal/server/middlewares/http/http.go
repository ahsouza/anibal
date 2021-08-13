package ani

import (
	"fmt"

	"github.com/soheilhy/args"
)

var ToListen = args.NewInt(args.Default(4200))
var BufferSize = args.NewUint64(args.Default(uint64(1024 * 1024)))
var StateDir = args.NewString(args.Flag("test.state.dir", "/tmp", "state dir"))

//var RoundTripper = args.New(Default(http.DefaultTransport))
//var Timeout = args.NewDuration(Flag("timeout", 10*time.Second, "timeout"))

func Server(opts ...args.V) {
	port := ToListen.Get(opts)
	//bufs := BufferSize.Get(opts)
	//sdir := StateDir.Get(opts)
	//rt := RoundTripper.Get(args)
	//to := Timeout.Get(args)

	fmt.Printf("Server listening on the address: http://localhost:%d\n", port)
}
