const net = require('net');

const client = new net.Socket();
client.connect(8080, '127.0.0.1', () => {
    console.log('Connected');
    client.write(JSON.stringify({
        user: 'admin',
        password: 'password',
        action: 'create_table',
        name: 'test_table',
        columns: ['id', 'name', 'age']
    }) + '\n');

    client.write(JSON.stringify({
        user: 'admin',
        password: 'password',
        action: 'insert_row',
        name: 'test_table',
        row: ['1', 'Alice', '30']
    }) + '\n');

    client.write(JSON.stringify({
        user: 'admin',
        password: 'password',
        action: 'select_all',
        name: 'test_table'
    }) + '\n');
});

client.on('data', (data) => {
    console.log('Received: ' + data);
    // close the connection if the server sends 'exit'
    if (data.toString().includes('exit')) {
        client.destroy();
    }
});

client.on('close', () => {
    console.log('Connection closed');
});
