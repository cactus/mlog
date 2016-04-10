// Copyright (c) 2012-2016 Eli Janssen
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

/*
Package mlog provides a purposefully basic logging library for Go.

mlog only has 3 logging levels: Debug, Info, and Fatal.

*   Debugm and Debugf only emit when the debug flag is set.

*   Fatalm and Fatalf call os.Exit(1) after emitting the associated message data.

*   Infom and Infof just emit the associated message data.

Example usage:

    import (
        "bytes"

        "github.com/cactus/mlog"
    )

    func main() {
        mlog.Infom("this is a log", mlog.Map{
            "interesting": "data",
            "something": 42,
        })

        mlog.Debugm("this won't print")

        // set flags for the default logger
        // alternatively, you can create your own logger
        // and supply flags at creation time
        mlog.SetFlags(mlog.Ldebug)

        mlog.Debugm("this will print!")

        mlog.Debugm("can it print?", mlog.Map{
            "how_fancy": []byte{'v', 'e', 'r', 'y', '!'},
            "this_too": bytes.NewBuffer([]byte("if fmt.Print can print it!")),
        })

        // you can use a more classical Printf type log method too.
        mlog.Debugf("a printf style debug log: %s", "here!")
        mlog.Infof("a printf style info log: %s", "here!")

        mlog.Fatalm("time for a nap", mlog.Map{"cleanup": false})
    }
*/
package mlog
