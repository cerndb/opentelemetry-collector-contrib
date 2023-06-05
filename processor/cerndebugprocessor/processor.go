package cerndebugprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/cerndebugprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/davecgh/go-spew/spew"
)

type debugProcessor struct {
	logger *zap.Logger
}

func (dp *debugProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	dp.logger.Debug("CERN",
		zap.String("context", spew.Sdump(ctx)),
		zap.String("trace", spew.Sdump(td)))
	return td, nil
}

func (dp *debugProcessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	dp.logger.Debug("CERN",
		zap.String("context", spew.Sdump(ctx)),
		zap.String("metrics", spew.Sdump(md)))
	return md, nil
}

func (dp *debugProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	dp.logger.Debug("CERN",
		zap.String("context", spew.Sdump(ctx)),
		zap.String("logs", spew.Sdump(ld)))
	return ld, nil
}
