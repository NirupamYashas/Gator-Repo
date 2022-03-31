import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from "@angular/router";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginForm!: FormGroup;

  constructor(private router: Router,private http: HttpClient) { }

  ngOnInit(): void {
    this.loginForm = new FormGroup({
      useremail: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required])
    })
  }

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
      // this.http.post<any>('http://localhost:10000/students/', { Email: email, Password: password }).subscribe(data => { })
      // var user = email + " " + password
      // this.http.get('https://localhost:10000/user', {}).subscribe(data => {
      //       console.log(data)
            
      // })
      this.router.navigate(['/projects'])
      
    } else {
      console.log('There is a problem with the login form');
    }
  }

  resetForm(){
    this.loginForm.reset();
  }

}
