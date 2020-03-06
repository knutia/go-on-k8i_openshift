Echo off
ECHO Start Loop of %1

FOR /L %%G IN (1,1,600) DO (
	curl %1
	echo .
	PING -n 2 127.0.0.1>nul
)