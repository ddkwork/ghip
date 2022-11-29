package main

import (
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/src/net/httpClient"
	"github.com/ddkwork/golibrary/src/stream"
	"github.com/ddkwork/golibrary/src/stream/tool"
)

// pkg install resolv-conf -y
func main() {
	uri := "https://hosts.gitcdn.top/hosts.txt"
	c := httpClient.New()
	if !c.Url(uri).Get().Request() {
		return
	}
	mylog.Json("", string(c.ResponseBuf()))
	hosts := stream.NewString(`
127.0.0.1 localhost
::1 ip6-localhost
`)
	hosts.NewLine()
	hosts.WriteBytesLn(c.ResponseBuf())
	path := "/data/data/com.termux/files/usr/etc/hosts"
	tool.File().WriteTruncate(path, hosts.String())
}
