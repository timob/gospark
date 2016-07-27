package gospark

import "github.com/timob/javabind"

type SchedulerClusterExecutorInfoInterface interface {
	JavaLangObjectInterface

	// public java.lang.String executorHost()
	ExecutorHost() string

	// public int totalCores()
	TotalCores() int

	// public boolean canEqual(java.lang.Object)
	CanEqual(a interface{}) bool
}

type SchedulerClusterExecutorInfo struct {
	JavaLangObject
}

// public java.lang.String executorHost()
func (jbobject *SchedulerClusterExecutorInfo) ExecutorHost() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "executorHost", "java/lang/String")
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

// public int totalCores()
func (jbobject *SchedulerClusterExecutorInfo) TotalCores() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "totalCores", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean canEqual(java.lang.Object)
func (jbobject *SchedulerClusterExecutorInfo) CanEqual(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "canEqual", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean equals(java.lang.Object)
func (jbobject *SchedulerClusterExecutorInfo) Equals(a interface{}) bool {
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
func (jbobject *SchedulerClusterExecutorInfo) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}


