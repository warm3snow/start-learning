/**
 * @Author: xueyanghan
 * @File: slog_test.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/10/11 17:50
 */

package v1_21

import (
	"context"
	"log/slog"
	"os"
	"testing"
)

func TestSlog(t *testing.T) {
	slog.Info("hello slog")
	slog.Info("hello slog", "name", "slog")

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Error("hello slog", "name", "slog")

	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Error("hello slog", "name", "slog")
}

func TestSlogFaster(t *testing.T) {
	slog.LogAttrs(context.Background(), slog.LevelInfo, "hello slog", slog.String("name", "slog"))

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.LogAttrs(context.Background(), slog.LevelInfo, "hello slog", slog.String("name", "slog"))

	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.LogAttrs(context.Background(), slog.LevelInfo, "hello slog", slog.String("name", "slog"))
}

func TestSlogKVAttr(t *testing.T) {
	slog.Info("hello slog", "name", "slog", "age", 18)

	slog.Info("hello slog", slog.String("name", "slog"), slog.Int("age", 18))
}
