	"github.com/m3db/m3db/clock"
// Options represents the options for filesystem persistence
	// SetClockOptions sets the clock options
	SetClockOptions(value clock.Options) Options

	// ClockOptions returns the clock options
	ClockOptions() clock.Options


	// SetThroughputCheckInterval sets the filesystem throughput check interval
	SetThroughputCheckInterval(value time.Duration) Options

	// ThroughputCheckInterval returns the filesystem throughput check interval
	ThroughputCheckInterval() time.Duration

	// SetThroughputLimitMbps sets the filesystem throughput limit
	SetThroughputLimitMbps(value float64) Options

	// ThroughputLimitMbps returns the filesystem throughput limit
	ThroughputLimitMbps() float64