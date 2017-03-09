# logrus error caller plugin

logrus log show caller

## Requirements

* [logrus](https://github.com/sirupsen/logrus)


## Install

`go get github.com/syhlion/greq`

## Usage

```
l := logrus.New()
l.AddHook(&logcaller.CallerHook{})
l.Warn("test")

//if you log it can stdout like this
//time="2017-02-10T14:26:19+08:00" level=warning msg="test" caller="api.go:200"

```
