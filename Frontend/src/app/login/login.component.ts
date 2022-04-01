import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router,ActivatedRoute } from "@angular/router";
import { first } from 'rxjs/operators';

import { AuthenticationService } from '../services/authentication.service';
import { AlertService } from '../services/alert.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginForm!: FormGroup;
  loginMsg: string;
  loading = false;
  submitted = false;
  returnUrl: string;

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
      useremail: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required])
    })

    // get return url from route parameters or default to '/'
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';
  }

  // convenience getter for easy access to form fields
  get f() { return this.loginForm.controls; }

  get loginEmailField(): any {
    return this.loginForm.get('useremail');
  }

  get loginPasswordField(): any {
    return this.loginForm.get('password');
  }

  loginFormSubmit(): void {

    if (this.loginForm.valid) {
      var email = this.loginForm.getRawValue().useremail;
      var password = this.loginForm.getRawValue().password;
      console.log(email,password)
      this.http.post<any>('http://localhost:8080/api/users/login', { email: email, password: password }).subscribe(data => {
        console.log(data);
        this.loginMsg = data.message;
        
        if(data.allow){
          alert(this.loginMsg);
          this.router.navigate(['/projects']);
        }else{
          alert(this.loginMsg);
          this.router.navigate(['/login']);
        }

      });
     
    } else {
      console.log('There is a problem with the login form');
    }
  }

  googleLogin(){
    this.authenticationService.login();
  }

  resetForm(){
    this.loginForm.reset();
  }

}
