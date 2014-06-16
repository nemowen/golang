@echo off

set thisPath=%cd%
set serviceName=RBGClient

net stop %serviceName%

.\ServiceTools\instsrv.exe %serviceName% remove

REG DELETE HKLM\SYSTEM\CurrentControlSet\services\%serviceName% /va /f

pause