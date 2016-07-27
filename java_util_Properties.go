package gospark

import "github.com/timob/javabind"

type JavaUtilPropertiesInterface interface {
	JavaUtilHashtableInterface

	// public synchronized java.lang.Object setProperty(java.lang.String, java.lang.String)
	SetProperty(a string, b string) *JavaLangObject

	// public java.lang.String getProperty(java.lang.String)
	GetProperty(a string) string

	// public java.lang.String getProperty(java.lang.String, java.lang.String)
	GetProperty2(a string, b string) string

	// public java.util.Enumeration<?> propertyNames()
	PropertyNames() *JavaUtilEnumeration

	// public java.util.Set<java.lang.String> stringPropertyNames()
	StringPropertyNames() []string
}

type JavaUtilProperties struct {
	JavaUtilHashtable
}

// public java.util.Properties()
func NewJavaUtilProperties() (*JavaUtilProperties) {

	obj, err := javabind.GetEnv().NewObject("java/util/Properties")
	if err != nil {
		panic(err)
	}
	x := &JavaUtilProperties{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.Properties(java.util.Properties)
func NewJavaUtilProperties2(a JavaUtilPropertiesInterface) (*JavaUtilProperties) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/util/Properties", conv_a.Value().Cast("java/util/Properties"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaUtilProperties{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public synchronized java.lang.Object setProperty(java.lang.String, java.lang.String)
func (jbobject *JavaUtilProperties) SetProperty(a string, b string) *JavaLangObject {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setProperty", "java/lang/Object", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
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

// public java.lang.String getProperty(java.lang.String)
func (jbobject *JavaUtilProperties) GetProperty(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProperty", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
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

// public java.lang.String getProperty(java.lang.String, java.lang.String)
func (jbobject *JavaUtilProperties) GetProperty2(a string, b string) string {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProperty", "java/lang/String", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.Enumeration<?> propertyNames()
func (jbobject *JavaUtilProperties) PropertyNames() *JavaUtilEnumeration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "propertyNames", "java/util/Enumeration")
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
	unique_x := &JavaUtilEnumeration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.Set<java.lang.String> stringPropertyNames()
func (jbobject *JavaUtilProperties) StringPropertyNames() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stringPropertyNames", "java/util/Set")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoSet(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}


