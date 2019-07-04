# yaegi-talk
talks and docs about yaegi

## Install

See the slides and play the examples on your machine (not tested on Windows)

Prerequisite: Go is installed, default GOPATH in `~/go`

Make sure that [present](https://godoc.org/golang.org/x/tools/cmd/present)
tool is installed:

```console
$ go get -u golang.org/x/tools/cmd/present
```

Clone [yaegi](https://github.com/containous/yaegi/) and
[yaegi-talk](https://github.com/containous/yaegi-talk) at the proper
place:

```console
$ mkdir -p ~/go/src/github.com/containous
$ cd ~/go/src/github.com/containous
$ git clone containous/yaegi
$ git clone containous/yaegi-talk
```

See talks with runnable/editable yaegi examples:

```console
$ cd ~/go/src/github.com/containous/yaegi-talk
$ ~/go/bin/present
```

Then open a web browser and open the link displayed, normally
http://127.0.0.1:3999/.

Runnable examples have a small `run` button
at the bottom right of code text area. Code text areas should be
editable from the browser.

