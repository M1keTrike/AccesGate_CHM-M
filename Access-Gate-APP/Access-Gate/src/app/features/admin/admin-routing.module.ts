import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AdminDashboardComponent } from './components/admin-dashboard/admin-dashboard.component';
import { CreateUserComponent } from './components/create-user/create-user.component';
import { CreateEventComponent } from './components/create-event/create-event.component';
import { AssignGuestsComponent } from './components/assign-guests/assign-guests.component';
import { ViewUsersComponent } from './components/view-users/view-users.component';
import { ScanNfcComponent } from './components/scan-nfc/scan-nfc.component';
import { LinkDeviceComponent } from './components/link-device/link-device.component';

const routes: Routes = [
  { path: '', component: AdminDashboardComponent },
  { path: 'create-user', component: CreateUserComponent },
  { path: 'scan-nfc', component: ScanNfcComponent },
  { path: 'create-event', component: CreateEventComponent },
  { path: 'assign-guests', component: AssignGuestsComponent },
  { path: 'view-users', component: ViewUsersComponent },
  { path: 'link-device', component: LinkDeviceComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class AdminRoutingModule {}
