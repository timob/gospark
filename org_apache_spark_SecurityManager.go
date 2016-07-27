package gospark

import "github.com/timob/javabind"

type SecurityManagerInterface interface {
	JavaLangObjectInterface

	// public static java.lang.String SECRET_LOOKUP_KEY()
	SECRET_LOOKUP_KEY() string

	// public static java.lang.String ENV_AUTH_SECRET()
	ENV_AUTH_SECRET() string

	// public static java.lang.String SPARK_AUTH_SECRET_CONF()
	SPARK_AUTH_SECRET_CONF() string

	// public static java.lang.String SPARK_AUTH_CONF()
	SPARK_AUTH_CONF() string

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public org.apache.spark.SSLOptions fileServerSSLOptions()
	FileServerSSLOptions() *SSLOptions

	// public org.apache.spark.SSLOptions akkaSSLOptions()
	AkkaSSLOptions() *SSLOptions

	// public void setViewAcls(java.lang.String, java.lang.String)
	SetViewAcls(a string, b string) 

	// public java.lang.String getViewAcls()
	GetViewAcls() string

	// public java.lang.String getModifyAcls()
	GetModifyAcls() string

	// public void setAdminAcls(java.lang.String)
	SetAdminAcls(a string) 

	// public void setAcls(boolean)
	SetAcls(a bool) 

	// public boolean aclsEnabled()
	AclsEnabled() bool

	// public boolean checkUIViewPermissions(java.lang.String)
	CheckUIViewPermissions(a string) bool

	// public boolean checkModifyPermissions(java.lang.String)
	CheckModifyPermissions(a string) bool

	// public boolean isAuthenticationEnabled()
	IsAuthenticationEnabled() bool

	// public boolean isSaslEncryptionEnabled()
	IsSaslEncryptionEnabled() bool

	// public java.lang.String getHttpUser()
	GetHttpUser() string

	// public java.lang.String getSaslUser()
	GetSaslUser() string

	// public java.lang.String getSecretKey()
	GetSecretKey() string

	// public java.lang.String getSaslUser(java.lang.String)
	GetSaslUser2(a string) string

	// public java.lang.String getSecretKey(java.lang.String)
	GetSecretKey2(a string) string
}

type SecurityManager struct {
	JavaLangObject
}

// public org.apache.spark.SecurityManager(org.apache.spark.SparkConf)
func NewSecurityManager(a SparkConfInterface) (*SecurityManager) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/SecurityManager", conv_a.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &SecurityManager{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static java.lang.String SECRET_LOOKUP_KEY()
func (jbobject *SecurityManager) SECRET_LOOKUP_KEY() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SecurityManager", "SECRET_LOOKUP_KEY", "java/lang/String")
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

// public static java.lang.String ENV_AUTH_SECRET()
func (jbobject *SecurityManager) ENV_AUTH_SECRET() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SecurityManager", "ENV_AUTH_SECRET", "java/lang/String")
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

// public static java.lang.String SPARK_AUTH_SECRET_CONF()
func (jbobject *SecurityManager) SPARK_AUTH_SECRET_CONF() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SecurityManager", "SPARK_AUTH_SECRET_CONF", "java/lang/String")
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

// public static java.lang.String SPARK_AUTH_CONF()
func (jbobject *SecurityManager) SPARK_AUTH_CONF() string {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SecurityManager", "SPARK_AUTH_CONF", "java/lang/String")
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

// public java.lang.String logName()
func (jbobject *SecurityManager) LogName() string {
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
func (jbobject *SecurityManager) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.SSLOptions fileServerSSLOptions()
func (jbobject *SecurityManager) FileServerSSLOptions() *SSLOptions {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "fileServerSSLOptions", "org/apache/spark/SSLOptions")
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
	unique_x := &SSLOptions{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SSLOptions akkaSSLOptions()
func (jbobject *SecurityManager) AkkaSSLOptions() *SSLOptions {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "akkaSSLOptions", "org/apache/spark/SSLOptions")
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
	unique_x := &SSLOptions{}
	unique_x.Callable = dst
	return unique_x
}

// public void setViewAcls(java.lang.String, java.lang.String)
func (jbobject *SecurityManager) SetViewAcls(a string, b string)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setViewAcls", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public java.lang.String getViewAcls()
func (jbobject *SecurityManager) GetViewAcls() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getViewAcls", "java/lang/String")
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

// public java.lang.String getModifyAcls()
func (jbobject *SecurityManager) GetModifyAcls() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getModifyAcls", "java/lang/String")
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

// public void setAdminAcls(java.lang.String)
func (jbobject *SecurityManager) SetAdminAcls(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAdminAcls", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setAcls(boolean)
func (jbobject *SecurityManager) SetAcls(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAcls", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public boolean aclsEnabled()
func (jbobject *SecurityManager) AclsEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "aclsEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean checkUIViewPermissions(java.lang.String)
func (jbobject *SecurityManager) CheckUIViewPermissions(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "checkUIViewPermissions", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean checkModifyPermissions(java.lang.String)
func (jbobject *SecurityManager) CheckModifyPermissions(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "checkModifyPermissions", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean isAuthenticationEnabled()
func (jbobject *SecurityManager) IsAuthenticationEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isAuthenticationEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isSaslEncryptionEnabled()
func (jbobject *SecurityManager) IsSaslEncryptionEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isSaslEncryptionEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String getHttpUser()
func (jbobject *SecurityManager) GetHttpUser() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHttpUser", "java/lang/String")
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

// public java.lang.String getSaslUser()
func (jbobject *SecurityManager) GetSaslUser() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSaslUser", "java/lang/String")
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

// public java.lang.String getSecretKey()
func (jbobject *SecurityManager) GetSecretKey() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSecretKey", "java/lang/String")
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

// public java.lang.String getSaslUser(java.lang.String)
func (jbobject *SecurityManager) GetSaslUser2(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSaslUser", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
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

// public java.lang.String getSecretKey(java.lang.String)
func (jbobject *SecurityManager) GetSecretKey2(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSecretKey", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
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


