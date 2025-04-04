import { Component, OnInit } from '@angular/core';
import { EventService } from '../../services/eventService.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Event } from '../../models/Event';
import { Router } from '@angular/router';

@Component({
  selector: 'app-my-events',
  templateUrl: './my-events.component.html',
  styleUrls: ['./my-events.component.css'],
  standalone: false
})
export class MyEventsComponent implements OnInit {
  events: Event[] = [];
  loading: boolean = true;

  constructor(
    private eventService: EventService,
    private router: Router
  ) {}

  ngOnInit() {
    this.loadEvents();
  }

  private getCurrentUserId(): number {
    const token = localStorage.getItem('Authorization');
    if (token) {
      const tokenPayload = JSON.parse(atob(token.split('.')[1]));
      return tokenPayload.user_id;
    }
    return 0;
  }

  private loadEvents() {
    const userId = this.getCurrentUserId();
    this.eventService.getEventsByCreator(userId).subscribe({
      next: (events) => {
        this.events = events;
        this.loading = false;
      },
      error: (error) => {
        console.error('Error loading events:', error);
        this.loading = false;
      }
    });
  }

  onUpdateEvent(eventId: number) {
    this.router.navigate(['/organizer/update-event', eventId]);
  }

  onDeleteEvent(eventId: number) {
    if (confirm('Are you sure you want to delete this event?')) {
      this.eventService.deleteEvent(eventId).subscribe({
        next: () => {
          this.events = this.events.filter(event => event.id !== eventId);
        },
        error: (error) => {
          console.error('Error deleting event:', error);
        }
      });
    }
  }
}
