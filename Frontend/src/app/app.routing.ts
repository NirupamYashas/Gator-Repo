import { ProjectsComponent } from "./projects/projects.component";
import { CreateComponent } from "./create/create.component";
import { HomepageComponent } from "./homepage/homepage.component";
import { RegisterComponent } from "./register/register.component";
import { LoginComponent } from "./login/login.component";
import { AdminUsersComponent } from "./admin/admin-users/admin-users.component";
import {ChangeAccountSettingsComponent} from "./change-account-settings/change-account-settings.component"
import {ProjectPageComponent} from "./project-page/project-page.component"
import { AdminProjectsComponent } from "./admin/admin-projects/admin-projects.component";
import { AuthGuard } from './_helpers/auth.guard';
import { UserProfileComponent } from "./user-profile/user-profile.component";

export const AvailableRoutes: any = [
    { path: "", component: HomepageComponent},
    { path: "register", component: RegisterComponent},
    { path: "login", component: LoginComponent},
    { path: "projects", component: ProjectsComponent,canActivate:[AuthGuard]},
    { path: "projects/project", component: ProjectPageComponent,canActivate:[AuthGuard]},
    { path: "create", component: CreateComponent, canActivate:[AuthGuard] },
    { path: "admin/users", component: AdminUsersComponent, canActivate:[AuthGuard]},
    { path: "admin/projects", component: AdminProjectsComponent, canActivate:[AuthGuard]},
    { path: "user-profile", component: UserProfileComponent, canActivate:[AuthGuard]},
    { path: "changeaccountsettings", component: ChangeAccountSettingsComponent, canActivate:[AuthGuard]}
];