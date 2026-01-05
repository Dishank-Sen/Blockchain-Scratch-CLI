//go:build !windows

package client

import "github.com/Dishank-Sen/Blockchain-Scratch-CLI/constants"

func NewClient() Client{
	return newUnixClient(constants.SocketPath)
}