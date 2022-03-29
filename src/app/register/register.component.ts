import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup ,Validators } from '@angular/forms';
import { UsersService } from '../services/users.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  addUserForm! : FormGroup;

  test : Date = new Date();
  focus: any;
  focus1: any;

  constructor(public usersService: UsersService) { }

  ngOnInit(): void {

    this.addUserForm = new FormGroup({
      firstname: new FormControl('', [Validators.required]),
      lastname: new FormControl('', [Validators.required]),
      email: new FormControl('', [Validators.required, Validators.email]),
      password: new FormControl('', [Validators.required])
    })
  }

  addUser(){
    console.log(this.addUserForm.value);
  }

  addUserSubmit(){
     var firstName = this.addUserForm.getRawValue().firstname;
     var lastName = this.addUserForm.getRawValue().lastname;
     var Email = this.addUserForm.getRawValue().email;
     var Password = this.addUserForm.getRawValue().password;
     const newFormData = { firstname: firstName, lastname: lastName, email: Email , password: Password };
     console.log(newFormData);
     this.usersService.addUser(newFormData).subscribe(data => {})
  }

}
