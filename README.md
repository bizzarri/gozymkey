This is a small sample application for the Zymbit zymkey 4i.  It was written and tested and used on a Raspberry pi 3+.

I've also included a small shell script (testz.sh) to invoke it multiple times for testing.  I use it to make sure the Zymbit library is still working after updates.

zymkey.go without the -debug flag (as used in the shell script) waits for a tap from the user.  The shell script takes two arguments - loops through zymkey and time to wait between loops. So

./testz.sh 5 2

will loop through zymkey -debug 5 times waiting 2 seconds between loops, which is the default if testz is not called with command line args.

./zymkey

without the debug flag will wait for a tap from the user after the screen prompt.  It will also exercise the random byte features and display random bytes as well as the current date and time.  It will create small files called test.in and test.out and testagain.out to test out the random number feature.



Maurice Bizzarri
Bizzarri Software
maurice@bizzarrisoftware.com
