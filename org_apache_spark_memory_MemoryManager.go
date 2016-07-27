package gospark

import "github.com/timob/javabind"

type MemoryMemoryManagerInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public org.apache.spark.memory.StorageMemoryPool storageMemoryPool()
	StorageMemoryPool() *MemoryStorageMemoryPool

	// public org.apache.spark.memory.ExecutionMemoryPool onHeapExecutionMemoryPool()
	OnHeapExecutionMemoryPool() *MemoryExecutionMemoryPool

	// public org.apache.spark.memory.ExecutionMemoryPool offHeapExecutionMemoryPool()
	OffHeapExecutionMemoryPool() *MemoryExecutionMemoryPool

	// public abstract long maxStorageMemory()
	MaxStorageMemory() int64

	// public final synchronized void setMemoryStore(org.apache.spark.storage.MemoryStore)
	SetMemoryStore(a StorageMemoryStoreInterface) 

	// public abstract long acquireExecutionMemory(long, long, org.apache.spark.memory.MemoryMode)
	AcquireExecutionMemory(a int64, b int64, c MemoryMemoryModeInterface) int64

	// public synchronized void releaseExecutionMemory(long, long, org.apache.spark.memory.MemoryMode)
	ReleaseExecutionMemory(a int64, b int64, c MemoryMemoryModeInterface) 

	// public synchronized long releaseAllExecutionMemoryForTask(long)
	ReleaseAllExecutionMemoryForTask(a int64) int64

	// public synchronized void releaseStorageMemory(long)
	ReleaseStorageMemory(a int64) 

	// public final synchronized void releaseAllStorageMemory()
	ReleaseAllStorageMemory() 

	// public final synchronized void releaseUnrollMemory(long)
	ReleaseUnrollMemory(a int64) 

	// public final synchronized long executionMemoryUsed()
	ExecutionMemoryUsed() int64

	// public final synchronized long storageMemoryUsed()
	StorageMemoryUsed() int64

	// public synchronized long getExecutionMemoryUsageForTask(long)
	GetExecutionMemoryUsageForTask(a int64) int64

	// public final org.apache.spark.memory.MemoryMode tungstenMemoryMode()
	TungstenMemoryMode() *MemoryMemoryMode

	// public long pageSizeBytes()
	PageSizeBytes() int64

	// public final org.apache.spark.unsafe.memory.MemoryAllocator tungstenMemoryAllocator()
	TungstenMemoryAllocator() *UnsafeMemoryMemoryAllocator
}

type MemoryMemoryManager struct {
	JavaLangObject
}

// public org.apache.spark.memory.MemoryManager(org.apache.spark.SparkConf, int, long, long)
func NewMemoryMemoryManager(a SparkConfInterface, b int, c int64, d int64) (*MemoryMemoryManager) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/memory/MemoryManager", conv_a.Value().Cast("org/apache/spark/SparkConf"), b, c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &MemoryMemoryManager{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *MemoryMemoryManager) LogName() string {
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
func (jbobject *MemoryMemoryManager) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.memory.StorageMemoryPool storageMemoryPool()
func (jbobject *MemoryMemoryManager) StorageMemoryPool() *MemoryStorageMemoryPool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "storageMemoryPool", "org/apache/spark/memory/StorageMemoryPool")
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
	unique_x := &MemoryStorageMemoryPool{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.memory.ExecutionMemoryPool onHeapExecutionMemoryPool()
func (jbobject *MemoryMemoryManager) OnHeapExecutionMemoryPool() *MemoryExecutionMemoryPool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "onHeapExecutionMemoryPool", "org/apache/spark/memory/ExecutionMemoryPool")
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
	unique_x := &MemoryExecutionMemoryPool{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.memory.ExecutionMemoryPool offHeapExecutionMemoryPool()
func (jbobject *MemoryMemoryManager) OffHeapExecutionMemoryPool() *MemoryExecutionMemoryPool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "offHeapExecutionMemoryPool", "org/apache/spark/memory/ExecutionMemoryPool")
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
	unique_x := &MemoryExecutionMemoryPool{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract long maxStorageMemory()
func (jbobject *MemoryMemoryManager) MaxStorageMemory() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "maxStorageMemory", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final synchronized void setMemoryStore(org.apache.spark.storage.MemoryStore)
func (jbobject *MemoryMemoryManager) SetMemoryStore(a StorageMemoryStoreInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMemoryStore", javabind.Void, conv_a.Value().Cast("org/apache/spark/storage/MemoryStore"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract long acquireExecutionMemory(long, long, org.apache.spark.memory.MemoryMode)
func (jbobject *MemoryMemoryManager) AcquireExecutionMemory(a int64, b int64, c MemoryMemoryModeInterface) int64 {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "acquireExecutionMemory", javabind.Long, a, b, conv_c.Value().Cast("org/apache/spark/memory/MemoryMode"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	return jret.(int64)
}

// public synchronized void releaseExecutionMemory(long, long, org.apache.spark.memory.MemoryMode)
func (jbobject *MemoryMemoryManager) ReleaseExecutionMemory(a int64, b int64, c MemoryMemoryModeInterface)  {
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "releaseExecutionMemory", javabind.Void, a, b, conv_c.Value().Cast("org/apache/spark/memory/MemoryMode"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()

}

// public synchronized long releaseAllExecutionMemoryForTask(long)
func (jbobject *MemoryMemoryManager) ReleaseAllExecutionMemoryForTask(a int64) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "releaseAllExecutionMemoryForTask", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public synchronized void releaseStorageMemory(long)
func (jbobject *MemoryMemoryManager) ReleaseStorageMemory(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "releaseStorageMemory", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public final synchronized void releaseAllStorageMemory()
func (jbobject *MemoryMemoryManager) ReleaseAllStorageMemory()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "releaseAllStorageMemory", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public final synchronized void releaseUnrollMemory(long)
func (jbobject *MemoryMemoryManager) ReleaseUnrollMemory(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "releaseUnrollMemory", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public final synchronized long executionMemoryUsed()
func (jbobject *MemoryMemoryManager) ExecutionMemoryUsed() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "executionMemoryUsed", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final synchronized long storageMemoryUsed()
func (jbobject *MemoryMemoryManager) StorageMemoryUsed() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "storageMemoryUsed", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public synchronized long getExecutionMemoryUsageForTask(long)
func (jbobject *MemoryMemoryManager) GetExecutionMemoryUsageForTask(a int64) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExecutionMemoryUsageForTask", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final org.apache.spark.memory.MemoryMode tungstenMemoryMode()
func (jbobject *MemoryMemoryManager) TungstenMemoryMode() *MemoryMemoryMode {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "tungstenMemoryMode", "org/apache/spark/memory/MemoryMode")
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
	unique_x := &MemoryMemoryMode{}
	unique_x.Callable = dst
	return unique_x
}

// public long pageSizeBytes()
func (jbobject *MemoryMemoryManager) PageSizeBytes() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "pageSizeBytes", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public final org.apache.spark.unsafe.memory.MemoryAllocator tungstenMemoryAllocator()
func (jbobject *MemoryMemoryManager) TungstenMemoryAllocator() *UnsafeMemoryMemoryAllocator {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "tungstenMemoryAllocator", "org/apache/spark/unsafe/memory/MemoryAllocator")
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
	unique_x := &UnsafeMemoryMemoryAllocator{}
	unique_x.Callable = dst
	return unique_x
}


