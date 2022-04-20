import { Component, OnChanges, OnInit, SimpleChanges } from '@angular/core';
import { FormBuilder, FormControl, FormGroup ,Validators } from '@angular/forms';
import { UsersService } from '../services/users.service';
import { Router,ActivatedRoute } from "@angular/router";
// import * as shajs from 'sha.js';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit{

  addUserForm! : FormGroup;
  returnUrl: string;

  constructor(public usersService: UsersService,private fb: FormBuilder,private router: Router,
    private route: ActivatedRoute,) { }

  ngOnInit(): void {

    
    /*
    this.addUserForm = new FormGroup({
      firstname: new FormControl('', [Validators.required]),
      lastname: new FormControl('', [Validators.required]),
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required])
    })
    */

    this.addUserForm = this.fb.group({
      firstname: new FormControl('', [Validators.required]),
      lastname: new FormControl('', [Validators.required]),
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required])
    });

    // get return url from route parameters or default to '/'
    this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/';

    //value changes method implementation for learning
    /*
    this.addUserForm.get('password')?.valueChanges.subscribe(data => {
      console.log(data);
    });
    */

    //method for setting values in the form right away when the page is loaded
    /*
    const newUserObj = {
      'firstname': 'yashas',
      'lastname': 'kuchimanchi',
      'email': 'example@gmail.com',
      'password':'password'
    }

    this.addUserForm.setValue(newUserObj);

    this.addUserForm.patchValue(newUserObj);
    */
  }

  addUser(){
    console.log(this.addUserForm.value);
    //Reading Form Values
    /*
    console.log(this.addUserForm.get('email')?.value);

    this.addUserForm.valueChanges.subscribe(data => {
       console.log(data);
    });
    */
  }

  addUserSubmit(credentials: any){

     console.log(credentials);
     this.usersService.addUser(credentials);

    //  var firstName = this.addUserForm.getRawValue().firstname;
    //  var lastName = this.addUserForm.getRawValue().lastname;
    //  var Email = this.addUserForm.getRawValue().email;
    //  var Password = this.addUserForm.getRawValue().password;
    // //  Password=shajs('shajs').update(Password).digest('hex')
    //  const newFormData = { firstname: firstName, lastname: lastName, email: Email , password: Password };
    //  console.log(newFormData);
    //  this.usersService.addUser(newFormData).subscribe(data => {})
  }

  resetForm(){
    this.addUserForm.reset();
  }

}
