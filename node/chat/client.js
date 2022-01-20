const net  = require('net')
const process = require('process')

process.stdin.setEncoding = 'utf8'
const client = net.connect({port: 888},() => {
    console.log('client start');
    process.stdin.on('data', data => {
        console.log('发送消息：',data.toString());
        client.write(data)
        if(data.toString() === 'exit\n') {
            console.log(123);
            client.end()
        }
    })
})


client.on('data', data => {
    console.log('收到消息：', data.toString());
})

client.on('end', () => {
    console.log('离线');
    process.exit()
})