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
import { HttpClientModule } from '@angular/common/http';
import { LinkDeviceComponent } from './components/link-device/link-device.component';
import { LinkStatusComponent } from './components/link-status/link-status.component';
import { DevicesListComponent } from './components/devices-list/devices-list.component';
import { WifiCredentialsFormComponent } from './components/wifi-credentials-form/wifi-credentials-form.component';

@NgModule({
  declarations: [
    AdminDashboardComponent,
    CreateUserComponent,
    CreateEventComponent,
    AssignGuestsComponent,
    ViewUsersComponent,
    ScanNfcComponent,
    LinkDeviceComponent,
    LinkStatusComponent,
    DevicesListComponent,
    WifiCredentialsFormComponent,
  ],
  imports: [CommonModule, AdminRoutingModule, FormsModule, HttpClientModule],
  exports: [
    AdminDashboardComponent,
    CreateUserComponent,
    CreateEventComponent,
    AssignGuestsComponent,
    ViewUsersComponent,
    ScanNfcComponent,
  ],
})
export class AdminModule {}
