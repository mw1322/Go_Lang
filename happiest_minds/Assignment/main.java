import java.util.*;
public class main
{
	public static void main(String[] args) {
	    int[] m = {10, 20, 30};
	    int[] n = {100, 200, 300, 400 };
	    int[][] ll = {m,n};
	    m[0] = 99;
	   	System.out.println( Arrays.toString(m));
 		System.out.println( Arrays.deepToString(ll));
	}
}

// Output :
// [99, 20, 30]
// [[99, 20, 30], [100, 200, 300, 400]]
