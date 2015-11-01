# grpcxlog
grpcxlog implements an adapter to use xlog as the logger for grpc

gRPC applications use glog as the defautl logger.

[xlog](https://github.com/rs/xlog) is a powerfull logger for HTTP based applications.

grpcxlog is and apdater to use xlog in gRPC based applications to have the same log format independently of the protocol used.

Usage

    // Install the logger handler with default output on the console
	  lh := xlog.NewHandler(xlog.LevelDebug)

	  // Set some global env fields
	  host, _ := os.Hostname()
	  lh.SetFields(xlog.F{
  		"svc":  serviceID,
  		"host": host,
  	})
  
  	// Plug the xlog handler's input to gRPC's default logger
  	grpclog.SetLogger(grpcxlog.Log{lh.NewLogger()})
