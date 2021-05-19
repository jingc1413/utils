package utils

import (
	"net"
	"testing"
)

func TestGetNotInternalIPv4(t *testing.T) {
	defIP4 := "7.7.7.7"

	for _, v := range []struct {
		in   string
		out  string
		def  string
		flag bool
	}{
		{"1.2.3.4", "1.2.3.4", defIP4, false},
		{"1.2.3.4", "1.2.3.4", defIP4, true},
		{"10.0.0.1", defIP4, defIP4, false},
		{"10.0.0.1", defIP4, defIP4, true},
		{"100.125.1.1", defIP4, defIP4, false},
		{"100.125.1.1", defIP4, defIP4, true},
		{"127.0.0.1", defIP4, defIP4, false},
		{"127.0.0.1", defIP4, defIP4, true},
		{"169.254.1.1", defIP4, defIP4, false},
		{"169.254.1.1", defIP4, defIP4, true},
		{"192.168.1.1", defIP4, defIP4, false},
		{"192.168.1.1", defIP4, defIP4, true},
		{"192.168.1.1", "192.168.1.1", "", true},
	} {
		AssertEqual(t, v.out, GetNotInternalIPv4String(v.in, v.def, v.flag))
		AssertEqual(t, net.ParseIP(v.out), GetNotInternalIPv4(net.ParseIP(v.in), net.ParseIP(v.def), v.flag))
	}
}

func TestIPv4AndLong(t *testing.T) {
	for _, v := range []struct {
		ipv4 string
		long int
	}{
		{"0.0.0.0", 0},
		{"0.0.0.1", 1},
		{"1.2.3.4", 16909060},
		{"10.0.0.0", 167772160},
		{"255.255.255.255", 4294967295},
		{"", -1},
	} {
		AssertEqual(t, v.long, IPv4String2Long(v.ipv4))
		AssertEqual(t, v.ipv4, Long2IPv4String(v.long))
	}

	AssertEqual(t, 151060737, IPv4String2Long("009.001.01.1"))
	AssertEqual(t, -1, IPv4String2Long("ff"))
	AssertEqual(t, -1, IPv4String2Long("255.255.255.256"))
	AssertEqual(t, "", Long2IPv4String(4294967296))
}

func BenchmarkGetNotInternalIPv4String(b *testing.B) {
	cip := "100.125.1.1"
	fip := "1.1.1.1,2.2.2.2"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetNotInternalIPv4String(cip, fip)
	}
}

func BenchmarkGetNotInternalIPv4StringTrue(b *testing.B) {
	cip := "100.125.1.1"
	fip := ""
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GetNotInternalIPv4String(cip, fip, true)
	}
}

// BenchmarkGetNotInternalIPv4String-8       	 7723636	       143.0 ns/op	      16 B/op	       1 allocs/op
// BenchmarkGetNotInternalIPv4String-8       	 8165636	       135.1 ns/op	      16 B/op	       1 allocs/op
// BenchmarkGetNotInternalIPv4String-8       	 8405780	       148.5 ns/op	      16 B/op	       1 allocs/op
// BenchmarkGetNotInternalIPv4StringTrue-8   	 9055933	       145.5 ns/op	      16 B/op	       1 allocs/op
// BenchmarkGetNotInternalIPv4StringTrue-8   	 7916724	       206.0 ns/op	      16 B/op	       1 allocs/op
// BenchmarkGetNotInternalIPv4StringTrue-8   	 7862268	       151.3 ns/op	      16 B/op	       1 allocs/op
