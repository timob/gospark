package gospark

import "github.com/timob/javabind"

type ExecutorShuffleReadMetricsInterface interface {
	JavaLangObjectInterface

	// public int remoteBlocksFetched()
	RemoteBlocksFetched() int

	// public void incRemoteBlocksFetched(int)
	IncRemoteBlocksFetched(a int) 

	// public void decRemoteBlocksFetched(int)
	DecRemoteBlocksFetched(a int) 

	// public int localBlocksFetched()
	LocalBlocksFetched() int

	// public void incLocalBlocksFetched(int)
	IncLocalBlocksFetched(a int) 

	// public void decLocalBlocksFetched(int)
	DecLocalBlocksFetched(a int) 

	// public long fetchWaitTime()
	FetchWaitTime() int64

	// public void incFetchWaitTime(long)
	IncFetchWaitTime(a int64) 

	// public void decFetchWaitTime(long)
	DecFetchWaitTime(a int64) 

	// public long remoteBytesRead()
	RemoteBytesRead() int64

	// public void incRemoteBytesRead(long)
	IncRemoteBytesRead(a int64) 

	// public void decRemoteBytesRead(long)
	DecRemoteBytesRead(a int64) 

	// public long localBytesRead()
	LocalBytesRead() int64

	// public void incLocalBytesRead(long)
	IncLocalBytesRead(a int64) 

	// public long totalBytesRead()
	TotalBytesRead() int64

	// public int totalBlocksFetched()
	TotalBlocksFetched() int

	// public long recordsRead()
	RecordsRead() int64

	// public void incRecordsRead(long)
	IncRecordsRead(a int64) 

	// public void decRecordsRead(long)
	DecRecordsRead(a int64) 
}

type ExecutorShuffleReadMetrics struct {
	JavaLangObject
}

// public org.apache.spark.executor.ShuffleReadMetrics()
func NewExecutorShuffleReadMetrics() (*ExecutorShuffleReadMetrics) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/executor/ShuffleReadMetrics")
	if err != nil {
		panic(err)
	}
	x := &ExecutorShuffleReadMetrics{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int remoteBlocksFetched()
func (jbobject *ExecutorShuffleReadMetrics) RemoteBlocksFetched() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "remoteBlocksFetched", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void incRemoteBlocksFetched(int)
func (jbobject *ExecutorShuffleReadMetrics) IncRemoteBlocksFetched(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "incRemoteBlocksFetched", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void decRemoteBlocksFetched(int)
func (jbobject *ExecutorShuffleReadMetrics) DecRemoteBlocksFetched(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "decRemoteBlocksFetched", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public int localBlocksFetched()
func (jbobject *ExecutorShuffleReadMetrics) LocalBlocksFetched() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "localBlocksFetched", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void incLocalBlocksFetched(int)
func (jbobject *ExecutorShuffleReadMetrics) IncLocalBlocksFetched(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "incLocalBlocksFetched", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void decLocalBlocksFetched(int)
func (jbobject *ExecutorShuffleReadMetrics) DecLocalBlocksFetched(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "decLocalBlocksFetched", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long fetchWaitTime()
func (jbobject *ExecutorShuffleReadMetrics) FetchWaitTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "fetchWaitTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void incFetchWaitTime(long)
func (jbobject *ExecutorShuffleReadMetrics) IncFetchWaitTime(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "incFetchWaitTime", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void decFetchWaitTime(long)
func (jbobject *ExecutorShuffleReadMetrics) DecFetchWaitTime(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "decFetchWaitTime", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long remoteBytesRead()
func (jbobject *ExecutorShuffleReadMetrics) RemoteBytesRead() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "remoteBytesRead", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void incRemoteBytesRead(long)
func (jbobject *ExecutorShuffleReadMetrics) IncRemoteBytesRead(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "incRemoteBytesRead", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void decRemoteBytesRead(long)
func (jbobject *ExecutorShuffleReadMetrics) DecRemoteBytesRead(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "decRemoteBytesRead", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long localBytesRead()
func (jbobject *ExecutorShuffleReadMetrics) LocalBytesRead() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "localBytesRead", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void incLocalBytesRead(long)
func (jbobject *ExecutorShuffleReadMetrics) IncLocalBytesRead(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "incLocalBytesRead", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long totalBytesRead()
func (jbobject *ExecutorShuffleReadMetrics) TotalBytesRead() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "totalBytesRead", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public int totalBlocksFetched()
func (jbobject *ExecutorShuffleReadMetrics) TotalBlocksFetched() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "totalBlocksFetched", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public long recordsRead()
func (jbobject *ExecutorShuffleReadMetrics) RecordsRead() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "recordsRead", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void incRecordsRead(long)
func (jbobject *ExecutorShuffleReadMetrics) IncRecordsRead(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "incRecordsRead", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void decRecordsRead(long)
func (jbobject *ExecutorShuffleReadMetrics) DecRecordsRead(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "decRecordsRead", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


