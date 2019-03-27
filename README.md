# Rugby Result

![travis](https://travis-ci.org/rugby-board/rugby-result.svg?branch=master)
![codecov](https://codecov.io/gh/rugby-board/rugby-result/branch/master/graph/badge.svg)

Rugby match result CLI retriever in Go.

It is rather simple by now, only allows query with EventID for results from last week, and output in Markdown table format.

## Installation

```shell
go get github.com/rugby-board/rugby-result
```

## Usage

```shell
# Need days, default 7
rugby-result -id=209 -days=3
# NRC need `round`
rugby-result -id=247 -round=1
# Show all events info
rugby-result -list-events
# Check all events
rugby-result -iter-events
```

## PID Mapping

### From planetrugby.com

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

### From rugby.com.au

* 247: National Rugby Championship

## License

MIT
