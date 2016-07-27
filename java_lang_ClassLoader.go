package gospark

import "github.com/timob/javabind"

type JavaLangClassLoaderInterface interface {
	JavaLangObjectInterface

	// public java.net.URL getResource(java.lang.String)
	GetResource(a string) *JavaNetURL

	// public static java.net.URL getSystemResource(java.lang.String)
	GetSystemResource(a string) *JavaNetURL

	// public final java.lang.ClassLoader getParent()
	GetParent() *JavaLangClassLoader

	// public static java.lang.ClassLoader getSystemClassLoader()
	GetSystemClassLoader() *JavaLangClassLoader

	// public synchronized void setDefaultAssertionStatus(boolean)
	SetDefaultAssertionStatus(a bool) 

	// public synchronized void setPackageAssertionStatus(java.lang.String, boolean)
	SetPackageAssertionStatus(a string, b bool) 

	// public synchronized void setClassAssertionStatus(java.lang.String, boolean)
	SetClassAssertionStatus(a string, b bool) 

	// public synchronized void clearAssertionStatus()
	ClearAssertionStatus() 
}

type JavaLangClassLoader struct {
	JavaLangObject
}

// public java.net.URL getResource(java.lang.String)
func (jbobject *JavaLangClassLoader) GetResource(a string) *JavaNetURL {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResource", "java/net/URL", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &JavaNetURL{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.Enumeration<java.net.URL> getResources(java.lang.String) throws java.io.IOException
func (jbobject *JavaLangClassLoader) GetResources(a string) (*JavaUtilEnumeration, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResources", "java/util/Enumeration", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *JavaUtilEnumeration
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaUtilEnumeration{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static java.net.URL getSystemResource(java.lang.String)
func (jbobject *JavaLangClassLoader) GetSystemResource(a string) *JavaNetURL {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/ClassLoader", "getSystemResource", "java/net/URL", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &JavaNetURL{}
	unique_x.Callable = dst
	return unique_x
}

// public static java.util.Enumeration<java.net.URL> getSystemResources(java.lang.String) throws java.io.IOException
func (jbobject *JavaLangClassLoader) GetSystemResources(a string) (*JavaUtilEnumeration, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/ClassLoader", "getSystemResources", "java/util/Enumeration", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *JavaUtilEnumeration
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaUtilEnumeration{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public final java.lang.ClassLoader getParent()
func (jbobject *JavaLangClassLoader) GetParent() *JavaLangClassLoader {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "java/lang/ClassLoader")
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
	unique_x := &JavaLangClassLoader{}
	unique_x.Callable = dst
	return unique_x
}

// public static java.lang.ClassLoader getSystemClassLoader()
func (jbobject *JavaLangClassLoader) GetSystemClassLoader() *JavaLangClassLoader {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/ClassLoader", "getSystemClassLoader", "java/lang/ClassLoader")
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
	unique_x := &JavaLangClassLoader{}
	unique_x.Callable = dst
	return unique_x
}

// public synchronized void setDefaultAssertionStatus(boolean)
func (jbobject *JavaLangClassLoader) SetDefaultAssertionStatus(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDefaultAssertionStatus", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public synchronized void setPackageAssertionStatus(java.lang.String, boolean)
func (jbobject *JavaLangClassLoader) SetPackageAssertionStatus(a string, b bool)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPackageAssertionStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized void setClassAssertionStatus(java.lang.String, boolean)
func (jbobject *JavaLangClassLoader) SetClassAssertionStatus(a string, b bool)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClassAssertionStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized void clearAssertionStatus()
func (jbobject *JavaLangClassLoader) ClearAssertionStatus()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearAssertionStatus", javabind.Void)
	if err != nil {
		panic(err)
	}

}


