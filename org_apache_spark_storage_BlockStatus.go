package gospark

import "github.com/timob/javabind"

type StorageBlockStatusInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.storage.BlockStatus empty()
	Empty() *StorageBlockStatus

	// public org.apache.spark.storage.StorageLevel storageLevel()
	StorageLevel() *StorageStorageLevel

	// public long memSize()
	MemSize() int64

	// public long diskSize()
	DiskSize() int64

	// public long externalBlockStoreSize()
	ExternalBlockStoreSize() int64

	// public boolean isCached()
	IsCached() bool

	// public org.apache.spark.storage.BlockStatus copy(org.apache.spark.storage.StorageLevel, long, long, long)
	Copy(a StorageStorageLevelInterface, b int64, c int64, d int64) *StorageBlockStatus

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type StorageBlockStatus struct {
	JavaLangObject
}

// public org.apache.spark.storage.BlockStatus(org.apache.spark.storage.StorageLevel, long, long, long)
func NewStorageBlockStatus(a StorageStorageLevelInterface, b int64, c int64, d int64) (*StorageBlockStatus) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/BlockStatus", conv_a.Value().Cast("org/apache/spark/storage/StorageLevel"), b, c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &StorageBlockStatus{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.apache.spark.storage.BlockStatus empty()
func (jbobject *StorageBlockStatus) Empty() *StorageBlockStatus {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/BlockStatus", "empty", "org/apache/spark/storage/BlockStatus")
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
	unique_x := &StorageBlockStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.storage.StorageLevel storageLevel()
func (jbobject *StorageBlockStatus) StorageLevel() *StorageStorageLevel {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "storageLevel", "org/apache/spark/storage/StorageLevel")
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

// public long memSize()
func (jbobject *StorageBlockStatus) MemSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memSize", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long diskSize()
func (jbobject *StorageBlockStatus) DiskSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "diskSize", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long externalBlockStoreSize()
func (jbobject *StorageBlockStatus) ExternalBlockStoreSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "externalBlockStoreSize", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public boolean isCached()
func (jbobject *StorageBlockStatus) IsCached() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isCached", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.storage.BlockStatus copy(org.apache.spark.storage.StorageLevel, long, long, long)
func (jbobject *StorageBlockStatus) Copy(a StorageStorageLevelInterface, b int64, c int64, d int64) *StorageBlockStatus {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/storage/BlockStatus", conv_a.Value().Cast("org/apache/spark/storage/StorageLevel"), b, c, d)
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
	unique_x := &StorageBlockStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *StorageBlockStatus) ProductPrefix() string {
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
func (jbobject *StorageBlockStatus) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *StorageBlockStatus) ProductElement(a int) *JavaLangObject {
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
func (jbobject *StorageBlockStatus) CanEqual(a interface{}) bool {
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
func (jbobject *StorageBlockStatus) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *StorageBlockStatus) ToString() string {
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
func (jbobject *StorageBlockStatus) Equals(a interface{}) bool {
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


