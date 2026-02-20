import { Router } from 'express';
import { registerUser } from '../controllers/user.controller.js';

const router = Router();
router.route('/register').post(registerUser);
// Assign Protocol 'HTTP' Method 'Post' for '../register' path/route

// HTTP methods 'GET'     =>  get data from server
//              'POST'    =>  create data in server/db
//              'PUT'     =>  update whole data
//              'PATCH'   =>  update some data
//              'DELETE'  =>  delete data

// HTTP status  '200'     =>  success                 =>  request was successful
//              '201'     =>  created                 =>  something new was made
//              '204'     =>  success*                =>  success but no response
//              '400'     =>  bad request             =>  wrong request or invalid input
//              '401'     =>  unathorized             =>  need to login in first
//              '403'     =>  forbidden               =>  user is not allowed to do this
//              '404'     =>  not found               =>  nothing here at this address
//              '409'     =>  conflict                =>  data already exists or clashes
//              '500'     =>  Internal Server Error   =>  server messed up (not user's fault)

export default router;
