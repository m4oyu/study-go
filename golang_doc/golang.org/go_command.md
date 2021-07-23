# Go Command 
src: https://pkg.go.dev/cmd/go#hdr-Environment_variables

## Environment variables


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
