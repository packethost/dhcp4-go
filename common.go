package dhcpv4

// Reply defines an interface implemented by DHCP replies.
type Reply interface {
	Validate() error
	ToBytes() ([]byte, error)
}

// ReplyWriter defines an interface for the object that writes a reply to the
// network to the intended received, be it via broadcast or unicast.
type ReplyWriter interface {
	WriteReply(r Reply) error
}