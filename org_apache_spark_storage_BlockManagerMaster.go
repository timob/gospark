package gospark

import "github.com/timob/javabind"

type StorageBlockManagerMasterInterface interface {
	JavaLangObjectInterface

	// public static java.lang.String DRIVER_ENDPOINT_NAME()
	DRIVER_ENDPOINT_NAME() string

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public org.apache.spark.rpc.RpcEndpointRef driverEndpoint()
	DriverEndpoint() *RpcRpcEndpointRef

	// public org.apache.spark.rpc.RpcTimeout timeout()
	Timeout() *RpcRpcTimeout

	// public void removeExecutor(java.lang.String)
	RemoveExecutor(a string) 

	// public void registerBlockManager(org.apache.spark.storage.BlockManagerId, long, org.apache.spark.rpc.RpcEndpointRef)
	RegisterBlockManager(a StorageBlockManagerIdInterface, b int64, c RpcRpcEndpointRefInterface) 

	// public boolean updateBlockInfo(org.apache.spark.storage.BlockManagerId, org.apache.spark.storage.BlockId, org.apache.spark.storage.StorageLevel, long, long, long)
	UpdateBlockInfo(a StorageBlockManagerIdInterface, b StorageBlockIdInterface, c StorageStorageLevelInterface, d int64, e int64, f int64) bool

	// public boolean contains(org.apache.spark.storage.BlockId)
	Contains(a StorageBlockIdInterface) bool

	// public void removeBlock(org.apache.spark.storage.BlockId)
	RemoveBlock(a StorageBlockIdInterface) 

	// public void removeRdd(int, boolean)
	RemoveRdd(a int, b bool) 

	// public void removeShuffle(int, boolean)
	RemoveShuffle(a int, b bool) 

	// public void removeBroadcast(long, boolean, boolean)
	RemoveBroadcast(a int64, b bool, c bool) 

	// public org.apache.spark.storage.StorageStatus[] getStorageStatus()
	GetStorageStatus() []*StorageStorageStatus

	// public boolean hasCachedBlocks(java.lang.String)
	HasCachedBlocks(a string) bool

	// public void stop()
	Stop() 
}

type StorageBlockManagerMaster struct {
	JavaLangObject
}

// public org.apache.spark.storage.BlockManagerMaster(org.apache.spark.rpc.RpcEndpointRef, org.apache.spark.SparkConf, boolean)
func NewStorageBlockManagerMaster(a RpcRpcEndpointRefInterface, b SparkConfInterface, c bool) (*StorageBlockManagerMaster) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/BlockManagerMaster", conv_a.Value().Cast("org/apache/spark/rpc/RpcEndpointRef"), conv_b.Value().Cast("org/apache/spark/SparkConf"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &StorageBlockManagerMaster{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static java.lang.String DRIVER_ENDPOINT_NAME()
func (jbobject *StorageBlockManagerMaster) DRIVER_ENDPOINT_NAME() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/BlockManagerMaster", "DRIVER_ENDPOINT_NAME", "java/lang/String")
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
func (jbobject *StorageBlockManagerMaster) LogName() string {
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
func (jbobject *StorageBlockManagerMaster) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.rpc.RpcEndpointRef driverEndpoint()
func (jbobject *StorageBlockManagerMaster) DriverEndpoint() *RpcRpcEndpointRef {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "driverEndpoint", "org/apache/spark/rpc/RpcEndpointRef")
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

// public org.apache.spark.rpc.RpcTimeout timeout()
func (jbobject *StorageBlockManagerMaster) Timeout() *RpcRpcTimeout {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "timeout", "org/apache/spark/rpc/RpcTimeout")
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
	unique_x := &RpcRpcTimeout{}
	unique_x.Callable = dst
	return unique_x
}

// public void removeExecutor(java.lang.String)
func (jbobject *StorageBlockManagerMaster) RemoveExecutor(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeExecutor", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void registerBlockManager(org.apache.spark.storage.BlockManagerId, long, org.apache.spark.rpc.RpcEndpointRef)
func (jbobject *StorageBlockManagerMaster) RegisterBlockManager(a StorageBlockManagerIdInterface, b int64, c RpcRpcEndpointRefInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "registerBlockManager", javabind.Void, conv_a.Value().Cast("org/apache/spark/storage/BlockManagerId"), b, conv_c.Value().Cast("org/apache/spark/rpc/RpcEndpointRef"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_c.CleanUp()

}

// public boolean updateBlockInfo(org.apache.spark.storage.BlockManagerId, org.apache.spark.storage.BlockId, org.apache.spark.storage.StorageLevel, long, long, long)
func (jbobject *StorageBlockManagerMaster) UpdateBlockInfo(a StorageBlockManagerIdInterface, b StorageBlockIdInterface, c StorageStorageLevelInterface, d int64, e int64, f int64) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "updateBlockInfo", javabind.Boolean, conv_a.Value().Cast("org/apache/spark/storage/BlockManagerId"), conv_b.Value().Cast("org/apache/spark/storage/BlockId"), conv_c.Value().Cast("org/apache/spark/storage/StorageLevel"), d, e, f)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	return jret.(bool)
}

// public boolean contains(org.apache.spark.storage.BlockId)
func (jbobject *StorageBlockManagerMaster) Contains(a StorageBlockIdInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "contains", javabind.Boolean, conv_a.Value().Cast("org/apache/spark/storage/BlockId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public void removeBlock(org.apache.spark.storage.BlockId)
func (jbobject *StorageBlockManagerMaster) RemoveBlock(a StorageBlockIdInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeBlock", javabind.Void, conv_a.Value().Cast("org/apache/spark/storage/BlockId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeRdd(int, boolean)
func (jbobject *StorageBlockManagerMaster) RemoveRdd(a int, b bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeRdd", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void removeShuffle(int, boolean)
func (jbobject *StorageBlockManagerMaster) RemoveShuffle(a int, b bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeShuffle", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void removeBroadcast(long, boolean, boolean)
func (jbobject *StorageBlockManagerMaster) RemoveBroadcast(a int64, b bool, c bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeBroadcast", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public org.apache.spark.storage.StorageStatus[] getStorageStatus()
func (jbobject *StorageBlockManagerMaster) GetStorageStatus() []*StorageStorageStatus {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStorageStatus", javabind.ObjectArrayType("org/apache/spark/storage/StorageStatus"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/storage/StorageStatus")
	dst := new([]*StorageStorageStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean hasCachedBlocks(java.lang.String)
func (jbobject *StorageBlockManagerMaster) HasCachedBlocks(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hasCachedBlocks", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public void stop()
func (jbobject *StorageBlockManagerMaster) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}


