const net = require('net')

const server = net.createServer(socket => {
    console.log('start');
    socket.on('data', data => {
        console.log(data.toString());
    })

    socket.write('hello')

    socket.on('end', () => {
        console.log('end');
    })
})

server.listen(888, () => console.log('up'))