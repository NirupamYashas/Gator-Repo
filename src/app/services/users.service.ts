import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class UsersService {

  private _registerUrl = "http://localhost:8080/api/users/signup";
  private _loginUrl = "http://localhost:8080/api/users/login";

  constructor() { }
}
