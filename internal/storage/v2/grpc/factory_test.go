// Copyright (c) 2025 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package grpc

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configauth"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"google.golang.org/grpc"

	"github.com/jaegertracing/jaeger/internal/telemetry"
	"github.com/jaegertracing/jaeger/internal/tenancy"
)

func TestNewFactory_NonEmptyAuthenticator(t *testing.T) {
	cfg := &Config{
		ClientConfig: configgrpc.ClientConfig{
			Auth: configoptional.Some(configauth.Config{}),
		},
	}
	_, err := NewFactory(context.Background(), *cfg, telemetry.NoopSettings())
	require.ErrorContains(t, err, "authenticator is not supported")
}

func TestNewFactory(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	require.NoError(t, err, "failed to listen")

	cfg := Config{
		ClientConfig: configgrpc.ClientConfig{
			Endpoint: lis.Addr().String(),
		},
		TimeoutConfig: exporterhelper.TimeoutConfig{
			Timeout: 1 * time.Second,
		},
		Tenancy: tenancy.Options{
			Enabled: true,
		},
	}
	telset := telemetry.NoopSettings()
	f, err := NewFactory(context.Background(), cfg, telset)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, f.Close()) })
	require.Equal(t, lis.Addr().String(), f.readerConn.Target())
	require.Equal(t, lis.Addr().String(), f.writerConn.Target())
}

func TestNewFactory_WriteEndpointOverride(t *testing.T) {
	readListener, err := net.Listen("tcp", ":0")
	require.NoError(t, err, "failed to listen")

	writeListener, err := net.Listen("tcp", ":0")
	require.NoError(t, err, "failed to listen")

	cfg := Config{
		ClientConfig: configgrpc.ClientConfig{
			Endpoint: readListener.Addr().String(),
		},
		Writer: configgrpc.ClientConfig{
			Endpoint: writeListener.Addr().String(),
		},
		TimeoutConfig: exporterhelper.TimeoutConfig{
			Timeout: 1 * time.Second,
		},
		Tenancy: tenancy.Options{
			Enabled: true,
		},
	}
	telset := telemetry.NoopSettings()
	f, err := NewFactory(context.Background(), cfg, telset)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, f.Close()) })
	require.Equal(t, readListener.Addr().String(), f.readerConn.Target())
	require.Equal(t, writeListener.Addr().String(), f.writerConn.Target())
}

func TestFactory(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	require.NoError(t, err, "failed to listen")

	s := grpc.NewServer()

	conn := startServer(t, s, lis)
	f := &Factory{
		readerConn: conn,
	}

	t.Run("CreateTraceReader", func(t *testing.T) {
		tr, err := f.CreateTraceReader()
		require.NoError(t, err)
		require.NotNil(t, tr)
	})

	t.Run("CreateTraceWriter", func(t *testing.T) {
		tr, err := f.CreateTraceWriter()
		require.NoError(t, err)
		require.NotNil(t, tr)
	})

	t.Run("CreateDependencyReader", func(t *testing.T) {
		tr, err := f.CreateDependencyReader()
		require.NoError(t, err)
		require.NotNil(t, tr)
	})
}

func TestInitializeConnections_ClientError(t *testing.T) {
	f, err := NewFactory(
		context.Background(),
		Config{
			ClientConfig: configgrpc.ClientConfig{
				Endpoint: ":0",
			},
		}, telemetry.NoopSettings())
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, f.Close()) })
	newClientFn := func(_ component.TelemetrySettings, _ *configgrpc.ClientConfig, _ ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
		return nil, assert.AnError
	}
	err = f.initializeConnections(
		component.TelemetrySettings{},
		component.TelemetrySettings{},
		&configgrpc.ClientConfig{},
		&configgrpc.ClientConfig{},
		newClientFn,
	)
	assert.ErrorContains(t, err, "error creating reader client connection")
}
