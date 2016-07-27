package gospark

import "github.com/timob/javabind"

type MemoryStorageMemoryPoolInterface interface {
	MemoryMemoryPoolInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public org.apache.spark.storage.MemoryStore memoryStore()
	MemoryStore() *StorageMemoryStore

	// public final void setMemoryStore(org.apache.spark.storage.MemoryStore)
	SetMemoryStore(a StorageMemoryStoreInterface) 

	// public void releaseMemory(long)
	ReleaseMemory(a int64) 

	// public void releaseAllMemory()
	ReleaseAllMemory() 

	// public long freeSpaceToShrinkPool(long)
	FreeSpaceToShrinkPool(a int64) int64
}

type MemoryStorageMemoryPool struct {
	MemoryMemoryPool
}

// public org.apache.spark.memory.StorageMemoryPool(java.lang.Object)
func NewMemoryStorageMemoryPool(a interface{}) (*MemoryStorageMemoryPool) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/memory/StorageMemoryPool", conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &MemoryStorageMemoryPool{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *MemoryStorageMemoryPool) LogName() string {
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
func (jbobject *MemoryStorageMemoryPool) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public long memoryUsed()
func (jbobject *MemoryStorageMemoryPool) MemoryUsed() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memoryUsed", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.storage.MemoryStore memoryStore()
func (jbobject *MemoryStorageMemoryPool) MemoryStore() *StorageMemoryStore {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memoryStore", "org/apache/spark/storage/MemoryStore")
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
	unique_x := &StorageMemoryStore{}
	unique_x.Callable = dst
	return unique_x
}

// public final void setMemoryStore(org.apache.spark.storage.MemoryStore)
func (jbobject *MemoryStorageMemoryPool) SetMemoryStore(a StorageMemoryStoreInterface)  {
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

// public void releaseMemory(long)
func (jbobject *MemoryStorageMemoryPool) ReleaseMemory(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "releaseMemory", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void releaseAllMemory()
func (jbobject *MemoryStorageMemoryPool) ReleaseAllMemory()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "releaseAllMemory", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public long freeSpaceToShrinkPool(long)
func (jbobject *MemoryStorageMemoryPool) FreeSpaceToShrinkPool(a int64) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "freeSpaceToShrinkPool", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}


