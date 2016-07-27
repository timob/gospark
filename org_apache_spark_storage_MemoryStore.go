package gospark

import "github.com/timob/javabind"

type StorageMemoryStoreInterface interface {
	StorageBlockStoreInterface

	// public void releaseUnrollMemoryForThisTask(long)
	ReleaseUnrollMemoryForThisTask(a int64) 

	// public void releasePendingUnrollMemoryForThisTask(long)
	ReleasePendingUnrollMemoryForThisTask(a int64) 

	// public long currentUnrollMemory()
	CurrentUnrollMemory() int64

	// public long currentUnrollMemoryForThisTask()
	CurrentUnrollMemoryForThisTask() int64
}

type StorageMemoryStore struct {
	StorageBlockStore
}

// public org.apache.spark.storage.MemoryStore(org.apache.spark.storage.BlockManager, org.apache.spark.memory.MemoryManager)
func NewStorageMemoryStore(a StorageBlockManagerInterface, b MemoryMemoryManagerInterface) (*StorageMemoryStore) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/MemoryStore", conv_a.Value().Cast("org/apache/spark/storage/BlockManager"), conv_b.Value().Cast("org/apache/spark/memory/MemoryManager"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &StorageMemoryStore{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public long getSize(org.apache.spark.storage.BlockId)
func (jbobject *StorageMemoryStore) GetSize(a StorageBlockIdInterface) int64 {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSize", javabind.Long, conv_a.Value().Cast("org/apache/spark/storage/BlockId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public org.apache.spark.storage.PutResult putArray(org.apache.spark.storage.BlockId, java.lang.Object[], org.apache.spark.storage.StorageLevel, boolean)
func (jbobject *StorageMemoryStore) PutArray(a StorageBlockIdInterface, b []*JavaLangObject, c StorageStorageLevelInterface, d bool) *StoragePutResult {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "java/lang/Object")
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
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "putArray", "org/apache/spark/storage/PutResult", conv_a.Value().Cast("org/apache/spark/storage/BlockId"), conv_b.Value().Cast("java/lang/Object"), conv_c.Value().Cast("org/apache/spark/storage/StorageLevel"), d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &StoragePutResult{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean remove(org.apache.spark.storage.BlockId)
func (jbobject *StorageMemoryStore) Remove(a StorageBlockIdInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "remove", javabind.Boolean, conv_a.Value().Cast("org/apache/spark/storage/BlockId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public void clear()
func (jbobject *StorageMemoryStore) Clear()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clear", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public boolean contains(org.apache.spark.storage.BlockId)
func (jbobject *StorageMemoryStore) Contains(a StorageBlockIdInterface) bool {
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

// public void releaseUnrollMemoryForThisTask(long)
func (jbobject *StorageMemoryStore) ReleaseUnrollMemoryForThisTask(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "releaseUnrollMemoryForThisTask", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void releasePendingUnrollMemoryForThisTask(long)
func (jbobject *StorageMemoryStore) ReleasePendingUnrollMemoryForThisTask(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "releasePendingUnrollMemoryForThisTask", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long currentUnrollMemory()
func (jbobject *StorageMemoryStore) CurrentUnrollMemory() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "currentUnrollMemory", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long currentUnrollMemoryForThisTask()
func (jbobject *StorageMemoryStore) CurrentUnrollMemoryForThisTask() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "currentUnrollMemoryForThisTask", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}


