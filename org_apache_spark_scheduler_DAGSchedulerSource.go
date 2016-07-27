package gospark

import "github.com/timob/javabind"

type SchedulerDAGSchedulerSourceInterface interface {
	JavaLangObjectInterface

	// public org.apache.spark.scheduler.DAGScheduler dagScheduler()
	DagScheduler() *SchedulerDAGScheduler

	// public java.lang.String sourceName()
	SourceName() string
}

type SchedulerDAGSchedulerSource struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.DAGSchedulerSource(org.apache.spark.scheduler.DAGScheduler)
func NewSchedulerDAGSchedulerSource(a SchedulerDAGSchedulerInterface) (*SchedulerDAGSchedulerSource) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/DAGSchedulerSource", conv_a.Value().Cast("org/apache/spark/scheduler/DAGScheduler"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &SchedulerDAGSchedulerSource{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public org.apache.spark.scheduler.DAGScheduler dagScheduler()
func (jbobject *SchedulerDAGSchedulerSource) DagScheduler() *SchedulerDAGScheduler {
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

// public java.lang.String sourceName()
func (jbobject *SchedulerDAGSchedulerSource) SourceName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "sourceName", "java/lang/String")
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


