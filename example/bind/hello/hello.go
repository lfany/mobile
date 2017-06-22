// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hello is a trivial package for gomobile bind example.
package hello

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"os"
	"log"
	"strings"
)

func init() {
	raven.SetDSN("http://d9c084a8db004c299c6ba7b63f29265f:e491bf7bb9204bd5a0ccae3f48bcd21f@sentry.ifnot.cc:9000/3")
}

func Greetings(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func Title(name string) string {
	return fmt.Sprintf("Go: %s", name)
}

func ReportError(what string) string {
	_, err := os.Open("test.txt")
	if err != nil {
		msg := raven.CaptureErrorAndWait(err, map[string]string{"what": what})
		log.Panic(err)
		return msg
	}
	return ""
}

func ReportPainc(what string) string {
	_, id := raven.CapturePanic(func() {
		//	 do something
		panic(nil)
	}, nil)

	id2 := raven.CaptureMessage("msg from gomobile!"+what, map[string]string{"env": strings.Join(os.Environ(), "-")})

	if len(id) != 0 && len(id2) != 0 {
		return id + "-" + id2
	} else if len(id) == 0 {
		return id2
	} else if len(id2) == 0 {
		return id
	} else {
		return id
	}
}
