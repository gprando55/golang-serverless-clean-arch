package merge

import (
	"reflect"
	"time"

	"github.com/imdario/mergo"
)

type timeTransformer struct {
}

func (t timeTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(time.Time{}) {
		return func(dst, src reflect.Value) error {
			if dst.CanSet() {
				isZero := dst.MethodByName("IsZero")
				result := isZero.Call([]reflect.Value{})
				if result[0].Bool() {
					dst.Set(src)
				}
			}
			return nil
		}
	}
	return nil
}

func MergeStruct(dest interface{}, src interface{}) error {
	err := mergo.Merge(&dest, src, mergo.WithTransformers(timeTransformer{}))
	return err
}

// type Snapshot struct {
// 	Time time.Time
// 	// ...
// }

// func main() {
// 	src := Snapshot{time.Now()}
// 	dest := Snapshot{}
// 	mergo.Merge(&dest, src, mergo.WithTransformers(timeTransformer{}))
// 	fmt.Println(dest)
// 	// Will print
// 	// { 2018-01-12 01:15:00 +0000 UTC m=+0.000000001 }
// }
