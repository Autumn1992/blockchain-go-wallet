package db

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
	mylog "walletserver/log"
)

var db *gorm.DB

// 记录ctx
var gCtx *context.Context

func init() {
	var err error
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", GetConfig().Mysql.Username, GetConfig().Mysql.Pwd, GetConfig().Mysql.Host, GetConfig().Mysql.DbName)
	newLogger := logger.New(
		log.New(mylog.ErrorWriter, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "t_",
		},
	})
	// 启用调试模式
	if gin.Mode() == gin.DebugMode {
		db = db.Debug()
	}
	fmt.Println("dsn", dsn)
	if err != nil {
		log.Fatal("数据库链接错误", err)
	}

}

type contextTxKey struct{}

type TxFn func(context.Context) error

func ExecTx(ctx context.Context, fn TxFn) error {
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func GetDB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return db
}
