const readLine = require('readline')
const process = require('process')
const http = require('http')
const cheerio = require('cheerio')

const rl = readLine.createInterface({
    input: process.stdin,
    output: process.stdout,
    prompt: 'search:'
})

rl.prompt()
rl.on('line', data => {
    if(data === 'exit') {
        process.exit(0)
    }
    search(data, () => rl.prompt())
}).on('close', () => {
    console.log('see you');
    process.exit(0)
})


function search(words, callback) {
    const req= http.request({
        hostname: 'www.baidu.com',
        port: 80,
        path: `/s?wd=${encodeURI(words)}`,
        method: 'GET'
    }, res => {
        let body=''
        res.setEncoding('utf8') // 设置编码
        res.on('data',chunk => {
            body+=chunk
        })
        res.on('end', ()=>{
            let $ = cheerio.load(body)
            $('a').each(function(i, el){
                console.log($(this).text(), $(this).attr('href'),'\n')
            })
            callback && callback()
        })
    })
    req.end()
}