package sql2struct

import "github.com/cascax/sql2gorm/parser"

func Sql2gorm(sql string) (*parser.ModelCodes, error) {
	data, err := parser.ParseSql(sql, parser.WithTablePrefix("t_"), parser.WithJsonTag())
	if err != nil {
		return nil, err
	}
	return &data, nil
}
