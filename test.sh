cd fsys;
sh pretest.sh;
go test -coverprofile=coverage.out -v;
go tool cover -func=coverage.out;
sh posttest.sh;
