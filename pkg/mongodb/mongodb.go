package mongodb

import (
	"context"
	"fwds/internal/conf"
	"fwds/pkg/log"
	"github.com/qiniu/qmgo"
)

func Open(db *conf.MongoDB) *qmgo.QmgoClient {
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: db.Uri, Database: db.Database, Coll: db.Coll})
	if err != nil {
		log.SugaredLogger.Panic("MongoDB open", err)
	}
	return cli
}

func Close(cli *qmgo.QmgoClient) {
	ctx := context.Background()
	if err := cli.Close(ctx); err != nil {
		log.SugaredLogger.Panic("MongoDB close", err)
	}
}
