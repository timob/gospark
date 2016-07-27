package gospark

import "github.com/timob/javabind"

type JobExecutionStatusInterface interface {
	JavaLangEnumInterface

	// public static org.apache.spark.JobExecutionStatus[] values()
	Values() []*JobExecutionStatus

	// public static org.apache.spark.JobExecutionStatus valueOf(java.lang.String)
	ValueOf(a string) *JobExecutionStatus

	// public static org.apache.spark.JobExecutionStatus fromString(java.lang.String)
	FromString(a string) *JobExecutionStatus
}

type JobExecutionStatus struct {
	JavaLangEnum
}

// public static org.apache.spark.JobExecutionStatus[] values()
func (jbobject *JobExecutionStatus) Values() []*JobExecutionStatus {
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/JobExecutionStatus", "values", javabind.ObjectArrayType("org/apache/spark/JobExecutionStatus"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/JobExecutionStatus")
	dst := new([]*JobExecutionStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static org.apache.spark.JobExecutionStatus valueOf(java.lang.String)
func (jbobject *JobExecutionStatus) ValueOf(a string) *JobExecutionStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/JobExecutionStatus", "valueOf", "org/apache/spark/JobExecutionStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &JobExecutionStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public static org.apache.spark.JobExecutionStatus fromString(java.lang.String)
func (jbobject *JobExecutionStatus) FromString(a string) *JobExecutionStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("org/apache/spark/JobExecutionStatus", "fromString", "org/apache/spark/JobExecutionStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &JobExecutionStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JobExecutionStatus) RUNNING() *JobExecutionStatus {
	jret, err := javabind.GetEnv().GetStaticField("org/apache/spark/JobExecutionStatus", "RUNNING", "org/apache/spark/JobExecutionStatus")
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
	unique_x := &JobExecutionStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JobExecutionStatus) SetFieldRUNNING(val JobExecutionStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/apache/spark/JobExecutionStatus", "RUNNING", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JobExecutionStatus) SUCCEEDED() *JobExecutionStatus {
	jret, err := javabind.GetEnv().GetStaticField("org/apache/spark/JobExecutionStatus", "SUCCEEDED", "org/apache/spark/JobExecutionStatus")
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
	unique_x := &JobExecutionStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JobExecutionStatus) SetFieldSUCCEEDED(val JobExecutionStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/apache/spark/JobExecutionStatus", "SUCCEEDED", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JobExecutionStatus) FAILED() *JobExecutionStatus {
	jret, err := javabind.GetEnv().GetStaticField("org/apache/spark/JobExecutionStatus", "FAILED", "org/apache/spark/JobExecutionStatus")
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
	unique_x := &JobExecutionStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JobExecutionStatus) SetFieldFAILED(val JobExecutionStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/apache/spark/JobExecutionStatus", "FAILED", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JobExecutionStatus) UNKNOWN() *JobExecutionStatus {
	jret, err := javabind.GetEnv().GetStaticField("org/apache/spark/JobExecutionStatus", "UNKNOWN", "org/apache/spark/JobExecutionStatus")
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
	unique_x := &JobExecutionStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JobExecutionStatus) SetFieldUNKNOWN(val JobExecutionStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("org/apache/spark/JobExecutionStatus", "UNKNOWN", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


