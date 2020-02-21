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

package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const namespace = "syseleven"

var (
	computeCoresTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "compute_cores_total",
	}, []string{"region"})

	computeCoresUsed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "compute_cores_used",
	}, []string{"region"})

	computeInstancesTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "compute_instances_total",
	}, []string{"region"})

	computeInstancesUsed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "compute_instances_used",
	}, []string{"region"})

	computeFlavorsUsed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "compute_flavors_used",
	}, []string{"region", "flavor"})

	computeRamTotalMegabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "compute_ram_total_megabytes",
	}, []string{"region"})

	computeRamUsedMegabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "compute_ram_used_megabytes",
	}, []string{"region"})

	dnsZonesTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "dns_zones_total",
	}, []string{"region"})

	dnsZonesUsed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "dns_zones_used",
	}, []string{"region"})

	networkFloatingIPsTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "network_floating_ips_total",
	}, []string{"region"})

	networkFloatingIPsUsed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "network_floating_ips_used",
	}, []string{"region"})

	networkLoadbalancersTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "network_loadbalancers_total",
	}, []string{"region"})

	networkLoadbalancersUsed = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "network_loadbalancers_used",
	}, []string{"region"})

	s3SpaceTotalBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "s3_space_total_bytes",
	}, []string{"region"})

	s3SpaceUsedBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "s3_space_used_bytes",
	}, []string{"region"})

	volumeSpaceTotalGigabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "volume_space_total_gigabytes",
	}, []string{"region"})

	volumeSpaceUsedGigabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "volume_space_used_gigabytes",
	}, []string{"region"})

	volumeVolumesTotalGigabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "volume_volumes_total",
	}, []string{"region"})

	volumeVolumesUsedGigabytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "volume_volumes_used",
	}, []string{"region"})
)
