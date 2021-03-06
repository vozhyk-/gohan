package autogen

// AUTO GENERATED CODE DO NOT MODIFY MANUALLY
import (
	"github.com/cloudwan/gohan/extension/gohanscript"
	"github.com/cloudwan/gohan/extension/gohanscript/lib"
)

func init() {

	gohanscript.RegisterStmtParser("uuid",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				result1 :=
					lib.UUID()

				return result1, nil

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("UUID",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			result1 :=
				lib.UUID()
			return []interface{}{
				result1}

		})

	gohanscript.RegisterStmtParser("format_uuid",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				var uuid string
				iuuid := stmt.Arg("uuid", context)
				if iuuid != nil {
					uuid = iuuid.(string)
				}

				result1,
					err :=
					lib.FormatUUID(
						uuid)

				return result1, err

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("FormatUUID",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			uuid, _ := args[0].(string)

			result1,
				err :=
				lib.FormatUUID(
					uuid)
			return []interface{}{
				result1,
				err}

		})

	gohanscript.RegisterStmtParser("env",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				result1 :=
					lib.Env()

				return result1, nil

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("Env",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			result1 :=
				lib.Env()
			return []interface{}{
				result1}

		})

	gohanscript.RegisterStmtParser("normalize_map",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			return func(context *gohanscript.Context) (interface{}, error) {

				var data map[string]interface{}
				idata := stmt.Arg("data", context)
				if idata != nil {
					data = idata.(map[string]interface{})
				}

				result1 :=
					lib.NormalizeMap(
						data)

				return result1, nil

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("NormalizeMap",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			data, _ := args[0].(map[string]interface{})

			result1 :=
				lib.NormalizeMap(
					data)
			return []interface{}{
				result1}

		})

}
