package gospark

import "github.com/timob/javabind"

type SchedulerMapStatusInterface interface {

	// public abstract org.apache.spark.storage.BlockManagerId location()
	Location() *StorageBlockManagerId

	// public abstract long getSizeForBlock(int)
	GetSizeForBlock(a int) int64
}

type SchedulerMapStatus struct {
	JavaLangObject
}

// public abstract org.apache.spark.storage.BlockManagerId location()
func (jbobject *SchedulerMapStatus) Location() *StorageBlockManagerId {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "location", "org/apache/spark/storage/BlockManagerId")
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
	unique_x := &StorageBlockManagerId{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract long getSizeForBlock(int)
func (jbobject *SchedulerMapStatus) GetSizeForBlock(a int) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSizeForBlock", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}


