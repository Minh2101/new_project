package handlers

import (
	"github.com/robfig/cron"
)

func (h *Hanler) SetTimeFeedData() {
	var (
		err error
		c   = cron.New()
	)
	// Đặt lịch vào lúc 9h sáng
	err = c.AddFunc("0 0 9 * * *", func() {
		err = h.service.ArticlesServier.ListAndSaveArticles()
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}
	// Đặt lịch 30 phút 1 lần
	err = c.AddFunc("0 */30 * * * *", func() {
		err = h.service.ArticlesServier.ListAndSaveArticles()
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}

	c.Start()
}
