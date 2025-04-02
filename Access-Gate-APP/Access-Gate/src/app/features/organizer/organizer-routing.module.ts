import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { OrganizerDashboardComponent } from './components/organizer-dashboard/organizer-dashboard.component';
import { CreateEventComponent } from './components/create-event/create-event.component';
import { MyEventsComponent } from './components/my-events/my-events.component';
import { ManageAttendeesComponent } from './components/manage-attendees/manage-attendees.component';
import { EventStatsComponent } from './components/event-stats/event-stats.component';

const routes: Routes = [
  { path: '', component: OrganizerDashboardComponent },
  { path: 'dashboard', component: OrganizerDashboardComponent },
  { path: 'create-event', component: CreateEventComponent },
  { path: 'my-events', component: MyEventsComponent },
  { path: 'manage-attendees', component: ManageAttendeesComponent },
  { path: 'event-stats', component: EventStatsComponent },
  { path: 'event/:id/edit', component: CreateEventComponent },
  { path: 'event/:id/attendees', component: ManageAttendeesComponent }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class OrganizerRoutingModule { }