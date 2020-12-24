package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"

	"github.com/Hoongeun/gogoatalk/client/app"
	"github.com/Hoongeun/gogoatalk/common"
	"github.com/Hoongeun/gogoatalk/common/util"
	"github.com/spf13/cobra"
)

func main() {
	var (
		flagPort int
		flagHost string
	)
	var rootCmd = &cobra.Command{
		Use:   "client",
		Short: "Gogoatalk-KakaoTalk replication(client)",
		Long:  `Simple terminal ui based chatapp`,
		Run: func(cmd *cobra.Command, args []string) {
			if !util.CheckPortRange(flagPort) {
				log.Fatal("Port should be between 1 to 65535")
			}
			if flagHost != "" {
				host, err := util.GetLocalIP()
				if err != nil {
					log.Fatal(err)
				}
				flagHost = host
			}
			sigctx := common.SignalContext(context.Background())
			app := app.NewApp(sigctx)
			err := app.Start(fmt.Sprintf("%s:%d", flagHost, flagPort))
			if err != nil {
				log.Fatal(fmt.Sprintf("Cannot connect to %s:%d", flagHost, flagPort))
			}
		},
	}
	rootCmd.Flags().IntVarP(&flagPort, "port", "p", 6655, "server port")
	rootCmd.Flags().StringVarP(&flagHost, "server", "s", "127.0.0.1", "server host")
	rootCmd.Execute()
}

var (
	ErrParseFail = errors.New("Fail to parse url")
	ErrHasNoHost = errors.New("Fail to get error")
)

func parseUrl(host string) (string, error) {
	fmt.Println(host)
	u, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	if u.Host == "" {
		return "", ErrHasNoHost
	}
	return u.Host, nil
}

func checkPortRange(port int) bool {
	return port >= 1 && port <= 65535
}

func handleConnection(conn net.Conn) {
	fmt.Println("Connected")
}
