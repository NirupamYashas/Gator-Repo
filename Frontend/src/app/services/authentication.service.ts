import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { AngularFireAuth } from '@angular/fire/compat/auth';
import { ActivatedRoute, Router } from '@angular/router';
import firebase from 'firebase/compat/app';

import { User } from '../_models/user';

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

  public get currentUserValue(): User {
      return this.currentUserSubject.value;
  }

  login() {
    this.afAuth.signInWithRedirect(new firebase.auth.GoogleAuthProvider());
    this.router.navigateByUrl('/projects');
  }

  logout() {
    this.afAuth.signOut();
    this.router.navigateByUrl('/');
  }
}
