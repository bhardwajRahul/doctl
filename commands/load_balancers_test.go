package commands

import (
	"testing"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/doctl/do"
	"github.com/digitalocean/godo"

	"github.com/stretchr/testify/assert"
)

var (
	testLoadBalancer = do.LoadBalancer{
		LoadBalancer: &godo.LoadBalancer{
			Algorithm: "round_robin",
			Region: &godo.Region{
				Slug: "nyc1",
			},
			SizeSlug:       "lb-small",
			StickySessions: &godo.StickySessions{},
			HealthCheck:    &godo.HealthCheck{},
		}}

	testLoadBalancerList = do.LoadBalancers{
		testLoadBalancer,
	}
)

func TestLoadBalancerCommand(t *testing.T) {
	cmd := LoadBalancer()
	assert.NotNil(t, cmd)
	assertCommandNames(t, cmd, "get", "list", "create", "update", "delete", "add-droplets", "remove-droplets", "add-forwarding-rules", "remove-forwarding-rules", "purge-cache")
}

func TestLoadBalancerGet(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		tm.loadBalancers.EXPECT().Get(lbID).Return(&testLoadBalancer, nil)

		config.Args = append(config.Args, lbID)

		err := RunLoadBalancerGet(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerGetNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerGet(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerList(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		tm.loadBalancers.EXPECT().List().Return(testLoadBalancerList, nil)

		err := RunLoadBalancerList(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerCreateWithInvalidDropletIDsArgs(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		config.Doit.Set(config.NS, doctl.ArgDropletIDs, []string{"bogus"})

		err := RunLoadBalancerCreate(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerCreateWithMalformedForwardingRulesArgs(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		config.Doit.Set(config.NS, doctl.ArgForwardingRules, "something,something")

		err := RunLoadBalancerCreate(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerCreate(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		vpcUUID := "00000000-0000-4000-8000-000000000000"
		timeout := uint64(120)
		r := godo.LoadBalancerRequest{
			Name:       "lb-name",
			Region:     "nyc1",
			SizeSlug:   "lb-small",
			Type:       "REGIONAL",
			DropletIDs: []int{1, 2},
			StickySessions: &godo.StickySessions{
				Type: "none",
			},
			HealthCheck: &godo.HealthCheck{
				Protocol:               "http",
				Port:                   80,
				CheckIntervalSeconds:   4,
				ResponseTimeoutSeconds: 23,
				HealthyThreshold:       5,
				UnhealthyThreshold:     10,
			},
			ForwardingRules: []godo.ForwardingRule{
				{
					EntryProtocol:  "tcp",
					EntryPort:      3306,
					TargetProtocol: "tcp",
					TargetPort:     3306,
					TlsPassthrough: true,
				},
			},
			VPCUUID:                vpcUUID,
			ProjectID:              "project-id-uuid",
			HTTPIdleTimeoutSeconds: &timeout,
			Firewall: &godo.LBFirewall{
				Deny:  []string{"cidr:1.2.0.0/16"},
				Allow: []string{"ip:1.2.3.4", "ip:1.2.3.5"},
			},
			Network:         "EXTERNAL",
			NetworkStack:    "IPV4",
			TLSCipherPolicy: "STRONG",
		}
		disableLetsEncryptDNSRecords := true
		r.DisableLetsEncryptDNSRecords = &disableLetsEncryptDNSRecords
		tm.loadBalancers.EXPECT().Create(&r).Return(&testLoadBalancer, nil)

		config.Doit.Set(config.NS, doctl.ArgRegionSlug, "nyc1")
		config.Doit.Set(config.NS, doctl.ArgSizeSlug, "lb-small")
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerName, "lb-name")
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerType, "REGIONAL")
		config.Doit.Set(config.NS, doctl.ArgVPCUUID, vpcUUID)
		config.Doit.Set(config.NS, doctl.ArgDropletIDs, []string{"1", "2"})
		config.Doit.Set(config.NS, doctl.ArgStickySessions, "type:none")
		config.Doit.Set(config.NS, doctl.ArgHealthCheck, "protocol:http,port:80,check_interval_seconds:4,response_timeout_seconds:23,healthy_threshold:5,unhealthy_threshold:10")
		config.Doit.Set(config.NS, doctl.ArgForwardingRules, "entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306,tls_passthrough:true")
		config.Doit.Set(config.NS, doctl.ArgDisableLetsEncryptDNSRecords, true)
		config.Doit.Set(config.NS, doctl.ArgProjectID, "project-id-uuid")
		config.Doit.Set(config.NS, doctl.ArgHTTPIdleTimeoutSeconds, 120)
		config.Doit.Set(config.NS, doctl.ArgDenyList, []string{"cidr:1.2.0.0/16"})
		config.Doit.Set(config.NS, doctl.ArgAllowList, []string{"ip:1.2.3.4", "ip:1.2.3.5"})
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerNetwork, "EXTERNAL")
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerNetworkStack, "IPV4")
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerTLSCipherPolicy, "STRONG")

		err := RunLoadBalancerCreate(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerCreateGLB(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		r := godo.LoadBalancerRequest{
			Name: "lb-name",
			Type: "GLOBAL",
			StickySessions: &godo.StickySessions{
				Type: "none",
			},
			HealthCheck: &godo.HealthCheck{
				Protocol:               "http",
				Port:                   80,
				Path:                   "/",
				CheckIntervalSeconds:   10,
				ResponseTimeoutSeconds: 5,
				HealthyThreshold:       5,
				UnhealthyThreshold:     3,
			},
			GLBSettings: &godo.GLBSettings{
				TargetProtocol:    "http",
				TargetPort:        80,
				CDN:               &godo.CDNSettings{IsEnabled: true},
				RegionPriorities:  map[string]uint32{"nyc1": 1, "nyc2": 2},
				FailoverThreshold: 10,
			},
			DropletIDs: []int{1, 2},
			Domains: []*godo.LBDomain{
				{
					Name:          "test-domain-1",
					IsManaged:     true,
					CertificateID: "test-cert-id-1",
				},
				{
					Name:          "test-domain-2",
					IsManaged:     false,
					CertificateID: "test-cert-id-2",
				},
			},
			TargetLoadBalancerIDs: []string{
				"019cb059-603f-4828-8be4-641a20f25006",
				"023da268-bc81-468f-aa4d-9abdc4f69935",
			},
			Network:      "EXTERNAL",
			NetworkStack: "IPV4",
		}
		disableLetsEncryptDNSRecords := true
		r.DisableLetsEncryptDNSRecords = &disableLetsEncryptDNSRecords
		tm.loadBalancers.EXPECT().Create(&r).Return(&testLoadBalancer, nil)

		config.Doit.Set(config.NS, doctl.ArgLoadBalancerName, "lb-name")
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerType, "GLOBAL")
		config.Doit.Set(config.NS, doctl.ArgStickySessions, "type:none")
		config.Doit.Set(config.NS, doctl.ArgHealthCheck, "protocol:http,port:80,path:/,check_interval_seconds:10,response_timeout_seconds:5,healthy_threshold:5,unhealthy_threshold:3")
		config.Doit.Set(config.NS, doctl.ArgGlobalLoadBalancerSettings, "target_protocol:http,target_port:80,region_priorities:nyc1=1 nyc2=2,failover_threshold:10")
		config.Doit.Set(config.NS, doctl.ArgGlobalLoadBalancerCDNSettings, "is_enabled:true")
		config.Doit.Set(config.NS, doctl.ArgDropletIDs, []string{"1", "2"})
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerDomains, []string{
			"name:test-domain-1 is_managed:true certificate_id:test-cert-id-1",
			"name:test-domain-2 is_managed:false certificate_id:test-cert-id-2",
		})
		config.Doit.Set(config.NS, doctl.ArgDisableLetsEncryptDNSRecords, true)
		config.Doit.Set(config.NS, doctl.ArgTargetLoadBalancerIDs, []string{
			"019cb059-603f-4828-8be4-641a20f25006",
			"023da268-bc81-468f-aa4d-9abdc4f69935",
		})
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerNetwork, "EXTERNAL")
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerNetworkStack, "IPV4")

		err := RunLoadBalancerCreate(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerUpdate(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		timeout := uint64(120)
		r := godo.LoadBalancerRequest{
			Name:       "lb-name",
			Region:     "nyc1",
			DropletIDs: []int{1, 2},
			SizeUnit:   4,
			StickySessions: &godo.StickySessions{
				Type:             "cookies",
				CookieName:       "DO-LB",
				CookieTtlSeconds: 5,
			},
			HealthCheck: &godo.HealthCheck{
				Protocol:               "http",
				Port:                   80,
				CheckIntervalSeconds:   4,
				ResponseTimeoutSeconds: 23,
				HealthyThreshold:       5,
				UnhealthyThreshold:     10,
			},
			ForwardingRules: []godo.ForwardingRule{
				{
					EntryProtocol:  "http",
					EntryPort:      80,
					TargetProtocol: "http",
					TargetPort:     80,
				},
			},
			ProjectID:              "project-id-uuid",
			HTTPIdleTimeoutSeconds: &timeout,
			Firewall: &godo.LBFirewall{
				Deny:  []string{"cidr:1.2.0.0/16"},
				Allow: []string{"ip:1.2.3.4", "ip:1.2.3.5"},
			},
		}
		disableLetsEncryptDNSRecords := true
		r.DisableLetsEncryptDNSRecords = &disableLetsEncryptDNSRecords
		tm.loadBalancers.EXPECT().Update(lbID, &r).Return(&testLoadBalancer, nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, doctl.ArgRegionSlug, "nyc1")
		config.Doit.Set(config.NS, doctl.ArgSizeSlug, "")
		config.Doit.Set(config.NS, doctl.ArgSizeUnit, 4)
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerName, "lb-name")
		config.Doit.Set(config.NS, doctl.ArgDropletIDs, []string{"1", "2"})
		config.Doit.Set(config.NS, doctl.ArgStickySessions, "type:cookies,cookie_name:DO-LB,cookie_ttl_seconds:5")
		config.Doit.Set(config.NS, doctl.ArgHealthCheck, "protocol:http,port:80,check_interval_seconds:4,response_timeout_seconds:23,healthy_threshold:5,unhealthy_threshold:10")
		config.Doit.Set(config.NS, doctl.ArgForwardingRules, "entry_protocol:http,entry_port:80,target_protocol:http,target_port:80")
		config.Doit.Set(config.NS, doctl.ArgDisableLetsEncryptDNSRecords, true)
		config.Doit.Set(config.NS, doctl.ArgProjectID, "project-id-uuid")
		config.Doit.Set(config.NS, doctl.ArgHTTPIdleTimeoutSeconds, 120)
		config.Doit.Set(config.NS, doctl.ArgDenyList, []string{"cidr:1.2.0.0/16"})
		config.Doit.Set(config.NS, doctl.ArgAllowList, []string{"ip:1.2.3.4", "ip:1.2.3.5"})

		err := RunLoadBalancerUpdate(config)
		assert.NoError(t, err)
	})
}
func TestLoadBalancerUpdateGLB(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		r := godo.LoadBalancerRequest{
			Name: "lb-name",
			Type: "GLOBAL",
			StickySessions: &godo.StickySessions{
				Type: "none",
			},
			HealthCheck: &godo.HealthCheck{
				Protocol:               "http",
				Port:                   80,
				Path:                   "/",
				CheckIntervalSeconds:   10,
				ResponseTimeoutSeconds: 5,
				HealthyThreshold:       5,
				UnhealthyThreshold:     3,
			},
			GLBSettings: &godo.GLBSettings{
				TargetProtocol:    "http",
				TargetPort:        80,
				CDN:               &godo.CDNSettings{IsEnabled: true},
				RegionPriorities:  map[string]uint32{"nyc1": 1, "nyc2": 2},
				FailoverThreshold: 10,
			},
			DropletIDs: []int{1, 2},
			Domains: []*godo.LBDomain{
				{
					Name:          "test-domain-1",
					IsManaged:     true,
					CertificateID: "test-cert-id-1",
				},
				{
					Name:          "test-domain-2",
					IsManaged:     false,
					CertificateID: "test-cert-id-2",
				},
			},
			TargetLoadBalancerIDs: []string{
				"019cb059-603f-4828-8be4-641a20f25006",
				"023da268-bc81-468f-aa4d-9abdc4f69935",
			},
		}
		disableLetsEncryptDNSRecords := true
		r.DisableLetsEncryptDNSRecords = &disableLetsEncryptDNSRecords
		tm.loadBalancers.EXPECT().Update(lbID, &r).Return(&testLoadBalancer, nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerName, "lb-name")
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerType, "GLOBAL")
		config.Doit.Set(config.NS, doctl.ArgStickySessions, "type:none")
		config.Doit.Set(config.NS, doctl.ArgHealthCheck, "protocol:http,port:80,path:/,check_interval_seconds:10,response_timeout_seconds:5,healthy_threshold:5,unhealthy_threshold:3")
		config.Doit.Set(config.NS, doctl.ArgGlobalLoadBalancerSettings, "target_protocol:http,target_port:80,region_priorities:nyc1=1 nyc2=2,failover_threshold:10")
		config.Doit.Set(config.NS, doctl.ArgGlobalLoadBalancerCDNSettings, "is_enabled:true")
		config.Doit.Set(config.NS, doctl.ArgDropletIDs, []string{"1", "2"})
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerDomains, []string{
			"name:test-domain-1 is_managed:true certificate_id:test-cert-id-1",
			"name:test-domain-2 is_managed:false certificate_id:test-cert-id-2",
		})
		config.Doit.Set(config.NS, doctl.ArgDisableLetsEncryptDNSRecords, true)
		config.Doit.Set(config.NS, doctl.ArgTargetLoadBalancerIDs, []string{
			"019cb059-603f-4828-8be4-641a20f25006",
			"023da268-bc81-468f-aa4d-9abdc4f69935",
		})

		err := RunLoadBalancerUpdate(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerUpdateNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerUpdate(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerUpdateNoFirewallFlags(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		timeout := uint64(120)
		r := godo.LoadBalancerRequest{
			Name:       "lb-name",
			Region:     "nyc1",
			DropletIDs: []int{1, 2},
			SizeUnit:   4,
			StickySessions: &godo.StickySessions{
				Type:             "cookies",
				CookieName:       "DO-LB",
				CookieTtlSeconds: 5,
			},
			HealthCheck: &godo.HealthCheck{
				Protocol:               "http",
				Port:                   80,
				CheckIntervalSeconds:   4,
				ResponseTimeoutSeconds: 23,
				HealthyThreshold:       5,
				UnhealthyThreshold:     10,
			},
			ForwardingRules: []godo.ForwardingRule{
				{
					EntryProtocol:  "http",
					EntryPort:      80,
					TargetProtocol: "http",
					TargetPort:     80,
				},
			},
			ProjectID:              "project-id-uuid",
			HTTPIdleTimeoutSeconds: &timeout,
			Firewall:               nil,
		}
		disableLetsEncryptDNSRecords := true
		r.DisableLetsEncryptDNSRecords = &disableLetsEncryptDNSRecords
		tm.loadBalancers.EXPECT().Update(lbID, &r).Return(&testLoadBalancer, nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, doctl.ArgRegionSlug, "nyc1")
		config.Doit.Set(config.NS, doctl.ArgSizeSlug, "")
		config.Doit.Set(config.NS, doctl.ArgSizeUnit, 4)
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerName, "lb-name")
		config.Doit.Set(config.NS, doctl.ArgDropletIDs, []string{"1", "2"})
		config.Doit.Set(config.NS, doctl.ArgStickySessions, "type:cookies,cookie_name:DO-LB,cookie_ttl_seconds:5")
		config.Doit.Set(config.NS, doctl.ArgHealthCheck, "protocol:http,port:80,check_interval_seconds:4,response_timeout_seconds:23,healthy_threshold:5,unhealthy_threshold:10")
		config.Doit.Set(config.NS, doctl.ArgForwardingRules, "entry_protocol:http,entry_port:80,target_protocol:http,target_port:80")
		config.Doit.Set(config.NS, doctl.ArgDisableLetsEncryptDNSRecords, true)
		config.Doit.Set(config.NS, doctl.ArgProjectID, "project-id-uuid")
		config.Doit.Set(config.NS, doctl.ArgHTTPIdleTimeoutSeconds, 120)

		err := RunLoadBalancerUpdate(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerUpdateEmptyFirewallFlags(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		timeout := uint64(120)
		r := godo.LoadBalancerRequest{
			Name:       "lb-name",
			Region:     "nyc1",
			DropletIDs: []int{1, 2},
			SizeUnit:   4,
			StickySessions: &godo.StickySessions{
				Type:             "cookies",
				CookieName:       "DO-LB",
				CookieTtlSeconds: 5,
			},
			HealthCheck: &godo.HealthCheck{
				Protocol:               "http",
				Port:                   80,
				CheckIntervalSeconds:   4,
				ResponseTimeoutSeconds: 23,
				HealthyThreshold:       5,
				UnhealthyThreshold:     10,
			},
			ForwardingRules: []godo.ForwardingRule{
				{
					EntryProtocol:  "http",
					EntryPort:      80,
					TargetProtocol: "http",
					TargetPort:     80,
				},
			},
			ProjectID:              "project-id-uuid",
			HTTPIdleTimeoutSeconds: &timeout,
			Firewall: &godo.LBFirewall{
				Deny:  []string{},
				Allow: []string{},
			},
		}
		disableLetsEncryptDNSRecords := true
		r.DisableLetsEncryptDNSRecords = &disableLetsEncryptDNSRecords
		tm.loadBalancers.EXPECT().Update(lbID, &r).Return(&testLoadBalancer, nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, doctl.ArgRegionSlug, "nyc1")
		config.Doit.Set(config.NS, doctl.ArgSizeSlug, "")
		config.Doit.Set(config.NS, doctl.ArgSizeUnit, 4)
		config.Doit.Set(config.NS, doctl.ArgLoadBalancerName, "lb-name")
		config.Doit.Set(config.NS, doctl.ArgDropletIDs, []string{"1", "2"})
		config.Doit.Set(config.NS, doctl.ArgStickySessions, "type:cookies,cookie_name:DO-LB,cookie_ttl_seconds:5")
		config.Doit.Set(config.NS, doctl.ArgHealthCheck, "protocol:http,port:80,check_interval_seconds:4,response_timeout_seconds:23,healthy_threshold:5,unhealthy_threshold:10")
		config.Doit.Set(config.NS, doctl.ArgForwardingRules, "entry_protocol:http,entry_port:80,target_protocol:http,target_port:80")
		config.Doit.Set(config.NS, doctl.ArgDisableLetsEncryptDNSRecords, true)
		config.Doit.Set(config.NS, doctl.ArgProjectID, "project-id-uuid")
		config.Doit.Set(config.NS, doctl.ArgHTTPIdleTimeoutSeconds, 120)
		config.Doit.Set(config.NS, doctl.ArgDenyList, []string{})
		config.Doit.Set(config.NS, doctl.ArgAllowList, []string{})

		err := RunLoadBalancerUpdate(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerDelete(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		tm.loadBalancers.EXPECT().Delete(lbID).Return(nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, doctl.ArgForce, true)

		err := RunLoadBalancerDelete(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerPurgeCache(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		tm.loadBalancers.EXPECT().PurgeCache(lbID).Return(nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, doctl.ArgForce, true)

		err := RunLoadBalancerPurgeCache(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerDeleteNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerDelete(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerAddDroplets(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		tm.loadBalancers.EXPECT().AddDroplets(lbID, 1, 23).Return(nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, doctl.ArgDropletIDs, []string{"1", "23"})

		err := RunLoadBalancerAddDroplets(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerAddDropletsNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerAddDroplets(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerRemoveDroplets(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		tm.loadBalancers.EXPECT().RemoveDroplets(lbID, 321).Return(nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, doctl.ArgDropletIDs, []string{"321"})

		err := RunLoadBalancerRemoveDroplets(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerRemoveDropletsNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerRemoveDroplets(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerAddForwardingRules(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		forwardingRule := godo.ForwardingRule{
			EntryProtocol:  "http",
			EntryPort:      80,
			TargetProtocol: "http",
			TargetPort:     80,
		}
		tm.loadBalancers.EXPECT().AddForwardingRules(lbID, forwardingRule).Return(nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, doctl.ArgForwardingRules, "entry_protocol:http,entry_port:80,target_protocol:http,target_port:80")

		err := RunLoadBalancerAddForwardingRules(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerAddForwardingRulesNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerAddForwardingRules(config)
		assert.Error(t, err)
	})
}

func TestLoadBalancerRemoveForwardingRules(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		lbID := "cde2c0d6-41e3-479e-ba60-ad971227232c"
		forwardingRules := []godo.ForwardingRule{
			{
				EntryProtocol:  "http",
				EntryPort:      80,
				TargetProtocol: "http",
				TargetPort:     80,
			},
			{
				EntryProtocol:  "tcp",
				EntryPort:      3306,
				TargetProtocol: "tcp",
				TargetPort:     3306,
				TlsPassthrough: true,
			},
		}
		tm.loadBalancers.EXPECT().RemoveForwardingRules(lbID, forwardingRules[0], forwardingRules[1]).Return(nil)

		config.Args = append(config.Args, lbID)
		config.Doit.Set(config.NS, doctl.ArgForwardingRules, "entry_protocol:http,entry_port:80,target_protocol:http,target_port:80 entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306,tls_passthrough:true")

		err := RunLoadBalancerRemoveForwardingRules(config)
		assert.NoError(t, err)
	})
}

func TestLoadBalancerRemoveForwardingRulesNoID(t *testing.T) {
	withTestClient(t, func(config *CmdConfig, tm *tcMocks) {
		err := RunLoadBalancerRemoveForwardingRules(config)
		assert.Error(t, err)
	})
}
