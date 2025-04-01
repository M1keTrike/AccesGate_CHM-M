import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

// Angular Material Modules
import { MatCardModule } from '@angular/material/card';
import { MatDividerModule } from '@angular/material/divider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatSelectModule } from '@angular/material/select';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatInputModule } from '@angular/material/input';

// Components
import { CreateUserComponent } from './components/create-user/create-user.component';
import { CreateEventComponent } from './components/create-event/create-event.component';
import { AssignGuestsComponent } from './components/assign-guests/assign-guests.component';
import { ViewUsersComponent } from './components/view-users/view-users.component';
import { AdminDashboardComponent } from './components/admin-dashboard/admin-dashboard.component';
import { AdminRoutingModule } from './admin-routing.module';
import { ScanNfcComponent } from './components/scan-nfc/scan-nfc.component';
import { LinkDeviceComponent } from './components/link-device/link-device.component';
import { LinkStatusComponent } from './components/link-status/link-status.component';
import { DevicesListComponent } from './components/devices-list/devices-list.component';
import { WifiCredentialsFormComponent } from './components/wifi-credentials-form/wifi-credentials-form.component';

import { FingerprintRegisterComponent } from './components/fingerprint-register/fingerprint-register.component';
import { FingerprintStatusComponent } from './components/fingerprint-status/fingerprint-status.component';
import { FingerprintActionComponent } from './components/fingerprint-action/fingerprint-action.component';

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
    FingerprintActionComponent,
    FingerprintStatusComponent,
    FingerprintRegisterComponent,
    FingerprintStatusComponent,
    FingerprintActionComponent,
  ],
  imports: [
    CommonModule,
    AdminRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    
    // Angular Material Modules
    MatCardModule,
    MatDividerModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatButtonModule,
    MatIconModule,
    MatListModule,
    MatProgressSpinnerModule,
    MatTooltipModule,
    MatSnackBarModule
  ],
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