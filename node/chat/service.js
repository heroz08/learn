const { Socket } = require('dgram');
const net = require('net')

const server = net.createServer()

let sockets = []

server.on('connection', socket => {
    console.log('new connect join!');

    sockets.push(socket)

    socket.on('data', data => {
        console.log('get data');
        sockets.forEach(other => {
            if(other !== socket)  {
                other.write(data)
            }
        })
    })

    socket.on('end', () => {
        console.log('somebody off');
        sockets.forEach(other => {
            if(other !== socket)  {
                other.write('one client closed')
            }
        })
        sockets = sockets.filter(other => other !== socket)
    })
})

server.on('error', err => {
    console.log(err);
})

server.on('close', () => {
    console.log('service closed');
})

server.listen(888, () => {
    console.log('service up');
})