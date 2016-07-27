package gospark

import "github.com/timob/javabind"

type SchedulerActiveJobInterface interface {
	JavaLangObjectInterface

	// public int jobId()
	JobId() int

	// public org.apache.spark.scheduler.Stage finalStage()
	FinalStage() *SchedulerStage

	// public org.apache.spark.util.CallSite callSite()
	CallSite() *UtilCallSite

	// public org.apache.spark.scheduler.JobListener listener()
	Listener() *SchedulerJobListener

	// public java.util.Properties properties()
	Properties() *JavaUtilProperties

	// public int numPartitions()
	NumPartitions() int

	// public boolean[] finished()
	Finished() []bool

	// public int numFinished()
	NumFinished() int
}

type SchedulerActiveJob struct {
	JavaLangObject
}

// public org.apache.spark.scheduler.ActiveJob(int, org.apache.spark.scheduler.Stage, org.apache.spark.util.CallSite, org.apache.spark.scheduler.JobListener, java.util.Properties)
func NewSchedulerActiveJob(a int, b SchedulerStageInterface, c UtilCallSiteInterface, d SchedulerJobListenerInterface, e JavaUtilPropertiesInterface) (*SchedulerActiveJob) {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	conv_e := javabind.NewGoToJavaCallable()
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

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/scheduler/ActiveJob", a, conv_b.Value().Cast("org/apache/spark/scheduler/Stage"), conv_c.Value().Cast("org/apache/spark/util/CallSite"), conv_d.Value().Cast("org/apache/spark/scheduler/JobListener"), conv_e.Value().Cast("java/util/Properties"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	x := &SchedulerActiveJob{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int jobId()
func (jbobject *SchedulerActiveJob) JobId() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "jobId", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public org.apache.spark.scheduler.Stage finalStage()
func (jbobject *SchedulerActiveJob) FinalStage() *SchedulerStage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "finalStage", "org/apache/spark/scheduler/Stage")
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
	unique_x := &SchedulerStage{}
	unique_x.Callable = dst
	return unique_x
}

// public org.apache.spark.util.CallSite callSite()
func (jbobject *SchedulerActiveJob) CallSite() *UtilCallSite {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "callSite", "org/apache/spark/util/CallSite")
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

// public org.apache.spark.scheduler.JobListener listener()
func (jbobject *SchedulerActiveJob) Listener() *SchedulerJobListener {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "listener", "org/apache/spark/scheduler/JobListener")
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
	unique_x := &SchedulerJobListener{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.Properties properties()
func (jbobject *SchedulerActiveJob) Properties() *JavaUtilProperties {
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

// public int numPartitions()
func (jbobject *SchedulerActiveJob) NumPartitions() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numPartitions", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean[] finished()
func (jbobject *SchedulerActiveJob) Finished() []bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "finished", javabind.ObjectArrayType("boolean"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "boolean")
	dst := new([]bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int numFinished()
func (jbobject *SchedulerActiveJob) NumFinished() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "numFinished", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}


