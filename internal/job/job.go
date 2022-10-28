package job

import (
	"github.com/robfig/cron/v3"
)

//https://mp.weixin.qq.com/s/Ak7RBv1NuS-VBeDNo8_fww

// Init
// @Description: 启动
//
func Init() {
	c := cron.New()
	AddFunc(c)
	AddJob(c)
	c.Start()
}

func AddFunc(c *cron.Cron) {
	//c.AddFunc("@every 1s", func() {
	//	fmt.Println("tick every 1 second")
	//})
	//c.AddFunc("0 0 1 1 *", func() {
	//	fmt.Println("Jun 1 every year")
	//})
}

func AddJob(c *cron.Cron) {
	//Recover：捕获内部Job产生的 panic；
	//DelayIfStillRunning：触发时，如果上一次任务还未执行完成（耗时太长），则等待上一次任务完成之后再执行；
	//SkipIfStillRunning：触发时，如果上一次任务还未完成，则跳过此次执行。
	//c.AddJob("@every 1s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&example.PanicJob{}))
	//c.AddJob("@every 1s", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&example.DelayJob{}))
	//c.AddJob("@every 1s", cron.NewChain(cron.SkipIfStillRunning(cron.DefaultLogger)).Then(&example.SkipJob{}))

}
