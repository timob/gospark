package gospark

import "github.com/timob/javabind"

type SchedulerSchedulerBackendInterface interface {

	// public abstract void start()
	Start() 

	// public abstract void stop()
	Stop() 

	// public abstract void reviveOffers()
	ReviveOffers() 

	// public abstract int defaultParallelism()
	DefaultParallelism() int

	// public abstract void killTask(long, java.lang.String, boolean)
	KillTask(a int64, b string, c bool) 

	// public abstract boolean isReady()
	IsReady() bool

	// public abstract java.lang.String applicationId()
	ApplicationId() string
}

type SchedulerSchedulerBackend struct {
	JavaLangObject
}

// public abstract void start()
func (jbobject *SchedulerSchedulerBackend) Start()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "start", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract void stop()
func (jbobject *SchedulerSchedulerBackend) Stop()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "stop", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract void reviveOffers()
func (jbobject *SchedulerSchedulerBackend) ReviveOffers()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "reviveOffers", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public abstract int defaultParallelism()
func (jbobject *SchedulerSchedulerBackend) DefaultParallelism() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "defaultParallelism", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract void killTask(long, java.lang.String, boolean)
func (jbobject *SchedulerSchedulerBackend) KillTask(a int64, b string, c bool)  {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "killTask", javabind.Void, a, conv_b.Value().Cast("java/lang/String"), c)
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()

}

// public abstract boolean isReady()
func (jbobject *SchedulerSchedulerBackend) IsReady() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isReady", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract java.lang.String applicationId()
func (jbobject *SchedulerSchedulerBackend) ApplicationId() string {
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


