package gospark

import "github.com/timob/javabind"

type StorageBlockUpdatedInfoInterface interface {
	JavaLangObjectInterface

	// public org.apache.spark.storage.BlockManagerId blockManagerId()
	BlockManagerId() *StorageBlockManagerId

	// public org.apache.spark.storage.BlockId blockId()
	BlockId() *StorageBlockId

	// public org.apache.spark.storage.StorageLevel storageLevel()
	StorageLevel() *StorageStorageLevel

	// public long memSize()
	MemSize() int64

	// public long diskSize()
	DiskSize() int64

	// public long externalBlockStoreSize()
	ExternalBlockStoreSize() int64

	// public org.apache.spark.storage.BlockUpdatedInfo copy(org.apache.spark.storage.BlockManagerId, org.apache.spark.storage.BlockId, org.apache.spark.storage.StorageLevel, long, long, long)
	Copy(a StorageBlockManagerIdInterface, b StorageBlockIdInterface, c StorageStorageLevelInterface, d int64, e int64, f int64) *StorageBlockUpdatedInfo

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type StorageBlockUpdatedInfo struct {
	JavaLangObject
}

// public org.apache.spark.storage.BlockUpdatedInfo(org.apache.spark.storage.BlockManagerId, org.apache.spark.storage.BlockId, org.apache.spark.storage.StorageLevel, long, long, long)
func NewStorageBlockUpdatedInfo(a StorageBlockManagerIdInterface, b StorageBlockIdInterface, c StorageStorageLevelInterface, d int64, e int64, f int64) (*StorageBlockUpdatedInfo) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
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

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/BlockUpdatedInfo", conv_a.Value().Cast("org/apache/spark/storage/BlockManagerId"), conv_b.Value().Cast("org/apache/spark/storage/BlockId"), conv_c.Value().Cast("org/apache/spark/storage/StorageLevel"), d, e, f)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &StorageBlockUpdatedInfo{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.storage.BlockManagerId blockManagerId()
func (jbobject *StorageBlockUpdatedInfo) BlockManagerId() *StorageBlockManagerId {
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

// public org.apache.spark.storage.BlockId blockId()
func (jbobject *StorageBlockUpdatedInfo) BlockId() *StorageBlockId {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "blockId", "org/apache/spark/storage/BlockId")
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
	unique_x := &StorageBlockId{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.storage.StorageLevel storageLevel()
func (jbobject *StorageBlockUpdatedInfo) StorageLevel() *StorageStorageLevel {
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
func (jbobject *StorageBlockUpdatedInfo) MemSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memSize", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long diskSize()
func (jbobject *StorageBlockUpdatedInfo) DiskSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "diskSize", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long externalBlockStoreSize()
func (jbobject *StorageBlockUpdatedInfo) ExternalBlockStoreSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "externalBlockStoreSize", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.storage.BlockUpdatedInfo copy(org.apache.spark.storage.BlockManagerId, org.apache.spark.storage.BlockId, org.apache.spark.storage.StorageLevel, long, long, long)
func (jbobject *StorageBlockUpdatedInfo) Copy(a StorageBlockManagerIdInterface, b StorageBlockIdInterface, c StorageStorageLevelInterface, d int64, e int64, f int64) *StorageBlockUpdatedInfo {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
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
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/storage/BlockUpdatedInfo", conv_a.Value().Cast("org/apache/spark/storage/BlockManagerId"), conv_b.Value().Cast("org/apache/spark/storage/BlockId"), conv_c.Value().Cast("org/apache/spark/storage/StorageLevel"), d, e, f)
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
	unique_x := &StorageBlockUpdatedInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *StorageBlockUpdatedInfo) ProductPrefix() string {
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
func (jbobject *StorageBlockUpdatedInfo) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *StorageBlockUpdatedInfo) ProductElement(a int) *JavaLangObject {
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
func (jbobject *StorageBlockUpdatedInfo) CanEqual(a interface{}) bool {
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
func (jbobject *StorageBlockUpdatedInfo) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *StorageBlockUpdatedInfo) ToString() string {
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
func (jbobject *StorageBlockUpdatedInfo) Equals(a interface{}) bool {
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


