package main

import (
	"context"
	"fmt"
	"github.com/eqkez0r/lesta_matchmaker/pkg/object/player"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	defer stop()
	client := resty.New()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	counter := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("exit")
				wg.Done()
				return
			default:
				{
					name := fmt.Sprintf("Player-%d", counter)
					skill := rand.Float32() * 10
					latency := rand.Float32() * 100
					pl := &player.Player{
						Name:    name,
						Skill:   skill,
						Latency: latency,
					}
					resp, err := client.R().
						SetHeader("Content-Type", "application/json").
						SetBody(pl).
						Post("http://your-server-endpoint")

					if err != nil {
						fmt.Println("Error:", err)
						return
					}

					fmt.Println("Player", pl.Name, "sent. Response:", resp.Status())
					counter++
				}
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
	wg.Wait()
}
