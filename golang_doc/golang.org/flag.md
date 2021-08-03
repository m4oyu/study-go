# flag [¶](https://pkg.go.dev/flag)
　package内でflagの宣言->flag.Parse()で渡して使用する。

渡し方
```shell
-flag
-flag=x
-flag x (non-boolean flag only)
```

## example
exampleで示されていた例（flag.Stringを例に）
- flag.Stringを変数に代入する
  スタンダード
- flag.Stringをinit関数で呼び出す例
  細かい設定するならこっちの方がきれいにできそう
- 配列にflagを入れる例
  よくわからない