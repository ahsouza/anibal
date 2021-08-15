package ani

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
	"github.com/soheilhy/args"
)

var ToListen = args.NewString(args.Default(":4200"))
var SetPort = args.NewInt(args.Default(8080))
var Framework = args.NewString(args.Default("ani"))
var BufferSize = args.NewUint64(args.Default(uint64(1024 * 1024)))
var StateDir = args.NewString(args.Flag("test.state.dir", "/tmp", "state dir"))

//var RoundTripper = args.New(Default(http.DefaultTransport))
//var Timeout = args.NewDuration(Flag("timeout", 10*time.Second, "timeout"))

func Default(opts ...args.V) {
	port := SetPort.Get(opts)
	//bufs := BufferSize.Get(opts)
	//sdir := StateDir.Get(opts)
	//rt := RoundTripper.Get(args)
	//to := Timeout.Get(args)

	fmt.Printf("Aníbal Barca Server listening on the address: http://localhost:%d\n", port)
}

func NewServer(opts ...args.V) {
	port := ToListen.Get(opts)
	framework := Framework.Get(opts)

	if framework == "fiber" {
		fmt.Printf("Fiber Framework")
		fi := fiber.New()
		fi.Listen(port)
	}
	if framework == "echo" {
		fmt.Printf("ECHO V4 Framework")
		//e := echo.new()
		//e.Start(port)
	}
	if framework == "gorilla" {
		fmt.Printf("Gorilla Mux Framework")
		gomx := mux.NewRouter()

		http.ListenAndServe(port, gomx)
	}
	if framework == "gin" {
		fmt.Printf("Gin Framework")
		g := gin.Default()
		g.Run(port)
	}
}
