package gospark

import "github.com/timob/javabind"

type JavaLangBooleanInterface interface {
	JavaLangObjectInterface

	// public static boolean parseBoolean(java.lang.String)
	ParseBoolean(a string) bool

	// public boolean booleanValue()
	BooleanValue() bool

	// public static java.lang.Boolean valueOf(boolean)
	ValueOf(a bool) bool

	// public static java.lang.Boolean valueOf(java.lang.String)
	ValueOf2(a string) bool

	// public static java.lang.String toString(boolean)
	ToString2(a bool) string

	// public static boolean getBoolean(java.lang.String)
	GetBoolean(a string) bool

	// public int compareTo(java.lang.Boolean)
	CompareTo(a bool) int

	// public int compareTo(java.lang.Object)
	CompareTo2(a interface{}) int
}

type JavaLangBoolean struct {
	JavaLangObject
}

// public java.lang.Boolean(boolean)
func NewJavaLangBoolean(a bool) (*JavaLangBoolean) {

	obj, err := javabind.GetEnv().NewObject("java/lang/Boolean", a)
	if err != nil {
		panic(err)
	}
	x := &JavaLangBoolean{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.Boolean(java.lang.String)
func NewJavaLangBoolean2(a string) (*JavaLangBoolean) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/lang/Boolean", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaLangBoolean{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static boolean parseBoolean(java.lang.String)
func (jbobject *JavaLangBoolean) ParseBoolean(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/Boolean", "parseBoolean", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean booleanValue()
func (jbobject *JavaLangBoolean) BooleanValue() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "booleanValue", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static java.lang.Boolean valueOf(boolean)
func (jbobject *JavaLangBoolean) ValueOf(a bool) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/Boolean", "valueOf", "java/lang/Boolean", a)
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static java.lang.Boolean valueOf(java.lang.String)
func (jbobject *JavaLangBoolean) ValueOf2(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/Boolean", "valueOf", "java/lang/Boolean", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static java.lang.String toString(boolean)
func (jbobject *JavaLangBoolean) ToString2(a bool) string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/Boolean", "toString", "java/lang/String", a)
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
func (jbobject *JavaLangBoolean) ToString() string {
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

// public int hashCode()
func (jbobject *JavaLangBoolean) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean equals(java.lang.Object)
func (jbobject *JavaLangBoolean) Equals(a interface{}) bool {
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

// public static boolean getBoolean(java.lang.String)
func (jbobject *JavaLangBoolean) GetBoolean(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/Boolean", "getBoolean", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int compareTo(java.lang.Boolean)
func (jbobject *JavaLangBoolean) CompareTo(a bool) int {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compareTo", javabind.Int, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int compareTo(java.lang.Object)
func (jbobject *JavaLangBoolean) CompareTo2(a interface{}) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compareTo", javabind.Int, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

func (jbobject *JavaLangBoolean) TRUE() bool {
	jret, err := javabind.GetEnv().GetStaticField("java/lang/Boolean", "TRUE", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

func (jbobject *JavaLangBoolean) SetFieldTRUE(val bool) {
	conv_val := javabind.NewGoToJavaBoolean()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/lang/Boolean", "TRUE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaLangBoolean) FALSE() bool {
	jret, err := javabind.GetEnv().GetStaticField("java/lang/Boolean", "FALSE", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

func (jbobject *JavaLangBoolean) SetFieldFALSE(val bool) {
	conv_val := javabind.NewGoToJavaBoolean()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/lang/Boolean", "FALSE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


