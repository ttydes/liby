package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	account "github.com/ttydes/liby/services/account"
	branch "github.com/ttydes/liby/services/branch"
	"google.golang.org/grpc"
)

type server interface {
	Run(int) error
}

func main() {
	var (
		port         = flag.Int("port", 8080, "Serve all on one port")
		account_port = flag.String("account_port", "account:8001", "Account service port")
		branch_port  = flag.String("branch_port", "search:8002", "Branch service port")
	)
	flag.Parse()

	t, err := trace.New("search", *jaegeraddr)
	if err != nil {
		log.Fatalf("trace new error: %v", err)
	}

	var srv server
	var cmd = os.Args[1]

	switch cmd {
	case "geo":
		srv = geosrv.New(t)
	case "rate":
		srv = ratesrv.New(t)
	case "profile":
		srv = profilesrv.New(t)
	case "search":
		srv = searchsrv.New(
			t,
			dial(*geoaddr, t),
			dial(*rateaddr, t),
		)
	case "frontend":
		srv = frontendsrv.New(
			t,
			dial(*searchaddr, t),
			dial(*profileaddr, t),
		)
	default:
		log.Fatalf("unknown cmd: %s", cmd)
	}

	if err := srv.Run(*port); err != nil {
		log.Fatalf("run %s error: %v", cmd, err)
	}
}

func dial(addr string, t opentracing.Tracer) *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(t)),
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		panic(fmt.Sprintf("ERROR: dial error: %v", err))
	}

	return conn
}
