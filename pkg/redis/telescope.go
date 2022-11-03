package redis

import (
	"context"
	"fwds/pkg/util/timeutil"

	tel "fwds/pkg/telescope"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type telescope struct {
}

func (t *telescope) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {

	return ctx, nil
}

func (t *telescope) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return nil
	}
	var err error
	if cmd.Err() != redis.Nil {
		err = cmd.Err()
	}
	cmd.Args()
	r := tel.Redis{
		Timestamp: timeutil.NowMicrosecondString(),
		Handle:    cmd.FullName(),
		Cmd:       cmd.String(),
		Err:       err,
	}
	tel.GetTelescope(c).AppendRedis(&r)
	return nil
}

func (t *telescope) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil

}

func (t *telescope) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return nil
	}
	ts := tel.GetTelescope(c)
	var err error
	for i := range cmds {
		if cmds[i].Err() != redis.Nil {
			err = cmds[i].Err()
		} else {
			err = nil
		}
		ts.AppendRedis(&tel.Redis{
			Timestamp: timeutil.NowMicrosecondString(),
			Handle:    cmds[i].FullName(),
			Cmd:       cmds[i].String(),
			Err:       err,
		})
	}
	return nil
}
