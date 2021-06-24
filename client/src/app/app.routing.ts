import { Route } from '@angular/router';
import { LayoutComponent } from 'app/layout/layout.component';


export const appRoutes: Route[] = [
  {
    path: '', pathMatch: 'full', redirectTo:'home'
  },
  {
    path: '',
    component: LayoutComponent,
    children: [
      {path: 'home', loadChildren: () => import('app/modules/home/home.module').then(m=>m.HomeModule)}
    ]
  }
];