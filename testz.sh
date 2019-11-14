PASS=1
if [ -z $1 ]; then
    LOOPS=5
else
    LOOPS=$1
fi

if [ -z $2 ]; then
    SLP=5
else
    SLP=$2
fi

echo "Starting: " $LOOPS  "Loops with a "  $SLP "second delay"
date
while [ $PASS -le $LOOPS ]; do
    echo
    echo "Pass: "  $PASS
    zymkey -debug
    if [  $? == 0 ]; then
	exit -1
     fi
    sleep $SLP
    PASS=$(expr $PASS + 1)
done
echo


