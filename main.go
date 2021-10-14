package main

import (
	"net"
	"net/http"
	"os"

	"github.com/nikandfor/cli"
	"github.com/nikandfor/errors"
	"github.com/nikandfor/tlog"
)

func main() {
	cli.App = cli.Command{
		Name:   "simple HTTP local server",
		Action: run,
		Flags: []*cli.Flag{
			cli.NewFlag("listen,l", ":8000", "listen address"),
			cli.NewFlag("path,p", "./", "dir to serve"),
		},
	}

	cli.RunAndExit(os.Args)
}

func run(c *cli.Command) (err error) {
	dir := http.Dir(c.String("path"))

	l, err := net.Listen("tcp", c.String("listen"))
	if err != nil {
		return errors.Wrap(err, "listen")
	}

	tlog.Printw("listening", "listen", l.Addr(), "path", dir)

	fs := http.FileServer(dir)

	err = http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		tlog.Printw("request", "method", req.Method, "uri", req.RequestURI, "remote_addr", req.RemoteAddr)

		fs.ServeHTTP(w, req)
	}))

	return errors.Wrap(err, "serve")
}
