import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { EventSecurityListComponent } from './components/event-security-list/event-security-list.component';
import { EventAttendeesControlComponent } from './components/event-attendees-control/event-attendees-control.component';

const routes: Routes = [
  {
    path: 'events',
    component: EventSecurityListComponent
  },
  {
    path: 'event-attendees/:id',
    component: EventAttendeesControlComponent
  },
  {
    path: '',
    redirectTo: 'events',
    pathMatch: 'full'
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SecurityRoutingModule { }