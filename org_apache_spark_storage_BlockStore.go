package gospark

import "github.com/timob/javabind"

type StorageBlockStoreInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public org.apache.spark.storage.BlockManager blockManager()
	BlockManager() *StorageBlockManager

	// public abstract org.apache.spark.storage.PutResult putArray(org.apache.spark.storage.BlockId, java.lang.Object[], org.apache.spark.storage.StorageLevel, boolean)
	PutArray(a StorageBlockIdInterface, b []*JavaLangObject, c StorageStorageLevelInterface, d bool) *StoragePutResult

	// public abstract long getSize(org.apache.spark.storage.BlockId)
	GetSize(a StorageBlockIdInterface) int64

	// public abstract boolean remove(org.apache.spark.storage.BlockId)
	Remove(a StorageBlockIdInterface) bool

	// public abstract boolean contains(org.apache.spark.storage.BlockId)
	Contains(a StorageBlockIdInterface) bool

	// public void clear()
	Clear() 
}

type StorageBlockStore struct {
	JavaLangObject
}

// public org.apache.spark.storage.BlockStore(org.apache.spark.storage.BlockManager)
func NewStorageBlockStore(a StorageBlockManagerInterface) (*StorageBlockStore) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/BlockStore", conv_a.Value().Cast("org/apache/spark/storage/BlockManager"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &StorageBlockStore{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *StorageBlockStore) LogName() string {
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
func (jbobject *StorageBlockStore) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.storage.BlockManager blockManager()
func (jbobject *StorageBlockStore) BlockManager() *StorageBlockManager {
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

// public abstract org.apache.spark.storage.PutResult putArray(org.apache.spark.storage.BlockId, java.lang.Object[], org.apache.spark.storage.StorageLevel, boolean)
func (jbobject *StorageBlockStore) PutArray(a StorageBlockIdInterface, b []*JavaLangObject, c StorageStorageLevelInterface, d bool) *StoragePutResult {
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

// public abstract long getSize(org.apache.spark.storage.BlockId)
func (jbobject *StorageBlockStore) GetSize(a StorageBlockIdInterface) int64 {
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

// public abstract boolean remove(org.apache.spark.storage.BlockId)
func (jbobject *StorageBlockStore) Remove(a StorageBlockIdInterface) bool {
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

// public abstract boolean contains(org.apache.spark.storage.BlockId)
func (jbobject *StorageBlockStore) Contains(a StorageBlockIdInterface) bool {
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

// public void clear()
func (jbobject *StorageBlockStore) Clear()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clear", javabind.Void)
	if err != nil {
		panic(err)
	}

}


