package main

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/DilemaFixer/Mi-Mi-Mi-Bot/src/consts"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const BOT_TOKEN = "SET_YOUR_TOKEN"

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(BOT_TOKEN, opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/next", bot.MatchTypeExact, nextHandler)
	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "next", bot.MatchTypeExact, nextButtomHandler)
	b.Start(ctx)
}

func getRandomPhrase() string {
	return consts.Phrases[rand.Intn(len(consts.Phrases))]
}

func nextHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        getRandomPhrase(),
		ReplyMarkup: consts.KB,
	})
}

func nextButtomHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
		Text:        getRandomPhrase(),
		ReplyMarkup: consts.KB,
	})
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Write /next for geting result",
	})
}
