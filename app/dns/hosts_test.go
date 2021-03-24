package dns_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	. "github.com/xtls/xray-core/app/dns"
	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/matcher/domain"
	"github.com/xtls/xray-core/common/net"
	"github.com/xtls/xray-core/features/dns"
)

func TestStaticHosts(t *testing.T) {
	pb := []*Config_HostMapping{
		{
			Type:   domain.MatchingType_Full,
			Domain: "example.com",
			Ip: [][]byte{
				{1, 1, 1, 1},
			},
		},
		{
			Type:   domain.MatchingType_Subdomain,
			Domain: "example.cn",
			Ip: [][]byte{
				{2, 2, 2, 2},
			},
		},
		{
			Type:   domain.MatchingType_Subdomain,
			Domain: "baidu.com",
			Ip: [][]byte{
				{127, 0, 0, 1},
			},
		},
	}

	hosts, err := NewStaticHosts(pb, nil)
	common.Must(err)

	{
		ips := hosts.Lookup("example.com", dns.IPOption{
			IPv4Enable: true,
			IPv6Enable: true,
		})
		if len(ips) != 1 {
			t.Error("expect 1 IP, but got ", len(ips))
		}
		if diff := cmp.Diff([]byte(ips[0].IP()), []byte{1, 1, 1, 1}); diff != "" {
			t.Error(diff)
		}
	}

	{
		ips := hosts.Lookup("www.example.cn", dns.IPOption{
			IPv4Enable: true,
			IPv6Enable: true,
		})
		if len(ips) != 1 {
			t.Error("expect 1 IP, but got ", len(ips))
		}
		if diff := cmp.Diff([]byte(ips[0].IP()), []byte{2, 2, 2, 2}); diff != "" {
			t.Error(diff)
		}
	}

	{
		ips := hosts.Lookup("baidu.com", dns.IPOption{
			IPv4Enable: false,
			IPv6Enable: true,
		})
		if len(ips) != 1 {
			t.Error("expect 1 IP, but got ", len(ips))
		}
		if diff := cmp.Diff([]byte(ips[0].IP()), []byte(net.LocalHostIPv6.IP())); diff != "" {
			t.Error(diff)
		}
	}
}
