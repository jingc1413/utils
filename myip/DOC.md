<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# myip

```go
import "github.com/fufuok/utils/myip"
```

## Index

- [func ExternalIP(v ...string) string](<#func-externalip>)
- [func ExternalIPAny(retries ...int) string](<#func-externalipany>)
- [func ExternalIPv4() string](<#func-externalipv4>)
- [func ExternalIPv6() string](<#func-externalipv6>)
- [func InternalIP(dstAddr, network string) string](<#func-internalip>)
- [func InternalIPv4() string](<#func-internalipv4>)
- [func InternalIPv6() string](<#func-internalipv6>)
- [func LocalIP() string](<#func-localip>)
- [func LocalIPv4s() (ips []string)](<#func-localipv4s>)


## func [ExternalIP](<https://gitee.com/fufuok/utils/blob/master/myip/myip.go#L52>)

```go
func ExternalIP(v ...string) string
```

ExternalIP 获取外网地址 \(出口公网地址\)

## func [ExternalIPAny](<https://gitee.com/fufuok/utils/blob/master/myip/myip.go#L30>)

```go
func ExternalIPAny(retries ...int) string
```

ExternalIPAny 获取外网地址

## func [ExternalIPv4](<https://gitee.com/fufuok/utils/blob/master/myip/myip.go#L61>)

```go
func ExternalIPv4() string
```

ExternalIPv4 获取外网地址 \(IPv4\)

## func [ExternalIPv6](<https://gitee.com/fufuok/utils/blob/master/myip/myip.go#L66>)

```go
func ExternalIPv6() string
```

ExternalIPv6 获取外网地址 \(IPv6\)

## func [InternalIP](<https://gitee.com/fufuok/utils/blob/master/myip/myip.go#L127>)

```go
func InternalIP(dstAddr, network string) string
```

InternalIP 获取内网地址 \(出口本地地址\)

## func [InternalIPv4](<https://gitee.com/fufuok/utils/blob/master/myip/myip.go#L117>)

```go
func InternalIPv4() string
```

InternalIPv4 获取内网地址 \(IPv4\)

## func [InternalIPv6](<https://gitee.com/fufuok/utils/blob/master/myip/myip.go#L122>)

```go
func InternalIPv6() string
```

InternalIPv6 获取内网地址 \(临时 IPv6 地址\)

## func [LocalIP](<https://gitee.com/fufuok/utils/blob/master/myip/myip.go#L154>)

```go
func LocalIP() string
```

LocalIP 获取本地地址 \(第一个\)

## func [LocalIPv4s](<https://gitee.com/fufuok/utils/blob/master/myip/myip.go#L169>)

```go
func LocalIPv4s() (ips []string)
```

LocalIPv4s 获取所有本地地址 IPv4



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)