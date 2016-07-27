package gospark

import "github.com/timob/javabind"

type StorageRDDInfoInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.storage.RDDInfo fromRdd(org.apache.spark.rdd.RDD<?>)
	FromRdd(a RddRDDInterface) *StorageRDDInfo

	// public int compareTo(java.lang.Object)
	CompareTo(a interface{}) int

	// public int id()
	Id() int

	// public java.lang.String name()
	Name() string

	// public int numPartitions()
	NumPartitions() int

	// public org.apache.spark.storage.StorageLevel storageLevel()
	StorageLevel() *StorageStorageLevel

	// public java.lang.String callSite()
	CallSite() string

	// public int numCachedPartitions()
	NumCachedPartitions() int

	// public long memSize()
	MemSize() int64

	// public long diskSize()
	DiskSize() int64

	// public long externalBlockStoreSize()
	ExternalBlockStoreSize() int64

	// public boolean isCached()
	IsCached() bool

	// public int compare(org.apache.spark.storage.RDDInfo)
	Compare2(a StorageRDDInfoInterface) int

	// public int compare(java.lang.Object)
	Compare(a interface{}) int
}

type StorageRDDInfo struct {
	JavaLangObject
}

// public static org.apache.spark.storage.RDDInfo fromRdd(org.apache.spark.rdd.RDD<?>)
func (jbobject *StorageRDDInfo) FromRdd(a RddRDDInterface) *StorageRDDInfo {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/RDDInfo", "fromRdd", "org/apache/spark/storage/RDDInfo", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
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
	unique_x := &StorageRDDInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public int compareTo(java.lang.Object)
func (jbobject *StorageRDDInfo) CompareTo(a interface{}) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compareTo", javabind.Int, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int id()
func (jbobject *StorageRDDInfo) Id() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "id", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String name()
func (jbobject *StorageRDDInfo) Name() string {
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

// public int numPartitions()
func (jbobject *StorageRDDInfo) NumPartitions() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numPartitions", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.storage.StorageLevel storageLevel()
func (jbobject *StorageRDDInfo) StorageLevel() *StorageStorageLevel {
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

// public java.lang.String callSite()
func (jbobject *StorageRDDInfo) CallSite() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "callSite", "java/lang/String")
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

// public int numCachedPartitions()
func (jbobject *StorageRDDInfo) NumCachedPartitions() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numCachedPartitions", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public long memSize()
func (jbobject *StorageRDDInfo) MemSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "memSize", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long diskSize()
func (jbobject *StorageRDDInfo) DiskSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "diskSize", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long externalBlockStoreSize()
func (jbobject *StorageRDDInfo) ExternalBlockStoreSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "externalBlockStoreSize", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public boolean isCached()
func (jbobject *StorageRDDInfo) IsCached() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isCached", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *StorageRDDInfo) ToString() string {
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

// public int compare(org.apache.spark.storage.RDDInfo)
func (jbobject *StorageRDDInfo) Compare2(a StorageRDDInfoInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compare", javabind.Int, conv_a.Value().Cast("org/apache/spark/storage/RDDInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int compare(java.lang.Object)
func (jbobject *StorageRDDInfo) Compare(a interface{}) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compare", javabind.Int, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}


