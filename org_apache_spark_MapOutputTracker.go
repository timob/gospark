package gospark

import "github.com/timob/javabind"

type MapOutputTrackerInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.scheduler.MapStatus[] deserializeMapStatuses(byte[])
	DeserializeMapStatuses(a []byte) []*SchedulerMapStatus

	// public static byte[] serializeMapStatuses(org.apache.spark.scheduler.MapStatus[])
	SerializeMapStatuses(a []*SchedulerMapStatus) []byte

	// public static java.lang.String ENDPOINT_NAME()
	ENDPOINT_NAME() string

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public org.apache.spark.rpc.RpcEndpointRef trackerEndpoint()
	TrackerEndpoint() *RpcRpcEndpointRef

	// public long epoch()
	Epoch() int64

	// public java.lang.Object epochLock()
	EpochLock() *JavaLangObject

	// public void sendTracker(java.lang.Object)
	SendTracker(a interface{}) 

	// public org.apache.spark.MapOutputStatistics getStatistics(org.apache.spark.ShuffleDependency<?, ?, ?>)
	GetStatistics(a ShuffleDependencyInterface) *MapOutputStatistics

	// public long getEpoch()
	GetEpoch() int64

	// public void updateEpoch(long)
	UpdateEpoch(a int64) 

	// public void unregisterShuffle(int)
	UnregisterShuffle(a int) 

	// public void stop()
	Stop() 
}

type MapOutputTracker struct {
	JavaLangObject
}

// public org.apache.spark.MapOutputTracker(org.apache.spark.SparkConf)
func NewMapOutputTracker(a SparkConfInterface) (*MapOutputTracker) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/MapOutputTracker", conv_a.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &MapOutputTracker{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.apache.spark.scheduler.MapStatus[] deserializeMapStatuses(byte[])
func (jbobject *MapOutputTracker) DeserializeMapStatuses(a []byte) []*SchedulerMapStatus {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/MapOutputTracker", "deserializeMapStatuses", javabind.ObjectArrayType("org/apache/spark/scheduler/MapStatus"), a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/scheduler/MapStatus")
	dst := new([]*SchedulerMapStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static byte[] serializeMapStatuses(org.apache.spark.scheduler.MapStatus[])
func (jbobject *MapOutputTracker) SerializeMapStatuses(a []*SchedulerMapStatus) []byte {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/apache/spark/scheduler/MapStatus")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/MapOutputTracker", "serializeMapStatuses", javabind.Byte | javabind.Array, conv_a.Value().Cast("org/apache/spark/scheduler/MapStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.([]byte)
}

// public static java.lang.String ENDPOINT_NAME()
func (jbobject *MapOutputTracker) ENDPOINT_NAME() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/MapOutputTracker", "ENDPOINT_NAME", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String logName()
func (jbobject *MapOutputTracker) LogName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "logName", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean isTraceEnabled()
func (jbobject *MapOutputTracker) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.rpc.RpcEndpointRef trackerEndpoint()
func (jbobject *MapOutputTracker) TrackerEndpoint() *RpcRpcEndpointRef {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "trackerEndpoint", "org/apache/spark/rpc/RpcEndpointRef")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &RpcRpcEndpointRef{}
	unique_x.Callable = dst
	return unique_x
}

// public long epoch()
func (jbobject *MapOutputTracker) Epoch() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "epoch", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public java.lang.Object epochLock()
func (jbobject *MapOutputTracker) EpochLock() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "epochLock", "java/lang/Object")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public void sendTracker(java.lang.Object)
func (jbobject *MapOutputTracker) SendTracker(a interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "sendTracker", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public org.apache.spark.MapOutputStatistics getStatistics(org.apache.spark.ShuffleDependency<?, ?, ?>)
func (jbobject *MapOutputTracker) GetStatistics(a ShuffleDependencyInterface) *MapOutputStatistics {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatistics", "org/apache/spark/MapOutputStatistics", conv_a.Value().Cast("org/apache/spark/ShuffleDependency"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &MapOutputStatistics{}
	unique_x.Callable = dst
	return unique_x
}

// public long getEpoch()
func (jbobject *MapOutputTracker) GetEpoch() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEpoch", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public void updateEpoch(long)
func (jbobject *MapOutputTracker) UpdateEpoch(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "updateEpoch", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void unregisterShuffle(int)
func (jbobject *MapOutputTracker) UnregisterShuffle(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "unregisterShuffle", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void stop()
func (jbobject *MapOutputTracker) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}


