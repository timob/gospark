package gospark

import "github.com/timob/javabind"

type SchedulerSparkListenerBlockManagerRemovedInterface interface {
	JavaLangObjectInterface

	// public long time()
	Time() int64

	// public org.apache.spark.storage.BlockManagerId blockManagerId()
	BlockManagerId() *StorageBlockManagerId

	// public org.apache.spark.scheduler.SparkListenerBlockManagerRemoved copy(long, org.apache.spark.storage.BlockManagerId)
	Copy(a int64, b StorageBlockManagerIdInterface) *SchedulerSparkListenerBlockManagerRemoved

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type SchedulerSparkListenerBlockManagerRemoved struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.SparkListenerBlockManagerRemoved(long, org.apache.spark.storage.BlockManagerId)
func NewSchedulerSparkListenerBlockManagerRemoved(a int64, b StorageBlockManagerIdInterface) (*SchedulerSparkListenerBlockManagerRemoved) {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/SparkListenerBlockManagerRemoved", a, conv_b.Value().Cast("org/apache/spark/storage/BlockManagerId"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	x := &SchedulerSparkListenerBlockManagerRemoved{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public long time()
func (jbobject *SchedulerSparkListenerBlockManagerRemoved) Time() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "time", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.storage.BlockManagerId blockManagerId()
func (jbobject *SchedulerSparkListenerBlockManagerRemoved) BlockManagerId() *StorageBlockManagerId {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "blockManagerId", "org/apache/spark/storage/BlockManagerId")
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

// public org.apache.spark.scheduler.SparkListenerBlockManagerRemoved copy(long, org.apache.spark.storage.BlockManagerId)
func (jbobject *SchedulerSparkListenerBlockManagerRemoved) Copy(a int64, b StorageBlockManagerIdInterface) *SchedulerSparkListenerBlockManagerRemoved {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/scheduler/SparkListenerBlockManagerRemoved", a, conv_b.Value().Cast("org/apache/spark/storage/BlockManagerId"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &SchedulerSparkListenerBlockManagerRemoved{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *SchedulerSparkListenerBlockManagerRemoved) ProductPrefix() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productPrefix", "java/lang/String")
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

// public int productArity()
func (jbobject *SchedulerSparkListenerBlockManagerRemoved) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *SchedulerSparkListenerBlockManagerRemoved) ProductElement(a int) *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productElement", "java/lang/Object", a)
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

// public boolean canEqual(java.lang.Object)
func (jbobject *SchedulerSparkListenerBlockManagerRemoved) CanEqual(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "canEqual", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *SchedulerSparkListenerBlockManagerRemoved) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *SchedulerSparkListenerBlockManagerRemoved) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
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

// public boolean equals(java.lang.Object)
func (jbobject *SchedulerSparkListenerBlockManagerRemoved) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}


