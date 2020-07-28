package server

import (
	"saver/config"
	"saver/store/status"
	"saver/subscriber"

	"github.com/nats-io/go-nats"
	"github.com/spf13/cobra"
)

func Register(root *cobra.Command, n *nats.Conn, natsCfg config.Nats,
	r status.RedisStatus, redisCfg config.Redis, s status.SQLStatus) {
	c := cobra.Command{
		Use:   "server",
		Short: "Subscribes to the right topic",
		Run: func(cmd *cobra.Command, args []string) {
			s := subscriber.New(n, natsCfg, &r, redisCfg, s)
			s.Subscribe()
		},
	}

	root.AddCommand(
		&c,
	)
}
