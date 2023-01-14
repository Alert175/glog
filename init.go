package glog

type GLog struct {
	FolderPath string
}

var config GLog

// Установить обязательные параметры для логгера
func Config(args *GLog) error {
	if args == nil {
		return nil
	}
	config.FolderPath = args.FolderPath
	return nil
}

func New() *GLog {
	return &config
}
