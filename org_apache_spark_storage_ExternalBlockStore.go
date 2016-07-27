package gospark

import "github.com/timob/javabind"

type StorageExternalBlockStoreInterface interface {
	StorageBlockStoreInterface

	// public static java.lang.String DEFAULT_BLOCK_MANAGER_NAME()
	DEFAULT_BLOCK_MANAGER_NAME() string

	// public static java.lang.String BLOCK_MANAGER_NAME()
	BLOCK_MANAGER_NAME() string

	// public static java.lang.String MASTER_URL()
	MASTER_URL() string

	// public static java.lang.String FOLD_NAME()
	FOLD_NAME() string

	// public static java.lang.String BASE_DIR()
	BASE_DIR() string

	// public static java.lang.String SUB_DIRS_PER_DIR()
	SUB_DIRS_PER_DIR() string

	// public static int MAX_DIR_CREATION_ATTEMPTS()
	MAX_DIR_CREATION_ATTEMPTS() int
}

type StorageExternalBlockStore struct {
	StorageBlockStore
}

// public org.apache.spark.storage.ExternalBlockStore(org.apache.spark.storage.BlockManager, java.lang.String)
func NewStorageExternalBlockStore(a StorageBlockManagerInterface, b string) (*StorageExternalBlockStore) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/ExternalBlockStore", conv_a.Value().Cast("org/apache/spark/storage/BlockManager"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &StorageExternalBlockStore{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static java.lang.String DEFAULT_BLOCK_MANAGER_NAME()
func (jbobject *StorageExternalBlockStore) DEFAULT_BLOCK_MANAGER_NAME() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/ExternalBlockStore", "DEFAULT_BLOCK_MANAGER_NAME", "java/lang/String")
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

// public static java.lang.String BLOCK_MANAGER_NAME()
func (jbobject *StorageExternalBlockStore) BLOCK_MANAGER_NAME() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/ExternalBlockStore", "BLOCK_MANAGER_NAME", "java/lang/String")
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

// public static java.lang.String MASTER_URL()
func (jbobject *StorageExternalBlockStore) MASTER_URL() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/ExternalBlockStore", "MASTER_URL", "java/lang/String")
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

// public static java.lang.String FOLD_NAME()
func (jbobject *StorageExternalBlockStore) FOLD_NAME() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/ExternalBlockStore", "FOLD_NAME", "java/lang/String")
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

// public static java.lang.String BASE_DIR()
func (jbobject *StorageExternalBlockStore) BASE_DIR() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/ExternalBlockStore", "BASE_DIR", "java/lang/String")
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

// public static java.lang.String SUB_DIRS_PER_DIR()
func (jbobject *StorageExternalBlockStore) SUB_DIRS_PER_DIR() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/ExternalBlockStore", "SUB_DIRS_PER_DIR", "java/lang/String")
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

// public static int MAX_DIR_CREATION_ATTEMPTS()
func (jbobject *StorageExternalBlockStore) MAX_DIR_CREATION_ATTEMPTS() int {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/ExternalBlockStore", "MAX_DIR_CREATION_ATTEMPTS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public long getSize(org.apache.spark.storage.BlockId)
func (jbobject *StorageExternalBlockStore) GetSize(a StorageBlockIdInterface) int64 {
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

// public org.apache.spark.storage.PutResult putArray(org.apache.spark.storage.BlockId, java.lang.Object[], org.apache.spark.storage.StorageLevel, boolean)
func (jbobject *StorageExternalBlockStore) PutArray(a StorageBlockIdInterface, b []*JavaLangObject, c StorageStorageLevelInterface, d bool) *StoragePutResult {
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

// public boolean remove(org.apache.spark.storage.BlockId)
func (jbobject *StorageExternalBlockStore) Remove(a StorageBlockIdInterface) bool {
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

// public boolean contains(org.apache.spark.storage.BlockId)
func (jbobject *StorageExternalBlockStore) Contains(a StorageBlockIdInterface) bool {
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


