package gospark

import "github.com/timob/javabind"

type SparkConfInterface interface {
	JavaLangObjectInterface

	// public static void logDeprecationWarning(java.lang.String)
	LogDeprecationWarning(a string) 

	// public static boolean isSparkPortConf(java.lang.String)
	IsSparkPortConf(a string) bool

	// public static boolean isExecutorStartupConf(java.lang.String)
	IsExecutorStartupConf(a string) bool

	// public static boolean isAkkaConf(java.lang.String)
	IsAkkaConf(a string) bool

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public org.apache.spark.SparkConf set(java.lang.String, java.lang.String)
	Set(a string, b string) *SparkConf

	// public org.apache.spark.SparkConf setMaster(java.lang.String)
	SetMaster(a string) *SparkConf

	// public org.apache.spark.SparkConf setAppName(java.lang.String)
	SetAppName(a string) *SparkConf

	// public org.apache.spark.SparkConf setJars(java.lang.String[])
	SetJars(a []string) *SparkConf

	// public org.apache.spark.SparkConf setExecutorEnv(java.lang.String, java.lang.String)
	SetExecutorEnv(a string, b string) *SparkConf

	// public org.apache.spark.SparkConf setSparkHome(java.lang.String)
	SetSparkHome(a string) *SparkConf

	// public org.apache.spark.SparkConf setIfMissing(java.lang.String, java.lang.String)
	SetIfMissing(a string, b string) *SparkConf

	// public org.apache.spark.SparkConf remove(java.lang.String)
	Remove(a string) *SparkConf

	// public java.lang.String get(java.lang.String)
	Get(a string) string

	// public java.lang.String get(java.lang.String, java.lang.String)
	Get2(a string, b string) string

	// public long getTimeAsSeconds(java.lang.String)
	GetTimeAsSeconds(a string) int64

	// public long getTimeAsSeconds(java.lang.String, java.lang.String)
	GetTimeAsSeconds2(a string, b string) int64

	// public long getTimeAsMs(java.lang.String)
	GetTimeAsMs(a string) int64

	// public long getTimeAsMs(java.lang.String, java.lang.String)
	GetTimeAsMs2(a string, b string) int64

	// public long getSizeAsBytes(java.lang.String)
	GetSizeAsBytes(a string) int64

	// public long getSizeAsBytes(java.lang.String, java.lang.String)
	GetSizeAsBytes3(a string, b string) int64

	// public long getSizeAsBytes(java.lang.String, long)
	GetSizeAsBytes2(a string, b int64) int64

	// public long getSizeAsKb(java.lang.String)
	GetSizeAsKb(a string) int64

	// public long getSizeAsKb(java.lang.String, java.lang.String)
	GetSizeAsKb2(a string, b string) int64

	// public long getSizeAsMb(java.lang.String)
	GetSizeAsMb(a string) int64

	// public long getSizeAsMb(java.lang.String, java.lang.String)
	GetSizeAsMb2(a string, b string) int64

	// public long getSizeAsGb(java.lang.String)
	GetSizeAsGb(a string) int64

	// public long getSizeAsGb(java.lang.String, java.lang.String)
	GetSizeAsGb2(a string, b string) int64

	// public int getInt(java.lang.String, int)
	GetInt(a string, b int) int

	// public long getLong(java.lang.String, long)
	GetLong(a string, b int64) int64

	// public double getDouble(java.lang.String, double)
	GetDouble(a string, b float64) float64

	// public boolean getBoolean(java.lang.String, boolean)
	GetBoolean(a string, b bool) bool

	// public java.lang.String getAppId()
	GetAppId() string

	// public boolean contains(java.lang.String)
	Contains(a string) bool

	// public org.apache.spark.SparkConf clone()
	Clone2() *SparkConf

	// public java.lang.String getenv(java.lang.String)
	Getenv(a string) string

	// public void validateSettings()
	ValidateSettings() 

	// public java.lang.String toDebugString()
	ToDebugString() string

	// public java.lang.Object clone()
	Clone() *JavaLangObject
}

type SparkConf struct {
	JavaLangObject
}

// public org.apache.spark.SparkConf(boolean)
func NewSparkConf2(a bool) (*SparkConf) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/SparkConf", a)
	if err != nil {
		panic(err)
	}
	x := &SparkConf{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.SparkConf()
func NewSparkConf() (*SparkConf) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/SparkConf")
	if err != nil {
		panic(err)
	}
	x := &SparkConf{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static void logDeprecationWarning(java.lang.String)
func (jbobject *SparkConf) LogDeprecationWarning(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkConf", "logDeprecationWarning", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public static boolean isSparkPortConf(java.lang.String)
func (jbobject *SparkConf) IsSparkPortConf(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkConf", "isSparkPortConf", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public static boolean isExecutorStartupConf(java.lang.String)
func (jbobject *SparkConf) IsExecutorStartupConf(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkConf", "isExecutorStartupConf", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public static boolean isAkkaConf(java.lang.String)
func (jbobject *SparkConf) IsAkkaConf(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkConf", "isAkkaConf", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public java.lang.String logName()
func (jbobject *SparkConf) LogName() string {
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
func (jbobject *SparkConf) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.SparkConf set(java.lang.String, java.lang.String)
func (jbobject *SparkConf) Set(a string, b string) *SparkConf {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "set", "org/apache/spark/SparkConf", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkConf setMaster(java.lang.String)
func (jbobject *SparkConf) SetMaster(a string) *SparkConf {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setMaster", "org/apache/spark/SparkConf", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkConf setAppName(java.lang.String)
func (jbobject *SparkConf) SetAppName(a string) *SparkConf {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setAppName", "org/apache/spark/SparkConf", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkConf setJars(java.lang.String[])
func (jbobject *SparkConf) SetJars(a []string) *SparkConf {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setJars", "org/apache/spark/SparkConf", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkConf setExecutorEnv(java.lang.String, java.lang.String)
func (jbobject *SparkConf) SetExecutorEnv(a string, b string) *SparkConf {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setExecutorEnv", "org/apache/spark/SparkConf", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkConf setSparkHome(java.lang.String)
func (jbobject *SparkConf) SetSparkHome(a string) *SparkConf {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setSparkHome", "org/apache/spark/SparkConf", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkConf setIfMissing(java.lang.String, java.lang.String)
func (jbobject *SparkConf) SetIfMissing(a string, b string) *SparkConf {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setIfMissing", "org/apache/spark/SparkConf", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkConf remove(java.lang.String)
func (jbobject *SparkConf) Remove(a string) *SparkConf {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "remove", "org/apache/spark/SparkConf", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String get(java.lang.String)
func (jbobject *SparkConf) Get(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "get", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
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

// public java.lang.String get(java.lang.String, java.lang.String)
func (jbobject *SparkConf) Get2(a string, b string) string {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "get", "java/lang/String", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
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

// public long getTimeAsSeconds(java.lang.String)
func (jbobject *SparkConf) GetTimeAsSeconds(a string) int64 {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTimeAsSeconds", javabind.Long, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public long getTimeAsSeconds(java.lang.String, java.lang.String)
func (jbobject *SparkConf) GetTimeAsSeconds2(a string, b string) int64 {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTimeAsSeconds", javabind.Long, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(int64)
}

// public long getTimeAsMs(java.lang.String)
func (jbobject *SparkConf) GetTimeAsMs(a string) int64 {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTimeAsMs", javabind.Long, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public long getTimeAsMs(java.lang.String, java.lang.String)
func (jbobject *SparkConf) GetTimeAsMs2(a string, b string) int64 {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTimeAsMs", javabind.Long, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(int64)
}

// public long getSizeAsBytes(java.lang.String)
func (jbobject *SparkConf) GetSizeAsBytes(a string) int64 {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSizeAsBytes", javabind.Long, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public long getSizeAsBytes(java.lang.String, java.lang.String)
func (jbobject *SparkConf) GetSizeAsBytes3(a string, b string) int64 {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSizeAsBytes", javabind.Long, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(int64)
}

// public long getSizeAsBytes(java.lang.String, long)
func (jbobject *SparkConf) GetSizeAsBytes2(a string, b int64) int64 {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSizeAsBytes", javabind.Long, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public long getSizeAsKb(java.lang.String)
func (jbobject *SparkConf) GetSizeAsKb(a string) int64 {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSizeAsKb", javabind.Long, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public long getSizeAsKb(java.lang.String, java.lang.String)
func (jbobject *SparkConf) GetSizeAsKb2(a string, b string) int64 {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSizeAsKb", javabind.Long, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(int64)
}

// public long getSizeAsMb(java.lang.String)
func (jbobject *SparkConf) GetSizeAsMb(a string) int64 {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSizeAsMb", javabind.Long, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public long getSizeAsMb(java.lang.String, java.lang.String)
func (jbobject *SparkConf) GetSizeAsMb2(a string, b string) int64 {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSizeAsMb", javabind.Long, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(int64)
}

// public long getSizeAsGb(java.lang.String)
func (jbobject *SparkConf) GetSizeAsGb(a string) int64 {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSizeAsGb", javabind.Long, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public long getSizeAsGb(java.lang.String, java.lang.String)
func (jbobject *SparkConf) GetSizeAsGb2(a string, b string) int64 {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSizeAsGb", javabind.Long, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(int64)
}

// public int getInt(java.lang.String, int)
func (jbobject *SparkConf) GetInt(a string, b int) int {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInt", javabind.Int, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public long getLong(java.lang.String, long)
func (jbobject *SparkConf) GetLong(a string, b int64) int64 {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLong", javabind.Long, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int64)
}

// public double getDouble(java.lang.String, double)
func (jbobject *SparkConf) GetDouble(a string, b float64) float64 {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDouble", javabind.Double, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(float64)
}

// public boolean getBoolean(java.lang.String, boolean)
func (jbobject *SparkConf) GetBoolean(a string, b bool) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBoolean", javabind.Boolean, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public java.lang.String getAppId()
func (jbobject *SparkConf) GetAppId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAppId", "java/lang/String")
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

// public boolean contains(java.lang.String)
func (jbobject *SparkConf) Contains(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "contains", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public org.apache.spark.SparkConf clone()
func (jbobject *SparkConf) Clone2() *SparkConf {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "org/apache/spark/SparkConf")
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
	unique_x := &SparkConf{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getenv(java.lang.String)
func (jbobject *SparkConf) Getenv(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getenv", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
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

// public void validateSettings()
func (jbobject *SparkConf) ValidateSettings()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "validateSettings", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public java.lang.String toDebugString()
func (jbobject *SparkConf) ToDebugString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toDebugString", "java/lang/String")
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

// public java.lang.Object clone()
func (jbobject *SparkConf) Clone() *JavaLangObject {
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


