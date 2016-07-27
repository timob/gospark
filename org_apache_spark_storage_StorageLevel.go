package gospark

import "github.com/timob/javabind"

type StorageStorageLevelInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.storage.StorageLevel apply(int, int)
	Apply(a int, b int) *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel apply(boolean, boolean, boolean, int)
	Apply2(a bool, b bool, c bool, d int) *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel apply(boolean, boolean, boolean, boolean, int)
	Apply3(a bool, b bool, c bool, d bool, e int) *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel fromString(java.lang.String)
	FromString(a string) *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel OFF_HEAP()
	OFF_HEAP() *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel MEMORY_AND_DISK_SER_2()
	MEMORY_AND_DISK_SER_2() *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel MEMORY_AND_DISK_SER()
	MEMORY_AND_DISK_SER() *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel MEMORY_AND_DISK_2()
	MEMORY_AND_DISK_2() *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel MEMORY_AND_DISK()
	MEMORY_AND_DISK() *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel MEMORY_ONLY_SER_2()
	MEMORY_ONLY_SER_2() *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel MEMORY_ONLY_SER()
	MEMORY_ONLY_SER() *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel MEMORY_ONLY_2()
	MEMORY_ONLY_2() *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel MEMORY_ONLY()
	MEMORY_ONLY() *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel DISK_ONLY_2()
	DISK_ONLY_2() *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel DISK_ONLY()
	DISK_ONLY() *StorageStorageLevel

	// public static org.apache.spark.storage.StorageLevel NONE()
	NONE() *StorageStorageLevel

	// public boolean useDisk()
	UseDisk() bool

	// public boolean useMemory()
	UseMemory() bool

	// public boolean useOffHeap()
	UseOffHeap() bool

	// public boolean deserialized()
	Deserialized() bool

	// public int replication()
	Replication() int

	// public org.apache.spark.storage.StorageLevel clone()
	Clone2() *StorageStorageLevel

	// public boolean isValid()
	IsValid() bool

	// public int toInt()
	ToInt() int

	// public java.lang.String description()
	Description() string

	// public java.lang.Object clone()
	Clone() *JavaLangObject
}

type StorageStorageLevel struct {
	JavaLangObject
}

// public org.apache.spark.storage.StorageLevel(boolean, boolean, boolean, boolean, int)
func NewStorageStorageLevel3(a bool, b bool, c bool, d bool, e int) (*StorageStorageLevel) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/StorageLevel", a, b, c, d, e)
	if err != nil {
		panic(err)
	}
	x := &StorageStorageLevel{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.storage.StorageLevel(int, int)
func NewStorageStorageLevel2(a int, b int) (*StorageStorageLevel) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/StorageLevel", a, b)
	if err != nil {
		panic(err)
	}
	x := &StorageStorageLevel{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.storage.StorageLevel()
func NewStorageStorageLevel() (*StorageStorageLevel) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/StorageLevel")
	if err != nil {
		panic(err)
	}
	x := &StorageStorageLevel{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.apache.spark.storage.StorageLevel apply(int, int)
func (jbobject *StorageStorageLevel) Apply(a int, b int) *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "apply", "org/apache/spark/storage/StorageLevel", a, b)
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

// public static org.apache.spark.storage.StorageLevel apply(boolean, boolean, boolean, int)
func (jbobject *StorageStorageLevel) Apply2(a bool, b bool, c bool, d int) *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "apply", "org/apache/spark/storage/StorageLevel", a, b, c, d)
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

// public static org.apache.spark.storage.StorageLevel apply(boolean, boolean, boolean, boolean, int)
func (jbobject *StorageStorageLevel) Apply3(a bool, b bool, c bool, d bool, e int) *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "apply", "org/apache/spark/storage/StorageLevel", a, b, c, d, e)
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

// public static org.apache.spark.storage.StorageLevel fromString(java.lang.String)
func (jbobject *StorageStorageLevel) FromString(a string) *StorageStorageLevel {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "fromString", "org/apache/spark/storage/StorageLevel", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &StorageStorageLevel{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.storage.StorageLevel OFF_HEAP()
func (jbobject *StorageStorageLevel) OFF_HEAP() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "OFF_HEAP", "org/apache/spark/storage/StorageLevel")
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

// public static org.apache.spark.storage.StorageLevel MEMORY_AND_DISK_SER_2()
func (jbobject *StorageStorageLevel) MEMORY_AND_DISK_SER_2() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "MEMORY_AND_DISK_SER_2", "org/apache/spark/storage/StorageLevel")
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

// public static org.apache.spark.storage.StorageLevel MEMORY_AND_DISK_SER()
func (jbobject *StorageStorageLevel) MEMORY_AND_DISK_SER() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "MEMORY_AND_DISK_SER", "org/apache/spark/storage/StorageLevel")
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

// public static org.apache.spark.storage.StorageLevel MEMORY_AND_DISK_2()
func (jbobject *StorageStorageLevel) MEMORY_AND_DISK_2() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "MEMORY_AND_DISK_2", "org/apache/spark/storage/StorageLevel")
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

// public static org.apache.spark.storage.StorageLevel MEMORY_AND_DISK()
func (jbobject *StorageStorageLevel) MEMORY_AND_DISK() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "MEMORY_AND_DISK", "org/apache/spark/storage/StorageLevel")
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

// public static org.apache.spark.storage.StorageLevel MEMORY_ONLY_SER_2()
func (jbobject *StorageStorageLevel) MEMORY_ONLY_SER_2() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "MEMORY_ONLY_SER_2", "org/apache/spark/storage/StorageLevel")
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

// public static org.apache.spark.storage.StorageLevel MEMORY_ONLY_SER()
func (jbobject *StorageStorageLevel) MEMORY_ONLY_SER() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "MEMORY_ONLY_SER", "org/apache/spark/storage/StorageLevel")
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

// public static org.apache.spark.storage.StorageLevel MEMORY_ONLY_2()
func (jbobject *StorageStorageLevel) MEMORY_ONLY_2() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "MEMORY_ONLY_2", "org/apache/spark/storage/StorageLevel")
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

// public static org.apache.spark.storage.StorageLevel MEMORY_ONLY()
func (jbobject *StorageStorageLevel) MEMORY_ONLY() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "MEMORY_ONLY", "org/apache/spark/storage/StorageLevel")
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

// public static org.apache.spark.storage.StorageLevel DISK_ONLY_2()
func (jbobject *StorageStorageLevel) DISK_ONLY_2() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "DISK_ONLY_2", "org/apache/spark/storage/StorageLevel")
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

// public static org.apache.spark.storage.StorageLevel DISK_ONLY()
func (jbobject *StorageStorageLevel) DISK_ONLY() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "DISK_ONLY", "org/apache/spark/storage/StorageLevel")
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

// public static org.apache.spark.storage.StorageLevel NONE()
func (jbobject *StorageStorageLevel) NONE() *StorageStorageLevel {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/storage/StorageLevel", "NONE", "org/apache/spark/storage/StorageLevel")
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

// public boolean useDisk()
func (jbobject *StorageStorageLevel) UseDisk() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "useDisk", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean useMemory()
func (jbobject *StorageStorageLevel) UseMemory() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "useMemory", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean useOffHeap()
func (jbobject *StorageStorageLevel) UseOffHeap() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "useOffHeap", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean deserialized()
func (jbobject *StorageStorageLevel) Deserialized() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "deserialized", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int replication()
func (jbobject *StorageStorageLevel) Replication() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "replication", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.storage.StorageLevel clone()
func (jbobject *StorageStorageLevel) Clone2() *StorageStorageLevel {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "org/apache/spark/storage/StorageLevel")
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

// public boolean equals(java.lang.Object)
func (jbobject *StorageStorageLevel) Equals(a interface{}) bool {
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

// public boolean isValid()
func (jbobject *StorageStorageLevel) IsValid() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isValid", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int toInt()
func (jbobject *StorageStorageLevel) ToInt() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toInt", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *StorageStorageLevel) ToString() string {
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
func (jbobject *StorageStorageLevel) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String description()
func (jbobject *StorageStorageLevel) Description() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "description", "java/lang/String")
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

// public java.lang.Object clone()
func (jbobject *StorageStorageLevel) Clone() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
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


