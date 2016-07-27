package gospark

import "github.com/timob/javabind"

type UtilMetadataCleanerInterface interface {
	JavaLangObjectInterface

	// public static void setDelaySeconds(org.apache.spark.SparkConf, int, boolean)
	SetDelaySeconds(a SparkConfInterface, b int, c bool) 

	// public static int getDelaySeconds(org.apache.spark.SparkConf)
	GetDelaySeconds(a SparkConfInterface) int

	// public java.lang.String logName()
	LogName() string

	// public boolean isTraceEnabled()
	IsTraceEnabled() bool

	// public java.lang.String name()
	Name() string

	// public void cancel()
	Cancel() 
}

type UtilMetadataCleaner struct {
	JavaLangObject
}

// public static void setDelaySeconds(org.apache.spark.SparkConf, int, boolean)
func (jbobject *UtilMetadataCleaner) SetDelaySeconds(a SparkConfInterface, b int, c bool)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/util/MetadataCleaner", "setDelaySeconds", javabind.Void, conv_a.Value().Cast("org/apache/spark/SparkConf"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public static int getDelaySeconds(org.apache.spark.SparkConf)
func (jbobject *UtilMetadataCleaner) GetDelaySeconds(a SparkConfInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/util/MetadataCleaner", "getDelaySeconds", javabind.Int, conv_a.Value().Cast("org/apache/spark/SparkConf"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public java.lang.String logName()
func (jbobject *UtilMetadataCleaner) LogName() string {
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
func (jbobject *UtilMetadataCleaner) IsTraceEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTraceEnabled", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String name()
func (jbobject *UtilMetadataCleaner) Name() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "name", "java/lang/String")
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

// public void cancel()
func (jbobject *UtilMetadataCleaner) Cancel()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "cancel", javabind.Void)
	if err != nil {
		panic(err)
	}

}


