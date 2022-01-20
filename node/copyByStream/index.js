const fs = require('fs')

const sourceStream = fs.createReadStream('./data.txt')
const targetStream = fs.createWriteStream('./target.txt')
sourceStream.pipe(targetStream)

const { size } = fs.statSync('./data.txt')
let temp = 0
const current = Date.now()
console.log(size);
sourceStream.on('data', chunk => {
    temp+=chunk.length
    
})

function show() {
    if(temp < size) {
        const htemp = temp / 1024 / 1024 + 'M'
        const pro = Math.floor((temp / size)* 100) + '%'
        console.log(pro, htemp);
        setTimeout(show, 15)
    } else {
        console.log('copy success');
        const drs = Date.now()-current
        console.log(drs+ 'ms');
    }
}

setTimeout(() => {
    show()
}, 5)