package gospark

import "github.com/timob/javabind"

type SchedulerStageInfoInterface interface {
	JavaLangObjectInterface

	// public int stageId()
	StageId() int

	// public int attemptId()
	AttemptId() int

	// public java.lang.String name()
	Name() string

	// public int numTasks()
	NumTasks() int

	// public java.lang.String details()
	Details() string

	// public void stageFailed(java.lang.String)
	StageFailed(a string) 

	// public java.lang.String getStatusString()
	GetStatusString() string
}

type SchedulerStageInfo struct {
	JavaLangObject
}

// public int stageId()
func (jbobject *SchedulerStageInfo) StageId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int attemptId()
func (jbobject *SchedulerStageInfo) AttemptId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "attemptId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String name()
func (jbobject *SchedulerStageInfo) Name() string {
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

// public int numTasks()
func (jbobject *SchedulerStageInfo) NumTasks() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numTasks", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String details()
func (jbobject *SchedulerStageInfo) Details() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "details", "java/lang/String")
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

// public void stageFailed(java.lang.String)
func (jbobject *SchedulerStageInfo) StageFailed(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stageFailed", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatusString()
func (jbobject *SchedulerStageInfo) GetStatusString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatusString", "java/lang/String")
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


