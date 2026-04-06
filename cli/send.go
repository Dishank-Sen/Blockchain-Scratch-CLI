package cli

import (
	"encoding/json"
	"fmt"

	"github.com/Dishank-Sen/Blockchain-Scratch-CLI/client"
	"github.com/Dishank-Sen/Blockchain-Scratch-CLI/utils/logger"
	"github.com/spf13/cobra"
)

func init() {
	Register("send", Send)
}

func Send() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [peer-id] [message]",
		Short: "Send a message to a connected peer",
		Args:  cobra.ExactArgs(2),
		RunE:  sendRunE,
		Example: `  bloc send abc123def456 "Hello from peer!"
  bloc send peer-id-hash "This is a test message"`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	return cmd
}

func sendRunE(cmd *cobra.Command, args []string) error {
	peerID := args[0]
	message := args[1]

	c := client.NewClient()

	payload := struct {
		To      string `json:"to"`
		Message string `json:"message"`
	}{
		To:      peerID,
		Message: message,
	}

	byteData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := c.Post("/send", byteData)
	if err != nil {
		if isDaemonNotRunning(err) {
			logger.Error("daemon is not running")
			fmt.Println("Run `bloc connect` to start the daemon.")
			return nil
		}
		return err
	}

	if resp.StatusCode != 200 {
		logger.Error(string(resp.Body))
		return fmt.Errorf("failed to send message")
	}

	logger.Info(string(resp.Body))
	return nil
}
