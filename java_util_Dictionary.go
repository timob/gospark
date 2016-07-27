package gospark

import "github.com/timob/javabind"

type JavaUtilDictionaryInterface interface {
	JavaLangObjectInterface

	// public abstract int size()
	Size() int

	// public abstract boolean isEmpty()
	IsEmpty() bool

	// public abstract java.util.Enumeration<K> keys()
	Keys() *JavaUtilEnumeration

	// public abstract java.util.Enumeration<V> elements()
	Elements() *JavaUtilEnumeration

	// public abstract V get(java.lang.Object)
	Get(a interface{}) *JavaLangObject

	// public abstract V put(K, V)
	Put(a JavaLangObjectInterface, b JavaLangObjectInterface) *JavaLangObject

	// public abstract V remove(java.lang.Object)
	Remove(a interface{}) *JavaLangObject
}

type JavaUtilDictionary struct {
	JavaLangObject
}

// public java.util.Dictionary()
func NewJavaUtilDictionary() (*JavaUtilDictionary) {

	obj, err := javabind.GetEnv().NewObject("java/util/Dictionary")
	if err != nil {
		panic(err)
	}
	x := &JavaUtilDictionary{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public abstract int size()
func (jbobject *JavaUtilDictionary) Size() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "size", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract boolean isEmpty()
func (jbobject *JavaUtilDictionary) IsEmpty() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEmpty", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract java.util.Enumeration<K> keys()
func (jbobject *JavaUtilDictionary) Keys() *JavaUtilEnumeration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "keys", "java/util/Enumeration")
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
	unique_x := &JavaUtilEnumeration{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract java.util.Enumeration<V> elements()
func (jbobject *JavaUtilDictionary) Elements() *JavaUtilEnumeration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "elements", "java/util/Enumeration")
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
	unique_x := &JavaUtilEnumeration{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract V get(java.lang.Object)
func (jbobject *JavaUtilDictionary) Get(a interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "get", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract V put(K, V)
func (jbobject *JavaUtilDictionary) Put(a JavaLangObjectInterface, b JavaLangObjectInterface) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "put", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract V remove(java.lang.Object)
func (jbobject *JavaUtilDictionary) Remove(a interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "remove", "java/lang/Object", conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}


