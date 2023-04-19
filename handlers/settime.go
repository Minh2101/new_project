package handlers

import (
	"github.com/robfig/cron"
)

var c = cron.New()

func (h *Handler) SetTimeFeedData() {
	// Đặt lịch vào lúc 9h sáng
	err := c.AddFunc("0 0 9 * * *", func() {
		if err := h.service.ArticlesServier.ListAndSaveArticles(); err != nil {
			return
		}
	})
	if err != nil {
		return
	}
	// Đặt lịch 30 phút 1 lần
	err = c.AddFunc("0 */30 * * * *", func() {
		if err = h.service.ArticlesServier.ListAndSaveArticles(); err != nil {
			return
		}
	})
	if err != nil {
		return
	}

	c.Start()
}
