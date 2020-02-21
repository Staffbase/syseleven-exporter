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
	"errors"
	"os"
	"time"

	"github.com/Staffbase/syseleven-exporter/pkg/api"
	"github.com/Staffbase/syseleven-exporter/pkg/auth"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type Exporter struct {
	Username  string
	Password  string
	ProjectID string
}

func New() (*Exporter, error) {
	if os.Getenv("OS_USERNAME") == "" {
		return nil, errors.New("OS_USERNAME is missing")
	}

	if os.Getenv("OS_PASSWORD") == "" {
		return nil, errors.New("OS_PASSWORD is missing")
	}

	if os.Getenv("OS_PROJECT_ID") == "" {
		return nil, errors.New("OS_PROJECT_ID is missing")
	}

	return &Exporter{
		Username:  os.Getenv("OS_USERNAME"),
		Password:  os.Getenv("OS_PASSWORD"),
		ProjectID: os.Getenv("OS_PROJECT_ID"),
	}, nil
}

func Run(interval int64, exporter *Exporter) {
	for {
		token, err := auth.GetToken(exporter.Username, exporter.Password)
		if err != nil {
			log.WithError(err).Error("Could not get API Token")
		}

		quota, err := api.GetQuota(exporter.ProjectID, token)
		if err != nil {
			log.WithError(err).Error("Could not get quota")
		}

		usage, err := api.GetCurrentUsage(exporter.ProjectID, token)
		if err != nil {
			log.WithError(err).Error("Could not get current usage")
		}

		for k, v := range quota {
			computeCoresTotal.With(prometheus.Labels{"region": k}).Set(v.ComputeCores)
			computeInstancesTotal.With(prometheus.Labels{"region": k}).Set(v.ComputeInstances)
			computeRamTotalMegabytes.With(prometheus.Labels{"region": k}).Set(v.ComputeRAMMb)
			dnsZonesTotal.With(prometheus.Labels{"region": k}).Set(v.DNSZones)
			networkFloatingIPsTotal.With(prometheus.Labels{"region": k}).Set(v.NetworkFloatingips)
			networkLoadbalancersTotal.With(prometheus.Labels{"region": k}).Set(v.NetworkLoadbalancers)
			s3SpaceTotalBytes.With(prometheus.Labels{"region": k}).Set(v.S3SpaceBytes)
			volumeSpaceTotalGigabytes.With(prometheus.Labels{"region": k}).Set(v.VolumeSpaceGb)
			volumeVolumesTotalGigabytes.With(prometheus.Labels{"region": k}).Set(v.VolumeVolumes)
		}

		for k, v := range usage {
			computeCoresUsed.With(prometheus.Labels{"region": k}).Set(v.ComputeCores)
			computeInstancesUsed.With(prometheus.Labels{"region": k}).Set(v.ComputeInstances)
			computeRamUsedMegabytes.With(prometheus.Labels{"region": k}).Set(v.ComputeRAMMb)
			dnsZonesUsed.With(prometheus.Labels{"region": k}).Set(v.DNSZones)
			networkFloatingIPsUsed.With(prometheus.Labels{"region": k}).Set(v.NetworkFloatingips)
			networkLoadbalancersUsed.With(prometheus.Labels{"region": k}).Set(v.NetworkLoadbalancers)
			s3SpaceUsedBytes.With(prometheus.Labels{"region": k}).Set(v.S3SpaceBytes)
			volumeSpaceUsedGigabytes.With(prometheus.Labels{"region": k}).Set(v.VolumeSpaceGb)
			volumeVolumesUsedGigabytes.With(prometheus.Labels{"region": k}).Set(v.VolumeVolumes)

			for flavor := range v.ComputeFlavors {
				computeFlavorsUsed.With(prometheus.Labels{"region": k, "flavor": flavor}).Set(v.ComputeFlavors[flavor])
			}
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
