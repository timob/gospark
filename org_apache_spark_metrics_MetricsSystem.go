package gospark

import "github.com/timob/javabind"

type MetricsMetricsSystemInterface interface {
	JavaLangObjectInterface

	// public static org.apache.spark.metrics.MetricsSystem createMetricsSystem(java.lang.String, org.apache.spark.SparkConf, org.apache.spark.SecurityManager)
	CreateMetricsSystem(a string, b SparkConfInterface, c SecurityManagerInterface) *MetricsMetricsSystem

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public java.lang.String instance()
	Instance() string

	// public void start()
	Start() 

	// public void stop()
	Stop() 

	// public void report()
	Report() 

	// public java.lang.String buildRegistryName(org.apache.spark.metrics.source.Source)
	BuildRegistryName(a MetricsSourceSourceInterface) string

	// public void registerSource(org.apache.spark.metrics.source.Source)
	RegisterSource(a MetricsSourceSourceInterface) 

	// public void removeSource(org.apache.spark.metrics.source.Source)
	RemoveSource(a MetricsSourceSourceInterface) 
}

type MetricsMetricsSystem struct {
	JavaLangObject
}

// public org.apache.spark.metrics.MetricsSystem(java.lang.String, org.apache.spark.SparkConf, org.apache.spark.SecurityManager)
func NewMetricsMetricsSystem(a string, b SparkConfInterface, c SecurityManagerInterface) (*MetricsMetricsSystem) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
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

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/metrics/MetricsSystem", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("org/apache/spark/SparkConf"), conv_c.Value().Cast("org/apache/spark/SecurityManager"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &MetricsMetricsSystem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static org.apache.spark.metrics.MetricsSystem createMetricsSystem(java.lang.String, org.apache.spark.SparkConf, org.apache.spark.SecurityManager)
func (jbobject *MetricsMetricsSystem) CreateMetricsSystem(a string, b SparkConfInterface, c SecurityManagerInterface) *MetricsMetricsSystem {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
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
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/metrics/MetricsSystem", "createMetricsSystem", "org/apache/spark/metrics/MetricsSystem", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("org/apache/spark/SparkConf"), conv_c.Value().Cast("org/apache/spark/SecurityManager"))
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
	unique_x := &MetricsMetricsSystem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String logName()
func (jbobject *MetricsMetricsSystem) LogName() string {
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
func (jbobject *MetricsMetricsSystem) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String instance()
func (jbobject *MetricsMetricsSystem) Instance() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "instance", "java/lang/String")
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

// public void start()
func (jbobject *MetricsMetricsSystem) Start()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "start", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void stop()
func (jbobject *MetricsMetricsSystem) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void report()
func (jbobject *MetricsMetricsSystem) Report()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "report", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public java.lang.String buildRegistryName(org.apache.spark.metrics.source.Source)
func (jbobject *MetricsMetricsSystem) BuildRegistryName(a MetricsSourceSourceInterface) string {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "buildRegistryName", "java/lang/String", conv_a.Value().Cast("org/apache/spark/metrics/source/Source"))
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

// public void registerSource(org.apache.spark.metrics.source.Source)
func (jbobject *MetricsMetricsSystem) RegisterSource(a MetricsSourceSourceInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "registerSource", javabind.Void, conv_a.Value().Cast("org/apache/spark/metrics/source/Source"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void removeSource(org.apache.spark.metrics.source.Source)
func (jbobject *MetricsMetricsSystem) RemoveSource(a MetricsSourceSourceInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "removeSource", javabind.Void, conv_a.Value().Cast("org/apache/spark/metrics/source/Source"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


