package gospark

import "github.com/timob/javabind"

type JavaNetURLConnectionInterface interface {
	JavaLangObjectInterface

	// public static synchronized java.net.FileNameMap getFileNameMap()
	GetFileNameMap() *JavaNetFileNameMap

	// public static void setFileNameMap(java.net.FileNameMap)
	SetFileNameMap(a JavaNetFileNameMapInterface) 

	// public void setConnectTimeout(int)
	SetConnectTimeout(a int) 

	// public int getConnectTimeout()
	GetConnectTimeout() int

	// public void setReadTimeout(int)
	SetReadTimeout(a int) 

	// public int getReadTimeout()
	GetReadTimeout() int

	// public java.net.URL getURL()
	GetURL() *JavaNetURL

	// public int getContentLength()
	GetContentLength() int

	// public java.lang.String getContentType()
	GetContentType() string

	// public java.lang.String getContentEncoding()
	GetContentEncoding() string

	// public long getExpiration()
	GetExpiration() int64

	// public long getDate()
	GetDate() int64

	// public long getLastModified()
	GetLastModified() int64

	// public java.lang.String getHeaderField(java.lang.String)
	GetHeaderField2(a string) string

	// public java.util.Map<java.lang.String, java.util.List<java.lang.String>> getHeaderFields()
	GetHeaderFields() map[string]*[]string

	// public int getHeaderFieldInt(java.lang.String, int)
	GetHeaderFieldInt(a string, b int) int

	// public long getHeaderFieldDate(java.lang.String, long)
	GetHeaderFieldDate(a string, b int64) int64

	// public java.lang.String getHeaderFieldKey(int)
	GetHeaderFieldKey(a int) string

	// public java.lang.String getHeaderField(int)
	GetHeaderField(a int) string

	// public void setDoInput(boolean)
	SetDoInput(a bool) 

	// public boolean getDoInput()
	GetDoInput() bool

	// public void setDoOutput(boolean)
	SetDoOutput(a bool) 

	// public boolean getDoOutput()
	GetDoOutput() bool

	// public void setAllowUserInteraction(boolean)
	SetAllowUserInteraction(a bool) 

	// public boolean getAllowUserInteraction()
	GetAllowUserInteraction() bool

	// public static void setDefaultAllowUserInteraction(boolean)
	SetDefaultAllowUserInteraction(a bool) 

	// public static boolean getDefaultAllowUserInteraction()
	GetDefaultAllowUserInteraction() bool

	// public void setUseCaches(boolean)
	SetUseCaches(a bool) 

	// public boolean getUseCaches()
	GetUseCaches() bool

	// public void setIfModifiedSince(long)
	SetIfModifiedSince(a int64) 

	// public long getIfModifiedSince()
	GetIfModifiedSince() int64

	// public boolean getDefaultUseCaches()
	GetDefaultUseCaches() bool

	// public void setDefaultUseCaches(boolean)
	SetDefaultUseCaches(a bool) 

	// public void setRequestProperty(java.lang.String, java.lang.String)
	SetRequestProperty(a string, b string) 

	// public void addRequestProperty(java.lang.String, java.lang.String)
	AddRequestProperty(a string, b string) 

	// public java.lang.String getRequestProperty(java.lang.String)
	GetRequestProperty(a string) string

	// public java.util.Map<java.lang.String, java.util.List<java.lang.String>> getRequestProperties()
	GetRequestProperties() map[string]*[]string

	// public static void setDefaultRequestProperty(java.lang.String, java.lang.String)
	SetDefaultRequestProperty(a string, b string) 

	// public static java.lang.String getDefaultRequestProperty(java.lang.String)
	GetDefaultRequestProperty(a string) string

	// public static synchronized void setContentHandlerFactory(java.net.ContentHandlerFactory)
	SetContentHandlerFactory(a JavaNetContentHandlerFactoryInterface) 

	// public static java.lang.String guessContentTypeFromName(java.lang.String)
	GuessContentTypeFromName(a string) string
}

type JavaNetURLConnection struct {
	JavaLangObject
}

// public static synchronized java.net.FileNameMap getFileNameMap()
func (jbobject *JavaNetURLConnection) GetFileNameMap() *JavaNetFileNameMap {
	jret, err := javabind.GetEnv().CallStaticMethod("java/net/URLConnection", "getFileNameMap", "java/net/FileNameMap")
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
	unique_x := &JavaNetFileNameMap{}
	unique_x.Callable = dst
	return unique_x
}

// public static void setFileNameMap(java.net.FileNameMap)
func (jbobject *JavaNetURLConnection) SetFileNameMap(a JavaNetFileNameMapInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("java/net/URLConnection", "setFileNameMap", javabind.Void, conv_a.Value().Cast("java/net/FileNameMap"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void connect() throws java.io.IOException
func (jbobject *JavaNetURLConnection) Connect() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "connect", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}

// public void setConnectTimeout(int)
func (jbobject *JavaNetURLConnection) SetConnectTimeout(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setConnectTimeout", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public int getConnectTimeout()
func (jbobject *JavaNetURLConnection) GetConnectTimeout() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getConnectTimeout", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setReadTimeout(int)
func (jbobject *JavaNetURLConnection) SetReadTimeout(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReadTimeout", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public int getReadTimeout()
func (jbobject *JavaNetURLConnection) GetReadTimeout() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReadTimeout", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.net.URL getURL()
func (jbobject *JavaNetURLConnection) GetURL() *JavaNetURL {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getURL", "java/net/URL")
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
	unique_x := &JavaNetURL{}
	unique_x.Callable = dst
	return unique_x
}

// public int getContentLength()
func (jbobject *JavaNetURLConnection) GetContentLength() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getContentLength", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getContentType()
func (jbobject *JavaNetURLConnection) GetContentType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getContentType", "java/lang/String")
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

// public java.lang.String getContentEncoding()
func (jbobject *JavaNetURLConnection) GetContentEncoding() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getContentEncoding", "java/lang/String")
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

// public long getExpiration()
func (jbobject *JavaNetURLConnection) GetExpiration() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExpiration", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long getDate()
func (jbobject *JavaNetURLConnection) GetDate() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDate", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long getLastModified()
func (jbobject *JavaNetURLConnection) GetLastModified() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLastModified", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public java.lang.String getHeaderField(java.lang.String)
func (jbobject *JavaNetURLConnection) GetHeaderField2(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaderField", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
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

// public java.util.Map<java.lang.String, java.util.List<java.lang.String>> getHeaderFields()
func (jbobject *JavaNetURLConnection) GetHeaderFields() map[string]*[]string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaderFields", "java/util/Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoList(javabind.NewJavaToGoString()))
	dst := new(map[string]*[]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int getHeaderFieldInt(java.lang.String, int)
func (jbobject *JavaNetURLConnection) GetHeaderFieldInt(a string, b int) int {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaderFieldInt", javabind.Int, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public long getHeaderFieldDate(java.lang.String, long)
func (jbobject *JavaNetURLConnection) GetHeaderFieldDate(a string, b int64) int64 {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaderFieldDate", javabind.Long, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public java.lang.String getHeaderFieldKey(int)
func (jbobject *JavaNetURLConnection) GetHeaderFieldKey(a int) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaderFieldKey", "java/lang/String", a)
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

// public java.lang.String getHeaderField(int)
func (jbobject *JavaNetURLConnection) GetHeaderField(a int) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHeaderField", "java/lang/String", a)
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

// public java.lang.Object getContent() throws java.io.IOException
func (jbobject *JavaNetURLConnection) GetContent() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getContent", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
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
	return unique_x, nil
}

// public java.lang.String toString()
func (jbobject *JavaNetURLConnection) ToString() string {
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

// public void setDoInput(boolean)
func (jbobject *JavaNetURLConnection) SetDoInput(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDoInput", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean getDoInput()
func (jbobject *JavaNetURLConnection) GetDoInput() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDoInput", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setDoOutput(boolean)
func (jbobject *JavaNetURLConnection) SetDoOutput(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDoOutput", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean getDoOutput()
func (jbobject *JavaNetURLConnection) GetDoOutput() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDoOutput", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setAllowUserInteraction(boolean)
func (jbobject *JavaNetURLConnection) SetAllowUserInteraction(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAllowUserInteraction", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean getAllowUserInteraction()
func (jbobject *JavaNetURLConnection) GetAllowUserInteraction() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAllowUserInteraction", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static void setDefaultAllowUserInteraction(boolean)
func (jbobject *JavaNetURLConnection) SetDefaultAllowUserInteraction(a bool)  {
	_, err := javabind.GetEnv().CallStaticMethod("java/net/URLConnection", "setDefaultAllowUserInteraction", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public static boolean getDefaultAllowUserInteraction()
func (jbobject *JavaNetURLConnection) GetDefaultAllowUserInteraction() bool {
	jret, err := javabind.GetEnv().CallStaticMethod("java/net/URLConnection", "getDefaultAllowUserInteraction", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setUseCaches(boolean)
func (jbobject *JavaNetURLConnection) SetUseCaches(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUseCaches", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean getUseCaches()
func (jbobject *JavaNetURLConnection) GetUseCaches() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUseCaches", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setIfModifiedSince(long)
func (jbobject *JavaNetURLConnection) SetIfModifiedSince(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIfModifiedSince", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public long getIfModifiedSince()
func (jbobject *JavaNetURLConnection) GetIfModifiedSince() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIfModifiedSince", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public boolean getDefaultUseCaches()
func (jbobject *JavaNetURLConnection) GetDefaultUseCaches() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDefaultUseCaches", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setDefaultUseCaches(boolean)
func (jbobject *JavaNetURLConnection) SetDefaultUseCaches(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDefaultUseCaches", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setRequestProperty(java.lang.String, java.lang.String)
func (jbobject *JavaNetURLConnection) SetRequestProperty(a string, b string)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRequestProperty", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void addRequestProperty(java.lang.String, java.lang.String)
func (jbobject *JavaNetURLConnection) AddRequestProperty(a string, b string)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addRequestProperty", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public java.lang.String getRequestProperty(java.lang.String)
func (jbobject *JavaNetURLConnection) GetRequestProperty(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequestProperty", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
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

// public java.util.Map<java.lang.String, java.util.List<java.lang.String>> getRequestProperties()
func (jbobject *JavaNetURLConnection) GetRequestProperties() map[string]*[]string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequestProperties", "java/util/Map")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoMap(javabind.NewJavaToGoString(), javabind.NewJavaToGoList(javabind.NewJavaToGoString()))
	dst := new(map[string]*[]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static void setDefaultRequestProperty(java.lang.String, java.lang.String)
func (jbobject *JavaNetURLConnection) SetDefaultRequestProperty(a string, b string)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("java/net/URLConnection", "setDefaultRequestProperty", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public static java.lang.String getDefaultRequestProperty(java.lang.String)
func (jbobject *JavaNetURLConnection) GetDefaultRequestProperty(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/net/URLConnection", "getDefaultRequestProperty", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
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

// public static synchronized void setContentHandlerFactory(java.net.ContentHandlerFactory)
func (jbobject *JavaNetURLConnection) SetContentHandlerFactory(a JavaNetContentHandlerFactoryInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("java/net/URLConnection", "setContentHandlerFactory", javabind.Void, conv_a.Value().Cast("java/net/ContentHandlerFactory"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public static java.lang.String guessContentTypeFromName(java.lang.String)
func (jbobject *JavaNetURLConnection) GuessContentTypeFromName(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/net/URLConnection", "guessContentTypeFromName", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
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


