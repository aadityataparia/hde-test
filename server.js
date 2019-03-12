const server = require('express')();
server.use(require('serve-static')('public'));

server.listen(1111);
