# SQS-Marksheet-Exe
## Description
SQS(Shared Questionnaire System)は、千葉商科大学政策情報学部准教授 久保裕也氏により開発されたオープンソースの「共有アンケート実施支援システム」です．  
なぜか職場のPCではJarファイルの関連付けがされていなかったので、Golangでコマンドをラップすることにしました。  
裏でJarの実行コマンドを入力しているだけなので、JDKやJRE等が必要になります。

## Usage
[Release](https://github.com/NNAKAYAMA/sqs-marksheet-exe/releases/tag/v1.0.0)からダウンロードしてください。  
ダブルクリックでうごきます。  

## Install
```
go get github.com/rakyll/statik@latest
git clone https://github.com/NNAKAYAMA/sqs-marksheet-exe.git
cd sqs-marksheet-exe
```
### Build MarksheetReader
```
cd MarksheetReader
go build
```
### Build MarksheetEditer
```
cd MarksheetEditer
go build
```

## Licence

[Apache License Version 2.0](./LICENSE)

## Link
[Shared Questionnaire System](https://sqs2.net/projects/sqs/wiki/Overview_ja)