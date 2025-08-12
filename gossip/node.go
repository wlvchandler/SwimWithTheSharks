package gossip

type Node struct{}

func Start(cfg Config) (*Node, error) { return &Node{}, nil }
func (n *Node) Close() error          { return nil }
