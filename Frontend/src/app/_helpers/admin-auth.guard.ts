import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthenticationService } from '../services/authentication.service';

@Injectable({
  providedIn: 'root'
})
export class AdminAuthGuard implements CanActivate {


  constructor(
    private router: Router,
    private authenticationService: AuthenticationService 
  ) {}

  
  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {

    let user = this.authenticationService.CurrentUser;
    console.log(user.data.email);
    if(user && user.data.isadmin){
      return true;
    }else{
      this.router.navigate(['/no-access']);
      return false;
    }
  }
  
}
