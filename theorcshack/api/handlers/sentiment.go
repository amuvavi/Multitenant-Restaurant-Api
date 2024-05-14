package handlers

import (
	"sync"

	"github.com/cdipaolo/sentiment"
)

var (
	model   sentiment.Models
	modelMu sync.Mutex
)

// InitSentimentModel initializes the sentiment model
func InitSentimentModel() error {
	modelMu.Lock()
	defer modelMu.Unlock()

	var err error
	model, err = sentiment.Restore()
	return err
}

// analyzeSentiment analyzes the sentiment of the given text
func analyzeSentiment(text string) float64 {
	modelMu.Lock()
	defer modelMu.Unlock()

	analysis := model.SentimentAnalysis(text, sentiment.English)
	return float64(analysis.Score)
}
