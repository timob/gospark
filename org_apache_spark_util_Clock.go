package gospark

import "github.com/timob/javabind"

type UtilClockInterface interface {

	// public abstract long getTimeMillis()
	GetTimeMillis() int64

	// public abstract long waitTillTime(long)
	WaitTillTime(a int64) int64
}

type UtilClock struct {
	JavaLangObject
}

// public abstract long getTimeMillis()
func (jbobject *UtilClock) GetTimeMillis() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTimeMillis", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public abstract long waitTillTime(long)
func (jbobject *UtilClock) WaitTillTime(a int64) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "waitTillTime", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}


