package gospark

import "github.com/timob/javabind"

type SparkEnvInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.SparkEnv getThreadLocal()
	GetThreadLocal() *SparkEnv

	// public static org.apache.spark.SparkEnv get()
	Get() *SparkEnv

	// public static void set(org.apache.spark.SparkEnv)
	Set(a SparkEnvInterface) 

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public java.lang.String executorId()
	ExecutorId() string

	// public org.apache.spark.rpc.RpcEnv rpcEnv()
	RpcEnv() *RpcRpcEnv

	// public org.apache.spark.serializer.Serializer serializer()
	Serializer() *SerializerSerializer

	// public org.apache.spark.serializer.Serializer closureSerializer()
	ClosureSerializer() *SerializerSerializer

	// public org.apache.spark.CacheManager cacheManager()
	CacheManager() *CacheManager

	// public org.apache.spark.MapOutputTracker mapOutputTracker()
	MapOutputTracker() *MapOutputTracker

	// public org.apache.spark.shuffle.ShuffleManager shuffleManager()
	ShuffleManager() *ShuffleShuffleManager

	// public org.apache.spark.broadcast.BroadcastManager broadcastManager()
	BroadcastManager() *BroadcastBroadcastManager

	// public org.apache.spark.network.BlockTransferService blockTransferService()
	BlockTransferService() *NetworkBlockTransferService

	// public org.apache.spark.storage.BlockManager blockManager()
	BlockManager() *StorageBlockManager

	// public org.apache.spark.SecurityManager securityManager()
	SecurityManager() *SecurityManager

	// public java.lang.String sparkFilesDir()
	SparkFilesDir() string

	// public org.apache.spark.metrics.MetricsSystem metricsSystem()
	MetricsSystem() *MetricsMetricsSystem

	// public org.apache.spark.memory.MemoryManager memoryManager()
	MemoryManager() *MemoryMemoryManager

	// public org.apache.spark.scheduler.OutputCommitCoordinator outputCommitCoordinator()
	OutputCommitCoordinator() *SchedulerOutputCommitCoordinator

	// public org.apache.spark.SparkConf conf()
	Conf() *SparkConf

	// public boolean isStopped()
	IsStopped() bool

	// public void stop()
	Stop() 
}

type SparkEnv struct {
	JavaLangObject
}

// public static org.apache.spark.SparkEnv getThreadLocal()
func (jbobject *SparkEnv) GetThreadLocal() *SparkEnv {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkEnv", "getThreadLocal", "org/apache/spark/SparkEnv")
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
	unique_x := &SparkEnv{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.SparkEnv get()
func (jbobject *SparkEnv) Get() *SparkEnv {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkEnv", "get", "org/apache/spark/SparkEnv")
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
	unique_x := &SparkEnv{}
	unique_x.Callable = dst
	return unique_x
}

// public static void set(org.apache.spark.SparkEnv)
func (jbobject *SparkEnv) Set(a SparkEnvInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkEnv", "set", javabind.Void, conv_a.Value().Cast("org/apache/spark/SparkEnv"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String logName()
func (jbobject *SparkEnv) LogName() string {
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
func (jbobject *SparkEnv) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String executorId()
func (jbobject *SparkEnv) ExecutorId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "executorId", "java/lang/String")
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

// public org.apache.spark.rpc.RpcEnv rpcEnv()
func (jbobject *SparkEnv) RpcEnv() *RpcRpcEnv {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "rpcEnv", "org/apache/spark/rpc/RpcEnv")
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
	unique_x := &RpcRpcEnv{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.serializer.Serializer serializer()
func (jbobject *SparkEnv) Serializer() *SerializerSerializer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "serializer", "org/apache/spark/serializer/Serializer")
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
	unique_x := &SerializerSerializer{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.serializer.Serializer closureSerializer()
func (jbobject *SparkEnv) ClosureSerializer() *SerializerSerializer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "closureSerializer", "org/apache/spark/serializer/Serializer")
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
	unique_x := &SerializerSerializer{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.CacheManager cacheManager()
func (jbobject *SparkEnv) CacheManager() *CacheManager {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "cacheManager", "org/apache/spark/CacheManager")
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
	unique_x := &CacheManager{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.MapOutputTracker mapOutputTracker()
func (jbobject *SparkEnv) MapOutputTracker() *MapOutputTracker {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapOutputTracker", "org/apache/spark/MapOutputTracker")
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
	unique_x := &MapOutputTracker{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.shuffle.ShuffleManager shuffleManager()
func (jbobject *SparkEnv) ShuffleManager() *ShuffleShuffleManager {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "shuffleManager", "org/apache/spark/shuffle/ShuffleManager")
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
	unique_x := &ShuffleShuffleManager{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.broadcast.BroadcastManager broadcastManager()
func (jbobject *SparkEnv) BroadcastManager() *BroadcastBroadcastManager {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "broadcastManager", "org/apache/spark/broadcast/BroadcastManager")
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
	unique_x := &BroadcastBroadcastManager{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.network.BlockTransferService blockTransferService()
func (jbobject *SparkEnv) BlockTransferService() *NetworkBlockTransferService {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "blockTransferService", "org/apache/spark/network/BlockTransferService")
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
	unique_x := &NetworkBlockTransferService{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.storage.BlockManager blockManager()
func (jbobject *SparkEnv) BlockManager() *StorageBlockManager {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "blockManager", "org/apache/spark/storage/BlockManager")
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
	unique_x := &StorageBlockManager{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SecurityManager securityManager()
func (jbobject *SparkEnv) SecurityManager() *SecurityManager {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "securityManager", "org/apache/spark/SecurityManager")
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
	unique_x := &SecurityManager{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String sparkFilesDir()
func (jbobject *SparkEnv) SparkFilesDir() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sparkFilesDir", "java/lang/String")
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

// public org.apache.spark.metrics.MetricsSystem metricsSystem()
func (jbobject *SparkEnv) MetricsSystem() *MetricsMetricsSystem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "metricsSystem", "org/apache/spark/metrics/MetricsSystem")
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
	unique_x := &MetricsMetricsSystem{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.memory.MemoryManager memoryManager()
func (jbobject *SparkEnv) MemoryManager() *MemoryMemoryManager {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memoryManager", "org/apache/spark/memory/MemoryManager")
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
	unique_x := &MemoryMemoryManager{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.scheduler.OutputCommitCoordinator outputCommitCoordinator()
func (jbobject *SparkEnv) OutputCommitCoordinator() *SchedulerOutputCommitCoordinator {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "outputCommitCoordinator", "org/apache/spark/scheduler/OutputCommitCoordinator")
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
	unique_x := &SchedulerOutputCommitCoordinator{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkConf conf()
func (jbobject *SparkEnv) Conf() *SparkConf {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "conf", "org/apache/spark/SparkConf")
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean isStopped()
func (jbobject *SparkEnv) IsStopped() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isStopped", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void stop()
func (jbobject *SparkEnv) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}


