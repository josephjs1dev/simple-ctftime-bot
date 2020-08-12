package lineservice

import (
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/logger"
)

// HandleIncomingMessage finds command, executes command, and reply it to user
func (s *ImplService) HandleIncomingMessage(textMessageCtx *domain.LineTextMessageContext) error {
	cmd := s.findCommand(textMessageCtx)
	if cmd == nil {
		return nil
	}

	result, err := cmd.Process()
	if err != nil {
		logger.Errorf("Error processing command %v, error = %v", cmd, err)

		return err
	}

	if resp, err := s.bot.ReplyMessage(textMessageCtx.ReplyToken, result...).Do(); err != nil {
		logger.Errorf("Error reply message, error = %v", err)
	} else {
		logger.Infof("Successfully reply message Result = %v", *resp)
	}

	return nil
}
