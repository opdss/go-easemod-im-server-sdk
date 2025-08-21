package utils

import (
	"net"
	"net/url"
	"os"
	"syscall"
)

func PointerAny[T any](v T) *T {
	r := new(T)
	*r = v
	return r
}

// IsNetError 需要切换域名的网络错误
func IsNetError(err error) bool {
	netErr, ok := err.(net.Error)
	if !ok {
		return false
	}
	// 超时
	if netErr.Timeout() {
		return true
	}

	var opErr *net.OpError
	opErr, ok = netErr.(*net.OpError)
	if !ok {
		//  url 错误
		urlErr, ok := netErr.(*url.Error)
		if !ok {
			return false
		}
		opErr, ok = urlErr.Err.(*net.OpError)
		if !ok {
			return false
		}
	}

	switch t := opErr.Err.(type) {
	case *net.DNSError:
		return true
	case *os.SyscallError:
		if errno, ok := t.Err.(syscall.Errno); ok {
			switch errno {
			case syscall.ECONNREFUSED:
				return true
			case syscall.ETIMEDOUT:
				return true
			}
		}
	}

	return false
}
