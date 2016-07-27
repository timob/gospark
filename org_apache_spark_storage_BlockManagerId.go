package gospark

import "github.com/timob/javabind"

type StorageBlockManagerIdInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.storage.BlockManagerId getCachedBlockManagerId(org.apache.spark.storage.BlockManagerId)
	GetCachedBlockManagerId(a StorageBlockManagerIdInterface) *StorageBlockManagerId

	// public static org.apache.spark.storage.BlockManagerId apply(java.lang.String, java.lang.String, int)
	Apply(a string, b string, c int) *StorageBlockManagerId

	// public java.lang.String executorId()
	ExecutorId() string

	// public java.lang.String hostPort()
	HostPort() string

	// public java.lang.String host()
	Host() string

	// public int port()
	Port() int

	// public boolean isDriver()
	IsDriver() bool
}

type StorageBlockManagerId struct {
	JavaLangObject
}

// public org.apache.spark.storage.BlockManagerId(java.lang.String, java.lang.String, int)
func NewStorageBlockManagerId2(a string, b string, c int) (*StorageBlockManagerId) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/BlockManagerId", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &StorageBlockManagerId{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.storage.BlockManagerId()
func NewStorageBlockManagerId() (*StorageBlockManagerId) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/BlockManagerId")
	if err != nil {
		panic(err)
	}
	x := &StorageBlockManagerId{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.apache.spark.storage.BlockManagerId getCachedBlockManagerId(org.apache.spark.storage.BlockManagerId)
func (jbobject *StorageBlockManagerId) GetCachedBlockManagerId(a StorageBlockManagerIdInterface) *StorageBlockManagerId {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/BlockManagerId", "getCachedBlockManagerId", "org/apache/spark/storage/BlockManagerId", conv_a.Value().Cast("org/apache/spark/storage/BlockManagerId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
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

// public static org.apache.spark.storage.BlockManagerId apply(java.lang.String, java.lang.String, int)
func (jbobject *StorageBlockManagerId) Apply(a string, b string, c int) *StorageBlockManagerId {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/BlockManagerId", "apply", "org/apache/spark/storage/BlockManagerId", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
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

// public java.lang.String executorId()
func (jbobject *StorageBlockManagerId) ExecutorId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "executorId", "java/lang/String")
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

// public java.lang.String hostPort()
func (jbobject *StorageBlockManagerId) HostPort() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hostPort", "java/lang/String")
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

// public java.lang.String host()
func (jbobject *StorageBlockManagerId) Host() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "host", "java/lang/String")
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

// public int port()
func (jbobject *StorageBlockManagerId) Port() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "port", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean isDriver()
func (jbobject *StorageBlockManagerId) IsDriver() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDriver", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *StorageBlockManagerId) ToString() string {
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

// public int hashCode()
func (jbobject *StorageBlockManagerId) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean equals(java.lang.Object)
func (jbobject *StorageBlockManagerId) Equals(a interface{}) bool {
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


