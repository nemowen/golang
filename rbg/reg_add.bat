@echo off

set thisPath=%cd%
set serviceName=rbgServices

REG ADD HKLM\SYSTEM\CURRENTCONTROLSET\SERVICES\%serviceName% /f
REG ADD HKLM\SYSTEM\CURRENTCONTROLSET\SERVICES\%serviceName%\Parameters /f
REG ADD HKLM\SYSTEM\CURRENTCONTROLSET\SERVICES\%serviceName%\Parameters /f /v DisplayName /t REG_SZ /d %thisPath%\client.exe

echo "[Successful]"
pause