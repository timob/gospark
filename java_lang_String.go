package gospark

import "github.com/timob/javabind"

type JavaLangStringInterface interface {
	JavaLangObjectInterface

	// public int length()
	Length() int

	// public boolean isEmpty()
	IsEmpty() bool

	// public char charAt(int)
	CharAt(a int) uint16

	// public int codePointAt(int)
	CodePointAt(a int) int

	// public int codePointBefore(int)
	CodePointBefore(a int) int

	// public int codePointCount(int, int)
	CodePointCount(a int, b int) int

	// public int offsetByCodePoints(int, int)
	OffsetByCodePoints(a int, b int) int

	// public void getChars(int, int, char[], int)
	GetChars(a int, b int, c []uint16, d int) 

	// public void getBytes(int, int, byte[], int)
	GetBytes2(a int, b int, c []byte, d int) 

	// public byte[] getBytes()
	GetBytes() []byte

	// public boolean equalsIgnoreCase(java.lang.String)
	EqualsIgnoreCase(a string) bool

	// public int compareTo(java.lang.String)
	CompareTo2(a string) int

	// public int compareToIgnoreCase(java.lang.String)
	CompareToIgnoreCase(a string) int

	// public boolean regionMatches(int, java.lang.String, int, int)
	RegionMatches(a int, b string, c int, d int) bool

	// public boolean regionMatches(boolean, int, java.lang.String, int, int)
	RegionMatches2(a bool, b int, c string, d int, e int) bool

	// public boolean startsWith(java.lang.String, int)
	StartsWith2(a string, b int) bool

	// public boolean startsWith(java.lang.String)
	StartsWith(a string) bool

	// public boolean endsWith(java.lang.String)
	EndsWith(a string) bool

	// public int indexOf(int)
	IndexOf(a int) int

	// public int indexOf(int, int)
	IndexOf2(a int, b int) int

	// public int lastIndexOf(int)
	LastIndexOf(a int) int

	// public int lastIndexOf(int, int)
	LastIndexOf2(a int, b int) int

	// public int indexOf(java.lang.String)
	IndexOf3(a string) int

	// public int indexOf(java.lang.String, int)
	IndexOf4(a string, b int) int

	// public int lastIndexOf(java.lang.String)
	LastIndexOf3(a string) int

	// public int lastIndexOf(java.lang.String, int)
	LastIndexOf4(a string, b int) int

	// public java.lang.String substring(int)
	Substring(a int) string

	// public java.lang.String substring(int, int)
	Substring2(a int, b int) string

	// public java.lang.String concat(java.lang.String)
	Concat(a string) string

	// public java.lang.String replace(char, char)
	Replace(a uint16, b uint16) string

	// public boolean matches(java.lang.String)
	Matches(a string) bool

	// public java.lang.String replaceFirst(java.lang.String, java.lang.String)
	ReplaceFirst(a string, b string) string

	// public java.lang.String replaceAll(java.lang.String, java.lang.String)
	ReplaceAll(a string, b string) string

	// public java.lang.String[] split(java.lang.String, int)
	Split2(a string, b int) []string

	// public java.lang.String[] split(java.lang.String)
	Split(a string) []string

	// public java.lang.String toLowerCase()
	ToLowerCase() string

	// public java.lang.String toUpperCase()
	ToUpperCase() string

	// public java.lang.String trim()
	Trim() string

	// public char[] toCharArray()
	ToCharArray() []uint16

	// public static java.lang.String format(java.lang.String, java.lang.Object...)
	Format(a string, b ...*JavaLangObject) string

	// public static java.lang.String valueOf(java.lang.Object)
	ValueOf9(a interface{}) string

	// public static java.lang.String valueOf(char[])
	ValueOf3(a []uint16) string

	// public static java.lang.String valueOf(char[], int, int)
	ValueOf8(a []uint16, b int, c int) string

	// public static java.lang.String copyValueOf(char[], int, int)
	CopyValueOf2(a []uint16, b int, c int) string

	// public static java.lang.String copyValueOf(char[])
	CopyValueOf(a []uint16) string

	// public static java.lang.String valueOf(boolean)
	ValueOf(a bool) string

	// public static java.lang.String valueOf(char)
	ValueOf2(a uint16) string

	// public static java.lang.String valueOf(int)
	ValueOf6(a int) string

	// public static java.lang.String valueOf(long)
	ValueOf7(a int64) string

	// public static java.lang.String valueOf(float)
	ValueOf5(a float32) string

	// public static java.lang.String valueOf(double)
	ValueOf4(a float64) string

	// public native java.lang.String intern()
	Intern() string

	// public int compareTo(java.lang.Object)
	CompareTo(a interface{}) int
}

type JavaLangString struct {
	JavaLangObject
}

// public java.lang.String()
func NewJavaLangString() (*JavaLangString) {

	obj, err := javabind.GetEnv().NewObject("java/lang/String")
	if err != nil {
		panic(err)
	}
	x := &JavaLangString{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String(java.lang.String)
func NewJavaLangString9(a string) (*JavaLangString) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/lang/String", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaLangString{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String(char[])
func NewJavaLangString3(a []uint16) (*JavaLangString) {

	obj, err := javabind.GetEnv().NewObject("java/lang/String", a)
	if err != nil {
		panic(err)
	}
	x := &JavaLangString{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String(char[], int, int)
func NewJavaLangString6(a []uint16, b int, c int) (*JavaLangString) {

	obj, err := javabind.GetEnv().NewObject("java/lang/String", a, b, c)
	if err != nil {
		panic(err)
	}
	x := &JavaLangString{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String(int[], int, int)
func NewJavaLangString7(a []int, b int, c int) (*JavaLangString) {

	obj, err := javabind.GetEnv().NewObject("java/lang/String", a, b, c)
	if err != nil {
		panic(err)
	}
	x := &JavaLangString{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String(byte[], int, int, int)
func NewJavaLangString8(a []byte, b int, c int, d int) (*JavaLangString) {

	obj, err := javabind.GetEnv().NewObject("java/lang/String", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &JavaLangString{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String(byte[], int)
func NewJavaLangString4(a []byte, b int) (*JavaLangString) {

	obj, err := javabind.GetEnv().NewObject("java/lang/String", a, b)
	if err != nil {
		panic(err)
	}
	x := &JavaLangString{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String(byte[], int, int, java.lang.String) throws java.io.UnsupportedEncodingException
func NewJavaLangString11(a []byte, b int, c int, d string) (*JavaLangString, error) {
	conv_d := javabind.NewGoToJavaString()
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/lang/String", a, b, c, conv_d.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_d.CleanUp()
	x := &JavaLangString{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.lang.String(byte[], java.lang.String) throws java.io.UnsupportedEncodingException
func NewJavaLangString10(a []byte, b string) (*JavaLangString, error) {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/lang/String", a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_b.CleanUp()
	x := &JavaLangString{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.lang.String(byte[], int, int)
func NewJavaLangString5(a []byte, b int, c int) (*JavaLangString) {

	obj, err := javabind.GetEnv().NewObject("java/lang/String", a, b, c)
	if err != nil {
		panic(err)
	}
	x := &JavaLangString{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String(byte[])
func NewJavaLangString2(a []byte) (*JavaLangString) {

	obj, err := javabind.GetEnv().NewObject("java/lang/String", a)
	if err != nil {
		panic(err)
	}
	x := &JavaLangString{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int length()
func (jbobject *JavaLangString) Length() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "length", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean isEmpty()
func (jbobject *JavaLangString) IsEmpty() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEmpty", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public char charAt(int)
func (jbobject *JavaLangString) CharAt(a int) uint16 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "charAt", javabind.Char, a)
	if err != nil {
		panic(err)
	}
	return jret.(uint16)
}

// public int codePointAt(int)
func (jbobject *JavaLangString) CodePointAt(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "codePointAt", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int codePointBefore(int)
func (jbobject *JavaLangString) CodePointBefore(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "codePointBefore", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int codePointCount(int, int)
func (jbobject *JavaLangString) CodePointCount(a int, b int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "codePointCount", javabind.Int, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int offsetByCodePoints(int, int)
func (jbobject *JavaLangString) OffsetByCodePoints(a int, b int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "offsetByCodePoints", javabind.Int, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void getChars(int, int, char[], int)
func (jbobject *JavaLangString) GetChars(a int, b int, c []uint16, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getChars", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public void getBytes(int, int, byte[], int)
func (jbobject *JavaLangString) GetBytes2(a int, b int, c []byte, d int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "getBytes", javabind.Void, a, b, c, d)
	if err != nil {
		panic(err)
	}

}

// public byte[] getBytes(java.lang.String) throws java.io.UnsupportedEncodingException
func (jbobject *JavaLangString) GetBytes3(a string) ([]byte, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBytes", javabind.Byte | javabind.Array, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		var zero []byte
		return zero, err
	}
	conv_a.CleanUp()
	return jret.([]byte), nil
}

// public byte[] getBytes()
func (jbobject *JavaLangString) GetBytes() []byte {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBytes", javabind.Byte | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]byte)
}

// public boolean equals(java.lang.Object)
func (jbobject *JavaLangString) Equals(a interface{}) bool {
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

// public boolean equalsIgnoreCase(java.lang.String)
func (jbobject *JavaLangString) EqualsIgnoreCase(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equalsIgnoreCase", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int compareTo(java.lang.String)
func (jbobject *JavaLangString) CompareTo2(a string) int {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compareTo", javabind.Int, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int compareToIgnoreCase(java.lang.String)
func (jbobject *JavaLangString) CompareToIgnoreCase(a string) int {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compareToIgnoreCase", javabind.Int, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public boolean regionMatches(int, java.lang.String, int, int)
func (jbobject *JavaLangString) RegionMatches(a int, b string, c int, d int) bool {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "regionMatches", javabind.Boolean, a, conv_b.Value().Cast("java/lang/String"), c, d)
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	return jret.(bool)
}

// public boolean regionMatches(boolean, int, java.lang.String, int, int)
func (jbobject *JavaLangString) RegionMatches2(a bool, b int, c string, d int, e int) bool {
	conv_c := javabind.NewGoToJavaString()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "regionMatches", javabind.Boolean, a, b, conv_c.Value().Cast("java/lang/String"), d, e)
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	return jret.(bool)
}

// public boolean startsWith(java.lang.String, int)
func (jbobject *JavaLangString) StartsWith2(a string, b int) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "startsWith", javabind.Boolean, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean startsWith(java.lang.String)
func (jbobject *JavaLangString) StartsWith(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "startsWith", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean endsWith(java.lang.String)
func (jbobject *JavaLangString) EndsWith(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "endsWith", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *JavaLangString) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int indexOf(int)
func (jbobject *JavaLangString) IndexOf(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int indexOf(int, int)
func (jbobject *JavaLangString) IndexOf2(a int, b int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int lastIndexOf(int)
func (jbobject *JavaLangString) LastIndexOf(a int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "lastIndexOf", javabind.Int, a)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int lastIndexOf(int, int)
func (jbobject *JavaLangString) LastIndexOf2(a int, b int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "lastIndexOf", javabind.Int, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int indexOf(java.lang.String)
func (jbobject *JavaLangString) IndexOf3(a string) int {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int indexOf(java.lang.String, int)
func (jbobject *JavaLangString) IndexOf4(a string, b int) int {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "indexOf", javabind.Int, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int lastIndexOf(java.lang.String)
func (jbobject *JavaLangString) LastIndexOf3(a string) int {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "lastIndexOf", javabind.Int, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public int lastIndexOf(java.lang.String, int)
func (jbobject *JavaLangString) LastIndexOf4(a string, b int) int {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "lastIndexOf", javabind.Int, conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public java.lang.String substring(int)
func (jbobject *JavaLangString) Substring(a int) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "substring", "java/lang/String", a)
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

// public java.lang.String substring(int, int)
func (jbobject *JavaLangString) Substring2(a int, b int) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "substring", "java/lang/String", a, b)
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

// public java.lang.String concat(java.lang.String)
func (jbobject *JavaLangString) Concat(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "concat", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String replace(char, char)
func (jbobject *JavaLangString) Replace(a uint16, b uint16) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "replace", "java/lang/String", a, b)
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

// public boolean matches(java.lang.String)
func (jbobject *JavaLangString) Matches(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "matches", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public java.lang.String replaceFirst(java.lang.String, java.lang.String)
func (jbobject *JavaLangString) ReplaceFirst(a string, b string) string {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "replaceFirst", "java/lang/String", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String replaceAll(java.lang.String, java.lang.String)
func (jbobject *JavaLangString) ReplaceAll(a string, b string) string {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "replaceAll", "java/lang/String", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String[] split(java.lang.String, int)
func (jbobject *JavaLangString) Split2(a string, b int) []string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "split", javabind.ObjectArrayType("java/lang/String"), conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String[] split(java.lang.String)
func (jbobject *JavaLangString) Split(a string) []string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "split", javabind.ObjectArrayType("java/lang/String"), conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String toLowerCase()
func (jbobject *JavaLangString) ToLowerCase() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toLowerCase", "java/lang/String")
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

// public java.lang.String toUpperCase()
func (jbobject *JavaLangString) ToUpperCase() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toUpperCase", "java/lang/String")
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

// public java.lang.String trim()
func (jbobject *JavaLangString) Trim() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "trim", "java/lang/String")
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
func (jbobject *JavaLangString) ToString() string {
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

// public char[] toCharArray()
func (jbobject *JavaLangString) ToCharArray() []uint16 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toCharArray", javabind.Char | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]uint16)
}

// public static java.lang.String format(java.lang.String, java.lang.Object...)
func (jbobject *JavaLangString) Format(a string, b ...*JavaLangObject) string {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "java/lang/Object")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "format", "java/lang/String", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static java.lang.String valueOf(java.lang.Object)
func (jbobject *JavaLangString) ValueOf9(a interface{}) string {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "valueOf", "java/lang/String", conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static java.lang.String valueOf(char[])
func (jbobject *JavaLangString) ValueOf3(a []uint16) string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "valueOf", "java/lang/String", a)
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

// public static java.lang.String valueOf(char[], int, int)
func (jbobject *JavaLangString) ValueOf8(a []uint16, b int, c int) string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "valueOf", "java/lang/String", a, b, c)
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

// public static java.lang.String copyValueOf(char[], int, int)
func (jbobject *JavaLangString) CopyValueOf2(a []uint16, b int, c int) string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "copyValueOf", "java/lang/String", a, b, c)
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

// public static java.lang.String copyValueOf(char[])
func (jbobject *JavaLangString) CopyValueOf(a []uint16) string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "copyValueOf", "java/lang/String", a)
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

// public static java.lang.String valueOf(boolean)
func (jbobject *JavaLangString) ValueOf(a bool) string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "valueOf", "java/lang/String", a)
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

// public static java.lang.String valueOf(char)
func (jbobject *JavaLangString) ValueOf2(a uint16) string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "valueOf", "java/lang/String", a)
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

// public static java.lang.String valueOf(int)
func (jbobject *JavaLangString) ValueOf6(a int) string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "valueOf", "java/lang/String", a)
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

// public static java.lang.String valueOf(long)
func (jbobject *JavaLangString) ValueOf7(a int64) string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "valueOf", "java/lang/String", a)
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

// public static java.lang.String valueOf(float)
func (jbobject *JavaLangString) ValueOf5(a float32) string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "valueOf", "java/lang/String", a)
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

// public static java.lang.String valueOf(double)
func (jbobject *JavaLangString) ValueOf4(a float64) string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/lang/String", "valueOf", "java/lang/String", a)
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

// public native java.lang.String intern()
func (jbobject *JavaLangString) Intern() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "intern", "java/lang/String")
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

// public int compareTo(java.lang.Object)
func (jbobject *JavaLangString) CompareTo(a interface{}) int {
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


