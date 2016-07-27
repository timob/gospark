package gospark

import "github.com/timob/javabind"

type MemoryMemoryModeInterface interface {
	JavaLangEnumInterface

	// public static org.apache.spark.memory.MemoryMode[] values()
	Values() []*MemoryMemoryMode

	// public static org.apache.spark.memory.MemoryMode valueOf(java.lang.String)
	ValueOf(a string) *MemoryMemoryMode
}

type MemoryMemoryMode struct {
	JavaLangEnum
}

// public static org.apache.spark.memory.MemoryMode[] values()
func (jbobject *MemoryMemoryMode) Values() []*MemoryMemoryMode {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/memory/MemoryMode", "values", javabind.ObjectArrayType("org/apache/spark/memory/MemoryMode"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/memory/MemoryMode")
	dst := new([]*MemoryMemoryMode)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static org.apache.spark.memory.MemoryMode valueOf(java.lang.String)
func (jbobject *MemoryMemoryMode) ValueOf(a string) *MemoryMemoryMode {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/memory/MemoryMode", "valueOf", "org/apache/spark/memory/MemoryMode", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &MemoryMemoryMode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *MemoryMemoryMode) ON_HEAP() *MemoryMemoryMode {
	jret, err := javabind.GetEnv().GetStaticField("org/apache/spark/memory/MemoryMode", "ON_HEAP", "org/apache/spark/memory/MemoryMode")
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
	unique_x := &MemoryMemoryMode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *MemoryMemoryMode) SetFieldON_HEAP(val MemoryMemoryModeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/apache/spark/memory/MemoryMode", "ON_HEAP", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *MemoryMemoryMode) OFF_HEAP() *MemoryMemoryMode {
	jret, err := javabind.GetEnv().GetStaticField("org/apache/spark/memory/MemoryMode", "OFF_HEAP", "org/apache/spark/memory/MemoryMode")
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
	unique_x := &MemoryMemoryMode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *MemoryMemoryMode) SetFieldOFF_HEAP(val MemoryMemoryModeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/apache/spark/memory/MemoryMode", "OFF_HEAP", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


