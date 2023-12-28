// Install required packages:
// npm install express axios body-parser

const express = require('express');
const bodyParser = require('body-parser');
const axios = require('axios');

const app = express();
const port = 3000;

app.use(bodyParser.json());

app.post('/process', async (req, res) => {
  const data = req.body;

  // Forward data to Python service
  await axios.post('http://ip:9001/process', data);

  // Forward data to Go service
  await axios.post('http://ip:5000/process', data);

  res.json({ message: 'Request forwarded to Python and Go services.' });
});

app.listen(port, () => {
  console.log(`Node.js service is listening at http://3.108.218.107:${port}`);
});

