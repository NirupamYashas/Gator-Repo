import { Component } from '@angular/core';
import { LocationStrategy, PlatformLocation, Location } from '@angular/common';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Gator-Repo';

  constructor(public location: Location){
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
