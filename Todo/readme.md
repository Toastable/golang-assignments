## ** What is it? **

This simple Todo app is split into two components; a restful API and a web application. The API offers endpoints that enable the client to manage a list of todo items and the web application offers a nicer to look at view of said todo items. 

### **Things left to do**

There are two areas in which the application could be improved overall. The first being the way that static files are handled on the Web Server. There is a bug in Windows 10 and 11 which causes Golang to always override its internal MIME type map with values taken from the Registry. Those values are always text/plain for file types like .css and .js/.mjs etc which causes modern browsers to reject the files when they are sent by the server due to MIME type mismatches. 

See https://github.com/golang/go/issues/32350 for more information.

The second area in which the application could be improved is switching out the use of Mutexes in the in memory database with channels. Ideally the functions exposed by the in memory database should take in at least two channels to communicate over with the calling function which would eliminate the need to manage lock and unlock directly as it does in the API handlers. One channel would handle the communication of response data and the other would handle the communication of any errors.