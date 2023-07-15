package main

import (
	"context"
	"io"
	"log"
	"net"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()

	log.Println("setting up ngrok tunnel, this may take a moment.")

	tun, err := ngrok.Listen(ctx,
		config.TCPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)

	if err != nil {
		panic(err)
	}

	log.Println("your tunnel is up, you can access your server through the following IP:", tun.URL()[6:])

	for {
		conn, err := tun.Accept()

		if err != nil {
			log.Printf("failed to accept connection from %s. err:", err)
			continue
		}

		log.Println("accepted connection from", conn.RemoteAddr())

		go func() {
			err := handleConn(ctx, "minecraft-server.railway.internal:25565", conn)
			if err != nil {
				log.Printf("connection closed with error from %s. ERR: %d\n", conn.RemoteAddr(), err)
				return
			}

			log.Printf("connection closed from %s\n", conn.RemoteAddr())
		}()
	}
}

func handleConn(ctx context.Context, dest string, conn net.Conn) error {
	next, err := net.Dial("tcp", dest)
	if err != nil {
		return err
	}

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		_, err := io.Copy(next, conn)
		return err
	})
	g.Go(func() error {
		_, err := io.Copy(conn, next)
		return err
	})

	return g.Wait()
}
