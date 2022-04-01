import { Injectable } from '@angular/core';
import { HttpClient,HttpHeaders, HttpParams } from '@angular/common/http';

import { User } from '../_models/user';

@Injectable({
  providedIn: 'root'
})
export class UsersService {

  private _registerUrl = "http://localhost:8080/api/users/signup";
  private _loginUrl = "http://localhost:8080/api/users/login";

  constructor(private httpclient: HttpClient) { }

  addUser(createResource: any){

    //Headers
    const httpHeaders = new HttpHeaders();
    httpHeaders.append('content-type','application/json')

    return this.httpclient.post(this._registerUrl, createResource, {headers: httpHeaders});
  }

  authenticateUser(createResource: any){
    //Headers
    const httpHeaders = new HttpHeaders();
    httpHeaders.append('content-type','application/json')

    return this.httpclient.post(this._loginUrl, createResource, {headers: httpHeaders});
  }

  getAll() {
    return this.httpclient.get<User[]>('http://localhost:8080/api/users');
  }

  delete(id: number) {
    return this.httpclient.delete('http://localhost:8080/api/users/'+id);
  }
}
