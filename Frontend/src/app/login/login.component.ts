import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router,ActivatedRoute } from "@angular/router";
import { first } from 'rxjs/operators';

import { AuthenticationService } from '../services/authentication.service';
import { AlertService } from '../services/alert.service';
import { JwtHelperService } from '@auth0/angular-jwt';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginForm!: FormGroup;
  
  loading = false;
  submitted = false;
  returnUrl: string;
  invalidLogin = false;

  constructor(
    private router: Router,
    private route: ActivatedRoute,
    private http: HttpClient,
    private fb: FormBuilder,
    private authenticationService: AuthenticationService,
    private alertService: AlertService
    ) { }

  ngOnInit(): void {
    this.loginForm = this.fb.group({
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required])
    })

    // get return url from route parameters or default to '/'
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';
  }

  // convenience getter for easy access to form fields
  get f() { return this.loginForm.controls; }

  get loginEmailField(): any {
    return this.loginForm.get('email');
  }

  get loginPasswordField(): any {
    return this.loginForm.get('password');
  }

  loginFormSubmit(credentials: any): void {

    console.log(credentials);

    this.authenticationService.loginService(credentials);
    // .subscribe(result => {
    //   if(result){
    //     this.router.navigate(['/projects'])
    //   }else{
    //     this.invalidLogin = true;
    //     console.log('There is a problem with the login form');
    //   }
    // })
  }

  googleLogin(){
    this.authenticationService.googleLoginService();
  }

  resetForm(){
    this.loginForm.reset();
  }

}
