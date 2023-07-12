const SerialPort = require('serialport')
const {DelimiterParser} = require('@serialport/parser-delimiter')

const port = new SerialPort({
    path:'COM3',
    baudRate: 9600
}
);
/*
const parser = new SerialPort.parser.Readline()

port.pipe(parser)

parser.on('data', (line)=>{
    console.log("el arduino dice: "+line)
    port.write('ñfljsdl ksdjflskd akjsdnlksaj')
})*/