import { Component, OnInit } from '@angular/core';
import { first } from 'rxjs/operators';

import { User } from '../../_models/user';
import { AuthenticationService } from '../../services/authentication.service';
import { UsersService } from '../../services/users.service';

@Component({
  selector: 'app-admin-users',
  templateUrl: './admin-users.component.html',
  styleUrls: ['./admin-users.component.css']
})
export class AdminUsersComponent implements OnInit {
  currentUser: User;
  users:any = [];

  constructor(
    private authenticationService: AuthenticationService,
    private userService: UsersService
  ) { 
    // this.currentUser = this.authenticationService.currentUserValue;
  }

  ngOnInit(): void {
    this.loadAllUsers();
  }

  deleteUser(id: number) {
    this.userService.delete(id)
            .pipe(first())
            .subscribe(() => this.loadAllUsers());
  }

  public loadAllUsers() {
    this.userService.getAll()
        .pipe(first())
        .subscribe(users => {
          console.log(users);
          this.users = users;
        });
  }

}
