/*
Copyright 2018 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
	http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	_ "embed"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/doctl/commands/displayers"
	"github.com/digitalocean/doctl/do"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

var (
	//go:embed forwarding_detail.txt
	forwardingDetail string

	//go:embed lb_detail.txt
	lbDetail string
)

// LoadBalancer creates the load balancer command.
func LoadBalancer() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:     "load-balancer",
			Short:   "Display commands to manage load balancers",
			Aliases: []string{"lb"},
			Long: `The sub-commands of ` + "`" + `doctl compute load-balancer` + "`" + ` manage your load balancers.

With the load-balancer command, you can list, create, or delete load balancers, and manage their configuration details.`,
		},
	}

	forwardingRulesTxt := "A comma-separated list of key-value pairs representing forwarding rules, which define how traffic is routed, e.g.: `entry_protocol:tcp,entry_port:3306,target_protocol:tcp,target_port:3306`."
	CmdBuilder(cmd, RunLoadBalancerGet, "get <load-balancer-id>", "Retrieve a load balancer", "Use this command to retrieve information about a load balancer instance, including:\n\n"+lbDetail, Writer,
		aliasOpt("g"), displayerType(&displayers.LoadBalancer{}))

	cmdLoadBalancerCreate := CmdBuilder(cmd, RunLoadBalancerCreate, "create",
		"Create a new load balancer", "Use this command to create a new load balancer on your account. Valid forwarding rules are:\n"+forwardingDetail, Writer, aliasOpt("c"))
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgLoadBalancerName, "", "",
		"The load balancer's name", requiredOpt())
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgRegionSlug, "", "",
		"The load balancer's region, e.g.: `nyc1`")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgSizeSlug, "", "",
		fmt.Sprintf("The load balancer's size, e.g.: `lb-small`. Only one of %s and %s should be used", doctl.ArgSizeSlug, doctl.ArgSizeUnit))
	AddIntFlag(cmdLoadBalancerCreate, doctl.ArgSizeUnit, "", 0,
		fmt.Sprintf("The load balancer's size, e.g.: 1. Only one of %s and %s should be used", doctl.ArgSizeUnit, doctl.ArgSizeSlug))
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgLoadBalancerType, "", "", "The type of load balancer, e.g.: `REGIONAL` or `GLOBAL`")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgVPCUUID, "", "", "The UUID of the VPC to create the load balancer in")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgLoadBalancerAlgorithm, "",
		"round_robin", "This field has been deprecated. You can no longer specify an algorithm for load balancers.")
	AddBoolFlag(cmdLoadBalancerCreate, doctl.ArgRedirectHTTPToHTTPS, "", false,
		"Redirects HTTP requests to the load balancer on port 80 to HTTPS on port 443")
	AddBoolFlag(cmdLoadBalancerCreate, doctl.ArgEnableProxyProtocol, "", false,
		"enable proxy protocol")
	AddBoolFlag(cmdLoadBalancerCreate, doctl.ArgEnableBackendKeepalive, "", false,
		"enable keepalive connections to backend target droplets")
	AddBoolFlag(cmdLoadBalancerCreate, doctl.ArgDisableLetsEncryptDNSRecords, "", false,
		"disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgTagName, "", "", "The name of a tag. All Droplets with this tag applied will be assigned to the load balancer.")
	AddStringSliceFlag(cmdLoadBalancerCreate, doctl.ArgDropletIDs, "", []string{},
		"A comma-separated list of Droplet IDs to add to the load balancer, e.g.: `12,33`")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgStickySessions, "", "",
		"A comma-separated list of key-value pairs representing a list of active sessions, e.g.: `type:cookies, cookie_name:DO-LB, cookie_ttl_seconds:5`")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgHealthCheck, "", "",
		"A comma-separated list of key-value pairs representing recent health check results, e.g.: `protocol:http,port:80,path:/index.html,check_interval_seconds:10,response_timeout_seconds:5,healthy_threshold:5,unhealthy_threshold:3`")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgForwardingRules, "", "",
		forwardingRulesTxt)
	AddBoolFlag(cmdLoadBalancerCreate, doctl.ArgCommandWait, "", false, "Boolean that specifies whether to wait for a load balancer to complete before returning control to the terminal")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgProjectID, "", "", "Indicates which project to associate the Load Balancer with. If not specified, the Load Balancer will be placed in your default project.")
	AddIntFlag(cmdLoadBalancerCreate, doctl.ArgHTTPIdleTimeoutSeconds, "", 0, "HTTP idle timeout that configures the idle timeout for http connections on the load balancer")
	AddStringSliceFlag(cmdLoadBalancerCreate, doctl.ArgAllowList, "", []string{},
		"A comma-separated list of ALLOW rules for the load balancer, e.g.: `ip:1.2.3.4,cidr:1.2.0.0/16`")
	AddStringSliceFlag(cmdLoadBalancerCreate, doctl.ArgDenyList, "", []string{},
		"A comma-separated list of DENY rules for the load balancer, e.g.: `ip:1.2.3.4,cidr:1.2.0.0/16`")
	AddStringSliceFlag(cmdLoadBalancerCreate, doctl.ArgLoadBalancerDomains, "", []string{},
		"A comma-separated list of domains required to ingress traffic to a global load balancer, e.g.: `name:test-domain-1 is_managed:true certificate_id:test-cert-id-1` ")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgGlobalLoadBalancerSettings, "", "",
		"Target protocol and port settings for ingressing traffic to a global load balancer, e.g.: `target_protocol:http,target_port:80` ")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgGlobalLoadBalancerCDNSettings, "", "",
		"CDN cache settings global load balancer, e.g.: `is_enabled:true` ")
	AddStringSliceFlag(cmdLoadBalancerCreate, doctl.ArgTargetLoadBalancerIDs, "", []string{},
		"A comma-separated list of Load Balancer IDs to add as target to the global load balancer ")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgLoadBalancerNetwork, "", "", "The type of network the load balancer is accessible from, e.g.: `EXTERNAL` or `INTERNAL`")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgLoadBalancerNetworkStack, "", "", "The network stack type determines the allocation of ipv4/ipv6 addresses to the load balancer, e.g.: `IPV4` or `DUALSTACK`"+
		" (NOTE: this feature is in private preview, contact DigitalOcean support to review its public availability.)")
	AddStringFlag(cmdLoadBalancerCreate, doctl.ArgLoadBalancerTLSCipherPolicy, "", "", "The tls cipher policy to be used for the load balancer, e.g.: `DEFAULT` or `STRONG`")

	cmdRecordUpdate := CmdBuilder(cmd, RunLoadBalancerUpdate, "update <load-balancer-id>",
		"Update a load balancer's configuration", `Use this command to update the configuration of a specified load balancer. Using all applicable flags, the command should contain a full representation of the load balancer including existing attributes, such as the load balancer's name, region, forwarding rules, and Droplet IDs. Any attribute that is not provided is reset to its default value.`, Writer, aliasOpt("u"))
	AddStringFlag(cmdRecordUpdate, doctl.ArgLoadBalancerName, "", "",
		"The load balancer's name")
	AddStringFlag(cmdRecordUpdate, doctl.ArgRegionSlug, "", "",
		"The load balancer's region, e.g.: `nyc1`")
	AddStringFlag(cmdRecordUpdate, doctl.ArgSizeSlug, "", "",
		fmt.Sprintf("The load balancer's size, e.g.: `lb-small`. Only one of %s and %s should be used", doctl.ArgSizeSlug, doctl.ArgSizeUnit))
	AddIntFlag(cmdRecordUpdate, doctl.ArgSizeUnit, "", 0,
		fmt.Sprintf("The load balancer's size, e.g.: 1. Only one of %s and %s should be used", doctl.ArgSizeUnit, doctl.ArgSizeSlug))
	AddStringFlag(cmdRecordUpdate, doctl.ArgVPCUUID, "", "", "The UUID of the VPC to create the load balancer in")
	AddStringFlag(cmdRecordUpdate, doctl.ArgLoadBalancerAlgorithm, "",
		"round_robin", "This field has been deprecated. You can no longer specify an algorithm for load balancers.")
	AddBoolFlag(cmdRecordUpdate, doctl.ArgRedirectHTTPToHTTPS, "", false,
		"Flag to redirect HTTP requests to the load balancer on port 80 to HTTPS on port 443")
	AddBoolFlag(cmdRecordUpdate, doctl.ArgEnableProxyProtocol, "", false,
		"enable proxy protocol")
	AddBoolFlag(cmdRecordUpdate, doctl.ArgEnableBackendKeepalive, "", false,
		"enable keepalive connections to backend target droplets")
	AddStringFlag(cmdRecordUpdate, doctl.ArgTagName, "", "", "Assigns Droplets with the specified tag to the load balancer")
	AddStringSliceFlag(cmdRecordUpdate, doctl.ArgDropletIDs, "", []string{},
		"A comma-separated list of Droplet IDs, e.g.: `215,378`")
	AddStringFlag(cmdRecordUpdate, doctl.ArgStickySessions, "", "",
		"A comma-separated list of key-value pairs representing a list of active sessions, e.g.: `type:cookies, cookie_name:DO-LB, cookie_ttl_seconds:5`")
	AddStringFlag(cmdRecordUpdate, doctl.ArgHealthCheck, "", "",
		"A comma-separated list of key-value pairs representing recent health check results, e.g.: `protocol:http, port:80, path:/index.html, check_interval_seconds:10, response_timeout_seconds:5, healthy_threshold:5, unhealthy_threshold:3`")
	AddStringFlag(cmdRecordUpdate, doctl.ArgForwardingRules, "", "", forwardingRulesTxt)
	AddBoolFlag(cmdRecordUpdate, doctl.ArgDisableLetsEncryptDNSRecords, "", false,
		"disable automatic DNS record creation for Let's Encrypt certificates that are added to the load balancer")
	AddStringFlag(cmdRecordUpdate, doctl.ArgProjectID, "", "",
		"Indicates which project to associate the Load Balancer with. If not specified, the Load Balancer will be placed in your default project.")
	AddStringSliceFlag(cmdRecordUpdate, doctl.ArgAllowList, "", nil,
		"A comma-separated list of ALLOW rules for the load balancer, e.g.: `ip:1.2.3.4,cidr:1.2.0.0/16`")
	AddStringSliceFlag(cmdRecordUpdate, doctl.ArgDenyList, "", nil,
		"A comma-separated list of DENY rules for the load balancer, e.g.: `ip:1.2.3.4,cidr:1.2.0.0/16`")
	AddStringSliceFlag(cmdRecordUpdate, doctl.ArgLoadBalancerDomains, "", []string{},
		"A comma-separated list of domains required to ingress traffic to a global load balancer, e.g.: `name:test-domain-1 is_managed:true certificate_id:test-cert-id-1` ")
	AddStringFlag(cmdRecordUpdate, doctl.ArgGlobalLoadBalancerSettings, "", "",
		"Target protocol and port settings for ingressing traffic to a global load balancer, e.g.: `target_protocol:http,target_port:80` ")
	AddStringFlag(cmdRecordUpdate, doctl.ArgGlobalLoadBalancerCDNSettings, "", "",
		"CDN cache settings global load balancer, e.g.: `is_enabled:true` ")
	AddStringSliceFlag(cmdRecordUpdate, doctl.ArgTargetLoadBalancerIDs, "", []string{},
		"A comma-separated list of Load Balancer IDs to add as target to the global load balancer ")
	AddStringFlag(cmdRecordUpdate, doctl.ArgLoadBalancerTLSCipherPolicy, "", "", "The tls cipher policy to be used for the load balancer, e.g.: `DEFAULT` or `STRONG`")

	CmdBuilder(cmd, RunLoadBalancerList, "list", "List load balancers", "Use this command to get a list of the load balancers on your account, including the following information for each:\n\n"+lbDetail, Writer,
		aliasOpt("ls"), displayerType(&displayers.LoadBalancer{}))

	cmdRunRecordDelete := CmdBuilder(cmd, RunLoadBalancerDelete, "delete <load-balancer-id>",
		"Permanently delete a load balancer", `Use this command to permanently delete the specified load balancer. This is irreversible.`, Writer, aliasOpt("d", "rm"))
	AddBoolFlag(cmdRunRecordDelete, doctl.ArgForce, doctl.ArgShortForce, false,
		"Delete the load balancer without a confirmation prompt")

	cmdAddDroplets := CmdBuilder(cmd, RunLoadBalancerAddDroplets, "add-droplets <load-balancer-id>",
		"Add Droplets to a load balancer", `Use this command to add Droplets to a load balancer.`, Writer)
	AddStringSliceFlag(cmdAddDroplets, doctl.ArgDropletIDs, "", []string{},
		"A comma-separated list of IDs of Droplet to add to the load balancer, example value: `12,33`")

	cmdRemoveDroplets := CmdBuilder(cmd, RunLoadBalancerRemoveDroplets,
		"remove-droplets <load-balancer-id>", "Remove Droplets from a load balancer", `Use this command to remove Droplets from a load balancer. This command does not destroy any Droplets.`, Writer)
	AddStringSliceFlag(cmdRemoveDroplets, doctl.ArgDropletIDs, "", []string{},
		"A comma-separated list of IDs of Droplets to remove from the load balancer, example value: `12,33`")

	cmdAddForwardingRules := CmdBuilder(cmd, RunLoadBalancerAddForwardingRules,
		"add-forwarding-rules <load-balancer-id>", "Add forwarding rules to a load balancer", "Use this command to add forwarding rules to a load balancer, specified with the `--forwarding-rules` flag. Valid rules include:\n"+forwardingDetail, Writer)
	AddStringFlag(cmdAddForwardingRules, doctl.ArgForwardingRules, "", "", forwardingRulesTxt)

	cmdRemoveForwardingRules := CmdBuilder(cmd, RunLoadBalancerRemoveForwardingRules,
		"remove-forwarding-rules <load-balancer-id>", "Remove forwarding rules from a load balancer", "Use this command to remove forwarding rules from a load balancer, specified with the `--forwarding-rules` flag. Valid rules include:\n"+forwardingDetail, Writer)
	AddStringFlag(cmdRemoveForwardingRules, doctl.ArgForwardingRules, "", "", forwardingRulesTxt)

	cmdRunCachePurge := CmdBuilder(cmd, RunLoadBalancerPurgeCache, "purge-cache <load-balancer-id>",
		"Purges CDN cache for a global load balancer", `Use this command to purge the CDN cache for specified global load balancer.`, Writer)
	AddBoolFlag(cmdRunCachePurge, doctl.ArgForce, doctl.ArgShortForce, false,
		"Purge the global load balancer CDN cache without a confirmation prompt ")

	return cmd
}

// RunLoadBalancerGet retrieves an existing load balancer by its identifier.
func RunLoadBalancerGet(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	id := c.Args[0]

	lbs := c.LoadBalancers()
	lb, err := lbs.Get(id)
	if err != nil {
		return err
	}

	item := &displayers.LoadBalancer{LoadBalancers: do.LoadBalancers{*lb}}
	return c.Display(item)
}

// RunLoadBalancerList lists load balancers.
func RunLoadBalancerList(c *CmdConfig) error {
	lbs := c.LoadBalancers()
	list, err := lbs.List()
	if err != nil {
		return err
	}

	item := &displayers.LoadBalancer{LoadBalancers: list}
	return c.Display(item)
}

// RunLoadBalancerCreate creates a new load balancer with a given configuration.
func RunLoadBalancerCreate(c *CmdConfig) error {
	r := new(godo.LoadBalancerRequest)
	if err := buildRequestFromArgs(c, r); err != nil {
		return err
	}

	lbs := c.LoadBalancers()
	lb, err := lbs.Create(r)
	if err != nil {
		return err
	}

	wait, err := c.Doit.GetBool(c.NS, doctl.ArgCommandWait)
	if err != nil {
		return err
	}

	if wait {
		lbs := c.LoadBalancers()
		notice("Load balancer creation is in progress, waiting for load balancer to become active")

		err := waitForActiveLoadBalancer(lbs, lb.ID)
		if err != nil {
			return fmt.Errorf(
				"load balancer couldn't enter `active` state: %v",
				err,
			)
		}

		lb, _ = lbs.Get(lb.ID)
	}

	notice("Load balancer created")

	item := &displayers.LoadBalancer{LoadBalancers: do.LoadBalancers{*lb}}
	return c.Display(item)
}

// RunLoadBalancerUpdate updates an existing load balancer with new configuration.
func RunLoadBalancerUpdate(c *CmdConfig) error {
	if len(c.Args) == 0 {
		return doctl.NewMissingArgsErr(c.NS)
	}
	lbID := c.Args[0]

	r := new(godo.LoadBalancerRequest)
	if err := buildRequestFromArgs(c, r); err != nil {
		return err
	}

	lbs := c.LoadBalancers()
	lb, err := lbs.Update(lbID, r)
	if err != nil {
		return err
	}

	item := &displayers.LoadBalancer{LoadBalancers: do.LoadBalancers{*lb}}
	return c.Display(item)
}

// RunLoadBalancerDelete deletes a load balancer by its identifier.
func RunLoadBalancerDelete(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	lbID := c.Args[0]

	force, err := c.Doit.GetBool(c.NS, doctl.ArgForce)
	if err != nil {
		return err
	}

	if force || AskForConfirmDelete("load balancer", 1) == nil {
		lbs := c.LoadBalancers()
		if err := lbs.Delete(lbID); err != nil {
			return err
		}
	} else {
		return errOperationAborted
	}

	return nil
}

// RunLoadBalancerAddDroplets adds droplets to a load balancer.
func RunLoadBalancerAddDroplets(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	lbID := c.Args[0]

	dropletIDsList, err := c.Doit.GetStringSlice(c.NS, doctl.ArgDropletIDs)
	if err != nil {
		return err
	}

	dropletIDs, err := extractDropletIDs(dropletIDsList)
	if err != nil {
		return err
	}

	return c.LoadBalancers().AddDroplets(lbID, dropletIDs...)
}

// RunLoadBalancerRemoveDroplets removes droplets from a load balancer.
func RunLoadBalancerRemoveDroplets(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	lbID := c.Args[0]

	dropletIDsList, err := c.Doit.GetStringSlice(c.NS, doctl.ArgDropletIDs)
	if err != nil {
		return err
	}

	dropletIDs, err := extractDropletIDs(dropletIDsList)
	if err != nil {
		return err
	}

	return c.LoadBalancers().RemoveDroplets(lbID, dropletIDs...)
}

// RunLoadBalancerAddForwardingRules adds forwarding rules to a load balancer.
func RunLoadBalancerAddForwardingRules(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	lbID := c.Args[0]

	fra, err := c.Doit.GetString(c.NS, doctl.ArgForwardingRules)
	if err != nil {
		return err
	}

	forwardingRules, err := extractForwardingRules(fra)
	if err != nil {
		return err
	}

	return c.LoadBalancers().AddForwardingRules(lbID, forwardingRules...)
}

// RunLoadBalancerRemoveForwardingRules removes forwarding rules from a load balancer.
func RunLoadBalancerRemoveForwardingRules(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	lbID := c.Args[0]

	fra, err := c.Doit.GetString(c.NS, doctl.ArgForwardingRules)
	if err != nil {
		return err
	}

	forwardingRules, err := extractForwardingRules(fra)
	if err != nil {
		return err
	}

	return c.LoadBalancers().RemoveForwardingRules(lbID, forwardingRules...)
}

// RunLoadBalancerPurgeCache purges cache for a global load balancer by its identifier.
func RunLoadBalancerPurgeCache(c *CmdConfig) error {
	err := ensureOneArg(c)
	if err != nil {
		return err
	}
	lbID := c.Args[0]

	force, err := c.Doit.GetBool(c.NS, doctl.ArgForce)
	if err != nil {
		return err
	}

	if force || AskForConfirm("purge CDN cache for global load balancer") == nil {
		lbs := c.LoadBalancers()
		if err := lbs.PurgeCache(lbID); err != nil {
			return err
		}
	} else {
		return errOperationAborted
	}

	return nil
}

func extractForwardingRules(s string) (forwardingRules []godo.ForwardingRule, err error) {
	if len(s) == 0 {
		return forwardingRules, err
	}

	list := strings.Split(s, " ")

	for _, v := range list {
		forwardingRule := new(godo.ForwardingRule)
		if err := fillStructFromStringSliceArgs(forwardingRule, v, ","); err != nil {
			return nil, err
		}

		forwardingRules = append(forwardingRules, *forwardingRule)
	}

	return forwardingRules, err
}

func extractDomains(s []string) (domains []*godo.LBDomain, err error) {
	if len(s) == 0 {
		return domains, err
	}

	for _, v := range s {
		domain := new(godo.LBDomain)
		if err := fillStructFromStringSliceArgs(domain, v, " "); err != nil {
			return nil, err
		}

		domains = append(domains, domain)
	}

	return domains, err
}

func fillStructFromStringSliceArgs(obj any, s string, delimiter string) error {
	if len(s) == 0 {
		return nil
	}

	kvs := strings.Split(s, delimiter)
	m := map[string]string{}

	for _, v := range kvs {
		p := strings.Split(v, ":")
		if len(p) == 2 {
			m[p[0]] = p[1]
		} else {
			return fmt.Errorf("Unexpected input value %v: must be a key:value pair", p)
		}
	}

	structValue := reflect.Indirect(reflect.ValueOf(obj))
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		f := structValue.Field(i)
		jv := strings.Split(structType.Field(i).Tag.Get("json"), ",")[0]

		if val, exists := m[jv]; exists {
			switch f.Kind() {
			case reflect.Bool:
				if v, err := strconv.ParseBool(val); err == nil {
					f.Set(reflect.ValueOf(v))
				}
			case reflect.Int:
				if v, err := strconv.Atoi(val); err == nil {
					f.Set(reflect.ValueOf(v))
				}
			case reflect.Uint32:
				if v64, err := strconv.ParseUint(val, 10, 32); err == nil {
					f.Set(reflect.ValueOf(uint32(v64)))
				}
			case reflect.String:
				f.Set(reflect.ValueOf(val))
			case reflect.Map:
				for _, kvPair := range strings.Split(val, " ") {
					kv := strings.Split(kvPair, "=")
					if len(kv) == 2 {
						if v32, err := strconv.ParseUint(kv[1], 10, 32); err == nil {
							if f.IsZero() {
								f.Set(reflect.MakeMap(f.Type()))
							}
							f.SetMapIndex(reflect.ValueOf(kv[0]), reflect.ValueOf(uint32(v32)))
						}
					}
				}
			default:
				return fmt.Errorf("Unexpected type for struct field %v", val)
			}
		}
	}

	return nil
}

func buildRequestFromArgs(c *CmdConfig, r *godo.LoadBalancerRequest) error {
	name, err := c.Doit.GetString(c.NS, doctl.ArgLoadBalancerName)
	if err != nil {
		return err
	}
	r.Name = name

	region, err := c.Doit.GetString(c.NS, doctl.ArgRegionSlug)
	if err != nil {
		return err
	}
	r.Region = region

	size, err := c.Doit.GetString(c.NS, doctl.ArgSizeSlug)
	if err != nil {
		return err
	}
	r.SizeSlug = size

	sizeUnit, err := c.Doit.GetInt(c.NS, doctl.ArgSizeUnit)
	if err != nil {
		return err
	}
	r.SizeUnit = uint32(sizeUnit)

	lbType, err := c.Doit.GetString(c.NS, doctl.ArgLoadBalancerType)
	if err != nil {
		return err
	}
	r.Type = strings.ToUpper(lbType)

	algorithm, err := c.Doit.GetString(c.NS, doctl.ArgLoadBalancerAlgorithm)
	if err != nil {
		return err
	}
	r.Algorithm = algorithm

	tag, err := c.Doit.GetString(c.NS, doctl.ArgTagName)
	if err != nil {
		return err
	}
	r.Tag = tag

	vpcUUID, err := c.Doit.GetString(c.NS, doctl.ArgVPCUUID)
	if err != nil {
		return err
	}
	r.VPCUUID = vpcUUID

	redirectHTTPToHTTPS, err := c.Doit.GetBool(c.NS, doctl.ArgRedirectHTTPToHTTPS)
	if err != nil {
		return err
	}
	r.RedirectHttpToHttps = redirectHTTPToHTTPS

	enableProxyProtocol, err := c.Doit.GetBool(c.NS, doctl.ArgEnableProxyProtocol)
	if err != nil {
		return err
	}
	r.EnableProxyProtocol = enableProxyProtocol

	enableBackendKeepalive, err := c.Doit.GetBool(c.NS, doctl.ArgEnableBackendKeepalive)
	if err != nil {
		return err
	}
	r.EnableBackendKeepalive = enableBackendKeepalive

	disableLetsEncryptDNSRecords, err := c.Doit.GetBool(c.NS, doctl.ArgDisableLetsEncryptDNSRecords)
	if err != nil {
		return err
	}
	r.DisableLetsEncryptDNSRecords = &disableLetsEncryptDNSRecords

	dropletIDsList, err := c.Doit.GetStringSlice(c.NS, doctl.ArgDropletIDs)
	if err != nil {
		return err
	}

	dropletIDs, err := extractDropletIDs(dropletIDsList)
	if err != nil {
		return err
	}
	r.DropletIDs = dropletIDs

	ssa, err := c.Doit.GetString(c.NS, doctl.ArgStickySessions)
	if err != nil {
		return err
	}

	stickySession := new(godo.StickySessions)
	if err := fillStructFromStringSliceArgs(stickySession, ssa, ","); err != nil {
		return err
	}
	r.StickySessions = stickySession

	hca, err := c.Doit.GetString(c.NS, doctl.ArgHealthCheck)
	if err != nil {
		return err
	}

	healthCheck := new(godo.HealthCheck)
	if err := fillStructFromStringSliceArgs(healthCheck, hca, ","); err != nil {
		return err
	}
	r.HealthCheck = healthCheck

	fra, err := c.Doit.GetString(c.NS, doctl.ArgForwardingRules)
	if err != nil {
		return err
	}

	forwardingRules, err := extractForwardingRules(fra)
	if err != nil {
		return err
	}
	r.ForwardingRules = forwardingRules

	projectID, err := c.Doit.GetString(c.NS, doctl.ArgProjectID)
	if err != nil {
		return err
	}

	r.ProjectID = projectID

	httpIdleTimeout, err := c.Doit.GetInt(c.NS, doctl.ArgHTTPIdleTimeoutSeconds)
	if err != nil {
		return err
	}

	if httpIdleTimeout != 0 {
		t := uint64(httpIdleTimeout)
		r.HTTPIdleTimeoutSeconds = &t
	}

	allowRules, allowflagSet, err := c.Doit.GetStringSliceIsFlagSet(c.NS, doctl.ArgAllowList)
	if err != nil {
		return err
	}

	denyRules, denyFlagSet, err := c.Doit.GetStringSliceIsFlagSet(c.NS, doctl.ArgDenyList)
	if err != nil {
		return err
	}

	if allowflagSet || denyFlagSet {
		firewall := new(godo.LBFirewall)
		firewall.Allow = allowRules
		firewall.Deny = denyRules
		r.Firewall = firewall
	}

	dms, err := c.Doit.GetStringSlice(c.NS, doctl.ArgLoadBalancerDomains)
	if err != nil {
		return err
	}

	domains, err := extractDomains(dms)
	if err != nil {
		return err
	}
	r.Domains = domains

	glbs, err := c.Doit.GetString(c.NS, doctl.ArgGlobalLoadBalancerSettings)
	if err != nil {
		return err
	}

	glbSettings := new(godo.GLBSettings)
	if err := fillStructFromStringSliceArgs(glbSettings, glbs, ","); err != nil {
		return err
	}
	if glbSettings.TargetProtocol != "" && glbSettings.TargetPort != 0 {
		r.GLBSettings = glbSettings
	}

	cdns, err := c.Doit.GetString(c.NS, doctl.ArgGlobalLoadBalancerCDNSettings)
	if err != nil {
		return err
	}

	cdnSettings := new(godo.CDNSettings)
	if err := fillStructFromStringSliceArgs(cdnSettings, cdns, ","); err != nil {
		return err
	}
	if r.GLBSettings != nil {
		r.GLBSettings.CDN = cdnSettings
	}

	lbIDsList, err := c.Doit.GetStringSlice(c.NS, doctl.ArgTargetLoadBalancerIDs)
	if err != nil {
		return err
	}
	r.TargetLoadBalancerIDs = lbIDsList

	network, err := c.Doit.GetString(c.NS, doctl.ArgLoadBalancerNetwork)
	if err != nil {
		return err
	}
	r.Network = strings.ToUpper(network)

	networkStack, err := c.Doit.GetString(c.NS, doctl.ArgLoadBalancerNetworkStack)
	if err != nil {
		return err
	}
	r.NetworkStack = strings.ToUpper(networkStack)

	tlsCipherPolicy, err := c.Doit.GetString(c.NS, doctl.ArgLoadBalancerTLSCipherPolicy)
	if err != nil {
		return err
	}
	if tlsCipherPolicy != "" {
		r.TLSCipherPolicy = strings.ToUpper(tlsCipherPolicy)
	}

	return nil
}

func waitForActiveLoadBalancer(lbs do.LoadBalancersService, lbID string) error {
	const maxAttempts = 180
	const wantStatus = "active"
	const errStatus = "errored"
	attempts := 0
	printNewLineSet := false

	for i := 0; i < maxAttempts; i++ {
		if attempts != 0 {
			fmt.Fprint(os.Stderr, ".")
			if !printNewLineSet {
				printNewLineSet = true
				defer fmt.Fprintln(os.Stderr)
			}
		}

		lb, err := lbs.Get(lbID)
		if err != nil {
			return err
		}

		if lb.Status == errStatus {
			return fmt.Errorf(
				"load balancer (%s) entered status `errored`",
				lbID,
			)
		}

		if lb.Status == wantStatus {
			return nil
		}

		attempts++
		time.Sleep(10 * time.Second)
	}

	return fmt.Errorf(
		"timeout waiting for load balancer (%s) to become active",
		lbID,
	)
}
