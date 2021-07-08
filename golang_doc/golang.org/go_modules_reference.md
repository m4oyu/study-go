# 

## go get
go get コマンドはモジュールの依存関係をアップデートする

## go install
このコマンドはパスで指定されたパッケージをビルドしてインストールする




[Go 1.16 Release Notes](https://golang.org/doc/go1.16#go-command)
によるGoModules関連の内容

- Go111MODULE=onがデフォルトになった
- go build や go testによるgo.modとgo.sumの書き換えはなくなった
- モジュールの要件はgo mod tidy か go get で調整できる。
- go install は GO111MODULE=on において go.mod ファイルを無視してパッケージのビルド及びインストールを行う。

- モジュールインストールの推奨方法が go install になった。
- パッケージのビルドをせずに go.mod の調整を行うには -d フラグが必要
  - 将来的には -d フラグがデフォルトで有効になる。

- exclude コマンドで除外されたモジュールは go コマンドでは無視されるようになり、go コマンドによる勝手な書き換えによるファイルの損壊を防ぐ




### go mod vendor

### go mod tidy