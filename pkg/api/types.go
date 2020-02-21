/*
Copyright 2020, Staffbase GmbH and contributors.

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

package api

type Quota struct {
	ComputeCores                   float64 `json:"compute.cores"`
	ComputeInstances               float64 `json:"compute.instances"`
	ComputeRAMMb                   float64 `json:"compute.ram_mb"`
	DNSZones                       float64 `json:"dns.zones"`
	NetworkFloatingips             float64 `json:"network.floatingips"`
	NetworkLbHealthmonitors        float64 `json:"network.lb_healthmonitors"`
	NetworkLbListeners             float64 `json:"network.lb_listeners"`
	NetworkLbMembers               float64 `json:"network.lb_members"`
	NetworkLbPools                 float64 `json:"network.lb_pools"`
	NetworkLoadbalancers           float64 `json:"network.loadbalancers"`
	NetworkVpnEndpointGroups       float64 `json:"network.vpn_endpoint_groups"`
	NetworkVpnIkepolicies          float64 `json:"network.vpn_ikepolicies"`
	NetworkVpnIpsecSiteConnections float64 `json:"network.vpn_ipsec_site_connections"`
	NetworkVpnIpsecpolicies        float64 `json:"network.vpn_ipsecpolicies"`
	NetworkVpnServices             float64 `json:"network.vpn_services"`
	S3SpaceBytes                   float64 `json:"s3.space_bytes"`
	VolumeSpaceGb                  float64 `json:"volume.space_gb"`
	VolumeVolumes                  float64 `json:"volume.volumes"`
}

type CurrentUsage struct {
	ComputeCores                   float64            `json:"compute.cores"`
	ComputeFlavors                 map[string]float64 `json:"compute.flavors"`
	ComputeInstances               float64            `json:"compute.instances"`
	ComputeRAMMb                   float64            `json:"compute.ram_mb"`
	DNSZones                       float64            `json:"dns.zones"`
	NetworkFloatingips             float64            `json:"network.floatingips"`
	NetworkLbHealthmonitors        float64            `json:"network.lb_healthmonitors"`
	NetworkLbListeners             float64            `json:"network.lb_listeners"`
	NetworkLbMembers               float64            `json:"network.lb_members"`
	NetworkLbPools                 float64            `json:"network.lb_pools"`
	NetworkLoadbalancers           float64            `json:"network.loadbalancers"`
	NetworkVpnEndpointGroups       float64            `json:"network.vpn_endpoint_groups"`
	NetworkVpnIkepolicies          float64            `json:"network.vpn_ikepolicies"`
	NetworkVpnIpsecSiteConnections float64            `json:"network.vpn_ipsec_site_connections"`
	NetworkVpnIpsecpolicies        float64            `json:"network.vpn_ipsecpolicies"`
	NetworkVpnServices             float64            `json:"network.vpn_services"`
	S3SpaceBytes                   float64            `json:"s3.space_bytes"`
	VolumeSpaceGb                  float64            `json:"volume.space_gb"`
	VolumeVolumes                  float64            `json:"volume.volumes"`
}
