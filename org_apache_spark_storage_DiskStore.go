package gospark

import "github.com/timob/javabind"

type StorageDiskStoreInterface interface {
	StorageBlockStoreInterface

	// public long minMemoryMapBytes()
	MinMemoryMapBytes() int64
}

type StorageDiskStore struct {
	StorageBlockStore
}

// public org.apache.spark.storage.DiskStore(org.apache.spark.storage.BlockManager, org.apache.spark.storage.DiskBlockManager)
func NewStorageDiskStore(a StorageBlockManagerInterface, b StorageDiskBlockManagerInterface) (*StorageDiskStore) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/DiskStore", conv_a.Value().Cast("org/apache/spark/storage/BlockManager"), conv_b.Value().Cast("org/apache/spark/storage/DiskBlockManager"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &StorageDiskStore{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public long minMemoryMapBytes()
func (jbobject *StorageDiskStore) MinMemoryMapBytes() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "minMemoryMapBytes", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long getSize(org.apache.spark.storage.BlockId)
func (jbobject *StorageDiskStore) GetSize(a StorageBlockIdInterface) int64 {
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
func (jbobject *StorageDiskStore) PutArray(a StorageBlockIdInterface, b []*JavaLangObject, c StorageStorageLevelInterface, d bool) *StoragePutResult {
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
func (jbobject *StorageDiskStore) Remove(a StorageBlockIdInterface) bool {
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

// public boolean contains(org.apache.spark.storage.BlockId)
func (jbobject *StorageDiskStore) Contains(a StorageBlockIdInterface) bool {
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


