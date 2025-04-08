import { Component, OnInit } from '@angular/core';
import { EventService } from '../../services/eventService.service';
import { Event } from '../../models/Event';

@Component({
  selector: 'app-organizer-dashboard',
  templateUrl: './organizer-dashboard.component.html',
  styleUrls: ['./organizer-dashboard.component.css'],
  standalone: false,
})
export class OrganizerDashboardComponent implements OnInit {
  events: Event[] = [];

  constructor(private eventService: EventService) {}

  ngOnInit() {
    this.loadEvents();
  }

  private loadEvents() {
    // Assuming we can get the current user ID from a service
    const userId = 1; // Replace with actual user ID from auth service
    this.eventService.getEventsByCreator(userId).subscribe({
      next: (events) => {
        this.events = events;
        console.log('Events:', this.events);
      },
      error: (error) => {
        console.error('Error loading events:', error);
      }
    });
  }
}