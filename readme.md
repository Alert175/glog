# GLog

Пакет для пользовательского предствления вывода информации.

Ps. Использует пакет `github.com/pkg/errors`

## Установка

`go get github.com/Alert175/glog`

## Настройка

Перед использованием, необходимо обязательно настроить логгер.

Описание конфигурации:
```
type GLog struct {
	FolderPath string // путь к папке, где будут храниться логи, будут разбиваться на типы логов (error.log и т.д.) 
}
```

Использование:
```
package app

import (
	"errors"

	"github.com/Alert175/glog"
)

func main() {
	// настройка логгера
	glog.Config(glog.GLog{
		FolderPath: ".logs",
	})
}

func someHandler() {
	// получение экземпляра логгера
	logger := glog.New()

	// использование логгера
	logger.Info("info log")
	logger.Warn("warning log")
	logger.Error("error log", errors.New("some err"), false)
}
```
