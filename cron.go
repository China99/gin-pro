package main

import (
	"gin-pro/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

//编写定时任务

func main() {
	log.Println("Starting")
	//创建实例
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Panicln("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Panicln("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
