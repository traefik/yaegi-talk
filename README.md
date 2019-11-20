# yaegi-talk
talks and docs about yaegi

* [GopherCon-2019] (lightning talk)
* [GoLab-2019] (30 min. talk)
* [Toulouse Software Crafters Meetup, 19-nov-2019] (20 min, same as GoLab)

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

[GopherCon-2019]: https://talks.godoc.org/github.com/containous/yaegi-talk/GopherCon-2019/lightning-talk.slide
[GoLab-2019]: https://talks.godoc.org/github.com/containous/yaegi-talk/GoLab-2019/30min-talk.slide
[Toulouse Software Crafters Meetup, 19-nov-2019]: https://talks.godoc.org/github.com/containous/yaegi-talk/Meetup-software-crafters-Toulouse/talk.slide
