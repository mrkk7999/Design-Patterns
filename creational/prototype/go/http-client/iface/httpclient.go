package iface

type Client interface {
	Clone() Client
	GetBaseUrl() string
}
