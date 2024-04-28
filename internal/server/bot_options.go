package server

import (
	"io"
	"time"
)

type BotOption func(bot *Bot)

// WithBotNetworkDelay 设置机器人网络延迟及波动范围
//   - delay 延迟
//   - fluctuation 波动范围
func WithBotNetworkDelay(delay, fluctuation time.Duration) BotOption {
	return func(bot *Bot) {
		bot.conn.delay = delay
		bot.conn.fluctuation = fluctuation
	}
}

// WithBotWriter 设置机器人写入器，默认为 os.Stdout
func WithBotWriter(construction func(bot *Bot) io.Writer) BotOption {
	return func(bot *Bot) {
		writer := construction(bot)
		bot.conn.botWriter.Store(&writer)
	}
}
