

## Module-aware commands
- go.modファイルに書かれた依存関係を利用してモジュールキャッシュ空パッケージをロードする
- 足りない場合はキャッシュにダウンロードする。

### go mod verify

### go clean -modcache
- すべてのmodule cacheを削除する。
- モジュールキャッシュを削除するための最善の方法


## Module proxies
### _GOPROXY_ protcol
- proxyサーバからダウンロードした方が高速
- _GOPROXY_変数でダウンロード先の設定
- proxyサーバへGetリクエストした際にどう振る舞うかが指定されている。
- proxyサーバは常に同じモジュールを返す必要性がある（go.sumファイルやchecksum databaseによる認証）
### Communicating with proxies


## Version control systems



## Module cache
- goコマンドがダウンロードしたモジュールファイルを保存するディレクトリ
- ビルド後のファイルではない
- デフォルトでは$GOPATH/pke/modに保存 (GOMODCACHEで指定できる)
- 同一マシン上にある複数のGoプロジェクトで共有される
- 作成後はreadonlyとなり、意図しない変更を防ぐ (削除したいときはgo clean -modcache)

## Authenticating modules
- goコマンドを使用したとき、ハッシュ暗号を使用して前回のダウンロード時からバージョンアップなどがないか確認する。
- 正しいハッシュ値が得られなかった場合はセキュリティエラーを吐く
- 認証後、go.sumファイルとModuleCacheに追加される

## Module proxies
