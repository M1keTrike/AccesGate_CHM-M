import { Component, OnInit } from '@angular/core';
import { EventService } from '../../services/eventService.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Event } from '../../models/Event';
import { Router } from '@angular/router';
import { UsersService } from '../../../../services/Users.Service';

@Component({
  selector: 'app-my-events',
  templateUrl: './my-events.component.html',
  styleUrls: ['./my-events.component.css'],
  standalone: false
})
export class MyEventsComponent implements OnInit {
  events: Event[] = [];
  loading: boolean = true;
  isAttendee: boolean = false;
  isOrganizerRole: boolean = false;

  constructor(
    private eventService: EventService,
    private router: Router,
    private usersService: UsersService
  ) {}

  ngOnInit() {
    this.checkUserRole();
    this.loadEvents();
  }

  private checkUserRole() {
    const token = localStorage.getItem('Authorization');
    if (token) {
      const tokenPayload = JSON.parse(atob(token.split('.')[1]));
      this.isAttendee = tokenPayload.role === 'attendee';
      this.isOrganizerRole = tokenPayload.role === 'organizer';
    }
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
    this.loading = true;
    const currentUserId = this.getCurrentUserId();

    if (this.isAttendee|| this.isOrganizerRole) {
      this.usersService.getUserById(currentUserId).subscribe({
        next: (user) => {
          const creatorId = user.created_by ?? currentUserId;
          this.eventService.getAllEvents().subscribe({
            next: (events) => {
              this.events = events.filter(event => event.created_by === creatorId);
              this.loading = false;
            },
            error: (error) => {
              console.error('Error loading events:', error);
              this.loading = false;
            }
          });
        },
        error: (error) => {
          console.error('Error loading user:', error);
          this.loading = false;
        }
      });
    } else {
      this.eventService.getAllEvents().subscribe({
        next: (events) => {
          this.events = events.filter(event => event.created_by === currentUserId);
          this.loading = false;
        },
        error: (error) => {
          console.error('Error loading events:', error);
          this.loading = false;
        }
      });
    }
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
  logout() {
    this.usersService.logout();
    this.router.navigate(['/login']);
  }
}
