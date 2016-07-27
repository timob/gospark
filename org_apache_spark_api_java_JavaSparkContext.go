package gospark

import "github.com/timob/javabind"

type ApiJavaJavaSparkContextInterface interface {
	ApiJavaJavaSparkContextVarargsWorkaroundInterface

	// public static java.lang.String[] jarOfObject(java.lang.Object)
	JarOfObject(a interface{}) []string

	// public static org.apache.spark.SparkContext toSparkContext(org.apache.spark.api.java.JavaSparkContext)
	ToSparkContext(a ApiJavaJavaSparkContextInterface) *SparkContext

	// public static org.apache.spark.api.java.JavaSparkContext fromSparkContext(org.apache.spark.SparkContext)
	FromSparkContext(a SparkContextInterface) *ApiJavaJavaSparkContext

	// public org.apache.spark.SparkContext sc()
	Sc() *SparkContext

	// public org.apache.spark.SparkEnv env()
	Env() *SparkEnv

	// public org.apache.spark.api.java.JavaSparkStatusTracker statusTracker()
	StatusTracker() *ApiJavaJavaSparkStatusTracker

	// public java.lang.Boolean isLocal()
	IsLocal() bool

	// public java.lang.String sparkUser()
	SparkUser() string

	// public java.lang.String master()
	Master() string

	// public java.lang.String appName()
	AppName() string

	// public java.util.List<java.lang.String> jars()
	Jars() []string

	// public java.lang.Long startTime()
	StartTime() int64

	// public java.lang.String version()
	Version() string

	// public java.lang.Integer defaultParallelism()
	DefaultParallelism() int

	// public java.lang.Integer defaultMinSplits()
	DefaultMinSplits() int

	// public java.lang.Integer defaultMinPartitions()
	DefaultMinPartitions() int

	// public <T> org.apache.spark.api.java.JavaRDD<T> parallelize(java.util.List<T>, int)
	Parallelize2(a []*JavaLangObject, b int) *ApiJavaJavaRDD

	// public <T> org.apache.spark.api.java.JavaRDD<T> emptyRDD()
	EmptyRDD() *ApiJavaJavaRDD

	// public <T> org.apache.spark.api.java.JavaRDD<T> parallelize(java.util.List<T>)
	Parallelize(a []*JavaLangObject) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaDoubleRDD parallelizeDoubles(java.util.List<java.lang.Double>, int)
	ParallelizeDoubles2(a []float64, b int) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaDoubleRDD parallelizeDoubles(java.util.List<java.lang.Double>)
	ParallelizeDoubles(a []float64) *ApiJavaJavaDoubleRDD

	// public org.apache.spark.api.java.JavaRDD<java.lang.String> textFile(java.lang.String)
	TextFile(a string) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaRDD<java.lang.String> textFile(java.lang.String, int)
	TextFile2(a string, b int) *ApiJavaJavaRDD

	// public org.apache.spark.api.java.JavaPairRDD<java.lang.String, java.lang.String> wholeTextFiles(java.lang.String, int)
	WholeTextFiles2(a string, b int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<java.lang.String, java.lang.String> wholeTextFiles(java.lang.String)
	WholeTextFiles(a string) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<java.lang.String, org.apache.spark.input.PortableDataStream> binaryFiles(java.lang.String, int)
	BinaryFiles2(a string, b int) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaPairRDD<java.lang.String, org.apache.spark.input.PortableDataStream> binaryFiles(java.lang.String)
	BinaryFiles(a string) *ApiJavaJavaPairRDD

	// public org.apache.spark.api.java.JavaRDD<byte[]> binaryRecords(java.lang.String, int)
	BinaryRecords(a string, b int) *ApiJavaJavaRDD

	// public <T> org.apache.spark.api.java.JavaRDD<T> objectFile(java.lang.String, int)
	ObjectFile2(a string, b int) *ApiJavaJavaRDD

	// public <T> org.apache.spark.api.java.JavaRDD<T> objectFile(java.lang.String)
	ObjectFile(a string) *ApiJavaJavaRDD

	// public org.apache.spark.Accumulator<java.lang.Integer> intAccumulator(int)
	IntAccumulator(a int) *Accumulator

	// public org.apache.spark.Accumulator<java.lang.Integer> intAccumulator(int, java.lang.String)
	IntAccumulator2(a int, b string) *Accumulator

	// public org.apache.spark.Accumulator<java.lang.Double> doubleAccumulator(double)
	DoubleAccumulator(a float64) *Accumulator

	// public org.apache.spark.Accumulator<java.lang.Double> doubleAccumulator(double, java.lang.String)
	DoubleAccumulator2(a float64, b string) *Accumulator

	// public org.apache.spark.Accumulator<java.lang.Integer> accumulator(int)
	Accumulator2(a int) *Accumulator

	// public org.apache.spark.Accumulator<java.lang.Integer> accumulator(int, java.lang.String)
	Accumulator4(a int, b string) *Accumulator

	// public org.apache.spark.Accumulator<java.lang.Double> accumulator(double)
	Accumulator(a float64) *Accumulator

	// public org.apache.spark.Accumulator<java.lang.Double> accumulator(double, java.lang.String)
	Accumulator3(a float64, b string) *Accumulator

	// public <T> org.apache.spark.Accumulator<T> accumulator(T, org.apache.spark.AccumulatorParam<T>)
	Accumulator5(a interface{}, b AccumulatorParamInterface) *Accumulator

	// public <T> org.apache.spark.Accumulator<T> accumulator(T, java.lang.String, org.apache.spark.AccumulatorParam<T>)
	Accumulator6(a interface{}, b string, c AccumulatorParamInterface) *Accumulator

	// public <T, R> org.apache.spark.Accumulable<T, R> accumulable(T, org.apache.spark.AccumulableParam<T, R>)
	Accumulable(a interface{}, b AccumulableParamInterface) *Accumulable

	// public <T, R> org.apache.spark.Accumulable<T, R> accumulable(T, java.lang.String, org.apache.spark.AccumulableParam<T, R>)
	Accumulable2(a interface{}, b string, c AccumulableParamInterface) *Accumulable

	// public <T> org.apache.spark.broadcast.Broadcast<T> broadcast(T)
	Broadcast(a interface{}) *BroadcastBroadcast

	// public void stop()
	Stop() 

	// public void close()
	Close() 

	// public void addFile(java.lang.String)
	AddFile(a string) 

	// public void addJar(java.lang.String)
	AddJar(a string) 

	// public void clearJars()
	ClearJars() 

	// public void clearFiles()
	ClearFiles() 

	// public void setCheckpointDir(java.lang.String)
	SetCheckpointDir(a string) 

	// public <T> org.apache.spark.api.java.JavaRDD<T> checkpointFile(java.lang.String)
	CheckpointFile(a string) *ApiJavaJavaRDD

	// public org.apache.spark.SparkConf getConf()
	GetConf() *SparkConf

	// public void setCallSite(java.lang.String)
	SetCallSite(a string) 

	// public void clearCallSite()
	ClearCallSite() 

	// public void setLocalProperty(java.lang.String, java.lang.String)
	SetLocalProperty(a string, b string) 

	// public java.lang.String getLocalProperty(java.lang.String)
	GetLocalProperty(a string) string

	// public void setLogLevel(java.lang.String)
	SetLogLevel(a string) 

	// public void setJobGroup(java.lang.String, java.lang.String, boolean)
	SetJobGroup2(a string, b string, c bool) 

	// public void setJobGroup(java.lang.String, java.lang.String)
	SetJobGroup(a string, b string) 

	// public void clearJobGroup()
	ClearJobGroup() 

	// public void cancelJobGroup(java.lang.String)
	CancelJobGroup(a string) 

	// public void cancelAllJobs()
	CancelAllJobs() 
}

type ApiJavaJavaSparkContext struct {
	ApiJavaJavaSparkContextVarargsWorkaround
}

// public org.apache.spark.api.java.JavaSparkContext(org.apache.spark.SparkContext)
func NewApiJavaJavaSparkContext3(a SparkContextInterface) (*ApiJavaJavaSparkContext) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/JavaSparkContext", conv_a.Value().Cast("org/apache/spark/SparkContext"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ApiJavaJavaSparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.api.java.JavaSparkContext()
func NewApiJavaJavaSparkContext() (*ApiJavaJavaSparkContext) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/JavaSparkContext")
	if err != nil {
		panic(err)
	}
	x := &ApiJavaJavaSparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.api.java.JavaSparkContext(org.apache.spark.SparkConf)
func NewApiJavaJavaSparkContext2(a SparkConfInterface) (*ApiJavaJavaSparkContext) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/JavaSparkContext", conv_a.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ApiJavaJavaSparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.api.java.JavaSparkContext(java.lang.String, java.lang.String)
func NewApiJavaJavaSparkContext4(a string, b string) (*ApiJavaJavaSparkContext) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/JavaSparkContext", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ApiJavaJavaSparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.api.java.JavaSparkContext(java.lang.String, java.lang.String, org.apache.spark.SparkConf)
func NewApiJavaJavaSparkContext5(a string, b string, c SparkConfInterface) (*ApiJavaJavaSparkContext) {
	conv_a := javabind.NewGoToJavaString()
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

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/JavaSparkContext", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &ApiJavaJavaSparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.api.java.JavaSparkContext(java.lang.String, java.lang.String, java.lang.String, java.lang.String)
func NewApiJavaJavaSparkContext6(a string, b string, c string, d string) (*ApiJavaJavaSparkContext) {
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

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/JavaSparkContext", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &ApiJavaJavaSparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.api.java.JavaSparkContext(java.lang.String, java.lang.String, java.lang.String, java.lang.String[])
func NewApiJavaJavaSparkContext7(a string, b string, c string, d []string) (*ApiJavaJavaSparkContext) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
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

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/JavaSparkContext", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &ApiJavaJavaSparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.api.java.JavaSparkContext(java.lang.String, java.lang.String, java.lang.String, java.lang.String[], java.util.Map<java.lang.String, java.lang.String>)
func NewApiJavaJavaSparkContext8(a string, b string, c string, d []string, e map[string]string) (*ApiJavaJavaSparkContext) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	conv_e := javabind.NewGoToJavaMap(javabind.NewGoToJavaString(), javabind.NewGoToJavaString())
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

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/api/java/JavaSparkContext", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("java/lang/String"), conv_e.Value().Cast("java/util/Map"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	x := &ApiJavaJavaSparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static java.lang.String[] jarOfObject(java.lang.Object)
func (jbobject *ApiJavaJavaSparkContext) JarOfObject(a interface{}) []string {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/api/java/JavaSparkContext", "jarOfObject", javabind.ObjectArrayType("java/lang/String"), conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static org.apache.spark.SparkContext toSparkContext(org.apache.spark.api.java.JavaSparkContext)
func (jbobject *ApiJavaJavaSparkContext) ToSparkContext(a ApiJavaJavaSparkContextInterface) *SparkContext {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/api/java/JavaSparkContext", "toSparkContext", "org/apache/spark/SparkContext", conv_a.Value().Cast("org/apache/spark/api/java/JavaSparkContext"))
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
	unique_x := &SparkContext{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.api.java.JavaSparkContext fromSparkContext(org.apache.spark.SparkContext)
func (jbobject *ApiJavaJavaSparkContext) FromSparkContext(a SparkContextInterface) *ApiJavaJavaSparkContext {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/api/java/JavaSparkContext", "fromSparkContext", "org/apache/spark/api/java/JavaSparkContext", conv_a.Value().Cast("org/apache/spark/SparkContext"))
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
	unique_x := &ApiJavaJavaSparkContext{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkContext sc()
func (jbobject *ApiJavaJavaSparkContext) Sc() *SparkContext {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sc", "org/apache/spark/SparkContext")
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
	unique_x := &SparkContext{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkEnv env()
func (jbobject *ApiJavaJavaSparkContext) Env() *SparkEnv {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "env", "org/apache/spark/SparkEnv")
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
	unique_x := &SparkEnv{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaSparkStatusTracker statusTracker()
func (jbobject *ApiJavaJavaSparkContext) StatusTracker() *ApiJavaJavaSparkStatusTracker {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "statusTracker", "org/apache/spark/api/java/JavaSparkStatusTracker")
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
	unique_x := &ApiJavaJavaSparkStatusTracker{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isLocal()
func (jbobject *ApiJavaJavaSparkContext) IsLocal() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isLocal", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String sparkUser()
func (jbobject *ApiJavaJavaSparkContext) SparkUser() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sparkUser", "java/lang/String")
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

// public java.lang.String master()
func (jbobject *ApiJavaJavaSparkContext) Master() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "master", "java/lang/String")
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

// public java.lang.String appName()
func (jbobject *ApiJavaJavaSparkContext) AppName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "appName", "java/lang/String")
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

// public java.util.List<java.lang.String> jars()
func (jbobject *ApiJavaJavaSparkContext) Jars() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "jars", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Long startTime()
func (jbobject *ApiJavaJavaSparkContext) StartTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "startTime", "java/lang/Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String version()
func (jbobject *ApiJavaJavaSparkContext) Version() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "version", "java/lang/String")
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

// public java.lang.Integer defaultParallelism()
func (jbobject *ApiJavaJavaSparkContext) DefaultParallelism() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "defaultParallelism", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Integer defaultMinSplits()
func (jbobject *ApiJavaJavaSparkContext) DefaultMinSplits() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "defaultMinSplits", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Integer defaultMinPartitions()
func (jbobject *ApiJavaJavaSparkContext) DefaultMinPartitions() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "defaultMinPartitions", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public <T> org.apache.spark.api.java.JavaRDD<T> parallelize(java.util.List<T>, int)
func (jbobject *ApiJavaJavaSparkContext) Parallelize2(a []*JavaLangObject, b int) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "parallelize", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/util/List"), b)
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public <T> org.apache.spark.api.java.JavaRDD<T> emptyRDD()
func (jbobject *ApiJavaJavaSparkContext) EmptyRDD() *ApiJavaJavaRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "emptyRDD", "org/apache/spark/api/java/JavaRDD")
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public <T> org.apache.spark.api.java.JavaRDD<T> parallelize(java.util.List<T>)
func (jbobject *ApiJavaJavaSparkContext) Parallelize(a []*JavaLangObject) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "parallelize", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/util/List"))
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD parallelizeDoubles(java.util.List<java.lang.Double>, int)
func (jbobject *ApiJavaJavaSparkContext) ParallelizeDoubles2(a []float64, b int) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaDouble())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "parallelizeDoubles", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("java/util/List"), b)
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD parallelizeDoubles(java.util.List<java.lang.Double>)
func (jbobject *ApiJavaJavaSparkContext) ParallelizeDoubles(a []float64) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaDouble())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "parallelizeDoubles", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("java/util/List"))
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaRDD<java.lang.String> textFile(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) TextFile(a string) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "textFile", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaRDD<java.lang.String> textFile(java.lang.String, int)
func (jbobject *ApiJavaJavaSparkContext) TextFile2(a string, b int) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "textFile", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/lang/String"), b)
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<java.lang.String, java.lang.String> wholeTextFiles(java.lang.String, int)
func (jbobject *ApiJavaJavaSparkContext) WholeTextFiles2(a string, b int) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "wholeTextFiles", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("java/lang/String"), b)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<java.lang.String, java.lang.String> wholeTextFiles(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) WholeTextFiles(a string) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "wholeTextFiles", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<java.lang.String, org.apache.spark.input.PortableDataStream> binaryFiles(java.lang.String, int)
func (jbobject *ApiJavaJavaSparkContext) BinaryFiles2(a string, b int) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "binaryFiles", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("java/lang/String"), b)
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaPairRDD<java.lang.String, org.apache.spark.input.PortableDataStream> binaryFiles(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) BinaryFiles(a string) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "binaryFiles", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaRDD<byte[]> binaryRecords(java.lang.String, int)
func (jbobject *ApiJavaJavaSparkContext) BinaryRecords(a string, b int) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "binaryRecords", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/lang/String"), b)
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public <T> org.apache.spark.api.java.JavaRDD<T> objectFile(java.lang.String, int)
func (jbobject *ApiJavaJavaSparkContext) ObjectFile2(a string, b int) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "objectFile", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/lang/String"), b)
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public <T> org.apache.spark.api.java.JavaRDD<T> objectFile(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) ObjectFile(a string) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "objectFile", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public <T> org.apache.spark.api.java.JavaRDD<T> union(org.apache.spark.api.java.JavaRDD<T>, java.util.List<org.apache.spark.api.java.JavaRDD<T>>)
func (jbobject *ApiJavaJavaSparkContext) Union6(a ApiJavaJavaRDDInterface, b []*ApiJavaJavaRDD) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "union", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaRDD"), conv_b.Value().Cast("java/util/List"))
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public <K, V> org.apache.spark.api.java.JavaPairRDD<K, V> union(org.apache.spark.api.java.JavaPairRDD<K, V>, java.util.List<org.apache.spark.api.java.JavaPairRDD<K, V>>)
func (jbobject *ApiJavaJavaSparkContext) Union5(a ApiJavaJavaPairRDDInterface, b []*ApiJavaJavaPairRDD) *ApiJavaJavaPairRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "union", "org/apache/spark/api/java/JavaPairRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaPairRDD"), conv_b.Value().Cast("java/util/List"))
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
	unique_x := &ApiJavaJavaPairRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.api.java.JavaDoubleRDD union(org.apache.spark.api.java.JavaDoubleRDD, java.util.List<org.apache.spark.api.java.JavaDoubleRDD>)
func (jbobject *ApiJavaJavaSparkContext) Union4(a ApiJavaJavaDoubleRDDInterface, b []*ApiJavaJavaDoubleRDD) *ApiJavaJavaDoubleRDD {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "union", "org/apache/spark/api/java/JavaDoubleRDD", conv_a.Value().Cast("org/apache/spark/api/java/JavaDoubleRDD"), conv_b.Value().Cast("java/util/List"))
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
	unique_x := &ApiJavaJavaDoubleRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.Accumulator<java.lang.Integer> intAccumulator(int)
func (jbobject *ApiJavaJavaSparkContext) IntAccumulator(a int) *Accumulator {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intAccumulator", "org/apache/spark/Accumulator", a)
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
	unique_x := &Accumulator{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.Accumulator<java.lang.Integer> intAccumulator(int, java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) IntAccumulator2(a int, b string) *Accumulator {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intAccumulator", "org/apache/spark/Accumulator", a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &Accumulator{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.Accumulator<java.lang.Double> doubleAccumulator(double)
func (jbobject *ApiJavaJavaSparkContext) DoubleAccumulator(a float64) *Accumulator {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "doubleAccumulator", "org/apache/spark/Accumulator", a)
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
	unique_x := &Accumulator{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.Accumulator<java.lang.Double> doubleAccumulator(double, java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) DoubleAccumulator2(a float64, b string) *Accumulator {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "doubleAccumulator", "org/apache/spark/Accumulator", a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &Accumulator{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.Accumulator<java.lang.Integer> accumulator(int)
func (jbobject *ApiJavaJavaSparkContext) Accumulator2(a int) *Accumulator {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "accumulator", "org/apache/spark/Accumulator", a)
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
	unique_x := &Accumulator{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.Accumulator<java.lang.Integer> accumulator(int, java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) Accumulator4(a int, b string) *Accumulator {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "accumulator", "org/apache/spark/Accumulator", a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &Accumulator{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.Accumulator<java.lang.Double> accumulator(double)
func (jbobject *ApiJavaJavaSparkContext) Accumulator(a float64) *Accumulator {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "accumulator", "org/apache/spark/Accumulator", a)
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
	unique_x := &Accumulator{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.Accumulator<java.lang.Double> accumulator(double, java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) Accumulator3(a float64, b string) *Accumulator {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "accumulator", "org/apache/spark/Accumulator", a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &Accumulator{}
	unique_x.Callable = dst
	return unique_x
}

// public <T> org.apache.spark.Accumulator<T> accumulator(T, org.apache.spark.AccumulatorParam<T>)
func (jbobject *ApiJavaJavaSparkContext) Accumulator5(a interface{}, b AccumulatorParamInterface) *Accumulator {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "accumulator", "org/apache/spark/Accumulator", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/AccumulatorParam"))
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
	unique_x := &Accumulator{}
	unique_x.Callable = dst
	return unique_x
}

// public <T> org.apache.spark.Accumulator<T> accumulator(T, java.lang.String, org.apache.spark.AccumulatorParam<T>)
func (jbobject *ApiJavaJavaSparkContext) Accumulator6(a interface{}, b string, c AccumulatorParamInterface) *Accumulator {
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
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "accumulator", "org/apache/spark/Accumulator", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("org/apache/spark/AccumulatorParam"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &Accumulator{}
	unique_x.Callable = dst
	return unique_x
}

// public <T, R> org.apache.spark.Accumulable<T, R> accumulable(T, org.apache.spark.AccumulableParam<T, R>)
func (jbobject *ApiJavaJavaSparkContext) Accumulable(a interface{}, b AccumulableParamInterface) *Accumulable {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "accumulable", "org/apache/spark/Accumulable", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("org/apache/spark/AccumulableParam"))
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
	unique_x := &Accumulable{}
	unique_x.Callable = dst
	return unique_x
}

// public <T, R> org.apache.spark.Accumulable<T, R> accumulable(T, java.lang.String, org.apache.spark.AccumulableParam<T, R>)
func (jbobject *ApiJavaJavaSparkContext) Accumulable2(a interface{}, b string, c AccumulableParamInterface) *Accumulable {
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
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "accumulable", "org/apache/spark/Accumulable", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("org/apache/spark/AccumulableParam"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &Accumulable{}
	unique_x.Callable = dst
	return unique_x
}

// public <T> org.apache.spark.broadcast.Broadcast<T> broadcast(T)
func (jbobject *ApiJavaJavaSparkContext) Broadcast(a interface{}) *BroadcastBroadcast {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "broadcast", "org/apache/spark/broadcast/Broadcast", conv_a.Value().Cast("java/lang/Object"))
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
	unique_x := &BroadcastBroadcast{}
	unique_x.Callable = dst
	return unique_x
}

// public void stop()
func (jbobject *ApiJavaJavaSparkContext) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void close()
func (jbobject *ApiJavaJavaSparkContext) Close()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void addFile(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) AddFile(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addFile", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addJar(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) AddJar(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addJar", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void clearJars()
func (jbobject *ApiJavaJavaSparkContext) ClearJars()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearJars", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void clearFiles()
func (jbobject *ApiJavaJavaSparkContext) ClearFiles()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearFiles", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setCheckpointDir(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) SetCheckpointDir(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCheckpointDir", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public <T> org.apache.spark.api.java.JavaRDD<T> checkpointFile(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) CheckpointFile(a string) *ApiJavaJavaRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "checkpointFile", "org/apache/spark/api/java/JavaRDD", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ApiJavaJavaRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkConf getConf()
func (jbobject *ApiJavaJavaSparkContext) GetConf() *SparkConf {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getConf", "org/apache/spark/SparkConf")
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

// public void setCallSite(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) SetCallSite(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCallSite", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void clearCallSite()
func (jbobject *ApiJavaJavaSparkContext) ClearCallSite()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearCallSite", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setLocalProperty(java.lang.String, java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) SetLocalProperty(a string, b string)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLocalProperty", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public java.lang.String getLocalProperty(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) GetLocalProperty(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLocalProperty", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
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

// public void setLogLevel(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) SetLogLevel(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLogLevel", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setJobGroup(java.lang.String, java.lang.String, boolean)
func (jbobject *ApiJavaJavaSparkContext) SetJobGroup2(a string, b string, c bool)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setJobGroup", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void setJobGroup(java.lang.String, java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) SetJobGroup(a string, b string)  {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setJobGroup", javabind.Void, conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void clearJobGroup()
func (jbobject *ApiJavaJavaSparkContext) ClearJobGroup()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearJobGroup", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void cancelJobGroup(java.lang.String)
func (jbobject *ApiJavaJavaSparkContext) CancelJobGroup(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelJobGroup", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void cancelAllJobs()
func (jbobject *ApiJavaJavaSparkContext) CancelAllJobs()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelAllJobs", javabind.Void)
	if err != nil {
		panic(err)
	}

}


