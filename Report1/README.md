# 演習課題1

catコマンドを作ろう

## 起動方法

ex)
go run test.go test1.txt test2.txt
go run test.go -n test1.txt test2.txt

go run test.go "filename" "-n"

## 工夫した点

flag.Args()を用いることでos.Argsよりもスッキリした。