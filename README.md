# gin-project
原文详解地址[https://book.eddycjy.com/golang/gin/jwt.html]
##日志服务
## 优雅的重启
```$xslt
> ctrl + c 	SIGINT	强制进程结束

> ctrl + z	SIGTSTP	任务中断，进程挂起

> ctrl + \	SIGQUIT	进程结束 和 dump core

> ctrl + d		EOF

> 终止收到该信号的进程。若程序中没有捕捉该信号，当收到该信号时，进程就会退出（常用于 重启、重新加载进程）
```

### 实现优雅的重启
```$xslt
**endless**
Zero downtime restarts for golang HTTP and HTTPS servers. (for golang 1.3+)

我们借助 fvbock/endless <https://github.com/fvbock/endless>来实现
Golang HTTP/HTTPS 服务重新启动的零停机
endless server 监听以下几种信号量：
- syscall.SIGHUP：触发 fork 子进程和重新启动
- syscall.SIGUSR1/syscall.SIGTSTP：被监听，但不会触发任何动作
- syscall.SIGUSR2：触发 hammerTime
- syscall.SIGINT/syscall.SIGTERM：触发服务器关闭（会完成正在运行的请求）
- endless 正正是依靠监听这些信号量，完成管控的一系列动作
```



## 添加Swagger
## 修改Grom CallBack
## 添加定时任务
# 重构
> 优化conf 读取配置  