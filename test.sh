cd fsys;
sh pretest.sh;
go test -coverprofile=coverage.out -v;
status=$?;
sh posttest.sh;
if test $status -ne 0
then
	exit 1
fi
go tool cover -func=coverage.out;
status=$?;
if test $status -ne 0
then
	exit 1
fi
