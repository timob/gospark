package gospark

import "github.com/timob/javabind"

type SchedulerExecutorLossReasonInterface interface {
	JavaLangObjectInterface

	// public java.lang.String message()
	Message() string
}

type SchedulerExecutorLossReason struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.ExecutorLossReason(java.lang.String)
func NewSchedulerExecutorLossReason(a string) (*SchedulerExecutorLossReason) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/ExecutorLossReason", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &SchedulerExecutorLossReason{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String message()
func (jbobject *SchedulerExecutorLossReason) Message() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "message", "java/lang/String")
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

// public java.lang.String toString()
func (jbobject *SchedulerExecutorLossReason) ToString() string {
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


