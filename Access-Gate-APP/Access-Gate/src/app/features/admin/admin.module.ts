import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateUserComponent } from './components/create-user/create-user.component';
import { CreateEventComponent } from './components/create-event/create-event.component';
import { AssignGuestsComponent } from './components/assign-guests/assign-guests.component';
import { ViewUsersComponent } from './components/view-users/view-users.component';
import { AdminDashboardComponent } from './components/admin-dashboard/admin-dashboard.component';
import { AdminRoutingModule } from './admin-routing.module';
import { FormsModule } from '@angular/forms';
import { ScanNfcComponent } from './components/scan-nfc/scan-nfc.component';
import {  HttpClientModule } from '@angular/common/http';



@NgModule({
  declarations: [
    AdminDashboardComponent,
    CreateUserComponent,
    CreateEventComponent,
    AssignGuestsComponent,
    ViewUsersComponent,
    ScanNfcComponent
  ],
  imports: [
    CommonModule,
    AdminRoutingModule,
    FormsModule,
    HttpClientModule
  ]
})
export class AdminModule { }
