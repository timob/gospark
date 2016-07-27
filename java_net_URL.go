package gospark

import "github.com/timob/javabind"

type JavaNetURLInterface interface {
	JavaLangObjectInterface

	// public java.lang.String getQuery()
	GetQuery() string

	// public java.lang.String getPath()
	GetPath() string

	// public java.lang.String getUserInfo()
	GetUserInfo() string

	// public java.lang.String getAuthority()
	GetAuthority() string

	// public int getPort()
	GetPort() int

	// public int getDefaultPort()
	GetDefaultPort() int

	// public java.lang.String getProtocol()
	GetProtocol() string

	// public java.lang.String getHost()
	GetHost() string

	// public java.lang.String getFile()
	GetFile() string

	// public java.lang.String getRef()
	GetRef() string

	// public boolean sameFile(java.net.URL)
	SameFile(a JavaNetURLInterface) bool

	// public java.lang.String toExternalForm()
	ToExternalForm() string

	// public static void setURLStreamHandlerFactory(java.net.URLStreamHandlerFactory)
	SetURLStreamHandlerFactory(a JavaNetURLStreamHandlerFactoryInterface) 
}

type JavaNetURL struct {
	JavaLangObject
}

// public java.net.URL(java.lang.String, java.lang.String, int, java.lang.String) throws java.net.MalformedURLException
func NewJavaNetURL5(a string, b string, c int, d string) (*JavaNetURL, error) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/net/URL", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), c, conv_d.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_d.CleanUp()
	x := &JavaNetURL{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.net.URL(java.lang.String, java.lang.String, java.lang.String) throws java.net.MalformedURLException
func NewJavaNetURL3(a string, b string, c string) (*JavaNetURL, error) {
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

	obj, err := javabind.GetEnv().NewObject("java/net/URL", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &JavaNetURL{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.net.URL(java.lang.String, java.lang.String, int, java.lang.String, java.net.URLStreamHandler) throws java.net.MalformedURLException
func NewJavaNetURL6(a string, b string, c int, d string, e JavaNetURLStreamHandlerInterface) (*JavaNetURL, error) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaString()
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/net/URL", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), c, conv_d.Value().Cast("java/lang/String"), conv_e.Value().Cast("java/net/URLStreamHandler"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	x := &JavaNetURL{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.net.URL(java.lang.String) throws java.net.MalformedURLException
func NewJavaNetURL(a string) (*JavaNetURL, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/net/URL", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	x := &JavaNetURL{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.net.URL(java.net.URL, java.lang.String) throws java.net.MalformedURLException
func NewJavaNetURL2(a JavaNetURLInterface, b string) (*JavaNetURL, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/net/URL", conv_a.Value().Cast("java/net/URL"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaNetURL{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.net.URL(java.net.URL, java.lang.String, java.net.URLStreamHandler) throws java.net.MalformedURLException
func NewJavaNetURL4(a JavaNetURLInterface, b string, c JavaNetURLStreamHandlerInterface) (*JavaNetURL, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
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

	obj, err := javabind.GetEnv().NewObject("java/net/URL", conv_a.Value().Cast("java/net/URL"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/net/URLStreamHandler"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &JavaNetURL{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.lang.String getQuery()
func (jbobject *JavaNetURL) GetQuery() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getQuery", "java/lang/String")
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

// public java.lang.String getPath()
func (jbobject *JavaNetURL) GetPath() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPath", "java/lang/String")
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

// public java.lang.String getUserInfo()
func (jbobject *JavaNetURL) GetUserInfo() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserInfo", "java/lang/String")
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

// public java.lang.String getAuthority()
func (jbobject *JavaNetURL) GetAuthority() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAuthority", "java/lang/String")
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

// public int getPort()
func (jbobject *JavaNetURL) GetPort() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPort", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getDefaultPort()
func (jbobject *JavaNetURL) GetDefaultPort() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDefaultPort", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getProtocol()
func (jbobject *JavaNetURL) GetProtocol() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProtocol", "java/lang/String")
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

// public java.lang.String getHost()
func (jbobject *JavaNetURL) GetHost() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHost", "java/lang/String")
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

// public java.lang.String getFile()
func (jbobject *JavaNetURL) GetFile() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFile", "java/lang/String")
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

// public java.lang.String getRef()
func (jbobject *JavaNetURL) GetRef() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRef", "java/lang/String")
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

// public boolean equals(java.lang.Object)
func (jbobject *JavaNetURL) Equals(a interface{}) bool {
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

// public synchronized int hashCode()
func (jbobject *JavaNetURL) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean sameFile(java.net.URL)
func (jbobject *JavaNetURL) SameFile(a JavaNetURLInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sameFile", javabind.Boolean, conv_a.Value().Cast("java/net/URL"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *JavaNetURL) ToString() string {
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

// public java.lang.String toExternalForm()
func (jbobject *JavaNetURL) ToExternalForm() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toExternalForm", "java/lang/String")
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

// public java.net.URI toURI() throws java.net.URISyntaxException
func (jbobject *JavaNetURL) ToURI() (*JavaNetURI, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toURI", "java/net/URI")
	if err != nil {
		var zero *JavaNetURI
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaNetURI{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public java.net.URLConnection openConnection() throws java.io.IOException
func (jbobject *JavaNetURL) OpenConnection() (*JavaNetURLConnection, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "openConnection", "java/net/URLConnection")
	if err != nil {
		var zero *JavaNetURLConnection
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaNetURLConnection{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public java.net.URLConnection openConnection(java.net.Proxy) throws java.io.IOException
func (jbobject *JavaNetURL) OpenConnection2(a JavaNetProxyInterface) (*JavaNetURLConnection, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "openConnection", "java/net/URLConnection", conv_a.Value().Cast("java/net/Proxy"))
	if err != nil {
		var zero *JavaNetURLConnection
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
	unique_x := &JavaNetURLConnection{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public final java.lang.Object getContent() throws java.io.IOException
func (jbobject *JavaNetURL) GetContent() (*JavaLangObject, error) {
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

// public static void setURLStreamHandlerFactory(java.net.URLStreamHandlerFactory)
func (jbobject *JavaNetURL) SetURLStreamHandlerFactory(a JavaNetURLStreamHandlerFactoryInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("java/net/URL", "setURLStreamHandlerFactory", javabind.Void, conv_a.Value().Cast("java/net/URLStreamHandlerFactory"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


