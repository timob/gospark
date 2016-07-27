package gospark

import "github.com/timob/javabind"

type UtilTimeStampedHashMapInterface interface {
	JavaLangObjectInterface

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public void clear()
	Clear() 

	// public void sizeHint(int)
	SizeHint(a int) 

	// public boolean isEmpty()
	IsEmpty() bool

	// public boolean contains(A)
	Contains(a JavaLangObjectInterface) bool

	// public boolean isDefinedAt(A)
	IsDefinedAt(a JavaLangObjectInterface) bool

	// public B default(A)
	Default(a JavaLangObjectInterface) *JavaLangObject

	// public java.lang.String stringPrefix()
	StringPrefix() string

	// public <B> void copyToArray(java.lang.Object, int, int)
	CopyToArray3(a interface{}, b int, c int) 

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool

	// public java.lang.Object view()
	View() *JavaLangObject

	// public final boolean isTraversableAgain()
	IsTraversableAgain() bool

	// public boolean hasDefiniteSize()
	HasDefiniteSize() bool

	// public boolean nonEmpty()
	NonEmpty() bool

	// public <B> void copyToArray(java.lang.Object, int)
	CopyToArray2(a interface{}, b int) 

	// public <B> void copyToArray(java.lang.Object)
	CopyToArray(a interface{}) 

	// public java.lang.String mkString(java.lang.String, java.lang.String, java.lang.String)
	MkString3(a string, b string, c string) string

	// public java.lang.String mkString(java.lang.String)
	MkString2(a string) string

	// public java.lang.String mkString()
	MkString() string

	// public void update(A, B)
	Update(a JavaLangObjectInterface, b JavaLangObjectInterface) 

	// public B apply(A)
	Apply(a JavaLangObjectInterface) *JavaLangObject

	// public int size()
	Size() int

	// public void clearOldValues(long)
	ClearOldValues(a int64) 

	// public java.lang.Object result()
	Result() *JavaLangObject

	// public java.lang.Object clone()
	Clone() *JavaLangObject
}

type UtilTimeStampedHashMap struct {
	JavaLangObject
}

// public org.apache.spark.util.TimeStampedHashMap(boolean)
func NewUtilTimeStampedHashMap(a bool) (*UtilTimeStampedHashMap) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/util/TimeStampedHashMap", a)
	if err != nil {
		panic(err)
	}
	x := &UtilTimeStampedHashMap{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String logName()
func (jbobject *UtilTimeStampedHashMap) LogName() string {
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
func (jbobject *UtilTimeStampedHashMap) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void clear()
func (jbobject *UtilTimeStampedHashMap) Clear()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clear", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void sizeHint(int)
func (jbobject *UtilTimeStampedHashMap) SizeHint(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "sizeHint", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean isEmpty()
func (jbobject *UtilTimeStampedHashMap) IsEmpty() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEmpty", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean contains(A)
func (jbobject *UtilTimeStampedHashMap) Contains(a JavaLangObjectInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "contains", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean isDefinedAt(A)
func (jbobject *UtilTimeStampedHashMap) IsDefinedAt(a JavaLangObjectInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDefinedAt", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public B default(A)
func (jbobject *UtilTimeStampedHashMap) Default(a JavaLangObjectInterface) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "default", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String stringPrefix()
func (jbobject *UtilTimeStampedHashMap) StringPrefix() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stringPrefix", "java/lang/String")
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

// public java.lang.String toString()
func (jbobject *UtilTimeStampedHashMap) ToString() string {
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
func (jbobject *UtilTimeStampedHashMap) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean equals(java.lang.Object)
func (jbobject *UtilTimeStampedHashMap) Equals(a interface{}) bool {
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

// public <B> void copyToArray(java.lang.Object, int, int)
func (jbobject *UtilTimeStampedHashMap) CopyToArray3(a interface{}, b int, c int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copyToArray", javabind.Void, conv_a.Value().Cast("java/lang/Object"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean canEqual(java.lang.Object)
func (jbobject *UtilTimeStampedHashMap) CanEqual(a interface{}) bool {
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

// public java.lang.Object view()
func (jbobject *UtilTimeStampedHashMap) View() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "view", "java/lang/Object")
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

// public final boolean isTraversableAgain()
func (jbobject *UtilTimeStampedHashMap) IsTraversableAgain() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraversableAgain", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean hasDefiniteSize()
func (jbobject *UtilTimeStampedHashMap) HasDefiniteSize() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hasDefiniteSize", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean nonEmpty()
func (jbobject *UtilTimeStampedHashMap) NonEmpty() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "nonEmpty", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public <B> void copyToArray(java.lang.Object, int)
func (jbobject *UtilTimeStampedHashMap) CopyToArray2(a interface{}, b int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copyToArray", javabind.Void, conv_a.Value().Cast("java/lang/Object"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public <B> void copyToArray(java.lang.Object)
func (jbobject *UtilTimeStampedHashMap) CopyToArray(a interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "copyToArray", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String mkString(java.lang.String, java.lang.String, java.lang.String)
func (jbobject *UtilTimeStampedHashMap) MkString3(a string, b string, c string) string {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mkString", "java/lang/String", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String mkString(java.lang.String)
func (jbobject *UtilTimeStampedHashMap) MkString2(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mkString", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String mkString()
func (jbobject *UtilTimeStampedHashMap) MkString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mkString", "java/lang/String")
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

// public void update(A, B)
func (jbobject *UtilTimeStampedHashMap) Update(a JavaLangObjectInterface, b JavaLangObjectInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "update", javabind.Void, conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public B apply(A)
func (jbobject *UtilTimeStampedHashMap) Apply(a JavaLangObjectInterface) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "apply", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public int size()
func (jbobject *UtilTimeStampedHashMap) Size() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "size", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void clearOldValues(long)
func (jbobject *UtilTimeStampedHashMap) ClearOldValues(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearOldValues", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public java.lang.Object result()
func (jbobject *UtilTimeStampedHashMap) Result() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "result", "java/lang/Object")
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

// public java.lang.Object clone()
func (jbobject *UtilTimeStampedHashMap) Clone() *JavaLangObject {
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


