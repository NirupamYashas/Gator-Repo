import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, never, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { AngularFireAuth } from '@angular/fire/compat/auth';
import { ActivatedRoute, Router } from '@angular/router';
import firebase from 'firebase/compat/app';

import { User } from '../_models/user';
import { JwtHelperService } from '@auth0/angular-jwt';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private currentUserSubject: BehaviorSubject<User>;
  public currentUser: Observable<User>;

  user$: Observable<firebase.User|null>;

  constructor(private http: HttpClient,private afAuth: AngularFireAuth,private route: ActivatedRoute,public router: Router) {
    this.user$ = afAuth.authState;
  }

  // public get currentUserValue(): User {
  //     return this.currentUserSubject.value;
  // }

  googleLoginService() {
    let returnUrl = this.route.snapshot.queryParamMap.get('returnUrl') || '/';
    localStorage.setItem('returnUrl',returnUrl);

    this.afAuth.signInWithRedirect(new firebase.auth.GoogleAuthProvider());
    // this.router.navigateByUrl('/projects');
  }

  loginService(credentials: any){
    let returnUrl = this.route.snapshot.queryParamMap.get('returnUrl') || '/';
    localStorage.setItem('returnUrl',returnUrl);
    this.http.post<any>('http://localhost:8080/api/users/login',credentials)
    .subscribe(data => {
        console.log(data);
        if(data && data.token){
          localStorage.setItem('token',data.token);
          alert("Login Successful");
          this.router.navigate([returnUrl || '/']);
        }else{
          alert("Login Unsuccessful");
          this.router.navigate(['/login']);
        }
      });
  }

  logout() {
    localStorage.removeItem('token');
    // this.afAuth.signOut();
    this.router.navigateByUrl('/');
  }

  isLoggedIn(){
    let jwtHelper = new JwtHelperService();
    let token: any = localStorage.getItem('token');

    if(!token){
      return false;
    }
    
    let isExpired = jwtHelper.isTokenExpired(token);

    return !isExpired;
  }

  get CurrentUser(){
    let token = localStorage.getItem('token');
    if(!token) return null;

    return new JwtHelperService().decodeToken(token);

  }
}
