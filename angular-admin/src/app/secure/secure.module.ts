import { CUSTOM_ELEMENTS_SCHEMA, NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { NavComponent } from './nav/nav.component';
import { MenuComponent } from './menu/menu.component';
import { SecureComponent } from './secure.component';
import { RouterModule } from '@angular/router';
import { StarRatingComponent } from '../public/star-rating/star-rating.component';



@NgModule({
  declarations: [
    StarRatingComponent,
    NavComponent,
    MenuComponent,
    SecureComponent
  ],
  imports: [
    CommonModule,
    RouterModule
  ],
  exports: [
    StarRatingComponent
  ],
  schemas: [
    CUSTOM_ELEMENTS_SCHEMA
  ],
})
export class SecureModule { }
