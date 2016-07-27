package gospark

import "github.com/timob/javabind"

type JavaUtilHashtableInterface interface {
	JavaUtilDictionaryInterface

	// public synchronized boolean contains(java.lang.Object)
	Contains(a interface{}) bool

	// public boolean containsValue(java.lang.Object)
	ContainsValue(a interface{}) bool

	// public synchronized boolean containsKey(java.lang.Object)
	ContainsKey(a interface{}) bool

	// public synchronized void putAll(java.util.Map<? extends K, ? extends V>)
	PutAll(a map[*JavaLangObject]*JavaLangObject) 

	// public synchronized void clear()
	Clear() 

	// public synchronized java.lang.Object clone()
	Clone() *JavaLangObject

	// public java.util.Set<K> keySet()
	KeySet() []*JavaLangObject

	// public java.util.Collection<V> values()
	Values() []*JavaLangObject
}

type JavaUtilHashtable struct {
	JavaUtilDictionary
}

// public java.util.Hashtable(int, float)
func NewJavaUtilHashtable3(a int, b float32) (*JavaUtilHashtable) {

	obj, err := javabind.GetEnv().NewObject("java/util/Hashtable", a, b)
	if err != nil {
		panic(err)
	}
	x := &JavaUtilHashtable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.Hashtable(int)
func NewJavaUtilHashtable2(a int) (*JavaUtilHashtable) {

	obj, err := javabind.GetEnv().NewObject("java/util/Hashtable", a)
	if err != nil {
		panic(err)
	}
	x := &JavaUtilHashtable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.Hashtable()
func NewJavaUtilHashtable() (*JavaUtilHashtable) {

	obj, err := javabind.GetEnv().NewObject("java/util/Hashtable")
	if err != nil {
		panic(err)
	}
	x := &JavaUtilHashtable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.Hashtable(java.util.Map<? extends K, ? extends V>)
func NewJavaUtilHashtable4(a map[*JavaLangObject]*JavaLangObject) (*JavaUtilHashtable) {
	conv_a := javabind.NewGoToJavaMap(javabind.NewGoToJavaCallable(), javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/util/Hashtable", conv_a.Value().Cast("java/util/Map"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaUtilHashtable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public synchronized int size()
func (jbobject *JavaUtilHashtable) Size() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "size", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public synchronized boolean isEmpty()
func (jbobject *JavaUtilHashtable) IsEmpty() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEmpty", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public synchronized java.util.Enumeration<K> keys()
func (jbobject *JavaUtilHashtable) Keys() *JavaUtilEnumeration {
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

// public synchronized java.util.Enumeration<V> elements()
func (jbobject *JavaUtilHashtable) Elements() *JavaUtilEnumeration {
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

// public synchronized boolean contains(java.lang.Object)
func (jbobject *JavaUtilHashtable) Contains(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "contains", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean containsValue(java.lang.Object)
func (jbobject *JavaUtilHashtable) ContainsValue(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "containsValue", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public synchronized boolean containsKey(java.lang.Object)
func (jbobject *JavaUtilHashtable) ContainsKey(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "containsKey", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public synchronized V get(java.lang.Object)
func (jbobject *JavaUtilHashtable) Get(a interface{}) *JavaLangObject {
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

// public synchronized V put(K, V)
func (jbobject *JavaUtilHashtable) Put(a JavaLangObjectInterface, b JavaLangObjectInterface) *JavaLangObject {
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

// public synchronized V remove(java.lang.Object)
func (jbobject *JavaUtilHashtable) Remove(a interface{}) *JavaLangObject {
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

// public synchronized void putAll(java.util.Map<? extends K, ? extends V>)
func (jbobject *JavaUtilHashtable) PutAll(a map[*JavaLangObject]*JavaLangObject)  {
	conv_a := javabind.NewGoToJavaMap(javabind.NewGoToJavaCallable(), javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "putAll", javabind.Void, conv_a.Value().Cast("java/util/Map"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized void clear()
func (jbobject *JavaUtilHashtable) Clear()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "clear", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public synchronized java.lang.Object clone()
func (jbobject *JavaUtilHashtable) Clone() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public synchronized java.lang.String toString()
func (jbobject *JavaUtilHashtable) ToString() string {
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

// public java.util.Set<K> keySet()
func (jbobject *JavaUtilHashtable) KeySet() []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "keySet", "java/util/Set")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoSet(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.Collection<V> values()
func (jbobject *JavaUtilHashtable) Values() []*JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "values", "java/util/Collection")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCollection(javabind.NewJavaToGoCallable())
	dst := new([]*JavaLangObject)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public synchronized boolean equals(java.lang.Object)
func (jbobject *JavaUtilHashtable) Equals(a interface{}) bool {
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

// public synchronized int hashCode()
func (jbobject *JavaUtilHashtable) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}


