package storageNew

import (
	"bannerService/pkg/utils"
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"gitlab.com/piorun102/lg"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
	SslMode  string
}

var sqlFC = 2

func NewDB(ctx context.Context, cfg *DBConfig) (*pgxpool.Pool, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Db, cfg.SslMode)
	return pgxpool.New(ctx, connectionUrl)
}

func SelectOneStruct[T any](ctx lg.CtxLogger, pool *pgxpool.Pool, sql string, args pgx.NamedArgs) (*T, error) {
	var rows pgx.Rows
	sp, _ := opentracing.StartSpanFromContext(ctx.Ctx(), dbOperation)
	defer sp.Finish()
	ctx.SpanLog("sql", utils.ResSql(sql, args))
	ctx.Tracef("SQL: " + utils.ResSql(sql, args))
	rows, err := pool.Query(ctx.Ctx(), sql, args)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	return &t, err
}

func Insert(ctx lg.CtxLogger, pool *pgxpool.Pool, sql string, args pgx.NamedArgs) (err error) {
	var rows pgx.Rows
	sp, _ := opentracing.StartSpanFromContext(ctx.Ctx(), dbOperation)
	sp.LogKV("sql REQ", utils.ResSql(sql, args))
	ctx.TracefFC(sqlFC, "SQL"+utils.ResSql(sql, args))
	rows, err = pool.Query(ctx.Ctx(), sql, args)
	rows.Close()
	sp.LogKV("sql RESP ERR", fmt.Sprintf("%v", err))
	sp.Finish()
	return err
}

func SelectStructs[T any](ctx lg.CtxLogger, pool *pgxpool.Pool, sql string, args pgx.NamedArgs) (t []T, err error) {
	var rows pgx.Rows
	sp, _ := opentracing.StartSpanFromContext(ctx.Ctx(), dbOperation)
	defer sp.Finish()
	ctx.SpanLog("sql", utils.ResSql(sql, args))
	ctx.Tracef("SQL" + utils.ResSql(sql, args))
	rows, err = pool.Query(ctx.Ctx(), sql, args)
	defer rows.Close()
	if err != nil {
		return
	}
	t, err = pgx.CollectRows(rows, pgx.RowToStructByName[T])
	return
}

func Get[T any](ctx lg.CtxLogger, pool *pgxpool.Pool, sql string, args ...any) (t T, err error) {
	sp, _ := opentracing.StartSpanFromContext(ctx.Ctx(), dbOperation)
	defer sp.Finish()
	ctx.SpanLog("sql", utils.ResSql(sql, args...))
	ctx.Tracef("SQL" + utils.ResSql(sql, args...))
	err = pool.QueryRow(ctx.Ctx(), sql, args...).Scan(&t)
	return
}

func SelectSimple[T any](ctx lg.CtxLogger, pool *pgxpool.Pool, sql string, args ...any) (t []T, err error) {
	var rows pgx.Rows
	sp, _ := opentracing.StartSpanFromContext(ctx.Ctx(), dbOperation)
	sp.LogKV("sql REQ", utils.ResSql(sql, args...))
	ctx.Tracef("SQL" + utils.ResSql(sql, args...))
	rows, err = pool.Query(ctx.Ctx(), sql, args...)
	if err != nil {
		sp.LogKV("sql RESP", fmt.Sprintf("%v %v", t, err))
		sp.Finish()
		rows.Close()
		return
	}
	t, err = pgx.CollectRows(rows, pgx.RowTo[T])
	sp.LogKV("sql RESP", fmt.Sprintf("%v %v", t, err))
	sp.Finish()
	rows.Close()
	return
}

const dbOperation = "dbOperation"
