package internal

import (
	"context"

	"knife-panel/internal/app/config"
	icontext "knife-panel/internal/app/context"
	"knife-panel/internal/app/model"
	"knife-panel/internal/app/schema"
	"knife-panel/pkg/util"
)

// GetRootUser 获取root用户
func GetRootUser() *schema.User {
	user := config.Global().Root
	return &schema.User{
		RecordID: user.UserName,
		UserName: user.UserName,
		RealName: user.RealName,
		Password: util.MD5HashString(user.Password),
	}
}

// CheckIsRootUser 检查是否是root用户
func CheckIsRootUser(ctx context.Context, userID string) bool {
	return GetRootUser().RecordID == userID
}

// TransFunc 定义事务执行函数
type TransFunc func(context.Context) error

// ExecTrans 执行事务
func ExecTrans(ctx context.Context, transModel model.ITrans, fn TransFunc) error {
	if _, ok := icontext.FromTrans(ctx); ok {
		return fn(ctx)
	}
	trans, err := transModel.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = transModel.Rollback(ctx, trans)
			panic(r)
		}
	}()

	err = fn(icontext.NewTrans(ctx, trans))
	if err != nil {
		_ = transModel.Rollback(ctx, trans)
		return err
	}
	return transModel.Commit(ctx, trans)
}

// ExecTransWithLock 执行事务（加锁）
func ExecTransWithLock(ctx context.Context, transModel model.ITrans, fn TransFunc) error {
	if !icontext.FromTransLock(ctx) {
		ctx = icontext.NewTransLock(ctx)
	}
	return ExecTrans(ctx, transModel, fn)
}
