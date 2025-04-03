import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';

// Angular Material Imports
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBarModule } from '@angular/material/snack-bar';

import { OrganizerRoutingModule } from './organizer-routing.module';
import { OrganizerDashboardComponent } from './components/organizer-dashboard/organizer-dashboard.component';
import { CreateEventComponent } from './components/create-event/create-event.component';
import { MyEventsComponent } from './components/my-events/my-events.component';
import { EventStatsComponent } from './components/event-stats/event-stats.component';
import { UpdateEventComponent } from './components/update-event/update-event.component';

@NgModule({
  declarations: [
    OrganizerDashboardComponent,
    CreateEventComponent,
    MyEventsComponent,
    EventStatsComponent,
    UpdateEventComponent
  ],
  imports: [
    CommonModule,
    RouterModule,
    ReactiveFormsModule,
    OrganizerRoutingModule,
    // Angular Material Modules
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatSnackBarModule
  ],
  exports: [
    OrganizerDashboardComponent,
    CreateEventComponent,
    MyEventsComponent,
    EventStatsComponent
  ]
})
export class OrganizerModule { }
