package gospark

import "github.com/timob/javabind"

type DependencyInterface interface {
	JavaLangObjectInterface

	// public abstract org.apache.spark.rdd.RDD<T> rdd()
	Rdd() *RddRDD
}

type Dependency struct {
	JavaLangObject
}

// public org.apache.spark.Dependency()
func NewDependency() (*Dependency) {

	obj, err := javabind.GetEnv().NewObject("org/apache/spark/Dependency")
	if err != nil {
		panic(err)
	}
	x := &Dependency{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public abstract org.apache.spark.rdd.RDD<T> rdd()
func (jbobject *Dependency) Rdd() *RddRDD {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "rdd", "org/apache/spark/rdd/RDD")
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
	unique_x := &RddRDD{}
	unique_x.Callable = dst
	return unique_x
}


