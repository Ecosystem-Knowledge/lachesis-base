package peerleecher

import (
	"time"

	"github.com/Ecosystem-Knowledge/lachesis-base/inter/dag"
)

type EpochDownloaderConfig struct {
	RecheckInterval        time.Duration
	DefaultChunkSize       dag.Metric
	ParallelChunksDownload int
}
