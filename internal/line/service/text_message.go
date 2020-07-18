package lineservice

import (
	"log"

	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
)

// HandleIncomingMessage finds command, executes command, and reply it to user
func (s *ImplService) HandleIncomingMessage(textMessageCtx *domain.LineTextMessageContext) error {
	cmd := s.findCommand(textMessageCtx)
	if cmd == nil {
		return nil
	}

	result, err := cmd.Process()
	if err != nil {
		log.Printf("Error processing command %v, error = %v\n", cmd, err)

		return err
	}

	if resp, err := s.bot.ReplyMessage(textMessageCtx.ReplyToken, result...).Do(); err != nil {
		log.Printf("Error reply message, error = %v\n", err)
	} else {
		log.Printf("Successfully reply message Result = %v\n", resp)
	}

	return nil
}
