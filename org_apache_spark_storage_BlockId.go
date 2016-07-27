package gospark

import "github.com/timob/javabind"

type StorageBlockIdInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.storage.BlockId apply(java.lang.String)
	Apply(a string) *StorageBlockId

	// public abstract java.lang.String name()
	Name() string

	// public boolean isRDD()
	IsRDD() bool

	// public boolean isShuffle()
	IsShuffle() bool

	// public boolean isBroadcast()
	IsBroadcast() bool
}

type StorageBlockId struct {
	JavaLangObject
}

// public org.apache.spark.storage.BlockId()
func NewStorageBlockId() (*StorageBlockId) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/BlockId")
	if err != nil {
		panic(err)
	}
	x := &StorageBlockId{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.apache.spark.storage.BlockId apply(java.lang.String)
func (jbobject *StorageBlockId) Apply(a string) *StorageBlockId {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/BlockId", "apply", "org/apache/spark/storage/BlockId", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &StorageBlockId{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract java.lang.String name()
func (jbobject *StorageBlockId) Name() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "name", "java/lang/String")
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

// public boolean isRDD()
func (jbobject *StorageBlockId) IsRDD() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isRDD", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isShuffle()
func (jbobject *StorageBlockId) IsShuffle() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isShuffle", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isBroadcast()
func (jbobject *StorageBlockId) IsBroadcast() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isBroadcast", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *StorageBlockId) ToString() string {
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
func (jbobject *StorageBlockId) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean equals(java.lang.Object)
func (jbobject *StorageBlockId) Equals(a interface{}) bool {
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


