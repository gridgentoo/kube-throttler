// Licensed to Shingo Omura under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Shingo Omura licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package controllers

import (
	"github.com/everpeace/kube-throttler/pkg/apis/schedule/v1alpha1"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/component-base/metrics"
	"k8s.io/component-base/metrics/legacyregistry"
)

type ThrottleMetricsRecorder struct {
	MetricsRecorder
	specThresholdResourceCountsGauge              *metrics.GaugeVec
	specThresholdResourceRequestsGauge            *metrics.GaugeVec
	statusThrottledResourceCountsGauge            *metrics.GaugeVec
	statusThrottledResourceRequstsGauge           *metrics.GaugeVec
	statusUsedResourceCountsGauge                 *metrics.GaugeVec
	statusUsedResourceRequestsGauge               *metrics.GaugeVec
	statusCalculatedThresholdResourceCountsGauge  *metrics.GaugeVec
	statusCalculatedThresholdResourceRequstsGauge *metrics.GaugeVec
}

func NewThrottleMetricsRecorder() *ThrottleMetricsRecorder {
	r := &ThrottleMetricsRecorder{
		MetricsRecorder: MetricsRecorder{},
		specThresholdResourceCountsGauge: metrics.NewGaugeVec(
			&metrics.GaugeOpts{
				Name: "throttle_spec_threshold_resourceCounts",
				Help: "threshold on specific resourceCounts of the throttle",
			},
			[]string{"namespace", "name", "uid", "resource"},
		),
		specThresholdResourceRequestsGauge: metrics.NewGaugeVec(
			&metrics.GaugeOpts{
				Name: "throttle_spec_threshold_resourceRequests",
				Help: "threshold on specific resourceRequests of the throttle",
			},
			[]string{"namespace", "name", "uid", "resource"},
		),
		statusThrottledResourceCountsGauge: metrics.NewGaugeVec(
			&metrics.GaugeOpts{
				Name: "throttle_status_throttled_resourceCounts",
				Help: "resourceCounts of the throttle is throttled or not on specific resource (1=throttled, 0=not throttled)",
			},
			[]string{"namespace", "name", "uid", "resource"},
		),
		statusThrottledResourceRequstsGauge: metrics.NewGaugeVec(
			&metrics.GaugeOpts{
				Name: "throttle_status_throttled_resourceRequests",
				Help: "resourceRequests of the throttle is throttled or not on specific resource (1=throttled, 0=not throttled)",
			},
			[]string{"namespace", "name", "uid", "resource"},
		),
		statusUsedResourceCountsGauge: metrics.NewGaugeVec(
			&metrics.GaugeOpts{
				Name: "throttle_status_used_resourceCounts",
				Help: "used resource counts of the throttle",
			},
			[]string{"namespace", "name", "uid", "resource"},
		),
		statusUsedResourceRequestsGauge: metrics.NewGaugeVec(
			&metrics.GaugeOpts{
				Name: "throttle_status_used_resourceRequests",
				Help: "used amount of resource requests of the throttle",
			},
			[]string{"namespace", "name", "uid", "resource"},
		),
		statusCalculatedThresholdResourceCountsGauge: metrics.NewGaugeVec(
			&metrics.GaugeOpts{
				Name: "throttle_status_calculated_threshold_resourceCounts",
				Help: "calculated threshold on specific resourceCounts of the throttle",
			},
			[]string{"namespace", "name", "uid", "resource"},
		),
		statusCalculatedThresholdResourceRequstsGauge: metrics.NewGaugeVec(
			&metrics.GaugeOpts{
				Name: "throttle_status_calculated_threshold_resourceRequests",
				Help: "calculated threshold on specific resourceRequests of the throttle",
			},
			[]string{"namespace", "name", "uid", "resource"},
		),
	}
	legacyregistry.MustRegister(
		r.specThresholdResourceCountsGauge,
		r.specThresholdResourceRequestsGauge,
		r.statusThrottledResourceCountsGauge,
		r.statusThrottledResourceRequstsGauge,
		r.statusUsedResourceCountsGauge,
		r.statusUsedResourceRequestsGauge,
		r.statusCalculatedThresholdResourceCountsGauge,
		r.statusCalculatedThresholdResourceRequstsGauge,
	)
	return r
}

func (r *ThrottleMetricsRecorder) recordThrottleMetrics(thr *v1alpha1.Throttle) {
	labels := prometheus.Labels{
		"namespace": thr.Namespace,
		"name":      thr.Name,
		"uid":       string(thr.UID),
	}

	r.recordResourceCounts(r.specThresholdResourceCountsGauge.MustCurryWith(labels), thr.Spec.Threshold.ResourceCounts)
	r.recordResourceRequests(r.specThresholdResourceRequestsGauge.MustCurryWith(labels), thr.Spec.Threshold.ResourceRequests)

	r.recordIsResourceCountThrottled(r.statusThrottledResourceCountsGauge.MustCurryWith(labels), thr.Status.Throttled.ResourceCounts)
	r.recordIsResourceRequestsThrottled(r.statusThrottledResourceRequstsGauge.MustCurryWith(labels), thr.Status.Throttled.ResourceRequests)

	r.recordResourceCounts(r.statusUsedResourceCountsGauge.MustCurryWith(labels), thr.Status.Used.ResourceCounts)
	r.recordResourceRequests(r.statusUsedResourceRequestsGauge.MustCurryWith(labels), thr.Status.Used.ResourceRequests)

	r.recordResourceCounts(r.statusCalculatedThresholdResourceCountsGauge.MustCurryWith(labels), thr.Status.CalculatedThreshold.Threshold.ResourceCounts)
	r.recordResourceRequests(r.statusCalculatedThresholdResourceRequstsGauge.MustCurryWith(labels), thr.Status.CalculatedThreshold.Threshold.ResourceRequests)
}
