import { NgModule } from '@angular/core';
import { SharedModule } from 'app/shared/shared.module';
import { MainLayoutModule } from 'app/layout/layouts/main-layout/main-layout.module';
import { LayoutComponent } from 'app/layout/layout.component';


@NgModule({
  declarations: [
    LayoutComponent
  ],
  imports: [
    SharedModule,
    MainLayoutModule
  ],
  exports: [
    MainLayoutModule
  ]
})
export class LayoutModule { }
