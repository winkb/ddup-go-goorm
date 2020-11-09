package where

import (
	"github.com/winkb/ddup-go-util/orm/where"
	"reflect"
	"testing"
)

func TestForGorm(t *testing.T) {
	wr := where.Where{}

	t.Log(reflect.ValueOf(wr).String())
	t.Log("测试通过")
}
