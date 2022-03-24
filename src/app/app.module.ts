import {RouterModule} from '@angular/router';
import { ProjectsService } from './projects/projects.service';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';

import { AppComponent } from './app.component';
import { ProjectsComponent } from './projects/projects.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { CreateComponent } from './create/create.component';
import { AvailableRoutes } from './app.routing';
import { NavbarComponent } from './navbar/navbar.component';
import {MatListModule} from '@angular/material/list';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { HomepageComponent } from './homepage/homepage.component';
import { ToolbarComponent } from './toolbar/toolbar.component';
import { AddprojectSuccessComponent } from './addproject-success/addproject-success.component';
import { MyProjectsComponent } from './my-projects/my-projects.component';

@NgModule({
  declarations: [
    AppComponent,
    ProjectsComponent,
    CreateComponent,
    NavbarComponent,
    LoginComponent,
    RegisterComponent,
    HomepageComponent,
    ToolbarComponent,
    AddprojectSuccessComponent,
    MyProjectsComponent,
  ],
  imports: [
    NgbModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    BrowserModule,
    BrowserAnimationsModule,
    MatListModule,
    RouterModule,
    RouterModule.forRoot(AvailableRoutes)
  ],
  providers: [
    ProjectsService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
