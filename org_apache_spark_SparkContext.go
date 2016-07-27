package gospark

import "github.com/timob/javabind"

type SparkContextInterface interface {
	JavaLangObjectInterface

	// public static <T extends org.apache.hadoop.io.Writable> org.apache.spark.WritableConverter<T> writableWritableConverter()
	WritableWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.String> stringWritableConverter()
	StringWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<byte[]> bytesWritableConverter()
	BytesWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.Object> booleanWritableConverter()
	BooleanWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.Object> floatWritableConverter()
	FloatWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.Object> doubleWritableConverter()
	DoubleWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.Object> longWritableConverter()
	LongWritableConverter() *WritableConverter

	// public static org.apache.spark.WritableConverter<java.lang.Object> intWritableConverter()
	IntWritableConverter() *WritableConverter

	// public static org.apache.spark.rdd.DoubleRDDFunctions doubleRDDToDoubleRDDFunctions(org.apache.spark.rdd.RDD<java.lang.Object>)
	DoubleRDDToDoubleRDDFunctions(a RddRDDInterface) *RddDoubleRDDFunctions

	// public static org.apache.spark.SparkContext getOrCreate()
	GetOrCreate() *SparkContext

	// public static org.apache.spark.SparkContext getOrCreate(org.apache.spark.SparkConf)
	GetOrCreate2(a SparkConfInterface) *SparkContext

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public long startTime()
	StartTime() int64

	// public org.apache.spark.SparkConf conf()
	Conf() *SparkConf

	// public org.apache.spark.SparkConf getConf()
	GetConf() *SparkConf

	// public java.lang.String master()
	Master() string

	// public java.lang.String appName()
	AppName() string

	// public boolean isEventLogEnabled()
	IsEventLogEnabled() bool

	// public java.lang.String externalBlockStoreFolderName()
	ExternalBlockStoreFolderName() string

	// public java.lang.String tachyonFolderName()
	TachyonFolderName() string

	// public boolean isLocal()
	IsLocal() bool

	// public boolean isStopped()
	IsStopped() bool

	// public org.apache.spark.scheduler.LiveListenerBus listenerBus()
	ListenerBus() *SchedulerLiveListenerBus

	// public org.apache.spark.SparkEnv createSparkEnv(org.apache.spark.SparkConf, boolean, org.apache.spark.scheduler.LiveListenerBus)
	CreateSparkEnv(a SparkConfInterface, b bool, c SchedulerLiveListenerBusInterface) *SparkEnv

	// public org.apache.spark.SparkEnv env()
	Env() *SparkEnv

	// public org.apache.spark.util.TimeStampedWeakValueHashMap<java.lang.Object, org.apache.spark.rdd.RDD<?>> persistentRdds()
	PersistentRdds() *UtilTimeStampedWeakValueHashMap

	// public org.apache.spark.util.MetadataCleaner metadataCleaner()
	MetadataCleaner() *UtilMetadataCleaner

	// public org.apache.spark.ui.jobs.JobProgressListener jobProgressListener()
	JobProgressListener() *UiJobsJobProgressListener

	// public org.apache.spark.SparkStatusTracker statusTracker()
	StatusTracker() *SparkStatusTracker

	// public int executorMemory()
	ExecutorMemory() int

	// public java.lang.String sparkUser()
	SparkUser() string

	// public org.apache.spark.scheduler.SchedulerBackend schedulerBackend()
	SchedulerBackend() *SchedulerSchedulerBackend

	// public org.apache.spark.scheduler.TaskScheduler taskScheduler()
	TaskScheduler() *SchedulerTaskScheduler

	// public org.apache.spark.scheduler.DAGScheduler dagScheduler()
	DagScheduler() *SchedulerDAGScheduler

	// public java.lang.String applicationId()
	ApplicationId() string

	// public org.apache.spark.metrics.MetricsSystem metricsSystem()
	MetricsSystem() *MetricsMetricsSystem

	// public void setLogLevel(java.lang.String)
	SetLogLevel(a string) 

	// public java.util.Properties getLocalProperties()
	GetLocalProperties() *JavaUtilProperties

	// public void setLocalProperties(java.util.Properties)
	SetLocalProperties(a JavaUtilPropertiesInterface) 

	// public void initLocalProperties()
	InitLocalProperties() 

	// public void setLocalProperty(java.lang.String, java.lang.String)
	SetLocalProperty(a string, b string) 

	// public java.lang.String getLocalProperty(java.lang.String)
	GetLocalProperty(a string) string

	// public void setJobDescription(java.lang.String)
	SetJobDescription(a string) 

	// public void setJobGroup(java.lang.String, java.lang.String, boolean)
	SetJobGroup(a string, b string, c bool) 

	// public void clearJobGroup()
	ClearJobGroup() 

	// public org.apache.spark.rdd.RDD<java.lang.Object> range(long, long, long, int)
	Range(a int64, b int64, c int64, d int) *RddRDD

	// public org.apache.spark.rdd.RDD<java.lang.String> textFile(java.lang.String, int)
	TextFile(a string, b int) *RddRDD

	// public <T> org.apache.spark.Accumulator<T> accumulator(T, org.apache.spark.AccumulatorParam<T>)
	Accumulator(a interface{}, b AccumulatorParamInterface) *Accumulator

	// public <T> org.apache.spark.Accumulator<T> accumulator(T, java.lang.String, org.apache.spark.AccumulatorParam<T>)
	Accumulator2(a interface{}, b string, c AccumulatorParamInterface) *Accumulator

	// public <R, T> org.apache.spark.Accumulable<R, T> accumulable(R, org.apache.spark.AccumulableParam<R, T>)
	Accumulable(a interface{}, b AccumulableParamInterface) *Accumulable

	// public <R, T> org.apache.spark.Accumulable<R, T> accumulable(R, java.lang.String, org.apache.spark.AccumulableParam<R, T>)
	Accumulable2(a interface{}, b string, c AccumulableParamInterface) *Accumulable

	// public void addFile(java.lang.String)
	AddFile(a string) 

	// public void addFile(java.lang.String, boolean)
	AddFile2(a string, b bool) 

	// public void addSparkListener(org.apache.spark.scheduler.SparkListener)
	AddSparkListener(a SchedulerSparkListenerInterface) 

	// public boolean requestExecutors(int)
	RequestExecutors(a int) bool

	// public boolean killExecutor(java.lang.String)
	KillExecutor(a string) bool

	// public boolean killAndReplaceExecutor(java.lang.String)
	KillAndReplaceExecutor(a string) bool

	// public java.lang.String version()
	Version() string

	// public org.apache.spark.storage.RDDInfo[] getRDDStorageInfo()
	GetRDDStorageInfo() []*StorageRDDInfo

	// public org.apache.spark.storage.StorageStatus[] getExecutorStorageStatus()
	GetExecutorStorageStatus() []*StorageStorageStatus

	// public void clearFiles()
	ClearFiles() 

	// public void persistRDD(org.apache.spark.rdd.RDD<?>)
	PersistRDD(a RddRDDInterface) 

	// public void unpersistRDD(int, boolean)
	UnpersistRDD(a int, b bool) 

	// public void addJar(java.lang.String)
	AddJar(a string) 

	// public void clearJars()
	ClearJars() 

	// public void stop()
	Stop() 

	// public void setCallSite(java.lang.String)
	SetCallSite(a string) 

	// public void setCallSite(org.apache.spark.util.CallSite)
	SetCallSite2(a UtilCallSiteInterface) 

	// public void clearCallSite()
	ClearCallSite() 

	// public org.apache.spark.util.CallSite getCallSite()
	GetCallSite() *UtilCallSite

	// public <K, V, C> org.apache.spark.SimpleFutureAction<org.apache.spark.MapOutputStatistics> submitMapStage(org.apache.spark.ShuffleDependency<K, V, C>)
	SubmitMapStage(a ShuffleDependencyInterface) *SimpleFutureAction

	// public void cancelJobGroup(java.lang.String)
	CancelJobGroup(a string) 

	// public void cancelAllJobs()
	CancelAllJobs() 

	// public void cancelJob(int)
	CancelJob(a int) 

	// public void cancelStage(int)
	CancelStage(a int) 

	// public <F> F clean(F, boolean)
	Clean(a interface{}, b bool) *JavaLangObject

	// public void setCheckpointDir(java.lang.String)
	SetCheckpointDir(a string) 

	// public int defaultParallelism()
	DefaultParallelism() int

	// public int defaultMinSplits()
	DefaultMinSplits() int

	// public int defaultMinPartitions()
	DefaultMinPartitions() int

	// public int newShuffleId()
	NewShuffleId() int

	// public int newRddId()
	NewRddId() int

	// public void cleanup(long)
	Cleanup(a int64) 
}

type SparkContext struct {
	JavaLangObject
}

// public org.apache.spark.SparkContext(org.apache.spark.SparkConf)
func NewSparkContext2(a SparkConfInterface) (*SparkContext) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/SparkContext", conv_a.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &SparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.SparkContext()
func NewSparkContext() (*SparkContext) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/SparkContext")
	if err != nil {
		panic(err)
	}
	x := &SparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.SparkContext(java.lang.String, java.lang.String, org.apache.spark.SparkConf)
func NewSparkContext5(a string, b string, c SparkConfInterface) (*SparkContext) {
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

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/SparkContext", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &SparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.SparkContext(java.lang.String, java.lang.String)
func NewSparkContext3(a string, b string) (*SparkContext) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/SparkContext", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &SparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.SparkContext(java.lang.String, java.lang.String, java.lang.String)
func NewSparkContext4(a string, b string, c string) (*SparkContext) {
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

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/SparkContext", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &SparkContext{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static <T extends org.apache.hadoop.io.Writable> org.apache.spark.WritableConverter<T> writableWritableConverter()
func (jbobject *SparkContext) WritableWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkContext", "writableWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.String> stringWritableConverter()
func (jbobject *SparkContext) StringWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkContext", "stringWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<byte[]> bytesWritableConverter()
func (jbobject *SparkContext) BytesWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkContext", "bytesWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.Object> booleanWritableConverter()
func (jbobject *SparkContext) BooleanWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkContext", "booleanWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.Object> floatWritableConverter()
func (jbobject *SparkContext) FloatWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkContext", "floatWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.Object> doubleWritableConverter()
func (jbobject *SparkContext) DoubleWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkContext", "doubleWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.Object> longWritableConverter()
func (jbobject *SparkContext) LongWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkContext", "longWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.WritableConverter<java.lang.Object> intWritableConverter()
func (jbobject *SparkContext) IntWritableConverter() *WritableConverter {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkContext", "intWritableConverter", "org/apache/spark/WritableConverter")
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
	unique_x := &WritableConverter{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.rdd.DoubleRDDFunctions doubleRDDToDoubleRDDFunctions(org.apache.spark.rdd.RDD<java.lang.Object>)
func (jbobject *SparkContext) DoubleRDDToDoubleRDDFunctions(a RddRDDInterface) *RddDoubleRDDFunctions {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkContext", "doubleRDDToDoubleRDDFunctions", "org/apache/spark/rdd/DoubleRDDFunctions", conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
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
	unique_x := &RddDoubleRDDFunctions{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.SparkContext getOrCreate()
func (jbobject *SparkContext) GetOrCreate() *SparkContext {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkContext", "getOrCreate", "org/apache/spark/SparkContext")
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

// public static org.apache.spark.SparkContext getOrCreate(org.apache.spark.SparkConf)
func (jbobject *SparkContext) GetOrCreate2(a SparkConfInterface) *SparkContext {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/SparkContext", "getOrCreate", "org/apache/spark/SparkContext", conv_a.Value().Cast("org/apache/spark/SparkConf"))
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

// public java.lang.String logName()
func (jbobject *SparkContext) LogName() string {
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
func (jbobject *SparkContext) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public long startTime()
func (jbobject *SparkContext) StartTime() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "startTime", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public org.apache.spark.SparkConf conf()
func (jbobject *SparkContext) Conf() *SparkConf {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "conf", "org/apache/spark/SparkConf")
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

// public org.apache.spark.SparkConf getConf()
func (jbobject *SparkContext) GetConf() *SparkConf {
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

// public java.lang.String master()
func (jbobject *SparkContext) Master() string {
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
func (jbobject *SparkContext) AppName() string {
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

// public boolean isEventLogEnabled()
func (jbobject *SparkContext) IsEventLogEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEventLogEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String externalBlockStoreFolderName()
func (jbobject *SparkContext) ExternalBlockStoreFolderName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "externalBlockStoreFolderName", "java/lang/String")
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

// public java.lang.String tachyonFolderName()
func (jbobject *SparkContext) TachyonFolderName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "tachyonFolderName", "java/lang/String")
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

// public boolean isLocal()
func (jbobject *SparkContext) IsLocal() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isLocal", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isStopped()
func (jbobject *SparkContext) IsStopped() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isStopped", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public org.apache.spark.scheduler.LiveListenerBus listenerBus()
func (jbobject *SparkContext) ListenerBus() *SchedulerLiveListenerBus {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "listenerBus", "org/apache/spark/scheduler/LiveListenerBus")
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
	unique_x := &SchedulerLiveListenerBus{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkEnv createSparkEnv(org.apache.spark.SparkConf, boolean, org.apache.spark.scheduler.LiveListenerBus)
func (jbobject *SparkContext) CreateSparkEnv(a SparkConfInterface, b bool, c SchedulerLiveListenerBusInterface) *SparkEnv {
	conv_a := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createSparkEnv", "org/apache/spark/SparkEnv", conv_a.Value().Cast("org/apache/spark/SparkConf"), b, conv_c.Value().Cast("org/apache/spark/scheduler/LiveListenerBus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_c.CleanUp()
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

// public org.apache.spark.SparkEnv env()
func (jbobject *SparkContext) Env() *SparkEnv {
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

// public org.apache.spark.util.TimeStampedWeakValueHashMap<java.lang.Object, org.apache.spark.rdd.RDD<?>> persistentRdds()
func (jbobject *SparkContext) PersistentRdds() *UtilTimeStampedWeakValueHashMap {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "persistentRdds", "org/apache/spark/util/TimeStampedWeakValueHashMap")
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
	unique_x := &UtilTimeStampedWeakValueHashMap{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.util.MetadataCleaner metadataCleaner()
func (jbobject *SparkContext) MetadataCleaner() *UtilMetadataCleaner {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "metadataCleaner", "org/apache/spark/util/MetadataCleaner")
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
	unique_x := &UtilMetadataCleaner{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.ui.jobs.JobProgressListener jobProgressListener()
func (jbobject *SparkContext) JobProgressListener() *UiJobsJobProgressListener {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "jobProgressListener", "org/apache/spark/ui/jobs/JobProgressListener")
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
	unique_x := &UiJobsJobProgressListener{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.SparkStatusTracker statusTracker()
func (jbobject *SparkContext) StatusTracker() *SparkStatusTracker {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "statusTracker", "org/apache/spark/SparkStatusTracker")
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
	unique_x := &SparkStatusTracker{}
	unique_x.Callable = dst
	return unique_x
}

// public int executorMemory()
func (jbobject *SparkContext) ExecutorMemory() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "executorMemory", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String sparkUser()
func (jbobject *SparkContext) SparkUser() string {
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

// public org.apache.spark.scheduler.SchedulerBackend schedulerBackend()
func (jbobject *SparkContext) SchedulerBackend() *SchedulerSchedulerBackend {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "schedulerBackend", "org/apache/spark/scheduler/SchedulerBackend")
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
	unique_x := &SchedulerSchedulerBackend{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.scheduler.TaskScheduler taskScheduler()
func (jbobject *SparkContext) TaskScheduler() *SchedulerTaskScheduler {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "taskScheduler", "org/apache/spark/scheduler/TaskScheduler")
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
	unique_x := &SchedulerTaskScheduler{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.scheduler.DAGScheduler dagScheduler()
func (jbobject *SparkContext) DagScheduler() *SchedulerDAGScheduler {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "dagScheduler", "org/apache/spark/scheduler/DAGScheduler")
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
	unique_x := &SchedulerDAGScheduler{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String applicationId()
func (jbobject *SparkContext) ApplicationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "applicationId", "java/lang/String")
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

// public org.apache.spark.metrics.MetricsSystem metricsSystem()
func (jbobject *SparkContext) MetricsSystem() *MetricsMetricsSystem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "metricsSystem", "org/apache/spark/metrics/MetricsSystem")
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
	unique_x := &MetricsMetricsSystem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLogLevel(java.lang.String)
func (jbobject *SparkContext) SetLogLevel(a string)  {
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

// public java.util.Properties getLocalProperties()
func (jbobject *SparkContext) GetLocalProperties() *JavaUtilProperties {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLocalProperties", "java/util/Properties")
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
	unique_x := &JavaUtilProperties{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLocalProperties(java.util.Properties)
func (jbobject *SparkContext) SetLocalProperties(a JavaUtilPropertiesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLocalProperties", javabind.Void, conv_a.Value().Cast("java/util/Properties"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void initLocalProperties()
func (jbobject *SparkContext) InitLocalProperties()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "initLocalProperties", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setLocalProperty(java.lang.String, java.lang.String)
func (jbobject *SparkContext) SetLocalProperty(a string, b string)  {
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
func (jbobject *SparkContext) GetLocalProperty(a string) string {
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

// public void setJobDescription(java.lang.String)
func (jbobject *SparkContext) SetJobDescription(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setJobDescription", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setJobGroup(java.lang.String, java.lang.String, boolean)
func (jbobject *SparkContext) SetJobGroup(a string, b string, c bool)  {
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

// public void clearJobGroup()
func (jbobject *SparkContext) ClearJobGroup()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearJobGroup", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.apache.spark.rdd.RDD<java.lang.Object> range(long, long, long, int)
func (jbobject *SparkContext) Range(a int64, b int64, c int64, d int) *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "range", "org/apache/spark/rdd/RDD", a, b, c, d)
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.rdd.RDD<java.lang.String> textFile(java.lang.String, int)
func (jbobject *SparkContext) TextFile(a string, b int) *RddRDD {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "textFile", "org/apache/spark/rdd/RDD", conv_a.Value().Cast("java/lang/String"), b)
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}

// public <T> org.apache.spark.Accumulator<T> accumulator(T, org.apache.spark.AccumulatorParam<T>)
func (jbobject *SparkContext) Accumulator(a interface{}, b AccumulatorParamInterface) *Accumulator {
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
func (jbobject *SparkContext) Accumulator2(a interface{}, b string, c AccumulatorParamInterface) *Accumulator {
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

// public <R, T> org.apache.spark.Accumulable<R, T> accumulable(R, org.apache.spark.AccumulableParam<R, T>)
func (jbobject *SparkContext) Accumulable(a interface{}, b AccumulableParamInterface) *Accumulable {
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

// public <R, T> org.apache.spark.Accumulable<R, T> accumulable(R, java.lang.String, org.apache.spark.AccumulableParam<R, T>)
func (jbobject *SparkContext) Accumulable2(a interface{}, b string, c AccumulableParamInterface) *Accumulable {
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

// public void addFile(java.lang.String)
func (jbobject *SparkContext) AddFile(a string)  {
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

// public void addFile(java.lang.String, boolean)
func (jbobject *SparkContext) AddFile2(a string, b bool)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addFile", javabind.Void, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addSparkListener(org.apache.spark.scheduler.SparkListener)
func (jbobject *SparkContext) AddSparkListener(a SchedulerSparkListenerInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addSparkListener", javabind.Void, conv_a.Value().Cast("org/apache/spark/scheduler/SparkListener"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean requestExecutors(int)
func (jbobject *SparkContext) RequestExecutors(a int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "requestExecutors", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean killExecutor(java.lang.String)
func (jbobject *SparkContext) KillExecutor(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "killExecutor", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean killAndReplaceExecutor(java.lang.String)
func (jbobject *SparkContext) KillAndReplaceExecutor(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "killAndReplaceExecutor", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public java.lang.String version()
func (jbobject *SparkContext) Version() string {
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

// public org.apache.spark.storage.RDDInfo[] getRDDStorageInfo()
func (jbobject *SparkContext) GetRDDStorageInfo() []*StorageRDDInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRDDStorageInfo", javabind.ObjectArrayType("org/apache/spark/storage/RDDInfo"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/storage/RDDInfo")
	dst := new([]*StorageRDDInfo)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public org.apache.spark.storage.StorageStatus[] getExecutorStorageStatus()
func (jbobject *SparkContext) GetExecutorStorageStatus() []*StorageStorageStatus {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExecutorStorageStatus", javabind.ObjectArrayType("org/apache/spark/storage/StorageStatus"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/storage/StorageStatus")
	dst := new([]*StorageStorageStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void clearFiles()
func (jbobject *SparkContext) ClearFiles()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearFiles", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void persistRDD(org.apache.spark.rdd.RDD<?>)
func (jbobject *SparkContext) PersistRDD(a RddRDDInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "persistRDD", javabind.Void, conv_a.Value().Cast("org/apache/spark/rdd/RDD"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void unpersistRDD(int, boolean)
func (jbobject *SparkContext) UnpersistRDD(a int, b bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "unpersistRDD", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

// public void addJar(java.lang.String)
func (jbobject *SparkContext) AddJar(a string)  {
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
func (jbobject *SparkContext) ClearJars()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearJars", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void stop()
func (jbobject *SparkContext) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void setCallSite(java.lang.String)
func (jbobject *SparkContext) SetCallSite(a string)  {
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

// public void setCallSite(org.apache.spark.util.CallSite)
func (jbobject *SparkContext) SetCallSite2(a UtilCallSiteInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCallSite", javabind.Void, conv_a.Value().Cast("org/apache/spark/util/CallSite"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void clearCallSite()
func (jbobject *SparkContext) ClearCallSite()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clearCallSite", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public org.apache.spark.util.CallSite getCallSite()
func (jbobject *SparkContext) GetCallSite() *UtilCallSite {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCallSite", "org/apache/spark/util/CallSite")
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
	unique_x := &UtilCallSite{}
	unique_x.Callable = dst
	return unique_x
}

// public <K, V, C> org.apache.spark.SimpleFutureAction<org.apache.spark.MapOutputStatistics> submitMapStage(org.apache.spark.ShuffleDependency<K, V, C>)
func (jbobject *SparkContext) SubmitMapStage(a ShuffleDependencyInterface) *SimpleFutureAction {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "submitMapStage", "org/apache/spark/SimpleFutureAction", conv_a.Value().Cast("org/apache/spark/ShuffleDependency"))
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
	unique_x := &SimpleFutureAction{}
	unique_x.Callable = dst
	return unique_x
}

// public void cancelJobGroup(java.lang.String)
func (jbobject *SparkContext) CancelJobGroup(a string)  {
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
func (jbobject *SparkContext) CancelAllJobs()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelAllJobs", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void cancelJob(int)
func (jbobject *SparkContext) CancelJob(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelJob", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void cancelStage(int)
func (jbobject *SparkContext) CancelStage(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancelStage", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public <F> F clean(F, boolean)
func (jbobject *SparkContext) Clean(a interface{}, b bool) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clean", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"), b)
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

// public void setCheckpointDir(java.lang.String)
func (jbobject *SparkContext) SetCheckpointDir(a string)  {
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

// public int defaultParallelism()
func (jbobject *SparkContext) DefaultParallelism() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "defaultParallelism", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int defaultMinSplits()
func (jbobject *SparkContext) DefaultMinSplits() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "defaultMinSplits", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int defaultMinPartitions()
func (jbobject *SparkContext) DefaultMinPartitions() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "defaultMinPartitions", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int newShuffleId()
func (jbobject *SparkContext) NewShuffleId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "newShuffleId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int newRddId()
func (jbobject *SparkContext) NewRddId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "newRddId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void cleanup(long)
func (jbobject *SparkContext) Cleanup(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cleanup", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


