package gossip

type Node struct{}

func Start(cfg Config) (*Node, error) {
	_ = cfg.WithDefaults()
	return &Node{}, nil
}
func (n *Node) Close() error { return nil }
