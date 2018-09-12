lambci/lambdaを使ったDocker上でのLambda実行  
テストにかなり使えそう  

1. 対象の関数を含むファイルをbuildしておく(mainバイナリが必要)
2. docker run --rm -v "$PWD":/var/task lambci/lambda:go1.x main '{"some": "event"}'