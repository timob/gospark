package gospark

import "github.com/timob/javabind"

type PartitionInterface interface {

	// public abstract int index()
	Index() int

	// public abstract int hashCode()
	HashCode() int
}

type Partition struct {
	*javabind.Callable
}

// public abstract int index()
func (jbobject *Partition) Index() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "index", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract int hashCode()
func (jbobject *Partition) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}


