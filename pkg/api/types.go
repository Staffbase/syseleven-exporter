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
