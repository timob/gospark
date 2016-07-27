package gospark

import "github.com/timob/javabind"

type CacheManagerInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool
}

type CacheManager struct {
	JavaLangObject
}

// public org.apache.spark.CacheManager(org.apache.spark.storage.BlockManager)
func NewCacheManager(a StorageBlockManagerInterface) (*CacheManager) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/CacheManager", conv_a.Value().Cast("org/apache/spark/storage/BlockManager"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &CacheManager{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *CacheManager) LogName() string {
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
func (jbobject *CacheManager) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


