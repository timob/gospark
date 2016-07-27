package gospark

import "github.com/timob/javabind"

type JavaNetURIInterface interface {
	JavaLangObjectInterface

	// public static java.net.URI create(java.lang.String)
	Create(a string) *JavaNetURI

	// public java.net.URI normalize()
	Normalize() *JavaNetURI

	// public java.net.URI resolve(java.net.URI)
	Resolve2(a JavaNetURIInterface) *JavaNetURI

	// public java.net.URI resolve(java.lang.String)
	Resolve(a string) *JavaNetURI

	// public java.net.URI relativize(java.net.URI)
	Relativize(a JavaNetURIInterface) *JavaNetURI

	// public java.lang.String getScheme()
	GetScheme() string

	// public boolean isAbsolute()
	IsAbsolute() bool

	// public boolean isOpaque()
	IsOpaque() bool

	// public java.lang.String getRawSchemeSpecificPart()
	GetRawSchemeSpecificPart() string

	// public java.lang.String getSchemeSpecificPart()
	GetSchemeSpecificPart() string

	// public java.lang.String getRawAuthority()
	GetRawAuthority() string

	// public java.lang.String getAuthority()
	GetAuthority() string

	// public java.lang.String getRawUserInfo()
	GetRawUserInfo() string

	// public java.lang.String getUserInfo()
	GetUserInfo() string

	// public java.lang.String getHost()
	GetHost() string

	// public int getPort()
	GetPort() int

	// public java.lang.String getRawPath()
	GetRawPath() string

	// public java.lang.String getPath()
	GetPath() string

	// public java.lang.String getRawQuery()
	GetRawQuery() string

	// public java.lang.String getQuery()
	GetQuery() string

	// public java.lang.String getRawFragment()
	GetRawFragment() string

	// public java.lang.String getFragment()
	GetFragment() string

	// public int compareTo(java.net.URI)
	CompareTo2(a JavaNetURIInterface) int

	// public java.lang.String toASCIIString()
	ToASCIIString() string

	// public int compareTo(java.lang.Object)
	CompareTo(a interface{}) int
}

type JavaNetURI struct {
	JavaLangObject
}

// public java.net.URI(java.lang.String) throws java.net.URISyntaxException
func NewJavaNetURI(a string) (*JavaNetURI, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/net/URI", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	x := &JavaNetURI{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.net.URI(java.lang.String, java.lang.String, java.lang.String, int, java.lang.String, java.lang.String, java.lang.String) throws java.net.URISyntaxException
func NewJavaNetURI5(a string, b string, c string, d int, e string, f string, g string) (*JavaNetURI, error) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	conv_e := javabind.NewGoToJavaString()
	conv_f := javabind.NewGoToJavaString()
	conv_g := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/net/URI", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"), d, conv_e.Value().Cast("java/lang/String"), conv_f.Value().Cast("java/lang/String"), conv_g.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &JavaNetURI{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.net.URI(java.lang.String, java.lang.String, java.lang.String, java.lang.String, java.lang.String) throws java.net.URISyntaxException
func NewJavaNetURI4(a string, b string, c string, d string, e string) (*JavaNetURI, error) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaString()
	conv_e := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/net/URI", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("java/lang/String"), conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	x := &JavaNetURI{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.net.URI(java.lang.String, java.lang.String, java.lang.String, java.lang.String) throws java.net.URISyntaxException
func NewJavaNetURI3(a string, b string, c string, d string) (*JavaNetURI, error) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/net/URI", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &JavaNetURI{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.net.URI(java.lang.String, java.lang.String, java.lang.String) throws java.net.URISyntaxException
func NewJavaNetURI2(a string, b string, c string) (*JavaNetURI, error) {
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

	obj, err := javabind.GetEnv().NewObject("java/net/URI", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &JavaNetURI{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public static java.net.URI create(java.lang.String)
func (jbobject *JavaNetURI) Create(a string) *JavaNetURI {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/net/URI", "create", "java/net/URI", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &JavaNetURI{}
	unique_x.Callable = dst
	return unique_x
}

// public java.net.URI parseServerAuthority() throws java.net.URISyntaxException
func (jbobject *JavaNetURI) ParseServerAuthority() (*JavaNetURI, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "parseServerAuthority", "java/net/URI")
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

// public java.net.URI normalize()
func (jbobject *JavaNetURI) Normalize() *JavaNetURI {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "normalize", "java/net/URI")
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
	unique_x := &JavaNetURI{}
	unique_x.Callable = dst
	return unique_x
}

// public java.net.URI resolve(java.net.URI)
func (jbobject *JavaNetURI) Resolve2(a JavaNetURIInterface) *JavaNetURI {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "resolve", "java/net/URI", conv_a.Value().Cast("java/net/URI"))
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
	unique_x := &JavaNetURI{}
	unique_x.Callable = dst
	return unique_x
}

// public java.net.URI resolve(java.lang.String)
func (jbobject *JavaNetURI) Resolve(a string) *JavaNetURI {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "resolve", "java/net/URI", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &JavaNetURI{}
	unique_x.Callable = dst
	return unique_x
}

// public java.net.URI relativize(java.net.URI)
func (jbobject *JavaNetURI) Relativize(a JavaNetURIInterface) *JavaNetURI {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "relativize", "java/net/URI", conv_a.Value().Cast("java/net/URI"))
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
	unique_x := &JavaNetURI{}
	unique_x.Callable = dst
	return unique_x
}

// public java.net.URL toURL() throws java.net.MalformedURLException
func (jbobject *JavaNetURI) ToURL() (*JavaNetURL, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toURL", "java/net/URL")
	if err != nil {
		var zero *JavaNetURL
		return zero, err
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
	return unique_x, nil
}

// public java.lang.String getScheme()
func (jbobject *JavaNetURI) GetScheme() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getScheme", "java/lang/String")
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

// public boolean isAbsolute()
func (jbobject *JavaNetURI) IsAbsolute() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isAbsolute", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isOpaque()
func (jbobject *JavaNetURI) IsOpaque() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isOpaque", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String getRawSchemeSpecificPart()
func (jbobject *JavaNetURI) GetRawSchemeSpecificPart() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRawSchemeSpecificPart", "java/lang/String")
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

// public java.lang.String getSchemeSpecificPart()
func (jbobject *JavaNetURI) GetSchemeSpecificPart() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSchemeSpecificPart", "java/lang/String")
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

// public java.lang.String getRawAuthority()
func (jbobject *JavaNetURI) GetRawAuthority() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRawAuthority", "java/lang/String")
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
func (jbobject *JavaNetURI) GetAuthority() string {
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

// public java.lang.String getRawUserInfo()
func (jbobject *JavaNetURI) GetRawUserInfo() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRawUserInfo", "java/lang/String")
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
func (jbobject *JavaNetURI) GetUserInfo() string {
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

// public java.lang.String getHost()
func (jbobject *JavaNetURI) GetHost() string {
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

// public int getPort()
func (jbobject *JavaNetURI) GetPort() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPort", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getRawPath()
func (jbobject *JavaNetURI) GetRawPath() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRawPath", "java/lang/String")
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
func (jbobject *JavaNetURI) GetPath() string {
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

// public java.lang.String getRawQuery()
func (jbobject *JavaNetURI) GetRawQuery() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRawQuery", "java/lang/String")
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

// public java.lang.String getQuery()
func (jbobject *JavaNetURI) GetQuery() string {
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

// public java.lang.String getRawFragment()
func (jbobject *JavaNetURI) GetRawFragment() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRawFragment", "java/lang/String")
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

// public java.lang.String getFragment()
func (jbobject *JavaNetURI) GetFragment() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFragment", "java/lang/String")
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
func (jbobject *JavaNetURI) Equals(a interface{}) bool {
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

// public int hashCode()
func (jbobject *JavaNetURI) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int compareTo(java.net.URI)
func (jbobject *JavaNetURI) CompareTo2(a JavaNetURIInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compareTo", javabind.Int, conv_a.Value().Cast("java/net/URI"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *JavaNetURI) ToString() string {
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

// public java.lang.String toASCIIString()
func (jbobject *JavaNetURI) ToASCIIString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toASCIIString", "java/lang/String")
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

// public int compareTo(java.lang.Object)
func (jbobject *JavaNetURI) CompareTo(a interface{}) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compareTo", javabind.Int, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}


