/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package backend

import (
	"github.com/apache/incubator-servicecomb-service-center/server/metric"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	scCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metric.FamilyName,
			Subsystem: "db",
			Name:      "sc_total",
			Help:      "Counter of the Service Center instance",
		}, []string{"instance"})
)

func init() {
	prometheus.MustRegister(scCounter)
}

func ReportScInstance() {
	instance := metric.InstanceName()
	scCounter.WithLabelValues(instance).Add(1)
}
