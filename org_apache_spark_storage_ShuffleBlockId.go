package gospark

import "github.com/timob/javabind"

type StorageShuffleBlockIdInterface interface {
	StorageBlockIdInterface

	// public int shuffleId()
	ShuffleId() int

	// public int mapId()
	MapId() int

	// public int reduceId()
	ReduceId() int

	// public org.apache.spark.storage.ShuffleBlockId copy(int, int, int)
	Copy(a int, b int, c int) *StorageShuffleBlockId

	// public java.lang.String productPrefix()
	ProductPrefix() string

	// public int productArity()
	ProductArity() int

	// public java.lang.Object productElement(int)
	ProductElement(a int) *JavaLangObject

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type StorageShuffleBlockId struct {
	StorageBlockId
}

// public org.apache.spark.storage.ShuffleBlockId(int, int, int)
func NewStorageShuffleBlockId(a int, b int, c int) (*StorageShuffleBlockId) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/ShuffleBlockId", a, b, c)
	if err != nil {
		panic(err)
	}
	x := &StorageShuffleBlockId{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int shuffleId()
func (jbobject *StorageShuffleBlockId) ShuffleId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "shuffleId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int mapId()
func (jbobject *StorageShuffleBlockId) MapId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mapId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int reduceId()
func (jbobject *StorageShuffleBlockId) ReduceId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "reduceId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String name()
func (jbobject *StorageShuffleBlockId) Name() string {
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

// public org.apache.spark.storage.ShuffleBlockId copy(int, int, int)
func (jbobject *StorageShuffleBlockId) Copy(a int, b int, c int) *StorageShuffleBlockId {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "copy", "org/apache/spark/storage/ShuffleBlockId", a, b, c)
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
	unique_x := &StorageShuffleBlockId{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String productPrefix()
func (jbobject *StorageShuffleBlockId) ProductPrefix() string {
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
func (jbobject *StorageShuffleBlockId) ProductArity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "productArity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.Object productElement(int)
func (jbobject *StorageShuffleBlockId) ProductElement(a int) *JavaLangObject {
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
func (jbobject *StorageShuffleBlockId) CanEqual(a interface{}) bool {
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


