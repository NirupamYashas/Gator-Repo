import { Injectable } from '@angular/core';
import { HttpClient,HttpHeaders, HttpParams } from '@angular/common/http';
import { ActivatedRoute, Router } from '@angular/router';

import { User } from '../_models/user';

@Injectable({
  providedIn: 'root'
})
export class UsersService {

  private _registerUrl = "http://localhost:8080/api/users/signup";
  private _loginUrl = "http://localhost:8080/api/users/login";

  constructor(private httpclient: HttpClient,private route: ActivatedRoute) { }

  addUser(createResource: any){

    let returnUrl = this.route.snapshot.queryParamMap.get('returnUrl') || '/';
    localStorage.setItem('returnUrl',returnUrl);

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
