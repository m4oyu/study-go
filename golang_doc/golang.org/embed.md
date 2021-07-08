# Package embed
embedディレクティブを使用することで実行中のGoプログラムにム目こまれたファイルへのアクセスを提供する。

```
package main

import (
	_ "embed"
)

//go:embed hello.txt
var s string


func main() {

print(s)

}

-- hello.txt --
hello world

```
Future Tech Blogで説明されていた。勉強になりました。\
[link](./Future_Tech_Blog/go_embed.md)
