# grpcxlog
grpcxlog implements an adapter to use xlog as the logger for grpc

gRPC applications use glog as the defautl logger.

[xlog](https://github.com/rs/xlog) is a powerfull logger for HTTP based applications.

grpcxlog is and apdater to use xlog in gRPC based applications to have the same log format independently of the protocol used.

Usage

	host, _ := os.Hostname()
	conf := xlog.Config{
		// Log debug level and higher
		Level: xlog.LevelDebug,
		// Set some global env fields
		Fields: xlog.F{
			"svc":  "my-service",
			"host": host,
		},
		// Output everything on console
		Output: xlog.NewOutputChannel(xlog.NewConsoleOutput()),
	}
	
	log = xlog.New(conf)
	
	// Plug the xlog logger to gRPC logger
	grpclog.SetLogger(grpcxlog.Log{log})
