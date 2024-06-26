// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package leaderreceivercreator

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"github.com/skhalash/leaderreceivercreator/internal/sharedcomponent"
	"github.com/skhalash/leaderreceivercreator/internal/metadata"
)

var receivers = sharedcomponent.NewSharedComponents()

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithLogs(createLogsReceiver, metadata.LogsStability),
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability),
		receiver.WithTraces(createTracesReceiver, metadata.TracesStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createLogsReceiver(
	_ context.Context,
	params receiver.CreateSettings,
	cfg component.Config,
	consumer consumer.Logs,
) (receiver.Logs, error) {
	r := receivers.GetOrAdd(cfg, func() component.Component {
		return newLeaderReceiverCreator(params, cfg.(*Config))
	})
	r.Component.(*leaderReceiverCreator).nextLogsConsumer = consumer
	return r, nil
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.CreateSettings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	r := receivers.GetOrAdd(cfg, func() component.Component {
		return newLeaderReceiverCreator(params, cfg.(*Config))
	})
	r.Component.(*leaderReceiverCreator).nextMetricsConsumer = consumer
	return r, nil
}

func createTracesReceiver(
	_ context.Context,
	params receiver.CreateSettings,
	cfg component.Config,
	consumer consumer.Traces,
) (receiver.Traces, error) {
	r := receivers.GetOrAdd(cfg, func() component.Component {
		return newLeaderReceiverCreator(params, cfg.(*Config))
	})
	r.Component.(*leaderReceiverCreator).nextTracesConsumer = consumer
	return r, nil
}
