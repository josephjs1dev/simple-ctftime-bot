package lineservice

import (
	applog "github.com/josephsalimin/simple-ctftime-bot/internal/pkg/log"

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
		applog.Errorf("Error processing command %v, error = %v", cmd, err)

		return err
	}

	if resp, err := s.bot.ReplyMessage(textMessageCtx.ReplyToken, result...).Do(); err != nil {
		applog.Errorf("Error reply message, error = %v", err)
	} else {
		applog.Infof("Successfully reply message Result = %v", *resp)
	}

	return nil
}
