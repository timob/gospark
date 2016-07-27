gospark
=======

Go bindings for Apache Spark (https://spark.apache.org).

####Changes
* 2016-07-27: First very alpha version.

####Example
This is a Spark example from the quick start documentation: https://spark.apache.org/docs/latest/quick-start.html

```` go
func start(envPtr unsafe.Pointer) {
	javabind.SetupJVMFromEnv(envPtr)
	conf := gospark.NewSparkConf().SetAppName("Simple Application")
	inputFile := "YOUR_SPARK_HOME/README.md"
	sc := gospark.NewApiJavaJavaSparkContext2(conf)
	fileData := sc.TextFile(inputFile).Cache()
	numAs := fileData.Filter(newFuncStringBool(func(s string) bool {
		return strings.Contains(s, "a")
	})).Count()
	numBs := fileData.Filter(newFuncStringBool(func(s string) bool {
		return strings.Contains(s, "b")
	})).Count()
	fmt.Printf("Lines with a: %d, lines with b: %d\n", numAs, numBs)
} 
````

#####Building example
First edit the code in example/simple.go and change "YOUR_SPARK_HOME/README.md", to the path of the readme
in your spark install dir. The example counts lines in this file.

Steps to build shared library with gospark code that will be called from Java:
(note you man need to change the include path for jni.h)

````
cd example
go build --buildmode=c-shared
mv example libgospark.so
chmod +x libgospark.so
````

Steps to build Java jar with stub that calls above shared library and will be invoked by Spark. (requires maven)
````
cd spark_init
mvn package
````

Submit java jar to be run (spark-submit is in your Spark install bin directory):
````
export LD_LIBRARY_PATH=$(pwd)/example
export SPARK_CLASSPATH=$(pwd)/classpath
spark-submit  --class "GoSparkInit" --master 'local[2]' spark_init/target/gospark-init-1.0.jar
````

Should print "Lines with a: 58, lines with b: 26".