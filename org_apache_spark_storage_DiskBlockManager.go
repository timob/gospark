package gospark

import "github.com/timob/javabind"

type StorageDiskBlockManagerInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public int subDirsPerLocalDir()
	SubDirsPerLocalDir() int

	// public boolean containsBlock(org.apache.spark.storage.BlockId)
	ContainsBlock(a StorageBlockIdInterface) bool

	// public void stop()
	Stop() 
}

type StorageDiskBlockManager struct {
	JavaLangObject
}

// public org.apache.spark.storage.DiskBlockManager(org.apache.spark.storage.BlockManager, org.apache.spark.SparkConf)
func NewStorageDiskBlockManager(a StorageBlockManagerInterface, b SparkConfInterface) (*StorageDiskBlockManager) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/storage/DiskBlockManager", conv_a.Value().Cast("org/apache/spark/storage/BlockManager"), conv_b.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &StorageDiskBlockManager{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *StorageDiskBlockManager) LogName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "logName", "java/lang/String")
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

// public boolean isTraceEnabled()
func (jbobject *StorageDiskBlockManager) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int subDirsPerLocalDir()
func (jbobject *StorageDiskBlockManager) SubDirsPerLocalDir() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subDirsPerLocalDir", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean containsBlock(org.apache.spark.storage.BlockId)
func (jbobject *StorageDiskBlockManager) ContainsBlock(a StorageBlockIdInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "containsBlock", javabind.Boolean, conv_a.Value().Cast("org/apache/spark/storage/BlockId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public void stop()
func (jbobject *StorageDiskBlockManager) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}


