package gospark

import "github.com/timob/javabind"

type StorageBlockInfoInterface interface {
	JavaLangObjectInterface

	// public org.apache.spark.storage.StorageLevel level()
	Level() *StorageStorageLevel

	// public boolean tellMaster()
	TellMaster() bool

	// public long size()
	Size() int64

	// public boolean waitForReady()
	WaitForReady() bool

	// public void markReady(long)
	MarkReady(a int64) 

	// public void markFailure()
	MarkFailure() 
}

type StorageBlockInfo struct {
	JavaLangObject
}

// public org.apache.spark.storage.BlockInfo(org.apache.spark.storage.StorageLevel, boolean)
func NewStorageBlockInfo(a StorageStorageLevelInterface, b bool) (*StorageBlockInfo) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/BlockInfo", conv_a.Value().Cast("org/apache/spark/storage/StorageLevel"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &StorageBlockInfo{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.storage.StorageLevel level()
func (jbobject *StorageBlockInfo) Level() *StorageStorageLevel {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "level", "org/apache/spark/storage/StorageLevel")
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
	unique_x := &StorageStorageLevel{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean tellMaster()
func (jbobject *StorageBlockInfo) TellMaster() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "tellMaster", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public long size()
func (jbobject *StorageBlockInfo) Size() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "size", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public boolean waitForReady()
func (jbobject *StorageBlockInfo) WaitForReady() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "waitForReady", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void markReady(long)
func (jbobject *StorageBlockInfo) MarkReady(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "markReady", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void markFailure()
func (jbobject *StorageBlockInfo) MarkFailure()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "markFailure", javabind.Void)
	if err != nil {
		panic(err)
	}

}


