import { Component } from '@angular/core';
import { LocationStrategy, PlatformLocation, Location } from '@angular/common';
import { AuthenticationService } from './services/authentication.service';
import { Router } from '@angular/router';
import { UsersService } from './services/users.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Gator-Repo';

  constructor(public location: Location,private userService:  UsersService, private auth: AuthenticationService, router: Router){
    auth.user$.subscribe(user => {
      if(user){

        let returnUrl: string | any = localStorage.getItem('returnUrl');
        router.navigateByUrl(returnUrl);
      }
    })
  }

  removeNavbar() {
    var title = this.location.prepareExternalUrl(this.location.path());
    title = title.slice(1);
    if(title === 'projects'){
        return true;
    }
    else {
        return false;
    }
  }
}
