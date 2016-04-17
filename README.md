mlog
====

[![Build Status](https://travis-ci.org/cactus/mlog.png?branch=master)](https://travis-ci.org/cactus/mlog)
[![GoDoc](https://godoc.org/github.com/cactus/mlog?status.png)](https://godoc.org/github.com/cactus/mlog)

## About

A purposefully basic logging library for Go.

mlog only has 3 logging levels: Debug, Info, and Fatal.

### Why only 3 levels?

[Over here][1] Dave Cheney has a good explanation of the reasoning.

### How does it work?

*   `Debugm`, `Debugf`, and `Debug` only emit when the debug flag is set.
*   `Fatalm`, `Fatalf`, and `Fatal` call `os.Exit(1)` after emitting the
    associated message data.
*   `Infom`, `Infof`, and `Info` just emit the associated message data.
*   `Printm`, `Printf`, and `Print` are aliases of the Info methods above.

That's it!

For more info, check out the [docs][3].

## Usage

``` go
import (
    "bytes"

    "github.com/cactus/mlog"
)

func main() {
    mlog.Infom("this is a log")

    mlog.Infom("this is a log with more data", mlog.Map{
        "interesting": "data",
        "something":   42,
    })

    thing := mlog.Map(
        map[string]interface{}{
            "what‽":       "yup",
            "this-works?": "as long as it is a mlog.Map",
        },
    )

    mlog.Infom("this is a log with more data", thing)

    mlog.Debugm("this won't print")

    // set flags for the default logger
    // alternatively, you can create your own logger
    // and supply flags at creation time
    mlog.SetFlags(mlog.Ltimestamp | mlog.Ldebug)

    mlog.Debugm("now this will print!")

    mlog.Debugm("can it print?", mlog.Map{
        "how_fancy": []byte{'v', 'e', 'r', 'y', '!'},
        "this_too":  bytes.NewBuffer([]byte("if fmt.Print can print it!")),
    })

    // you can use a more classical Printf type log method too.
    mlog.Debugf("a printf style debug log: %s", "here!")
    mlog.Infof("a printf style info log: %s", "here!")

    mlog.Fatalm("time for a nap", mlog.Map{"cleanup": false})
}
```

Output:

```
time="2016-04-11T00:31:30Z" level="I" msg="this is a log"
time="2016-04-11T00:31:30Z" level="I" msg="this is a log with more data" interesting="data" something="42"
time="2016-04-11T00:31:30Z" level="I" msg="this is a log with more data" this-works?="as long as it is a mlog.Map" what‽="yup"
time="2016-04-11T00:31:30Z" level="D" msg="now this will print!"
time="2016-04-11T00:31:30Z" level="D" msg="can it print?" how_fancy="[118 101 114 121 33]" this_too="if fmt.Print can print it!"
time="2016-04-11T00:31:30Z" level="D" msg="a printf style debug log: here!"
time="2016-04-11T00:31:30Z" level="I" msg="a printf style info log: here!"
time="2016-04-11T00:31:30Z" level="F" msg="time for a nap" cleanup="false"
exit status 1
```

## License

Released under the [MIT license][2]. See `LICENSE.md` file for details.

[1]: http://dave.cheney.net/2015/11/05/lets-talk-about-logging
[2]: http://www.opensource.org/licenses/mit-license.php
[3]: https://godoc.org/github.com/cactus/mlog
