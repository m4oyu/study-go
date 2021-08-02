# Go Command [¶](https://pkg.go.dev/cmd/go)


## Compile packages and dependencies [¶](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies)
```shell
go build [-o output] [buld flags] [packages]
```
- -oフラグで出力先を指定できる
- Buildはimportで指定したパッケージとその依存関係をコンパイルする。
- installはされない
- build対象が複数or単一の非Mainパッケージの場合は結果を破棄する。buildできるかのチェックができる。
- _test.goは無視される
- buildコマンド用のフラグはclean, get, install, list, run, testで共有される

簡単な動作チェックように使えそう。docker環境であれば、同ディレクトリに実行ファイルが生成される（要確認）ので扱いやすそう。


## Environment variables [¶](https://pkg.go.dev/cmd/go#hdr-Environment_variables)

| variables | description |
|---|---|
| GO111MODULE | Goコマンドをmodule-awareモードで実行するか否か.デフォルトはon.offの場合はGOPATHモードになる。 |
| GOARCH | コンパイラのアーキテクチャやプロセッサ名 |
| GCCGO | `go build -compiler=gccgo`するため |
| GOBIN | `go install` した際のインストール先ディレクトリ |
| GOCACHE | go modules のキャッシュ関連情報? |
| GOMODCACHE | go コマンドがダウンロードしたモジュールを格納するところ |
| GODEBUG |  |
| GOPROXY | go module proxy の URL |
