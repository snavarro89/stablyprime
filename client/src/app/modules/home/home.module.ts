import { NgModule } from '@angular/core';
import { SharedModule } from 'app/shared/shared.module';
import { Route, RouterModule } from '@angular/router';
import { HomeComponent } from './home.component';

const homeRoutes: Route[] = [
  {
      path     : '',
      component: HomeComponent
  }
];

@NgModule({
  declarations: [
    HomeComponent
  ],
  imports: [
    SharedModule,
    RouterModule.forChild(homeRoutes), 
  ]
})
export class HomeModule { }
