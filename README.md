# go-struct-default-getter-codegen

This tool generates functions to get assigned or initial values for all structure pointer type fields defined in a given file.

## Usage
Installation:
```shell 
go install github.com/SHOWROOM-inc/go-struct-default-getter-codegen/cmd/default-getter-gen@latest 
`` 

Running the command: 
`` shell 
default-getter-gen --input model.go --output model.gen.go --package package_name 
``` 

## Generated contents
Returns the value or initial value of a pointer-type field for all structures defined in the file given by `--input`. Generate a function that returns the value or initial value of a field of pointer type for all structures defined in the file given by `--input`.

- The target is only pointer types.
- If the pointer field has a value, the value is returned; if it is nil, the initial value is returned.
- If there is a `default` tag, the initial value is the value specified by the default tag.
- If there is no `default` tag, then the initial value of the type is used.

For example, if the input [live.go](. /examples/live.go) for input [live.gen.go](. /examples/live.gen.go).

---
与えられたファイルに定義された全ての構造体のポインタ型のフィールドに対して、代入された値もしくは初期値を取得する関数を生成するツールです。

## 使い方
インストール:
```shell
go install github.com/SHOWROOM-inc/go-struct-defalut-getter-codegen/cmd/default-getter-gen@latest
```

コマンドの実行:
```shell 
default-getter-gen --input model.go --output model.gen.go --package package_name
```

## 生成内容
`--input`で与えられたファイルに定義された全構造体を対象に、ポインタ型のフィールドの値もしくは初期値を返す関数を生成します。

- 対象はポインタ型のみです。
- ポインタフィールドに値があればその値を返します。nilの場合は初期値を返します。
- `default`タグがあれば、デフォルトタグで指定された値を初期値とします。
- `default` タグがなければ、その型の初期値を使います。

例えば、入力 [live.go](./examples/live.go) に対して [live.gen.go](./examples/live.gen.go) を出力します。
