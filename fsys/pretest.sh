rm -rf tmp/;
mkdir tmp/;
mkdir tmp/readonly;
mkdir tmp/noaccess;
mkdir tmp/normal;
chmod 555 tmp/readonly;
chmod 400 tmp/noaccess;
chmod 777 tmp/normal;
