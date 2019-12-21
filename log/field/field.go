package field

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Field = zap.Field

func Binary(key string, value []byte) zap.Field {
	return zap.Binary(key, value)
}

func Bool(key string, value bool) zap.Field {
	return zap.Bool(key, value)
}

func ByteString(key string, value []byte) zap.Field {
	return zap.ByteString(key, value)
}

func Complex128(key string, value complex128) zap.Field {
	return zap.Complex128(key, value)
}

func Complex64(key string, value complex64) zap.Field {
	return zap.Complex64(key, value)
}

func Error(err error) zap.Field {
	return zap.Error(err)
}

func Float64(key string, value float64) zap.Field {
	return zap.Float64(key, value)
}

func Float32(key string, value float32) zap.Field {
	return zap.Float32(key, value)
}

func Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

func Int64(key string, value int64) zap.Field {
	return zap.Int64(key, value)
}

func Int32(key string, value int32) zap.Field {
	return zap.Int32(key, value)
}

func Int16(key string, value int16) zap.Field {
	return zap.Int16(key, value)
}

func Int8(key string, value int8) zap.Field {
	return zap.Int8(key, value)
}

func String(key string, value string) zap.Field {
	return zap.String(key, value)
}

func Uint(key string, value uint) zap.Field {
	return zap.Uint(key, value)
}

func Uint64(key string, value uint64) zap.Field {
	return zap.Uint64(key, value)
}

func Uint32(key string, value uint32) zap.Field {
	return zap.Uint32(key, value)
}

func Uint16(key string, value uint16) zap.Field {
	return zap.Uint16(key, value)
}

func Uint8(key string, value uint8) zap.Field {
	return zap.Uint8(key, value)
}

func Stringer(key string, value fmt.Stringer) zap.Field {
	return zap.Stringer(key, value)
}

func Time(key string, value time.Time) zap.Field {
	return zap.Time(key, value)
}

func Stack(key string) zap.Field {
	return zap.Stack(key)
}

func Duration(key string, value time.Duration) zap.Field {
	return zap.Duration(key, value)
}

func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func Package(name string) zap.Field {
	return zap.String("package", name)
}

func Entity(name string) zap.Field {
	return zap.String("entity", name)
}
