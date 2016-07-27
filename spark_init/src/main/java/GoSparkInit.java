/* GoSparkInit.java */
import java.lang.System;

public class GoSparkInit {
  public static void main(String[] args) {
//	System.out.println(System.getProperty("java.library.path"));
	System.loadLibrary("gospark");
	Start();
  }

  public native static int Start();
}

