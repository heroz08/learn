const net = require('net')

const client = net.connect({port: 888,host: '127.0.0.1'}, () => {
    console.log('connect');
    client.write('connect server')
})

client.on('data', data =>{
    console.log(data.toString());
    client.write('content')

    setTimeout(() => {
        client.end()
    }, 3000)
})

client.on('end', () => {
    console.log('close');
})