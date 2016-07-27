package gospark

import "github.com/timob/javabind"

type SchedulerTaskSetInterface interface {
	JavaLangObjectInterface

	// public org.apache.spark.scheduler.Task<?>[] tasks()
	Tasks() []*SchedulerTask

	// public int stageId()
	StageId() int

	// public int stageAttemptId()
	StageAttemptId() int

	// public int priority()
	Priority() int

	// public java.util.Properties properties()
	Properties() *JavaUtilProperties

	// public java.lang.String id()
	Id() string
}

type SchedulerTaskSet struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.TaskSet(org.apache.spark.scheduler.Task<?>[], int, int, int, java.util.Properties)
func NewSchedulerTaskSet(a []*SchedulerTask, b int, c int, d int, e JavaUtilPropertiesInterface) (*SchedulerTaskSet) {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "org/apache/spark/scheduler/Task<?>")
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/TaskSet", conv_a.Value().Cast("org/apache/spark/scheduler/Task"), b, c, d, conv_e.Value().Cast("java/util/Properties"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_e.CleanUp()
	x := &SchedulerTaskSet{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.scheduler.Task<?>[] tasks()
func (jbobject *SchedulerTaskSet) Tasks() []*SchedulerTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "tasks", javabind.ObjectArrayType("org/apache/spark/scheduler/Task<?>"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "org/apache/spark/scheduler/Task<?>")
	dst := new([]*SchedulerTask)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int stageId()
func (jbobject *SchedulerTaskSet) StageId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int stageAttemptId()
func (jbobject *SchedulerTaskSet) StageAttemptId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "stageAttemptId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int priority()
func (jbobject *SchedulerTaskSet) Priority() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "priority", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.util.Properties properties()
func (jbobject *SchedulerTaskSet) Properties() *JavaUtilProperties {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "properties", "java/util/Properties")
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

// public java.lang.String id()
func (jbobject *SchedulerTaskSet) Id() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "id", "java/lang/String")
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
func (jbobject *SchedulerTaskSet) ToString() string {
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


