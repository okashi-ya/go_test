package u

import "github.com/google/uuid"

func GetUid() string {
	return uuid.NewString()
}

// Version go-cqhttp的版本信息，在编译时使用ldflags进行覆盖
var Version = "unknown"

func V() string {
	return Version
}
