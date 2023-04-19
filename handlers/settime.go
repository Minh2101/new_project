package handlers

import (
	"github.com/robfig/cron"
)

func (h *Hanler) SetTimeFeedData() {
	c := cron.New()

	// Đặt lịch vào lúc 9h sáng
	c.AddFunc("0 0 9 * * *", func() {
		h.service.ArticlesServier.ListAndSaveArticles()
	})
	// Đặt lịch 30 phút 1 lần
	c.AddFunc("0 */30 * * * *", func() {
		h.service.ArticlesServier.ListAndSaveArticles()
	})

	c.Start()
}
