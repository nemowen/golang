@echo off

set thisPath=%cd%
set serviceName=SNRService

.\ServiceTools\instsrv.exe %serviceName% remove

REG DELETE HKLM\SYSTEM\CurrentControlSet\services\%serviceName% /va /f

pause