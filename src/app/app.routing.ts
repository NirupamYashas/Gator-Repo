import { ProjectsComponent } from "./projects/projects.component";
import { CreateComponent } from "./create/create.component";
import { HomepageComponent } from "./homepage/homepage.component";
import { RegisterComponent } from "./register/register.component";
import { LoginComponent } from "./login/login.component";
import {ChangeAccountSettingsComponent} from "./change-account-settings/change-account-settings.component"
import {ProjectPageComponent} from "./project-page/project-page.component"

export const AvailableRoutes: any = [
    { path: "", component: HomepageComponent},
    { path: "projects", component: ProjectsComponent},
    { path: "projects/project", component: ProjectPageComponent},
    { path: "create", component: CreateComponent },
    { path: "homepage", component: HomepageComponent},
    { path: "register", component: RegisterComponent},
    { path: "login", component: LoginComponent},
    { path: "changeaccountsettings", component: ChangeAccountSettingsComponent},
];