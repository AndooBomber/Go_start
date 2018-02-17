# 演習課題2

タイピングゲームを作ろう

## 起動方法

go run Typing.go

## 操作方法

出力された英語をタイピングする。制限時間は10秒。間違えると次の単語に進めない。

## 工夫した点

単語スライスの要素を30個用意した。
time.After(10 * time.Second)をSelectの外に定義することで制限時間が更新されないようにした。

scan := bufio.NewScanner(os.std)
for scan.Scan(){
  //ここがbreakしないと終了してくれなかった。解決できず。
}