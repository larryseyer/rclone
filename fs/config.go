package fs

import (
	"net"
	"time"
)

// Global
var (
	// Config is the global config
	Config = NewConfig()

	// Read a value from the config file
	//
	// This is a function pointer to decouple the config
	// implementation from the fs
	ConfigFileGet = func(section, key string, defaultVal ...string) string { return "" }

	// CountError counts an error.  If any errors have been
	// counted then it will exit with a non zero error code.
	//
	// This is a function pointer to decouple the config
	// implementation from the fs
	CountError = func(err error) {}

	// ConfigProvider is the config key used for provider options
	ConfigProvider = "provider"
)

// ConfigInfo is filesystem config options
type ConfigInfo struct {
	LogLevel              LogLevel
	StatsLogLevel         LogLevel
	DryRun                bool
	CheckSum              bool
	SizeOnly              bool
	IgnoreTimes           bool
	IgnoreExisting        bool
	IgnoreErrors          bool
	ModifyWindow          time.Duration
	Checkers              int
	Transfers             int
	ConnectTimeout        time.Duration // Connect timeout
	Timeout               time.Duration // Data channel timeout
	Dump                  DumpFlags
	InsecureSkipVerify    bool // Skip server certificate verification
	DeleteMode            DeleteMode
	MaxDelete             int64
	TrackRenames          bool // Track file renames.
	LowLevelRetries       int
	UpdateOlder           bool // Skip files that are newer on the destination
	NoGzip                bool // Disable compression
	MaxDepth              int
	IgnoreSize            bool
	IgnoreChecksum        bool
	NoUpdateModTime       bool
	DataRateUnit          string
	BackupDir             string
	Suffix                string
	UseListR              bool
	BufferSize            SizeSuffix
	BwLimit               BwTimetable
	TPSLimit              float64
	TPSLimitBurst         int
	BindAddr              net.IP
	DisableFeatures       []string
	UserAgent             string
	Immutable             bool
	AutoConfirm           bool
	StreamingUploadCutoff SizeSuffix
	StatsFileNameLength   int
	AskPassword           bool
	UseServerModTime      bool
}

// NewConfig creates a new config with everything set to the default
// value.  These are the ultimate defaults and are overriden by the
// config module.
func NewConfig() *ConfigInfo {
	c := new(ConfigInfo)

	// Set any values which aren't the zero for the type
	c.LogLevel = LogLevelNotice
	c.StatsLogLevel = LogLevelInfo
	c.ModifyWindow = time.Nanosecond
	c.Checkers = 8
	c.Transfers = 4
	c.ConnectTimeout = 60 * time.Second
	c.Timeout = 5 * 60 * time.Second
	c.DeleteMode = DeleteModeDefault
	c.MaxDelete = -1
	c.LowLevelRetries = 10
	c.MaxDepth = -1
	c.DataRateUnit = "bytes"
	c.BufferSize = SizeSuffix(16 << 20)
	c.UserAgent = "rclone/" + Version
	c.StreamingUploadCutoff = SizeSuffix(100 * 1024)
	c.StatsFileNameLength = 40
	c.AskPassword = true
	c.TPSLimitBurst = 1

	return c
}
