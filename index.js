const path = require('path');
const express = require('express');
const app = express();

app.get('/', (req, res) => {
    const filename = req.query.filename || 'index.html';
    res.sendFile(path.join(__dirname, '/public/', filename));
});

app.listen(3000, () => {
    console.log('3000');
});