Echo off
ECHO Start of Loop

FOR /L %%G IN (1,1,600) DO (
	curl http://localhost:8001/api/v1/namespaces/default/services/example-service/proxy/
	REM curl http://localhost:31081
	echo .
	PING -n 2 127.0.0.1>nul
)