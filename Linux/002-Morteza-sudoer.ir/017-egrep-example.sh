______________________________________________________________________________
Example 1::  egrep -v "^#|^$" /etc/services
egrep -v "^#|^$" /etc/services | head  ------> Dont show command lines Or dont show empty lines
tcpmux		1/tcp				# TCP port service multiplexer
echo		7/tcp
echo		7/udp
discard		9/tcp		sink null
discard		9/udp		sink null
systat		11/tcp		users
daytime		13/tcp
daytime		13/udp
netstat		15/tcp
qotd		17/tcp		quote
