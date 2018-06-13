# result-cli

![travis](https://travis-ci.org/rugby-board/result-cli.svg?branch=master)

Rugby match result retriever in Go.

It is rather simple by now, only allows query with EventID for results from last week, and output in Markdown table format.

## Usage

```
$ go get github.com/rugby-board/result-cli
$ result-cli -id=209
```

## API

Pattern: `http://kratos.365.co.za:9001/getresultsbycompidanddaterange/[PID]/[DATE-START]/[DATE-END]`

### PID Mapping

* 3: International Tests
* 201: Premiership
* 203: Top14
* 204: Pro14
* 205: Super Rugby
* 206: Anglo Welsh Cup
* 208: Mitre 10 Cup
* 209: Six Nations
* 210: Rugby World Cup
* 214: The Rugby Championship
* 221: British & Irish Lions
* 242: European Champion Cup
* 243: European Challenge Cup
* 303: Currie Cup Premier
