package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"path/filepath"

	"github.com/Hoongeun/gogoatalk/common"
	"github.com/Hoongeun/gogoatalk/common/util"
	"github.com/Hoongeun/gogoatalk/server/socket"
	"github.com/spf13/cobra"
)

func main() {
	var (
		flagPort   int
		flagJWTKey string
	)
	var rootCmd = &cobra.Command{
		Use:   "server",
		Short: "Cocoatalk-KakaoTalk replication",
		Long:  `Simple terminal ui based chatapp`,
		Run: func(cmd *cobra.Command, args []string) {
			if !util.CheckPortRange(flagPort) {
				log.Fatal("Port should be between 1 to 65535")
			}
			host, err := util.GetLocalIP()
			if err != nil {
				log.Fatal(err)
			}
			dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
			jwtKeyPath := filepath.Join(dir, flagJWTKey)

			if !util.DoesFileExist(jwtKeyPath) {
				log.Fatalf("jwtkey doesn't exist %s", jwtKeyPath)
			}

			if !util.DoesFileExist(jwtKeyPath + ".pub") {
				log.Fatalf("jwtkey.pub doesn't exist %s", jwtKeyPath)
			}

			sigctx := common.SignalContext(context.Background())
			server, err := socket.NewServer(jwtKeyPath)
			if err != nil {
				log.Fatalf("Failed to Initialized Server: %s", err)
			}

			server.Listen(fmt.Sprintf("%s:%d", host, flagPort), sigctx)
			if err != nil {
				log.Fatalf("Cannot Listen port number %d", flagPort)
			}
		},
	}
	rootCmd.Flags().IntVarP(&flagPort, "port", "p", 6655, "server port")
	rootCmd.Flags().StringVarP(&flagJWTKey, "key", "k", "jwtRS2048.key", "jwt ssh key")
	rootCmd.Execute()
}

var (
	ErrParseFail = errors.New("Fail to parse url")
	ErrHasNoHost = errors.New("Fail to get error")
)

func parseUrl(host string) (string, error) {
	u, err := url.Parse(host)
	if err != nil {
		return "", ErrParseFail
	}
	if u.Host == "" {
		return "", ErrHasNoHost
	}
	return u.Hostname(), nil
}

func handleConnection(conn net.Conn) {
	fmt.Println("Connected")
}
