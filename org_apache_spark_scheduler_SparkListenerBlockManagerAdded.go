package gospark

import "github.com/timob/javabind"

type SchedulerSparkListenerBlockManagerAddedInterface interface {
	JavaLangObjectInterface

	// public long time()
	Time() int64

	// public org.apache.spark.storage.BlockManagerId blockManagerId()
	BlockManagerId() *StorageBlockManagerId

	// public long maxMem()
	MaxMem() int64

	// public org.apache.spark.scheduler.SparkListenerBlockManagerAdded copy(long, org.apache.spark.storage.BlockManagerId, long)
	Copy(a int64, b StorageBlockManagerIdInterface, c int64) *SchedulerSparkListenerBlockManagerAdded

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type SchedulerSparkListenerBlockManagerAdded struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.SparkListenerBlockManagerAdded(long, org.apache.spark.storage.BlockManagerId, long)
func NewSchedulerSparkListenerBlockManagerAdded(a int64, b StorageBlockManagerIdInterface, c int64) (*SchedulerSparkListenerBlockManagerAdded) {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/SparkListenerBlockManagerAdded", a, conv_b.Value().Cast("org/apache/spark/storage/BlockManagerId"), c)
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	x := &SchedulerSparkListenerBlockManagerAdded{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public long time()
func (jbobject *SchedulerSparkListenerBlockManagerAdded) Time() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "time", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.storage.BlockManagerId blockManagerId()
func (jbobject *SchedulerSparkListenerBlockManagerAdded) BlockManagerId() *StorageBlockManagerId {
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

// public long maxMem()
func (jbobject *SchedulerSparkListenerBlockManagerAdded) MaxMem() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "maxMem", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.scheduler.SparkListenerBlockManagerAdded copy(long, org.apache.spark.storage.BlockManagerId, long)
func (jbobject *SchedulerSparkListenerBlockManagerAdded) Copy(a int64, b StorageBlockManagerIdInterface, c int64) *SchedulerSparkListenerBlockManagerAdded {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/scheduler/SparkListenerBlockManagerAdded", a, conv_b.Value().Cast("org/apache/spark/storage/BlockManagerId"), c)
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
	unique_x := &SchedulerSparkListenerBlockManagerAdded{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *SchedulerSparkListenerBlockManagerAdded) ProductPrefix() string {
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
func (jbobject *SchedulerSparkListenerBlockManagerAdded) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *SchedulerSparkListenerBlockManagerAdded) ProductElement(a int) *JavaLangObject {
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
func (jbobject *SchedulerSparkListenerBlockManagerAdded) CanEqual(a interface{}) bool {
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
func (jbobject *SchedulerSparkListenerBlockManagerAdded) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *SchedulerSparkListenerBlockManagerAdded) ToString() string {
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
func (jbobject *SchedulerSparkListenerBlockManagerAdded) Equals(a interface{}) bool {
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


