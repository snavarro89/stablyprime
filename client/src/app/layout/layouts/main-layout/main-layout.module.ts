import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { SharedModule } from 'app/shared/shared.module';
import { HttpClientModule } from '@angular/common/http';
import { MainLayoutComponent } from 'app/layout/layouts/main-layout/main-layout.component';



@NgModule({
  declarations: [
    MainLayoutComponent
  ],
  imports: [
    RouterModule,
    SharedModule,
    HttpClientModule
  ],
  exports: [
    MainLayoutComponent
  ]
})
export class MainLayoutModule { }
