package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/libgox/addr"
	"github.com/protocol-laboratory/zookeeper-client-go/zk"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	config := &zk.Config{
		Addresses: []addr.Address{
			{
				Host: "localhost",
				Port: 2181,
			},
		},
		Logger: logger,
	}
	client, err := zk.NewClient(config)
	if err != nil {
		panic(err)
	}

	queue := []string{"/"}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		children, err := client.GetChildren(currentNode)
		if err != nil {
			logger.Error("failed to get children", slog.String("node", currentNode), slog.Any("err", err))
			continue
		}

		logger.Info("visited node", slog.String("node", currentNode), slog.Any("children", children))

		for _, child := range children.Children {
			childPath := currentNode
			if childPath != "/" {
				childPath += "/"
			}
			childPath += child
			queue = append(queue, childPath)
		}

		time.Sleep(1 * time.Second)
	}
}
