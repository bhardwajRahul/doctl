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

package displayers

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/digitalocean/doctl/do"
)

type LoadBalancer struct {
	LoadBalancers do.LoadBalancers
}

var _ Displayable = &LoadBalancer{}

func (lb *LoadBalancer) JSON(out io.Writer) error {
	return writeJSON(lb.LoadBalancers, out)
}

func (lb *LoadBalancer) Cols() []string {
	return []string{
		"ID",
		"IP",
		"IPv6",
		"Name",
		"Status",
		"Created",
		"Region",
		"Size",
		"SizeUnit",
		"VPCUUID",
		"Tag",
		"DropletIDs",
		"RedirectHttpToHttps",
		"StickySessions",
		"HealthCheck",
		"ForwardingRules",
		"Firewall",
		"DisableLetsEncryptDNSRecords",
	}
}

func (lb *LoadBalancer) ColMap() map[string]string {
	return map[string]string{
		"ID":                           "ID",
		"IP":                           "IP",
		"IPv6":                         "IPv6",
		"Name":                         "Name",
		"Status":                       "Status",
		"Created":                      "Created At",
		"Region":                       "Region",
		"Size":                         "Size",
		"SizeUnit":                     "Size Unit",
		"VPCUUID":                      "VPC UUID",
		"Tag":                          "Tag",
		"DropletIDs":                   "Droplet IDs",
		"RedirectHttpToHttps":          "SSL",
		"StickySessions":               "Sticky Sessions",
		"HealthCheck":                  "Health Check",
		"ForwardingRules":              "Forwarding Rules",
		"Firewall":                     "Firewall Rules",
		"DisableLetsEncryptDNSRecords": "Disable Lets Encrypt DNS Records",
	}
}

func (lb *LoadBalancer) KV() []map[string]any {
	out := make([]map[string]any, 0, len(lb.LoadBalancers))

	for _, l := range lb.LoadBalancers {
		forwardingRules := make([]string, 0, len(l.ForwardingRules))
		for _, r := range l.ForwardingRules {
			forwardingRules = append(forwardingRules, prettyPrintStruct(r))
		}

		o := map[string]any{
			"ID":                           l.ID,
			"IP":                           l.IP,
			"IPv6":                         l.IPv6,
			"Name":                         l.Name,
			"Status":                       l.Status,
			"Created":                      l.Created,
			"VPCUUID":                      l.VPCUUID,
			"Tag":                          l.Tag,
			"DropletIDs":                   strings.Trim(strings.Replace(fmt.Sprint(l.DropletIDs), " ", ",", -1), "[]"),
			"RedirectHttpToHttps":          l.RedirectHttpToHttps,
			"StickySessions":               prettyPrintStruct(l.StickySessions),
			"HealthCheck":                  prettyPrintStruct(l.HealthCheck),
			"ForwardingRules":              strings.Join(forwardingRules, " "),
			"DisableLetsEncryptDNSRecords": toBool(l.DisableLetsEncryptDNSRecords),
		}
		if l.Region != nil {
			o["Region"] = l.Region.Slug
		}
		if l.SizeSlug != "" {
			o["Size"] = l.SizeSlug
		}
		if l.SizeUnit > 0 {
			o["SizeUnit"] = l.SizeUnit
		}
		if l.Firewall != nil {
			o["Firewall"] = prettyPrintStruct(l.Firewall)
		}
		out = append(out, o)
	}

	return out
}

func toBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func prettyPrintStruct(obj any) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered from %v", err)
		}
	}()

	val := reflect.Indirect(reflect.ValueOf(obj))
	numField := val.NumField()
	output := make([]string, 0, numField)
	for i := 0; i < numField; i++ {
		k := strings.Split(val.Type().Field(i).Tag.Get("json"), ",")[0]
		v := reflect.ValueOf(val.Field(i).Interface())
		output = append(output, fmt.Sprintf("%v:%v", k, v))
	}

	return strings.Join(output, ",")
}
