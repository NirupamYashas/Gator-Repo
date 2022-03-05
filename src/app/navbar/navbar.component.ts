import { Component, OnInit } from '@angular/core';
import { ProjectsComponent } from '../projects/projects.component';


@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  selectedValue: string = '';
  projects: any;

  // constructor(public service: ProjectsComponent) { }

  //event handler for the select element's change event
  selectChangeHandler (event: any) {
    //update the ui
    this.selectedValue = event.target.value;
    // this.service.refreshProjects(this.selectedValue);
  }

  ngOnInit(): void {
  }

}
