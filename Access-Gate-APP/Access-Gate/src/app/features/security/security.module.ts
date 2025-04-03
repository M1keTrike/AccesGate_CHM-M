import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { SecurityRoutingModule } from './security-routing.module';
import { EventSecurityListComponent } from './components/event-security-list/event-security-list.component';
import { EventAttendeesControlComponent } from './components/event-attendees-control/event-attendees-control.component';

@NgModule({
  declarations: [
    EventSecurityListComponent,
    EventAttendeesControlComponent
  ],
  imports: [
    CommonModule,
    HttpClientModule,
    SecurityRoutingModule
  ],
  exports:[
    EventAttendeesControlComponent,
    EventSecurityListComponent
  ]
})
export class SecurityModule { }
