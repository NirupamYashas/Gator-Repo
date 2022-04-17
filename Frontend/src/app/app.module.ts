import {RouterModule} from '@angular/router';
import { ProjectsService } from './projects/projects.service';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';

import { environment } from 'src/environments/environment';
import { AngularFireModule } from '@angular/fire/compat';
import { AngularFireDatabaseModule } from '@angular/fire/compat/database'
import { AngularFireAuthModule } from '@angular/fire/compat/auth';

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
import { UserProfileComponent } from './user-profile/user-profile.component';
import { ProjectDescriptionComponent } from './project-description/project-description.component';
import { AdminProjectsComponent } from './admin/admin-projects/admin-projects.component';
import { AlertComponent } from './alert/alert.component';
import { AdminUsersComponent } from './admin/admin-users/admin-users.component';
import { ChangeAccountSettingsComponent } from './change-account-settings/change-account-settings.component';
import { ProjectPageComponent } from './project-page/project-page.component';
import { NoAccessComponent } from './no-access/no-access.component';

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
    UserProfileComponent,
    ProjectDescriptionComponent,
    AdminProjectsComponent,
    AlertComponent,
    AdminUsersComponent,
    ChangeAccountSettingsComponent,
    ProjectPageComponent,
    NoAccessComponent,
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
    RouterModule.forRoot(AvailableRoutes),
    AngularFireModule.initializeApp(environment.firebase),
    AngularFireDatabaseModule,
    AngularFireAuthModule
  ],
  providers: [
    ProjectsService
  ],
  schemas: [
    CUSTOM_ELEMENTS_SCHEMA
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
