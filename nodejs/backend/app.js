// app.js receives the request

import express from 'express';
import userRouter from './routes/user.route.js';
import postRouter from './routes/post.route.js';

// Start an express app
const app = express();

// Make the server able to parse json
app.use(express.json());

// Extend request to route file
app.use('/api/v1/users', userRouter);
app.use('/api/v1/posts', postRouter);

export default app;
